package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Rw  http.ResponseWriter
	Req *http.Request

	Method string
	Path   string
	Params map[string]string

	statuscode int

	// 中间件
	handlers []HandleFunc
	index    int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Rw:     w,
		Req:    req,
		Method: req.Method,
		Path:   req.URL.Path,

		index: -1,
	}
}

func (c *Context) Next() {
	c.index++

	length := len(c.handlers)
	for ; c.index < length; c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.statuscode = code
	c.Rw.WriteHeader(code)
}

func (c *Context) SetHeader(key, val string) {
	c.Rw.Header().Set(key, val)
}

func (c *Context) String(statuscode int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(statuscode)
	c.Rw.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) Json(statuscode int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(statuscode)

	if b, err := json.Marshal(obj); err == nil {
		c.Rw.Write(b)
	} else {
		http.Error(c.Rw, err.Error(), 500)
	}
}

func (c *Context) Data(statuscode int, data []byte) {
	c.Status(statuscode)
	c.Rw.Write(data)
}

func (c *Context) Html(statuscode int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(statuscode)
	c.Rw.Write([]byte(html))
}
