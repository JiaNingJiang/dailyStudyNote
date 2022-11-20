## 一、RPC入门

RPC是远程过程调用的简称，是分布式系统中不同节点间流行的通信方式。在互联网时代，RPC已经和IPC一样成为一个不可或缺的基础构件。因此Go语言的标准库也提供了一个简单的RPC实现，我们将以此为入口学习RPC的各种用法。

### 1. RPC版”Hello, World”

Go语言的RPC包的路径为net/rpc，也就是放在了net包目录下面。因此我们可以猜测该RPC包是建立在net包基础之上的。

我们先构造一个HelloService类型，其中的Hello方法用于实现打印功能：

```go
type HelloService struct {}

func (p *HelloService) Hello(request string, reply *string) error {
    *reply = "hello:" + request
    return nil
}
```

其中Hello方法必须满足Go语言的RPC规则：**方法只能有两个可序列化的参数(一个是请求，一个是回应)**，其中**第二个参数(也就是回应)必须是指针类型**，并且**返回一个error类型**，同时**必须是公开(方法名大写)的方法**。

然后就可以将HelloService类型的对象注册为一个RPC服务：

```go
// RPC服务端代码
func main() {
    rpc.RegisterName("HelloService", new(HelloService))

    listener, err := net.Listen("tcp", ":1234")
    if err != nil {
        log.Fatal("ListenTCP error:", err)
    }

    conn, err := listener.Accept()
    if err != nil {
        log.Fatal("Accept error:", err)
    }

    rpc.ServeConn(conn)
}
```

其中**rpc.Register函数调用会将传入的对象类型中所有满足RPC规则的对象方法注册为RPC函数，所有注册的方法会放在 “HelloService”服务空间之下**。然后我们建立一个唯一的TCP链接，并且**通过rpc.ServeConn函数在该TCP链接上为对方提供RPC服务**。

下面是客户端请求HelloService服务的代码：

```go
// RPC客户端代码
func main() {
    client, err := rpc.Dial("tcp", "localhost:1234")
    if err != nil {
        log.Fatal("dialing:", err)
    }

    var reply string
    err = client.Call("HelloService.Hello", "hello", &reply)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(reply)
}
```

首先是**通过rpc.Dial拨号RPC服务**，然后**通过client.Call调用具体的RPC方法**。在调用client.Call时，**第一个参数是用点号链接的RPC服务名字和方法名字**，**第二和第三个参数分别我们定义RPC方法的两个参数**。

### 2. 更安全的RPC接口

在涉及RPC的应用中，作为开发人员一般至少有三种角色：首先是服务端实现RPC方法的开发人员，其次是客户端调用RPC方法的人员，最后也是最重要的是制定服务端和客户端RPC接口规范的设计人员。在前面的例子中我们为了简化将以上几种角色的工作全部放到了一起，虽然看似实现简单，但是不利于后期的维护和工作的切割。

如果要**重构HelloService服务**，第一步需要**明确服务的名字和接口**：

```go
//重构rpc server端
const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceInterface = interface {
    Hello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
    return rpc.RegisterName(HelloServiceName, svc)
}
```

我们将RPC服务的接口规范分为三个部分：**首先是服务的名字**，然后是**服务要实现的详细方法列表**，最后是**注册该类型服务的函数**。为了避免名字冲突，我们在RPC服务的名字中增加了包路径前缀（这个是RPC服务抽象的包路径，并非完全等价Go语言的包路径）。**RegisterHelloService注册服务时，编译器会要求传入的对象满足HelloServiceInterface接口**。

在定义了RPC服务接口规范之后，客户端就可以根据规范编写RPC调用的代码了：

```go
// 重构之前的rpc client
func main() {
    client, err := rpc.Dial("tcp", "localhost:1234")
    if err != nil {
        log.Fatal("dialing:", err)
    }

    var reply string
    err = client.Call(HelloServiceName+".Hello", "hello", &reply)
    if err != nil {
        log.Fatal(err)
    }
}
```

**对于客户端来说唯一的变化是client.Call的第一个参数用HelloServiceName+”.Hello”代替了”HelloService.Hello”**。然而通过client.Call函数调用RPC方法依然比较繁琐，同时参数的类型依然无法得到编译器提供的安全保障。

为了简化客户端用户调用RPC函数，我们在可以**在接口规范部分增加对客户端的简单包装**：

```go
// 重构rpc client
type HelloServiceClient struct {
    *rpc.Client
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)   //此行代码只是保证HelloServiceClient实现了HelloServiceInterface

func DialHelloService(network, address string) (*HelloServiceClient, error) {
    c, err := rpc.Dial(network, address)
    if err != nil {
        return nil, err
    }
    return &HelloServiceClient{Client: c}, nil
}

// 实现了HelloServiceInterface接口的Hello方法
func (p *HelloServiceClient) Hello(request string, reply *string) error {
    return p.Client.Call(HelloServiceName+".Hello", request, reply)
}
```

我们在接口规范中针对客户端新增加了HelloServiceClient类型，该类型也必须实现HelloServiceInterface接口，这样客户端用户就可以直接通过接口对应的方法调用RPC函数。同时**提供了一个DialHelloService方法，直接拨号HelloService服务**。

基于新的客户端接口，我们可以简化客户端用户的代码：

