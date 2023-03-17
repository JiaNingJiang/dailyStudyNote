package gee

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// 为用户准备的路由方法接口，用于自定义
type HandlerFunc func(*Context)

// Gee框架的核心,是一个路由表
type Engine struct {
	*RouterGroup                // engine本身也可以作为最上层的分组，因此拥有RouterGroup的全部功能
	router       *router        // 核心，用于插入和查询路由
	groups       []*RouterGroup // store all groups
}

func New() *Engine {

	engine := &Engine{router: newRouter()}             // 核心路由
	engine.RouterGroup = &RouterGroup{engine: engine}  // engine拥有的最上层的group
	engine.groups = []*RouterGroup{engine.RouterGroup} // 所有的group包含engine本身就包含的最上层group

	return engine
}

// 为本次HTTP请求创建一个Context,执行其对应的路由HandlerFunc
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	var middlewares []HandlerFunc
	for _, group := range e.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) { // 获取对应组的全部中间件
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := NewContext(w, req)
	c.handlers = middlewares // 按顺序存储所有本组中间件
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

// 分组路由对象
type RouterGroup struct {
	prefix      string        // 当前的分组路由前缀
	middlewares []HandlerFunc // 本组支持的中间件
	engine      *Engine       // 所有的RouterGroup对象公用一个engine,用于添加路由(整个框架的所有资源都是由Engine统一协调的)
}

// 在当前group之下创建一个新的子group
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix, // 组前缀追加
		engine: engine,                // 共享同一个engine
	}
	engine.groups = append(engine.groups, newGroup) // group集合新加成员
	return newGroup
}

// 为当前group添加一个新的路由
func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp // 路由路径 = 组前缀+当前资源标识
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRouter(method, pattern, handler) // 仍然使用engine.router进行路由插入
}

// 以GET方式为当前group添加新路由
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// 以POST方式为当前group添加新路由
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

// 为当前group添加中间件
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}
