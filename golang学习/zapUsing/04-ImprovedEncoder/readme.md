### 1.将 JSON Encoder 更改为普通的 Log Encoder
现在，我们希望将编码器从`JSON Encoder` 更改为普通 `Encoder`。为此，我们需要将`NewJSONEncoder()`更改为`NewConsoleEncoder()`。

修改前:
```go
func getEncoder() zapcore.Encoder {
    return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}
```
修改后:
```go
func getEncoder() zapcore.Encoder {
    return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
}
```
当使用这些修改过的 `logger` 配置调用上述部分的`main()`函数时，以下输出将打印在文件——`test.log`中。
```shell
1.572161051846623e+09    debug    Trying to hit GET request for www.baidu.com
1.572161051846828e+09    error    Error fetching URL www.sogo.com : Error = Get www.baidu.com: unsupported protocol scheme ""
1.5721610518468401e+09    debug    Trying to hit GET request for http://www.baidu.com
1.572161052068744e+09    info    Success! statusCode = 200 OK for URL http://www.baidu.com
```
### 2.更改时间编码并添加调用者详细信息
鉴于我们对配置所做的更改，有下面两个问题：

    时间是以非人类可读的方式展示，例如 1.572161051846623e+09
    调用方函数的详细信息没有显示在日志中

我们要做的第一件事是覆盖默认的`ProductionConfig()`，并进行以下更改:
1. 修改时间编码器
2. 在日志文件中使用大写字母记录日志级别
```go
func getEncoder() zapcore.Encoder {
    encoderConfig := zap.NewProductionEncoderConfig()
    encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder  //修改编码器时间格式
    encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder //修改编码器在日志文件中使用大写字母记录日志级别
    return zapcore.NewConsoleEncoder(encoderConfig)
}
```

接下来，我们将修改 `zap logger` 代码，**添加将调用函数信息记录到日志中的功能**。为此，我们将在`zap.New(..)`函数中添加一个`Option`。
修改前：
```go
logger := zap.New(core)
```
修改后：
```go
logger := zap.New(core, zap.AddCaller())
```
当使用这些修改过的 `logger` 配置调用上述部分的`main()`函数时，以下输出将打印在文件——`test.log`中。
```shell
2022-11-10T19:18:05.744+0800	DEBUG	04-ImprovedEncoder/ImprovedSugaredLogger.go:40	Trying to hit GET request for www.baidu.com
2022-11-10T19:18:05.749+0800	ERROR	04-ImprovedEncoder/ImprovedSugaredLogger.go:43	Error fetching URL www.baidu.com : Error = Get "www.baidu.com": unsupported protocol scheme ""
2022-11-10T19:18:05.749+0800	DEBUG	04-ImprovedEncoder/ImprovedSugaredLogger.go:40	Trying to hit GET request for http://www.baidu.com
2022-11-10T19:18:05.772+0800	INFO	04-ImprovedEncoder/ImprovedSugaredLogger.go:45	Success! statusCode = 200 OK for URL http://www.baidu.com
```