```go
// 重构之后的rpc client
func main() {
    client, err := DialHelloService("tcp", "localhost:1234")
    if err != nil {
        log.Fatal("dialing:", err)
    }

    var reply string
    err = client.Hello("hello", &reply)
    if err != nil {
        log.Fatal(err)
    }
}
```

现在**客户端用户不用再担心RPC方法名字或参数类型不匹配等低级错误的发生**。

最后是基于RPC接口规范编写真实的服务端代码：

```go
// 重构之后的rpc server
type HelloService struct {}

func (p *HelloService) Hello(request string, reply *string) error {
    *reply = "hello:" + request
    return nil
}

func main() {
    RegisterHelloService(new(HelloService))

    listener, err := net.Listen("tcp", ":1234")
    if err != nil {
        log.Fatal("ListenTCP error:", err)
    }

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Fatal("Accept error:", err)
        }

        go rpc.ServeConn(conn)
    }
}
```

在新的RPC服务端实现中，我们**用RegisterHelloService函数来注册函数**，这样不仅可以**避免命名服务名称的工作**，同时也**保证了传入的服务对象满足了RPC接口的定义**。最后我们**新的服务改为支持多个TCP链接，然后为每个TCP链接提供RPC服务。**

### 3. 跨语言的RPC

**标准库的RPC默认采用Go语言特有的gob编码**，因此**从其它语言调用Go语言实现的RPC服务将比较困难**。在互联网的微服务时代，每个RPC以及服务的使用者都可能采用不同的编程语言，因此跨语言是互联网时代RPC的一个首要条件。得益于RPC的框架设计，Go语言的RPC其实也是很容易实现跨语言支持的。

**Go语言的RPC框架有两个比较有特色的设计**：**一个是RPC数据打包时可以通过插件实现自定义的编码和解码**；**另一个是RPC建立在抽象的io.ReadWriteCloser接口之上的，我们可以将RPC架设在不同的通讯协议之上**。这里我们将尝试通过官方自带的net/rpc/jsonrpc扩展实现一个跨语言的RPC。

#### 3.1 基于jsonrpc的服务端

首先是**基于json编码重新实现RPC服务端**：

```go
func main() {
    rpc.RegisterName("HelloService", new(HelloService))

    listener, err := net.Listen("tcp", ":1234")
    if err != nil {
        log.Fatal("ListenTCP error:", err)
    }

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Fatal("Accept error:", err)
        }

        go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
    }
}
```

代码中最大的变化是**用rpc.ServeCodec函数替代了rpc.ServeConn函数**，传入的参数是**针对服务端的json编解码器**。

#### 3.2 基于jsonrpc的客户端

然后是**实现json版本的客户端**：

```go
func main() {
    conn, err := net.Dial("tcp", "localhost:1234")
    if err != nil {
        log.Fatal("net.Dial:", err)
    }

    client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

    var reply string
    err = client.Call("HelloService.Hello", "hello", &reply)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(reply)
}
```

先**手工调用net.Dial函数建立TCP链接**，然后**基于该链接建立针对客户端的json编解码器**。

#### 3.3 jsonRpc消息格式

##### 3.3.1 nc模拟rpc服务端接收rpc客户端消息

**在确保客户端可以正常调用RPC服务的方法之后**，我们**用一个普通的TCP服务代替Go语言版本的RPC服务**，这样**可以查看客户端调用时发送的数据格式**。比如**通过nc命令`nc -l 1234`在同样的端口启动一个TCP服务**。然后再次执行一次客户端的RPC调用将会发现nc输出了以下的信息：

```go
{"method":"HelloService.Hello","params":["hello"],"id":0}
```

这是一个json编码的数据，其中**method部分对应要调用的rpc服务和方法组合成的名字**，**params部分的第一个元素为参数**，**id是由调用端维护的一个唯一的调用编号**。

**发送的请求request**的json数据对象在内部对应两个结构体：**客户端是clientRequest，服务端是serverRequest**。clientRequest和serverRequest结构体的内容基本是一致的：

```go
type clientRequest struct {
    Method string         `json:"method"`
    Params [1]interface{} `json:"params"`
    Id     uint64         `json:"id"`
}

type serverRequest struct {
    Method string           `json:"method"`
    Params *json.RawMessage `json:"params"`
    Id     *json.RawMessage `json:"id"`
}
```

##### 3.3.2  nc模拟rpc服务端回复rpc客户端

在获取到RPC调用对应的json数据后，我们可以通过**直接向架设了RPC服务的TCP服务器发送json数据模拟RPC方法调用**：

```
$ echo -e '{"method":"HelloService.Hello","params":["hello"],"id":1}' | nc localhost 1234
```

rpc客户端收到的返回结果也是一个json格式的数据：

```json
{"id":1,"result":"hello:hello","error":null}
```

其中id对应输入的id参数，result为返回的结果，error部分在出问题时表示错误信息。**对于顺序调用来说，id不是必须的。但是Go语言的RPC框架支持异步调用，当返回结果的顺序和调用的顺序不一致时，可以通过id来识别对应的调用。**

**获取的回应response**的json数据也是对应内部的两个结构体：**客户端是clientResponse，服务端是serverResponse**。两个结构体的内容同样也是类似的：

