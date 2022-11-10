## zap的简单使用

### 1.安装
```go
go get -u go.uber.org/zap
```

### 2. Zap Logger的介绍
```
Zap 提供了两种类型的日志记录器—Sugared Logger和Logger。

在性能很好但不是很关键的上下文中，使用SugaredLogger。它比其他结构化日志记录包快 4-10 倍，并且支持结构化和 printf 风格的日志记录。
在每一微秒和每一次内存分配都很重要的上下文中，使用Logger。它甚至比SugaredLogger更快，内存分配次数也更少，但它只支持强类型的结构化日志记录。
```

### 3.简单使用Logger记录器
```go
var logger *zap.Logger

func main() {
    InitLogger()
  defer logger.Sync()
    simpleHttpGet("www.baidu.com")
    simpleHttpGet("http://www.baidu.com")
}

// 创建一个logger(其实是production logger ,它会 默认记录调用函数信息、日期和时间等)
func InitLogger() {
    logger, _ = zap.NewProduction()
}

func simpleHttpGet(url string) {
    resp, err := http.Get(url)
    if err != nil {
        logger.Error(   //使用logger的Error记录日志
            "Error fetching url..",
            zap.String("url", url),
            zap.Error(err))
    } else {
        logger.Info("Success..",  //使用logger的Info记录日志
            zap.String("statusCode", resp.Status),
            zap.String("url", url))
        resp.Body.Close()
    }
}
```

日志记录器方法的语法是这样的：
```go
func (log *Logger) XXX(msg string, fields ...Field)
```

其中XXX是一个可变参数函数，可以是 Info / Error/ Debug / Panic 等。每个方法都接受一个消息字符串和任意数量的zapcore.Field场参数。

每个`zapcore.Field`其实就是一组键值对参数。
以上述的代码为例:
```go
 logger.Error(   //使用logger的Error记录日志
            "Error fetching url..",  // msg字段
            zap.String("url", url),  //第一个zapcore.Field(key为"url")
            zap.Error(err))  //第二个zapcore.Field (key为"err")
```
我们执行上面的代码会得到如下输出结果：
```go
{"level":"error","ts":1668076546.6231768,"caller":"01_ZapLogger/ZapLogger.go:24"
,"msg":"Error fetching url..","url":"www.baidu.com","error":"Get \"www.baidu.com
\": unsupported protocol scheme \"\"","stacktrace":"main.simpleHttpGet\n\tC:/Use
rs/hp-pc/GolandProjects/zapUsing/01_ZapLogger/ZapLogger.go:24\nmain.main\n\tC:/U
sers/hp-pc/GolandProjects/zapUsing/01_ZapLogger/ZapLogger.go:13\nruntime.main\n\
tD:/Go/src/runtime/proc.go:250"}

{"level":"info","ts":1668076546.6443026,"caller":"01_ZapLogger/ZapLogger.go:29",
"msg":"Success..","statusCode":"200 OK","url":"http://www.baidu.com"}
```

