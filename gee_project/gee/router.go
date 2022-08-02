package gee

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node // 使用 roots 来存储每种请求方式的Trie 树根节点。
	handlers map[string]HandleFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node, 10),
		handlers: make(map[string]HandleFunc, 10),
	}
}

func parsePattern(pattern string) []string {
	s := strings.Split(pattern, "/") // /p/*doc   可以匹配 /p/static/aaa.css  /p/js/abc.js

	parts := make([]string, 0, 10)
	for _, item := range s {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}

		}
	}
	return parts
}

func (r *router) addRoute(method string, pattern string, handler HandleFunc) {
	log.Printf("register router: %s - %s\n", method, pattern)

	parts := parsePattern(pattern)

	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}

	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

/*
getRoute 函数中，还解析了:和*两种匹配符的参数，返回一个 map。
例如 /p/go/doc匹配到/p/:lang/doc，解析结果为：{lang: "go"}，
     /static/css/geektutu.css匹配到/static/*filepath，解析结果为{filepath: "css/geektutu.css"}
*/
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path) // /assets/abc.css  [ assets abc.css ]
	params := make(map[string]string)
	root, ok := r.roots[method] // roots[GET]
	if !ok {
		return nil, nil
	}
	n := root.search(searchParts, 0)
	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(parts) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	fmt.Printf("c.Path: %v\n", c.Path)

	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotExtended, "page not found\n")
		})
	}
	c.Next()
}