```go
type clientResponse struct {
    Id     uint64           `json:"id"`
    Result *json.RawMessage `json:"result"`
    Error  interface{}      `json:"error"`
}

type serverResponse struct {
    Id     *json.RawMessage `json:"id"`
    Result interface{}      `json:"result"`
    Error  interface{}      `json:"error"`
}
```

因此无论采用何种语言，只要遵循同样的json结构，以同样的流程就可以和Go语言编写的RPC服务进行通信。这样我们就实现了跨语言的RPC。

### 4. Http上的RPC

Go语言内在的RPC框架已经支持在Http协议上提供RPC服务。但是框架的http服务同样采用了内置的gob协议，并且没有提供采用其它协议的接口，因此从其它语言依然无法访问的。在前面的例子中，我们已经实现了在TCP协议之上运行jsonrpc服务，并且通过nc命令行工具成功实现了RPC方法调用。现在我们尝试在http协议上提供jsonrpc服务。

新的RPC服务其实是一个类似REST规范的接口，接收请求并采用相应处理流程：

```go
func main() {
    rpc.RegisterName("HelloService", new(HelloService))

    http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
        var conn io.ReadWriteCloser = struct {
            io.Writer
            io.ReadCloser
        }{
            ReadCloser: r.Body,
            Writer:     w,
        }

        rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
    })

    http.ListenAndServe(":1234", nil)
}
```

RPC的服务架设在“/jsonrpc”路径，在处理函数中基于http.ResponseWriter和http.Request类型的参数构造一个io.ReadWriteCloser类型的conn通道。然后**基于conn构建针对服务端的json编码解码器**。最后**通过rpc.ServeRequest函数为每次请求处理一次RPC方法调用**。

模拟一次RPC调用的过程就是向该链接发送一个json字符串：

```
$ curl localhost:1234/jsonrpc -X POST \
    --data '{"method":"HelloService.Hello","params":["hello"],"id":0}'
```

返回的结果依然是json字符串：

```json
{"id":0,"result":"hello:hello","error":null}
```

这样就可以很方便地从不同语言中访问RPC服务了。

## 二、Protobuf

Protobuf是Protocol  Buffers的简称，它是Google公司开发的一种数据描述语言，并于2008年对外开源。Protobuf刚开源时的定位类似于XML、JSON等数据描述语言，通过附带工具生成代码并实现将结构化数据序列化的功能。但是我们更关注的是Protobuf作为接口规范的描述语言，可以作为设计安全的跨语言PRC接口的基础工具。

### 1. Protobuf入门

针对Protobuf的具体使用，请见gRpc.md。

这里，我们创建了一个hello.proto文件，其中包装HelloService服务中用到的字符串类型：

```protobuf
syntax = "proto3";

package main;

message String {
    string value = 1;
}
```

使用protoc-gen-go工具将其编译为go文件(hello.pb.go)：

```go
type String struct {
    Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *String) Reset()         { *m = String{} }
func (m *String) String() string { return proto.CompactTextString(m) }
func (*String) ProtoMessage()    {}
func (*String) Descriptor() ([]byte, []int) {
    return fileDescriptor_hello_069698f99dd8f029, []int{0}
}

func (m *String) GetValue() string {
    if m != nil {
        return m.Value
    }
    return ""
}
```

生成的结构体中还会包含一些以`XXX_`为名字前缀的成员，这里隐藏了这些成员。

同时String类型还自动生成了一组方法，其中**ProtoMessage方法表示这是一个实现了proto.Message接口的方法**。

此外**Protobuf还为每个成员生成了一个Get方法**，Get方法不仅可以处理空指针类型，而且可以和Protobuf第二版的方法保持一致（第二版的自定义默认值特性依赖这类方法）。



基于新的String类型，我们可以**重新实现下面的HelloService服务**：

```go
type HelloService struct{}

func (p *HelloService) Hello(request *String, reply *String) error {
    reply.Value = "hello:" + request.GetValue()
    return nil
}
```

Hello方法的**输入参数和输出的参数均改用Protobuf定义的String类型**(message，其实就是一个结构体)表示。**因为新的输入参数为结构体类型，因此改用指针类型作为输入参数**，函数的内部代码同时也做了相应的调整。



下面更新hello.proto文件，通过Protobuf来定义HelloService服务：

```protobuf
service HelloService {
    rpc Hello (String) returns (String);
}
```

但是重新生成的Go代码并没有发生变化。这是因为**世界上的RPC实现有千万种，protoc编译器并不知道该如何为HelloService服务生成代码。**

因此我们必须指定使用哪一种rpc，在protoc-gen-go内部已经集成了一个名字为`grpc`的插件，可以针对gRPC生成代码：

```
$ protoc --go_out=plugins=grpc:. hello.proto
```

在生成的代码中多了一些类似HelloServiceServer、HelloServiceClient的新类型。这些类型是为gRPC服务的，并不符合我们的RPC要求。

不过gRPC插件为我们提供了改进的思路，下面我们将探索如何为我们的RPC生成安全的代码。

### 2. 定制代码生成插件   --- 没学懂

Protobuf的protoc编译器是通过插件机制实现对不同语言的支持。

