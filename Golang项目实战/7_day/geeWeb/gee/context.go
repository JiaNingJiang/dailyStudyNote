package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	// response info
	StatusCode int
}

func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// 用于在POST提交的表单Form中查询key对应的value
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// 从url的Path中查询key对应的value
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// 设置响应码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// 为响应头设置自定义 key-value
func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

// 以字符串格式返回响应报文
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)

	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// 以JSON格式返回响应报文
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)

	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil { // json编码
		http.Error(c.Writer, err.Error(), 500)
	}
}

// 直接将字节流写入到响应报文
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// 以HTML格式返回响应报文
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
