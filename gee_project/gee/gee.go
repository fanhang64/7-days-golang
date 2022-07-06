package gee

import (
	"net/http"
)

type HandlerFunc func(ctx *Context)

type Engine struct {
	router *Router
}

func New() *Engine {
	return &Engine{
		router: NewRouter(),
	}
}

func (e *Engine) GET(pattern string, handlerFunc HandlerFunc) {
	e.router.addRoute("GET", pattern, handlerFunc)
}

func (e *Engine) POST(pattern string, handlerFunc HandlerFunc) {
	e.router.addRoute("POST", pattern, handlerFunc)
}

func (e *Engine) PUT(pattern string, handlerFunc HandlerFunc) {
	e.router.addRoute("PUT", pattern, handlerFunc)
}

func (e *Engine) DELETE(pattern string, handlerFunc HandlerFunc) {
	e.router.addRoute("DELETE", pattern, handlerFunc)
}

func (e *Engine) Run(address string) error {
	return http.ListenAndServe(address, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := NewContext(w, req)
	e.router.handle(ctx)
}
