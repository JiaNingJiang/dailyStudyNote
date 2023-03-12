package gee

import (
	"fmt"
	"net/http"
)

// 为用户准备的路由方法接口，用于自定义
type HandlerFunc func(*Context)

// Gee框架的核心,是一个路由表
type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	c := NewContext(w, req)
	e.router.handle(c)
}

// 为路由表添加一个路由
func (e *Engine) addRouter(method string, pattern string, handler HandlerFunc) {
	e.router.addRouter(method, pattern, handler)
}

// 添加一个GET路由
func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRouter("GET", pattern, handler)
}

// 添加一个POST路由
func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRouter("POST", pattern, handler)
}

// 启动Engine,运行Gee框架
func (e *Engine) Run(addr string) error {
	fmt.Printf("Gee Server is running in: %s\n", addr)
	return http.ListenAndServe(addr, e)
}