比如**protoc命令出现`--xxx_out`格式的参数**，那么**protoc将首先查询是否有内置的xxx插件**，如果**没有内置的xxx插件那么将继续查询当前系统中是否存在protoc-gen-xxx命名的可执行程序**，最终通过查询到的插件生成代码。

对于Go语言的protoc-gen-go插件来说，里面又实现了一层静态插件系统。比如**protoc-gen-go内置了一个gRPC插件，用户可以通过`--go_out=plugins=grpc`参数来生成gRPC相关代码，否则只会针对message生成相关代码。**

参考gRPC插件的代码，可以发现generator.RegisterPlugin函数可以用来注册插件。插件是一个generator.Plugin接口：

#### 2.1 设计proto插件接口(实现自己的rpc生成插件) 

```go
// A Plugin provides functionality to add to the output during
// Go code generation, such as to produce RPC stubs.
type Plugin interface {
    // Name identifies the plugin.
    Name() string
    // Init is called once after data structures are built but before
    // code generation begins.
    Init(g *Generator)
    // Generate produces the code generated by the plugin for this file,
    // except for the imports, by calling the generator's methods P, In,
    // and Out.
    Generate(file *FileDescriptor)
    // GenerateImports produces the import declarations for this file.
    // It is called after Generate.
    GenerateImports(file *FileDescriptor)
}
```

其中Name方法返回插件的名字，这是Go语言的Protobuf实现的插件体系，和protoc插件的名字并无关系。

然后**Init函数是通过g参数对插件进行初始化，g参数中包含Proto文件的所有信息**。最后的**Generate和GenerateImports方法用于生成主体代码和对应的导入包代码**。

因此我们**可以设计一个netrpcPlugin插件，用于为标准库的RPC框架生成代码**：

```go
import (
    "github.com/golang/protobuf/protoc-gen-go/generator"
)

type netrpcPlugin struct{ *generator.Generator }

func (p *netrpcPlugin) Name() string                { return "netrpc" }
func (p *netrpcPlugin) Init(g *generator.Generator) { p.Generator = g }

func (p *netrpcPlugin) GenerateImports(file *generator.FileDescriptor) {
    if len(file.Service) > 0 {
        p.genImportCode(file)
    }
}

func (p *netrpcPlugin) Generate(file *generator.FileDescriptor) {
    for _, svc := range file.Service {
        p.genServiceCode(svc)
    }
}
```

首先Name方法返回插件的名字。

netrpcPlugin插件**内置了一个匿名的`*generator.Generator`成员**，然后**在Init初始化的时候用参数g进行初始化**，因此**插件是从g参数对象继承了全部的公有方法。**

其中GenerateImports方法调用自定义的genImportCode函数生成导入代码。Generate方法调用自定义的genServiceCode方法生成每个服务的代码。

目前，自定义的genImportCode和genServiceCode方法只是输出一行简单的注释，等待实现：

```go
func (p *netrpcPlugin) genImportCode(file *generator.FileDescriptor) {
    p.P("// TODO: import code")
}

func (p *netrpcPlugin) genServiceCode(svc *descriptor.ServiceDescriptorProto) {
    p.P("// TODO: service code, Name = " + svc.GetName())
}
```

#### 2.2 注册插件

要使用该插件需要先通过generator.RegisterPlugin函数注册插件，可以在init函数中完成：

```go
func init() {
    generator.RegisterPlugin(new(netrpcPlugin))
}
```

因为Go语言的包只能静态导入，我们**无法向已经安装的protoc-gen-go添加我们新编写的插件**。我们将**重新克隆使用protoc-gen-go对应的main函数**，然后：

```go
package main

import (
    "io/ioutil"
    "os"

    "github.com/golang/protobuf/proto"
    "github.com/golang/protobuf/protoc-gen-go/generator"
)

func main() {
    g := generator.New()

    data, err := ioutil.ReadAll(os.Stdin)
    if err != nil {
        g.Error(err, "reading input")
    }

    if err := proto.Unmarshal(data, g.Request); err != nil {
        g.Error(err, "parsing input proto")
    }

    if len(g.Request.FileToGenerate) == 0 {
        g.Fail("no files to generate")
    }

    g.CommandLineParameters(g.Request.GetParameter())

    // Create a wrapped version of the Descriptors and EnumDescriptors that
    // point to the file that defines them.
    g.WrapTypes()

    g.SetPackageNames()
    g.BuildTypeNameMap()

    g.GenerateAllFiles()

    // Send back the results.
    data, err = proto.Marshal(g.Response)
    if err != nil {
        g.Error(err, "failed to marshal output proto")
    }
    _, err = os.Stdout.Write(data)
    if err != nil {
        g.Error(err, "failed to write output proto")
    }
}
```

**为了避免对protoc-gen-go插件造成干扰，我们将我们的可执行程序命名为protoc-gen-go-netrpc**，表示**包含了nerpc插件**。然后用以下命令**重新编译hello.proto文件**：

```shell
$ protoc --go-netrpc_out=plugins=netrpc:. hello.proto
```

其中`--go-netrpc_out`参数告知protoc编译器加载名为protoc-gen-go-netrpc的插件，插件中的`plugins=netrpc`指示启用内部唯一的名为netrpc的netrpcPlugin插件。在新生成的hello.pb.go文件中将包含增加的注释代码。

