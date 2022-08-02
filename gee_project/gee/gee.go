package gee

import (
	"net/http"
	"strings"
)

type HandleFunc func(c *Context)

type Engine struct {
	*RouterGroup // 将Engine作为最顶层的分组，也就是说Engine拥有RouterGroup所有的能力
	routers      *router
	groups       []*RouterGroup // 存储所有分组
}

type RouterGroup struct {
	prefix      string       // 组前缀
	middlewares []HandleFunc // 支持中间件
	parent      *RouterGroup // 支持嵌套
	engine      *Engine      // 指向engine
}

func New() *Engine {
	engine := &Engine{
		routers: newRouter(),
	}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup} // 存储所有分组，包含这个顶层的分组
	return engine
}

// addRoute 通过engine中添加路由
func (e *Engine) addRoute(pattern string, method string, handleFunc HandleFunc) {
	e.routers.addRoute(method, pattern, handleFunc)
}

func (e *Engine) GET(pattern string, handleFunc HandleFunc) {
	e.addRoute(pattern, "GET", handleFunc)
}

func (e *Engine) POST(pattern string, handleFunc HandleFunc) {
	e.addRoute(pattern, "POST", handleFunc)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

// Group 用于创建新的RouterGroup
func (rg *RouterGroup) Group(prefix string) *RouterGroup {
	engine := rg.engine
	newGroup := &RouterGroup{
		prefix: rg.prefix + prefix,
		parent: rg,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	engine.middlewares = make([]HandleFunc, 0)
	return newGroup
}

func (rg *RouterGroup) addRoute(method string, pattern string, handler HandleFunc) {
	totalPattern := rg.prefix + pattern
	rg.engine.routers.addRoute(method, totalPattern, handler)
}

func (rg *RouterGroup) GET(pattern string, handler HandleFunc) {
	rg.addRoute("GET", pattern, handler)
}

func (rg *RouterGroup) POST(pattern string, handler HandleFunc) {
	rg.addRoute("POST", pattern, handler)
}

func (rg *RouterGroup) Use(middlewares ...HandleFunc) {
	rg.middlewares = append(rg.middlewares, middlewares...)
}

func (e *Engine) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	var middlewares []HandleFunc
	for _, group := range e.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) { // 如果是以group.prefix作为前缀，获取这个分组添加中间件
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(rw, req)
	c.handlers = middlewares
	e.routers.handle(c)
}
