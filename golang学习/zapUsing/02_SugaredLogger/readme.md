## Sugared Logger的简单使用
1. 大部分的实现与logger基本都相同。
2. 唯一的区别是，我们通过调用主 logger 的. Sugar()方法来获取一个SugaredLogger。
3. 然后使用SugaredLogger以printf(Debugf/Errorf/Infof)格式记录语句

下面是修改过后使用`SugaredLogger`代替`Logger`的代码：
```go
var sugarLogger *zap.SugaredLogger

func main() {
    InitLogger()
    defer sugarLogger.Sync()
    simpleHttpGet("www.baidu.com")
    simpleHttpGet("http://www.baidu.com")
}

// 通过调用主 logger 的.Sugar()方法来获取一个SugaredLogger。
func InitLogger() {
    logger, _ := zap.NewProduction()
    sugarLogger = logger.Sugar()
}

// 使用(Debugf/Errorf/Infof输出不同等级类型的日志)
func simpleHttpGet(url string) {
    sugarLogger.Debugf("Trying to hit GET request for %s", url)  //整体都将作为"msg"对应的key值
    resp, err := http.Get(url)
    if err != nil {
        sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
    } else {
        sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
        resp.Body.Close()
    }
}
```
当你执行上面的代码会得到如下输出：
```go
{"level":"error","ts":1668077414.3614414,"caller":"02_SugaredLogger/SugaredLogge
r.go:26","msg":"Error fetching URL www.baidu.com : Error = Get \"www.baidu.com\"
: unsupported protocol scheme \"\"","stacktrace":"main.simpleHttpGet\n\tC:/Users
/hp-pc/GolandProjects/zapUsing/02_SugaredLogger/SugaredLogger.go:26\nmain.main\n
\tC:/Users/hp-pc/GolandProjects/zapUsing/02_SugaredLogger/SugaredLogger.go:13\nr
untime.main\n\tD:/Go/src/runtime/proc.go:250"}

{"level":"info","ts":1668077414.3846388,"caller":"02_SugaredLogger/SugaredLogger
.go:28","msg":"Success! statusCode = 200 OK for URL http://www.baidu.com"}    
```
> 到目前为止这两个 logger 都打印输出 JSON 结构格式，这也是默认的输出格式