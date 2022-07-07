package gee

import (
	"strings"
)

type Router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		roots:    make(map[string]*node, 10),
		handlers: make(map[string]HandlerFunc, 10),
	}
}

// only one * allowed
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *Router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)
	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *Router) getRoutes(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string, 10)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	node := root.search(searchParts, 0)
	if node != nil {
		parts := parsePattern(node.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return node, params
	}
	return nil, nil
}

func (r *Router) handle(ctx *Context) {
	node, params := r.getRoutes(ctx.Method, ctx.Path)

	if node != nil {
		ctx.Params = params
		key := ctx.Method + "-" + node.pattern
		r.handlers[key](ctx)
	} else {
		ctx.Writer.WriteHeader(404)
		ctx.Writer.Write([]byte("page not found."))
	}
}
