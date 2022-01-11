package httprouter

import (
	"net/http"
)

// Handle is a function that can be registered to a route to handle HTTP
// requests. Like http.HandlerFunc, but has a third parameter for the values of
// wildcards (variables).
type Handle func(http.ResponseWriter, *http.Request, map[string]string)

// NotFound is the default HTTP handler func for routes that can't be matched
// with an existing route.
// NotFound tries to redirect to a canonical URL generated with CleanPath.
// Otherwise the request is delegated to http.NotFound.
func NotFound(w http.ResponseWriter, req *http.Request) {
	if req.Method != "CONNECT" {
		path := req.URL.Path
		if cp := CleanPath(path); cp != path && cp != req.Referer() {
			http.Redirect(w, req, cp, http.StatusMovedPermanently)
			return
		}
	}

	http.NotFound(w, req)
}

type Params interface {
}

// Router is a http.Handler which can be used to dispatch requests to different
// handler functions via configurable routes
type Router struct {
	node // embed the root node

	// Enables automatic redirection if the current route can't be matched but a
	// handler for the path with (without) the trailing slash exists.
	// For example if /foo/ is requested but a route only exists for /foo, the
	// client is redirected to /foo with http status code 301.
	RedirectTrailingSlash bool

	// Enables automatic redirection if the current route can't be matched but a
	// case-insensitive lookup of the path finds a handler.
	// The router then permanent redirects (http status code 301) to the
	// corrected path.
	// For example /FOO and /Foo could be redirected to /foo.
	RedirectCaseInsensitive bool

	// Configurable handler func which is used when no matching route is found.
	// Default is the NotFound func of this package.
	NotFound http.HandlerFunc

	// Handler func to handle panics recovered from http handlers.
	// It should be used to generate a error page and return the http error code
	// "500 - Internal Server Error".
	// The handler can be used to keep your server from crashing because of
	// unrecovered panics.
	PanicHandler func(http.ResponseWriter, *http.Request, interface{})
}

// Make sure the Router conforms with the http.Handler interface
var _ http.Handler = New()

// New returnes a new initialized Router.
// The router can be configured to also match the requested HTTP method or the
// requested Host.
func New() *Router {
	return &Router{
		RedirectTrailingSlash:   true,
		RedirectCaseInsensitive: true,
		NotFound:                NotFound,
	}
}

// GET is a shortcut for router.Handle("GET", path, handle)
func (r *Router) GET(path string, handle Handle) {
	r.Handle("GET", path, handle)
}

// POST is a shortcut for router.Handle("POST", path, handle)
func (r *Router) POST(path string, handle Handle) {
	r.Handle("POST", path, handle)
}

// PUT is a shortcut for router.Handle("PUT", path, handle)
func (r *Router) PUT(path string, handle Handle) {
	r.Handle("PUT", path, handle)
}

// PATCH is a shortcut for router.Handle("PATCH", path, handle)
func (r *Router) PATCH(path string, handle Handle) {
	r.Handle("PATCH", path, handle)
}

// DELETE is a shortcut for router.Handle("DELETE", path, handle)
func (r *Router) DELETE(path string, handle Handle) {
	r.Handle("DELETE", path, handle)
}

// Handle registers a new request handle with the given path and method.
//
// For GET, POST, PUT, PATCH and DELETE requests the respective shortcut
// functions can be used.
//
// This function is intended for bulk loading and to allow the usage of less
// frequently used, non-standardized or custom methods (e.g. for internal
// communication with a proxy).
func (r *Router) Handle(method, path string, handle Handle) {
	if path[0] != '/' {
		panic("path must begin with '/'")
	}
	r.addRoute(method, path, handle)
}

// HandlerFunc is an adapter which allows the usage of a http.HandlerFunc as a
// request handle.
func (r *Router) HandlerFunc(method, path string, handler http.HandlerFunc) {
	r.Handle(method, path,
		func(w http.ResponseWriter, req *http.Request, _ map[string]string) {
			handler(w, req)
		},
	)
}

// ServeFiles serves files from the given file system root.
// The path must end with "/*filepath", files are then served from the local
// path /defined/root/dir/*filepath.
// For example if root is "/etc" and *filepath is "passwd", the local file
// "/etc/passwd" would be served.
// Internally a http.FileServer is used, therefore http.NotFound is used instead
// of the Router's NotFound handler.
// To use the operating system's file system implementation,
// use http.Dir:
//     router.ServeFiles("/src/*filepath", http.Dir("/var/www"))
func (r *Router) ServeFiles(path string, root http.FileSystem) {
	if len(path) < 10 || path[len(path)-10:] != "/*filepath" {
		panic("path must end with /*filepath")
	}

	fileServer := http.FileServer(root)

	r.GET(path, func(w http.ResponseWriter, req *http.Request, vars map[string]string) {
		req.URL.Path = vars["filepath"]
		fileServer.ServeHTTP(w, req)
	})
}

func (r *Router) recv(w http.ResponseWriter, req *http.Request) {
	if rcv := recover(); rcv != nil {
		r.PanicHandler(w, req, rcv)
	}
}

// Make the router implement the http.Handler interface.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if r.PanicHandler != nil {
		defer r.recv(w, req)
	}

	path := req.URL.Path

	if handle, vars, tsr := r.getValue(req.Method, path); handle != nil {
		handle(w, req, vars)
		return
	} else if tsr && r.RedirectTrailingSlash && path != "/" {
		if path[len(path)-1] == '/' {
			path = path[:len(path)-1]
		} else {
			path = path + "/"
		}
		http.Redirect(w, req, path, http.StatusMovedPermanently)
		return
	} else if r.RedirectCaseInsensitive {
		fixedPath, found := r.findCaseInsensitivePath(req.Method, path, r.RedirectTrailingSlash)
		if found {
			http.Redirect(w, req, string(fixedPath), http.StatusMovedPermanently)
			return
		}
	}

	// Handle 404
	if r.NotFound != nil {
		r.NotFound(w, req)
	} else {
		http.NotFound(w, req)
	}
}
