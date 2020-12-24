package gee

import "net/http"

type HandlerFunc func(c *Context)

type router struct {
	handlers map[string]HandlerFunc
}

func NewRouter() *router {
	return &router{handlers: map[string]HandlerFunc{}}
}

func (r *router) addRouter(method, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 Not Find %s\n", c.Path)
	}
}
