## 使用自定义的logger

### 1.将日志写入文件而不是终端
我们将使用`zap.New(…)`方法来手动传递所有配置，而不是使用像`zap.NewProduction()`这样的预置方法来创建 logger。
```go
func New(core zapcore.Core, options ...Option) *Logger
```
`zapcore.Core`需要三个配置——`Encoder`，`WriteSyncer`，`LogLevel`。
1. `Encoder`: 编码器 (写入日志格式)。我们将使用开箱即用的`NewJSONEncoder()`，并使用预先设置的`ProductionEncoderConfig()`。
```go
zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())  //输出的日志将为json类型
```
2. `WriterSyncer` ：指定日志写到哪里去。我们使用`zapcore.AddSync()`函数并且将打开的文件句柄传进去。
```go
file, _ := os.Create("./test.log")
writeSyncer := zapcore.AddSync(file)
```
3. `Log Level`：决定哪种级别以上的日志才会被写入。

我们将修改 `Logger` 代码，并重写`InitLogger()`方法。其余方法—`main() SimpleHttpGet()`保持不变。
```go
func InitLogger() {
    writeSyncer := getLogWriter()
    encoder := getEncoder()
    core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)  //允许DebugLevel及以上等级的日志写入日志文件

    logger := zap.New(core)
    sugarLogger = logger.Sugar()
}

// 获取日志编码器(JSON格式)
func getEncoder() zapcore.Encoder {
    return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

// 获取日志写出设备
func getLogWriter() zapcore.WriteSyncer {
    file, _ := os.Create("./test.log")
    return zapcore.AddSync(file)
}
```
当使用这些修改过的 `logger` 配置调用上述部分的`main()`函数时，以下输出将打印在文件——`test.log`中。
```go
{"level":"debug","ts":1668078506.4109528,"msg":"Trying to hit GET request for www.baidu.com"}
{"level":"error","ts":1668078506.4109528,"msg":"Error fetching URL www.baidu.com : Error = Get \"www.baidu.com\": unsupported protocol scheme \"\""}
{"level":"debug","ts":1668078506.4109528,"msg":"Trying to hit GET request for http://www.baidu.com"}
{"level":"info","ts":1668078506.4332914,"msg":"Success! statusCode = 200 OK for URL http://www.baidu.com"}
```