至此，手工定制的Protobuf代码生成插件终于可以工作了。

#### 2.3 自动生成完整的RPC代码

TODO





## 三、RPC实践使用

### 1. 客户端RPC的实现原理

Go语言的RPC库最简单的使用方式是通过`Client.Call`方法进行同步阻塞调用，该方法的实现如下：

```go
func (client *Client) Call(
    serviceMethod string, args interface{},
    reply interface{},
) error {
    call := <-client.Go(serviceMethod, args, reply, make(chan *Call, 1)).Done
    return call.Error
}
```

首先通过`Client.Go`方法进行一次异步调用，返回一个表示这次调用的`Call`类型结构体。然后等待`Call`结构体的Done管道返回调用结果。

我们也可以通过`Client.Go`方法异步调用前面的HelloService服务：

```go
func doClientWork(client *rpc.Client) {
    helloCall := client.Go("HelloService.Hello", "hello", new(string), nil)

    // TODO: do some thing

    helloCall = <-helloCall.Done
    if err := helloCall.Error; err != nil {
        log.Fatal(err)
    }

    args := helloCall.Args.(string)
    reply := helloCall.Reply.(string)
    fmt.Println(args, reply)
}
```

在异步调用命令发出后，一般会执行其他的任务，因此异步调用的输入参数和返回值可以通过返回的Call变量进行获取。

执行异步调用的`Client.Go`方法实现如下：

```go
func (client *Client) Go(
    serviceMethod string, args interface{},
    reply interface{},
    done chan *Call,
) *Call {
    call := new(Call)
    call.ServiceMethod = serviceMethod
    call.Args = args
    call.Reply = reply
    call.Done = make(chan *Call, 10) // buffered.

    client.send(call)
    return call
}
```

首先是构造一个表示当前调用的call变量，然后通过`client.send`将call的完整参数发送到RPC框架。`client.send`方法调用是线程安全的，因此可以从多个Goroutine同时向同一个RPC链接发送调用指令。

当调用完成或者发生错误时，将调用`call.done`方法通知完成：

```go
func (call *Call) done() {
    select {
    case call.Done <- call:
        // ok
    default:
        // We don't want to block here. It is the caller's responsibility to make
        // sure the channel has enough buffer space. See comment in Go().
    }
}
```

从`Call.done`方法的实现可以得知`call.Done`管道会将处理后的call返回。

### 2. 基于RPC实现Watch功能

在很多系统中都提供了**Watch监视功能的接口**，**当系统满足某种条件时Watch方法返回监控的结果**。在这里我们可以尝试通过RPC框架实现一个基本的Watch功能。如前文所描述，因为`client.send`是线程安全的，我们也可以通过在不同的Goroutine中同时并发阻塞调用RPC方法。通过在一个独立的Goroutine中调用Watch函数进行监控。

为了便于演示，我们计划通过RPC构造一个简单的内存KV数据库。首先定义服务如下：

```go
type KVStoreService struct {
    m      map[string]string
    filter map[string]func(key string)
    mu     sync.Mutex
}

func NewKVStoreService() *KVStoreService {
    return &KVStoreService{
        m:      make(map[string]string),
        filter: make(map[string]func(key string)),
    }
}
```

其中`m`成员是一个map类型，用于存储KV数据。`filter`成员对应每个Watch调用时定义的过滤器函数列表。而`mu`成员为互斥锁，用于在多个Goroutine访问或修改时对其它成员提供保护。

然后就是Get和Set方法：

```go
func (p *KVStoreService) Get(key string, value *string) error {
    p.mu.Lock()
    defer p.mu.Unlock()

    if v, ok := p.m[key]; ok {
        *value = v
        return nil
    }

    return fmt.Errorf("not found")
}

func (p *KVStoreService) Set(kv [2]string, reply *struct{}) error {
    p.mu.Lock()
    defer p.mu.Unlock()

    key, value := kv[0], kv[1]

    if oldValue := p.m[key]; oldValue != value {
        for _, fn := range p.filter {
            fn(key)
        }
    }

    p.m[key] = value
    return nil
}
```

在Set方法中，输入参数是key和value组成的数组，用一个**匿名的空结构体**表示**忽略了输出参数**。**当修改某个key对应的值时会调用每一个过滤器函数**。

而**过滤器列表在Watch方法中提供**：

```go
func (p *KVStoreService) Watch(timeoutSecond int, keyChanged *string) error {
    id := fmt.Sprintf("watch-%s-%03d", time.Now(), rand.Int())
    ch := make(chan string, 10) // buffered

    p.mu.Lock()
    p.filter[id] = func(key string) { ch <- key }
    p.mu.Unlock()

    select {
    case <-time.After(time.Duration(timeoutSecond) * time.Second):
        return fmt.Errorf("timeout")
    case key := <-ch:
        *keyChanged = key
        return nil
    }

    return nil
}
```

Watch方法的输入参数是超时的秒数。**当有key变化时将key作为返回值返回**。如果**超过时间后依然没有key被修改，则返回超时的错误**。Watch的实现中，**用唯一的id表示每个Watch调用，然后根据id将自身对应的过滤器函数注册到`p.filter`列表**。

KVStoreService服务的注册和启动过程我们不再赘述。下面我们看看如何从客户端使用Watch方法：

