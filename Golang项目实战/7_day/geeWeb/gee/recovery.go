package gee

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

// 追踪当前出错gorotine的函数栈信息，返回调用关系
func trace(message string) string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:]) // 跳过前三个(第一个是Callers/第二个是trace函数/第三个是defer recovery,因此从第四个开始)，返回从第四个开始的goroutine栈上的函数(程序)计数器，将其保存在pcs中

	var str strings.Builder // 字符串构建器
	str.WriteString(message + "\nTraceback:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)   // 通过 runtime.FuncForPC(pc) 获取对应的函数
		file, line := fn.FileLine(pc) // 通过 fn.FileLine(pc) 获取到调用该函数的文件名和行号
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}

// 将错误处理作为一个中间件投入使用
func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				c.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()

		c.Next()
	}
}
