package gee

import (
	"fmt"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc, 0),
	}
}

func (r *router) addRouter(method string, pattern string, handler HandlerFunc) {
	fmt.Printf("Route %4s - %s \t", method, pattern)
	key := method + "-" + pattern

	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