```go
func doClientWork(client *rpc.Client) {
    go func() {
        var keyChanged string
        err := client.Call("KVStoreService.Watch", 30, &keyChanged)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println("watch:", keyChanged)
    } ()

    err := client.Call(
        "KVStoreService.Set", [2]string{"abc", "abc-value"},
        new(struct{}),
    )
    if err != nil {
        log.Fatal(err)
    }

    time.Sleep(time.Second*3)
}
```

首先启动一个独立的Goroutine监控key的变化。同步的watch调用会阻塞，直到有key发生变化或者超时。然后在通过Set方法修改KV值时，服务器会将变化的key通过Watch方法返回。这样我们就可以实现对某些状态的监控。

### 3. 反向RPC

通常的RPC是基于C/S结构，RPC的服务端对应网络的服务器，RPC的客户端也对应网络客户端。但是对于一些特殊场景，比如在公司内网提供一个RPC服务，但是在外网无法链接到内网的服务器。这种时候我们可以参考类似反向代理的技术，**首先从内网主动链接到外网的TCP服务器，然后基于TCP链接向外网提供RPC服务**。

以下是**启动反向RPC服务的代码**：

```go
func main() {
    rpc.Register(new(HelloService))

    for {
        conn, _ := net.Dial("tcp", "localhost:1234")
        if conn == nil {
            time.Sleep(time.Second)
            continue
        }

        rpc.ServeConn(conn)
        conn.Close()
    }
}
```

反向RPC的**内网服务将不再主动提供TCP监听服务，而是首先主动链接到对方的TCP服务器**。然后**基于每个建立的TCP链接向对方提供RPC服务**。

而**RPC客户端则需要在一个公网地址提供一个TCP服务，用于接受RPC服务器的链接请求**：

```go
func main() {
    listener, err := net.Listen("tcp", ":1234")
    if err != nil {
        log.Fatal("ListenTCP error:", err)
    }

    clientChan := make(chan *rpc.Client)

    go func() {
        for {
            conn, err := listener.Accept()
            if err != nil {
                log.Fatal("Accept error:", err)
            }

            clientChan <- rpc.NewClient(conn)
        }
    }()

    doClientWork(clientChan)
}
```

当每个**链接建立后，基于网络链接构造RPC客户端对象并发送到clientChan管道**。

**客户端执行RPC调用的操作在doClientWork函数完成**：

```go
func doClientWork(clientChan <-chan *rpc.Client) {
    client := <-clientChan
    defer client.Close()

    var reply string
    err = client.Call("HelloService.Hello", "hello", &reply)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(reply)
}
```

首先从管道去取一个RPC客户端对象，并且通过defer语句指定在函数退出前关闭客户端。然后是执行正常的RPC调用。

### 4. 上下文信息

基于上下文我们可以**针对不同客户端提供定制化的RPC服务**。我们可以**通过为每个链接提供独立的RPC服务来实现对上下文特性的支持**。

首**先改造HelloService，里面增加了对应链接的conn**：

```go
type HelloService struct {
    conn net.Conn
}
```

然后为每个链接启动**独立的RPC服务**：

```go
func main() {
    listener, err := net.Listen("tcp", ":1234")
    if err != nil {
        log.Fatal("ListenTCP error:", err)
    }

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Fatal("Accept error:", err)
        }

        go func() {
            defer conn.Close()

            p := rpc.NewServer()
            p.Register(&HelloService{conn: conn})
            p.ServeConn(conn)
        } ()
    }
}
```

Hello方法中就可以根据conn成员识别不同链接的RPC调用：

```go
func (p *HelloService) Hello(request string, reply *string) error {
    *reply = "hello:" + request + ", from" + p.conn.RemoteAddr().String()
    return nil
}
```

基于上下文信息，我们可以**方便地为RPC服务增加简单的登陆状态的验证**：

```go
type HelloService struct {
    conn    net.Conn
    isLogin bool    //用户的登录状态
}

func (p *HelloService) Login(request string, reply *string) error {
    if request != "user:password" {
        return fmt.Errorf("auth failed")
    }
    log.Println("login ok")
    p.isLogin = true
    return nil
}

func (p *HelloService) Hello(request string, reply *string) error {
    if !p.isLogin {
        return fmt.Errorf("please login")
    }
    *reply = "hello:" + request + ", from" + p.conn.RemoteAddr().String()
    return nil
}
```

## 四、gRPC入门

### 1. gRPC入门

创建hello.proto文件，定义HelloService接口：

```protobuf
syntax = "proto3";

package main;

message String {
    string value = 1;
}

service HelloService {
    rpc Hello (String) returns (String);
}
```

使用protoc-gen-go内置的gRPC插件生成gRPC代码：

```
$ protoc --go_out=plugins=grpc:. hello.proto
```

gRPC插件会为**服务端**和**客户端**生成不同的接口：

```go
type HelloServiceServer interface {
    Hello(context.Context, *String) (*String, error)
}

type HelloServiceClient interface {
    Hello(context.Context, *String, ...grpc.CallOption) (*String, error)
}
```

gRPC通过context.Context参数，为每个方法调用提供了上下文支持。**客户端**在调用方法的时候，可以通过**可选的grpc.CallOption类型的参数**提供额外的上下文信息。

