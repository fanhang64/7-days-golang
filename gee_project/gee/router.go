package gee

import (
	"fmt"
)

type Router struct {
	handlers map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		handlers: make(map[string]HandlerFunc, 10),
	}
}

func (r *Router) addRoute(method string, pattern string, handler HandlerFunc) {
	key := fmt.Sprintf("%s-%s", method, pattern)
	r.handlers[key] = handler
}

func (r *Router) handle(ctx *Context) {
	key := ctx.Req.Method + "-" + ctx.Req.URL.Path
	if v, ok := r.handlers[key]; ok {
		v(ctx)
	} else {
		ctx.Writer.WriteHeader(404)
		ctx.Writer.Write([]byte("page not found."))
	}
}
