package main

import (
	"geeWeb/gee"
	"log"
	"net/http"
	"time"
)

func onlyForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {

	r := gee.New()

	r.Use(gee.Logger())

	r.Use(gee.Recovery())

	r.GET("/index", func(c *gee.Context) { //  http://127.0.0.1:8001/index
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	r.GET("/panic", func(c *gee.Context) { // 检查错误处理中间件
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100]) //特意引入数组越界访问
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gee.Context) { //  http://127.0.0.1:8001/v1/
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello", func(c *gee.Context) { //  http://127.0.0.1:8001/v1/hello
			// expect /hello?name=geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name", func(c *gee.Context) { //  http://127.0.0.1:8001/v2/hello/:name
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *gee.Context) { //  http://127.0.0.1:8001/v2/login
			c.JSON(http.StatusOK, gee.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}
	r.Run(":8001")
}