基于**服务端**的HelloServiceServer**接口**可以**自行实现**HelloService服务：

```go
type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(
    ctx context.Context, args *String,   //args是String类型结构体指针
) (*String, error) {
    reply := &String{Value: "hello:" + args.GetValue()}
    return reply, nil
}
```

gRPC服务的启动流程和标准库的RPC服务启动流程类似：

```go
func main() {
    grpcServer := grpc.NewServer()
    RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

    lis, err := net.Listen("tcp", ":1234")
    if err != nil {
        log.Fatal(err)
    }
    grpcServer.Serve(lis)
}
```

首先是**通过`grpc.NewServer()`构造一个gRPC服务对象，然后通过gRPC插件生成的RegisterHelloServiceServer函数注册我们实现的HelloServiceImpl服务。然后通过`grpcServer.Serve(lis)`在一个监听端口上提供gRPC服务。**

然后就可以**通过客户端链接gRPC服务**了：

```go
func main() {
    conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    client := NewHelloServiceClient(conn)
    reply, err := client.Hello(context.Background(), &String{Value: "hello"})
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(reply.GetValue())
}
```

其中grpc.Dial负责和gRPC服务建立链接，然后NewHelloServiceClient函数基于已经建立的链接构造HelloServiceClient对象。**返回的client其实是一个HelloServiceClient接口对象，通过接口定义的方法就可以调用服务端对应的gRPC服务提供的方法。**

> gRPC和标准库的RPC框架有一个区别，gRPC生成的接口并不支持异步调用。不过我们可以在多个Goroutine之间安全地共享gRPC底层的HTTP/2链接，因此可以通过在另一个Goroutine阻塞调用的方式模拟异步调用。



### 2. gRPC流服务

RPC是远程函数调用，因此**每次调用的函数参数和返回值不能太大，否则将严重影响每次调用的响应时间**。因此**传统的RPC方法调用对于上传和下载较大数据量场景并不适合**。同时传统RPC模式**也不适用于对时间不确定的订阅和发布模式**。为此，**gRPC框架针对服务器端和客户端分别提供了流特性**。

服务端或客户端的单向流是双向流的特例，我们**在HelloService增加一个支持双向流的Channel方法**：

```protobuf
service HelloService {
    rpc Hello (String) returns (String);

    rpc Channel (stream String) returns (stream String);
}
```

**关键字stream指定启用流特性**，参数部分是接收客户端参数的流，返回值是返回给客户端的流。

重新生成代码可以看到接口中新增加的Channel方法的定义：

```go
type HelloServiceServer interface {
    Hello(context.Context, *String) (*String, error)
    Channel(HelloService_ChannelServer) error
}
type HelloServiceClient interface {
    Hello(ctx context.Context, in *String, opts ...grpc.CallOption) (
        *String, error,
    )
    Channel(ctx context.Context, opts ...grpc.CallOption) (
        HelloService_ChannelClient, error,
    )
}
```

在服务端的Channel方法参数是一个**新的HelloService_ChannelServer类型的参数**，可以**用于和客户端双向通信**。客户端的Channel方法返回一个**HelloService_ChannelClient类型的返回值**，可以**用于和服务端进行双向通信**。

**HelloService_ChannelServer和HelloService_ChannelClient均为接口类型(不需要手动实现这两个接口)：**

```go
type HelloService_ChannelServer interface {
    Send(*String) error
    Recv() (*String, error)
    grpc.ServerStream
}

type HelloService_ChannelClient interface {
    Send(*String) error
    Recv() (*String, error)
    grpc.ClientStream
}
```

可以发现**服务端和客户端的流辅助接口均定义了Send和Recv方法用于流数据的双向通信**。

现在我们可以实现**服务端的流服务**：

```go
func (p *HelloServiceImpl) Channel(stream HelloService_ChannelServer) error {
    for {
        args, err := stream.Recv()   //接收来自于客户端的流数据
        if err != nil {
            if err == io.EOF {
                return nil
            }
            return err
        }

        reply := &String{Value: "hello:" + args.GetValue()}

        err = stream.Send(reply)   //向客户端进行回复
        if err != nil {
            return err
        }
    }
}
```

服务端在循环中接收客户端发来的数据，如果遇到io.EOF表示客户端流被关闭，如果函数退出表示服务端流关闭。生成返回的数据通过流发送给客户端，**双向流数据的发送和接收都是完全独立的行为。**

**需要注意的是，发送和接收的操作并不需要一一对应，用户可以根据真实场景进行组织代码。**

**客户端**需要先**调用Channel方法获取返回的流对象**：

```go
stream, err := client.Channel(context.Background())
if err != nil {
    log.Fatal(err)
}
```

在客户端我们将发送和接收操作放到两个独立的Goroutine。首先是**向服务端发送数据**：

```go
go func() {
    for {
        if err := stream.Send(&String{Value: "hi"}); err != nil {
            log.Fatal(err)
        }
        time.Sleep(time.Second)
    }
}()
```

然后在循环中**接收服务端返回的数据**：

```go
for {
    reply, err := stream.Recv()
    if err != nil {
        if err == io.EOF {
            break
        }
        log.Fatal(err)
    }
    fmt.Println(reply.GetValue())
}
```

### 3. 发布和订阅模式

