package web

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type router struct {
	//  其中头部的信息就是key value 的形式， 然后可以使用interface 来进行记录， 然后处理方式可以采用  method +path 作为key
	roots map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots: make(map[string]*node),
		handlers: make(map[string]HandlerFunc)}
}
func parsePattern(pattern string)[]string{
	vs:=strings.Split(pattern,"/")
	parts:=make([]string,0)
	for _,item:=range vs{
		if item!=""{
			parts= append(parts, item)
			// 如果后边是 *  就不继续匹配

			if item[0]=='*'{
				break
			}
		}
	}
	return parts
}
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts:=parsePattern(pattern)
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	_,ok:=r.roots[method]
	if !ok{
		r.roots[method]=&node{}
	}
	fmt.Println("execute roots insert....")
	r.roots[method].insert(pattern,parts,0)

	r.handlers[key] = handler
}

func (r *router) getRoute( method string ,path string)(*node, map[string]string){
	searchparts:=parsePattern(path)
	params:=make(map[string]string)
	root,ok:=r.roots[method]
	if !ok{
		return nil, nil
	}
	n:=root.search(searchparts,0)
	if n!=nil{
		parts:=parsePattern(n.pattern)
		for index, part:=range parts{
			if part[0]==':'{
				params[part[1:]]=searchparts[index]
			}
			if part[0]=='*' && len(part)>1{
				params[part[1:]]=strings.Join(searchparts[index:],"/")
				break
			}
		}
		return  n,params
	}
	return nil, nil
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)

	if n != nil {
		key := c.Method + "-" + n.pattern
		c.Params = params
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	c.Next()}

