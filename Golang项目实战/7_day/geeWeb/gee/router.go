package gee

import (
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node, 0),
		handlers: make(map[string]HandlerFunc, 0),
	}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" { // pattern中非空的part才会进行存储
			parts = append(parts, item)
			if item[0] == '*' { // 如果当前part的首字符为*，那么当前part以及之后的part都不会存储(即是说当前的以通配符*开始的前一个part将作为叶子结点)
				break
			}
		}
	}
	return parts
}

func (r *router) addRouter(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

	key := method + "-" + pattern
	r.handlers[key] = handler

	_, ok := r.roots[method] // 查询对应的路由是否已经注册在roots(字典树)中
	if !ok {
		r.roots[method] = &node{} // 如果没有，则为该路由在roots中创建根节点
	}
	r.roots[method].insert(pattern, parts, 0) // 递归插入

}

func (r *router) handle(c *Context) {

	n, params := r.getRouter(c.Method, c.Path)

	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern // node节点的作用就是让路由路径的叶子结点保存路由的全路径(即pattern)
		//r.handlers[key](c)
		c.handlers = append(c.handlers, r.handlers[key]) // 并不是直接执行路由HandlerFunc,而是将HandlerFunc添加到中间件所在的c.handlers
	} else {
		//c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)

		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}

	c.Next() // 开始执行从第一个中间件开始的全部HandlerFunc
}

// 从r.roots中查询并返回指定方法(POST/GET)对应路由(path指定)的字典树node节点和模糊匹配映射关系
func (r *router) getRouter(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method] // 检查要查询的路由是否已经注册

	if !ok { // 没有注册，直接返回nil
		return nil, nil
	}

	n := root.search(searchParts, 0) // 递归查询

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index] // 记录模糊匹配映射关系，比如： /home/:Province/QingDao 与 /home/ShanDong/QingDao匹配了，那么映射为params[Province]QingDao
			}
			if part[0] == '*' && len(part) > 1 { // 记录模糊匹配映射关系, 比如：/home/*Province 与 /home/ShanDong/QingDao 匹配了，那么映射为params[Province]ShanDong/QingDao
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}
