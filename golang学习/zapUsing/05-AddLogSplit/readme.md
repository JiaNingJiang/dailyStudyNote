## 使用 Lumberjack 进行日志切割归档

`Zap` 本身不支持切割归档日志文件，为了添加日志切割归档功能，我们将使用第三方库 `Lumberjack` 来实现。

### 1.安装
```go
go get -u github.com/natefinch/lumberjack
```
### 2. zap logger中加入Lumberjack
要在 `zap` 中加入 `Lumberjack` 支持，我们需要修改`WriteSyncer`代码。我们将按照下面的代码修改`getLogWriter()`函数：

修改前:
```go
func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./test1.log")
	return zapcore.AddSync(file)
}
```

修改后：
```go
func getLogWriter() zapcore.WriteSyncer {
    lumberJackLogger := &lumberjack.Logger{
        Filename:   "./test.log",
        MaxSize:    1,
        MaxBackups: 5,
        MaxAge:     30,
        Compress:   false,
    }
    return zapcore.AddSync(lumberJackLogger)
}
```
`Lumberjack Logger` 采用以下属性作为输入:
1. `Filename`: 日志文件的位置
2. `MaxSize`：在进行切割之前，日志文件的最大大小（以 `MB` 为单位）
3. `MaxBackups`：保留旧文件的最大个数
4. `MaxAges`：保留旧文件的最大天数
5. `Compress`：是否压缩 / 归档旧文件

同时，可以在`main`函数中循环记录日志，测试日志文件是否会自动切割和归档（日志文件每 `1MB` 会切割并且在当前目录下最多保存 `5` 个备份）。