在前一节中，我们**基于Go内置的RPC库实现了一个简化版的Watch方法**。**基于Watch的思路虽然也可以构造发布和订阅系统，但是因为RPC缺乏流机制导致每次只能返回一个结果**。在发布和订阅模式中，由调用者主动发起的发布行为类似一个普通函数调用，而被动的订阅者则类似gRPC客户端单向流中的接收者。现在我们**可以尝试基于gRPC的流特性构造一个发布和订阅系统**。

发布订阅是一个常见的设计模式，开源社区中已经存在很多该模式的实现。其中docker项目中提供了一个pubsub的极简实现，下面是基于pubsub包实现的本地发布订阅代码：

```go
import (
    "github.com/moby/moby/pkg/pubsub"
)

func main() {
    p := pubsub.NewPublisher(100*time.Millisecond, 10)

    golang := p.SubscribeTopic(func(v interface{}) bool {  //返回的是一个管道
        if key, ok := v.(string); ok {
            if strings.HasPrefix(key, "golang:") {
                return true
            }
        }
        return false
    })
    docker := p.SubscribeTopic(func(v interface{}) bool {  //返回的是一个管道
        if key, ok := v.(string); ok {
            if strings.HasPrefix(key, "docker:") {
                return true
            }
        }
        return false
    })

    go p.Publish("hi")   //发布一个事件
    go p.Publish("golang: https://golang.org")
    go p.Publish("docker: https://www.docker.com/")
    time.Sleep(1)

    go func() {
        fmt.Println("golang topic:", <-golang)  //从管道中等待读取订阅的事件
    }()
    go func() {
        fmt.Println("docker topic:", <-docker)  //从管道中等待读取订阅的事件
    }()

    <-make(chan bool)
}
```

其中`pubsub.NewPublisher`构造一个发布对象，`p.SubscribeTopic()`可以通过自定义函数筛选感兴趣的主题进行订阅。

现在尝试基于gRPC和pubsub包，提供一个跨网络的发布和订阅系统。首先通过Protobuf定义一个发布订阅服务接口：

```protobuf
service PubsubService {
    rpc Publish (String) returns (String);
    rpc Subscribe (String) returns (stream String);
}
```

其中Publish是普通的RPC方法，Subscribe则是一个单向的流服务。然后gRPC插件会为服务端和客户端生成对应的接口：

```go
type PubsubServiceServer interface {
    Publish(context.Context, *String) (*String, error)
    Subscribe(*String, PubsubService_SubscribeServer) error
}
type PubsubServiceClient interface {
    Publish(context.Context, *String, ...grpc.CallOption) (*String, error)
    Subscribe(context.Context, *String, ...grpc.CallOption) (
        PubsubService_SubscribeClient, error,
    )
}

type PubsubService_SubscribeServer interface {
    Send(*String) error
    grpc.ServerStream
}
```

**因为Subscribe是服务端的单向流，因此生成的HelloService_SubscribeServer接口中只有Send方法。**

然后就可以实现发布和订阅服务了：

```go
type PubsubService struct {
    pub *pubsub.Publisher
}

func NewPubsubService() *PubsubService {
    return &PubsubService{
        pub: pubsub.NewPublisher(100*time.Millisecond, 10),
    }
}
```

然后是实现发布方法和订阅方法：

```go
// rpc服务端需要实现PubsubServiceServer服务接口(客户端不需要实现PubsubServiceClient接口，因为它是远程调用服务端的函数)
func (p *PubsubService) Publish(
    ctx context.Context, arg *String,
) (*String, error) {
    p.pub.Publish(arg.GetValue())
    return &String{}, nil
}

func (p *PubsubService) Subscribe(
    arg *String, stream PubsubService_SubscribeServer,
) error {
    ch := p.pub.SubscribeTopic(func(v interface{}) bool {
        if key, ok := v.(string); ok {
            if strings.HasPrefix(key,arg.GetValue()) {
                return true
            }
        }
        return false
    })

    for v := range ch {
        if err := stream.Send(&String{Value: v.(string)}); err != nil {
            return err
        }
    }

    return nil
}
```

这样就可以从客户端向服务器发布信息了：

```go
// 该客户端负责远程调用服务端，让服务端发布事件
func main() {
    conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    client := NewPubsubServiceClient(conn)

    _, err = client.Publish(
        context.Background(), &String{Value: "golang: hello Go"},
    )
    if err != nil {
        log.Fatal(err)
    }
    _, err = client.Publish(
        context.Background(), &String{Value: "docker: hello Docker"},
    )
    if err != nil {
        log.Fatal(err)
    }
}
```

然后就可以在另一个客户端进行订阅信息了：

```go
// 该客户端负责远程调用服务端，订阅指定的事件
func main() {
    conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    client := NewPubsubServiceClient(conn)
    stream, err := client.Subscribe(
        context.Background(), &String{Value: "golang:"},
    )
    if err != nil {
        log.Fatal(err)
    }

    for {
        reply, err := stream.Recv()
        if err != nil {
            if err == io.EOF {
                break
            }
            log.Fatal(err)
        }

        fmt.Println(reply.GetValue())
    }
}
```

到此我们就基于gRPC简单实现了一个跨网络的发布和订阅服务。

## 五、gRPC 进阶(待学习......)