package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H   map[string]interface{}
type  Context struct {
	Writer  http.ResponseWriter
	Req * http.Request
	Method string
	Path string
	Params map[string]string
	Statuscode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context{
	return  &Context{
		Writer : w,
		Req: req,
		Method: req.Method,
		Path: req.URL.Path,
	}
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}
func (c *Context) PostForm(key string)string{
	return c.Req.FormValue(key)
}
func (c *Context) Query(key string)string{
	return c.Req.URL.Query().Get(key)
}
func (c *Context)Status(code int){
	c.Statuscode=code
	c.Writer.WriteHeader(code)
}
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}
func (c *Context)String (code int  ,format string,value ...interface{}){
	c.SetHeader("Content-Type","text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format,value...)))
}

func (c *Context)JSON(code int ,obj interface{}){
	c.SetHeader("Content-Type","application/json")
	c.Status(code)
	encode:=json.NewEncoder(c.Writer)
	if err:=encode.Encode(obj);err!=nil{
		http.Error(c.Writer, err.Error(), 500)
	}


}
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}