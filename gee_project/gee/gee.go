package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, req *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := fmt.Sprintf("%s-%s", method, pattern)
	e.router[key] = handler
}

func (e *Engine) GET(pattern string, handlerFunc HandlerFunc) {
	e.addRoute("GET", pattern, handlerFunc)
}

func (e *Engine) POST(pattern string, handlerFunc HandlerFunc) {
	e.addRoute("POST", pattern, handlerFunc)
}

func (e *Engine) PUT(pattern string, handlerFunc HandlerFunc) {
	e.addRoute("PUT", pattern, handlerFunc)
}

func (e *Engine) DELETE(pattern string, handlerFunc HandlerFunc) {
	e.addRoute("DELETE", pattern, handlerFunc)
}

func New() *Engine {
	return &Engine{
		router: make(map[string]HandlerFunc, 10),
	}
}

func (e *Engine) Run(address string) error {
	return http.ListenAndServe(address, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if v, ok := e.router[key]; ok {
		v(w, req)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("page not found."))
	}
}
