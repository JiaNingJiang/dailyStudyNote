1. 要安装其他Go版本，请运行Go-install命令，指定要安装的版本的下载位置。以下示例说明了版本1.13.8：

```go
go install golang.org/dl/go1.13.8@latest
```

 完成上述版本管理工具的下载之后,在GOPATH/bin目录下会看见一个go1.13.8.exe的可执行文件

接着在当前bin目录下(可以使用gitbash或者powershell在当前目录下打开)运行以下命令：

```go
 .\go1.13.8 download
```

2. 要使用新下载的版本运行go命令，请将版本号附加到go命令中(依旧需要在GOPATH/bin目录下运行)，如下所示：

```go
$ .\go1.13.8 version
go version go1.13.8 windows/amd64
```

3. 当您安装了多个版本时，您可以发现每个版本的安装位置，查看版本的GOROOT值。例如，运行以下命令：

```go
$ .\go1.13.8 env GOROOT
```

可以得到以下结果:

```go
C:\Users\hp-pc\sdk\go1.13.8
```



4. 下载完成之后，可以使用Goland使用不同的GOROOT运行项目

打开Goland --> File --> Settings --> GO --> GOROOT