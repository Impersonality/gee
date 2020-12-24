package gee

import (
	"net/http"
	"strings"
)

type HandlerFunc func(c *Context)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func NewRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func (r *router) addRouter(method, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	r.handlers[key] = handler

	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}

	parts := parsePattern(pattern)
	r.roots[method].insert(parts, pattern, 0)
}

func (r *router) getRouter(method, pattern string) (*node, map[string]string) {
	params := make(map[string]string)

	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	searchParts := parsePattern(pattern)
	n := root.search(searchParts, 0)
	if n != nil {
		parts := parsePattern(n.pattern)
		for k, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[k]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[k:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}

func (r *router) handle(c *Context) {
	n, params := r.getRouter(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + c.Path
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 Not Find %s\n", c.Path)
	}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, part := range vs {
		if part != "" {
			parts = append(parts, part)
			if part[0] == '*' {
				break
			}
		}
	}
	return parts
}
