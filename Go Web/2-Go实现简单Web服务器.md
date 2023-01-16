## 一、Go搭建一个Web服务器

Go语言里面提供了一个完善的net/http包，通过http包可以很方便的搭建起来一个可以运行的Web服务。同时使用这个包能很简单地对Web的路由，静态文件，模版，cookie等数据进行设置和操作。

### 1. http包建立Web服务器

```go
package main

import (
    "fmt"
    "net/http"
    "strings"
    "log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()  //ParseForm解析URL中的查询字符串，并将解析结果更新到r.Form字段。
    fmt.Println(r.Form)  //是一个map，以k-v形式保存request中的查询数据
    fmt.Println("path", r.URL.Path)  //资源路径
    fmt.Println("scheme", r.URL.Scheme)   //使用的协议
    fmt.Println(r.Form["url_long"])   //输出key为"url_long"对应的value
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func main() {
    http.HandleFunc("/", sayhelloName) //设置访问的路由sayhelloName()
    err := http.ListenAndServe(":9090", nil) //设置监听的端口
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
```

上面这个代码，我们build之后，然后执行web.exe,这个时候其实已经在9090端口监听http链接请求了。

在浏览器输入`http://localhost:9090`

可以看到浏览器页面输出了`Hello astaxie!`

可以换一个地址试试：`http://localhost:9090/?url_long=111&url_long=222`

看看浏览器输出的是什么，服务器输出的是什么？

在服务器端输出的信息如下：

![image-20221225210858187](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\image-20221225210858187.png)

```
第一次输入http://localhost:9090，返回结果为：
map[]
path /
scheme
[]
因为url后面没有任何资源数据，因此r.Form为空

第二次输入http://localhost:9090/?url_long=111&url_long=222，返回结果为：
map[url long:[111 222]]
path /
scheme
[111 222]
key: url long
val: 111222
```

我们看到上面的代码，要编写一个Web服务器很简单，只要调用http包的两个函数就可以了。**不需要使用nginx、apache服务器，因为他直接就监听tcp端口了，做了nginx做的事情，然后sayhelloName这个其实就是我们写的逻辑函数了。**

### 2. Go 如何让Web工作

#### 2.1 web工作方式的几个概念

以下均是服务器端的几个概念

Request：用户请求的信息，用来解析用户的请求信息，包括post、get、cookie、url等信息

Response：服务器需要反馈给客户端的信息

Conn：用户的每次请求链接

Handler：处理请求和生成返回信息的处理逻辑

#### 2.2 分析http包运行机制

下图是Go实现Web服务的工作模式的流程图

![image-20221225211618926](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\image-20221225211618926.png)

1. Http服务器创建Listen Socket, 监听指定的端口, 等待客户端请求到来。
2. Listen Socket接受客户端的请求, 得到Client Socket, 接下来**通过Client Socket与客户端通信**。
3. 处理客户端的请求, 首先从Client Socket读取HTTP请求的协议头,  **如果是POST方法, 还可能要读取客户端提交的数据, 然后交给相应的handler处理请求**, handler处理完毕准备好客户端需要的数据,  通过Client Socket写给客户端。

这整个的过程里面我们**只要了解清楚下面三个问题，也就知道Go是如何让Web运行起来了**

- **如何监听端口？**
- **如何接收客户端请求？**
- **如何分配handler？**

##### 2.2.1 如何监听端口？

前面小节的代码里面我们可以看到，**Go是通过一个函数`ListenAndServe`来处理这些事情的**，这个底层其实这样处理的：**初始化一个server对象，然后调用了`net.Listen("tcp", addr)`，也就是底层用TCP协议搭建了一个服务，然后监控我们设置的端口。**

下面代码来自Go的http包的源码，通过下面的代码我们可以看到整个的http处理过程：

```go
func (srv *Server) Serve(l net.Listener) error {
    defer l.Close()
    var tempDelay time.Duration // how long to sleep on accept failure
    for {
        rw, e := l.Accept()    //接受客户端的TCP连接请求
        if e != nil {   //conn连接建立失败，进行错误处理(等待以重新建立连接)
            if ne, ok := e.(net.Error); ok && ne.Temporary() {
                if tempDelay == 0 {
                    tempDelay = 5 * time.Millisecond
                } else {
                    tempDelay *= 2
                }
                if max := 1 * time.Second; tempDelay > max {
                    tempDelay = max
                }
                log.Printf("http: Accept error: %v; retrying in %v", e, tempDelay)
                time.Sleep(tempDelay)
                continue
            }
            return e
        }
        tempDelay = 0
        c, err := srv.newConn(rw)   //为conn socket创建http服务对象
        if err != nil {
            continue
        }
        go c.serve()   //协程运行http服务对象
    }
}
```

监控之后如何接收客户端的请求呢？**上面代码执行监控端口之后(即ListenAndServe()函数)，调用了`srv.Serve(net.Listener)`函数**，这个函数就是处理接收客户端的请求信息。**这个函数里面起了一个`for{}`，首先通过Listener接收请求，其次创建一个Conn，最后单独开了一个goroutine，把这个请求的数据当做参数扔给这个conn去服务：`go c.serve()`**。这个就是高并发体现了，**用户的每一次请求都是在一个新的goroutine去服务，相互不影响。**



##### 2.2.2 如何接受客户端请求并分配handler？

那么如何具体分配到相应的函数来处理请求呢？

**conn首先会解析request : `c.readRequest()`,然后获取相应的handler : `handler := c.server.Handler`**，也就是我们刚才**在调用函数`ListenAndServe`时候的第二个参数**。

我们**前面例子传递的是nil，也就是为空，那么默认获取`handler = DefaultServeMux`**,那么这个变量用来做什么的呢？对，这个变量就**是一个路由器，它用来匹配url跳转到其相应的handle函数**。

那么这个我们有设置过吗?有，我们调用的**代码里面第一句调用了`http.HandleFunc("/", sayhelloName)`**。这个**作用就是注册了请求`/`的路由规则，当请求url为”/“，路由就会转到函数sayhelloName**，**DefaultServeMux会调用ServeHTTP方法，这个方法内部其实就是调用sayhelloName本身，最后通过写入response的信息反馈到客户端。**

详细的整个流程如下图所示：

![image-20221225212808046](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\image-20221225212808046.png)

### 3.Go的http包详解

Go的http有两个核心功能：Conn、ServeMux

#### 3.1 Conn的goroutine

与我们一般编写的http服务器不同, **Go为了实现高并发和高性能, 使用了goroutines来处理Conn的读写事件, 这样每个请求都能保持独立，相互不会阻塞，可以高效的响应网络事件**。这是Go高效的保证。

Go在等待客户端请求里面是这样写的：

```go
c, err := srv.newConn(rw)
if err != nil {
    continue
}
go c.serve()
```

客户端的每次请求都会创建一个Conn，这个Conn里面保存了该次请求的信息，然后再传递到对应的handler，该handler中便可以读取到相应的header信息，这样保证了每个请求的独立性。

#### 3.2 ServeMux的自定义

我们前面小节讲述conn.server的时候，其实**内部是调用了http包默认的路由器（ListenAndServe()第二个参数设置为nil，就会使用默认的路由器DefaultServeMux，他是ServeMux类型的对象），通过路由器把本次请求的信息传递到了后端的处理函数**。那么这个路由器是怎么实现的呢？

它的结构如下：

```go
type ServeMux struct {
    mu sync.RWMutex   //锁，由于请求涉及到并发处理，因此这里需要一个锁机制
    m  map[string]muxEntry  // 路由规则，一个string对应一个mux实体，这里的string就是注册的路由表达式
    hosts bool // 是否在任意的规则中带有host信息
}
```

下面看一下muxEntry

```go
type muxEntry struct {
    explicit bool   // 是否精确匹配
    h        Handler // 这个路由表达式对应哪个handler
    pattern  string  //匹配字符串
}
```

接着看一下Handler的定义

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)  // 路由实现器
}
```

Handler是一个接口，但是前一小节中的**`sayhelloName`函数并没有实现ServeHTTP这个接口**，为什么能添加呢？

原来在**http包里面还定义了一个类型`HandlerFunc`**,我们定义的函数`sayhelloName`就是这个HandlerFunc调用之后的结果，**这个类型默认就实现了ServeHTTP这个接口**，即我们调用了HandlerFunc(sayhelloName),强制类型转换sayhelloName成为HandlerFunc类型，这样sayhelloName就拥有了ServeHTTP方法。

```go
type HandlerFunc func(ResponseWriter, *Request)   //是一种参数为(ResponseWriter, *Request)的函数变量，sayhelloName明显是这种类型的变量

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {  //调用ServerHTTP其实就是调用函数本身
    f(w, r)
}
```

路由器里面**存储好了相应的路由规则之后**，那么**具体的请求又是怎么分发的**呢？请看下面的代码，**默认的路由器实现了`ServeHTTP`**：

```go
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
    if r.RequestURI == "*" {
        w.Header().Set("Connection", "close")
        w.WriteHeader(StatusBadRequest)
        return
    }
    h, _ := mux.Handler(r)   //返回对于request的Handler处理函数h
    h.ServeHTTP(w, r)   //h调用自身处理此次客户端的http请求
}
```

如上所示路由器接收到请求之后，如果是`*`那么关闭链接，不然调用`mux.Handler(r)`返回对应设置路由的Handler处理函数，然后执行`h.ServeHTTP(w, r)`

那么**mux.Handler(r)怎么处理的**呢？

```go
func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string) {
    if r.Method != "CONNECT" {
        if p := cleanPath(r.URL.Path); p != r.URL.Path {
            _, pattern = mux.handler(r.Host, p)
            return RedirectHandler(p, StatusMovedPermanently), pattern
        }
    }    
    return mux.handler(r.Host, r.URL.Path)
}

func (mux *ServeMux) handler(host, path string) (h Handler, pattern string) {
    mux.mu.RLock()
    defer mux.mu.RUnlock()

    // Host-specific pattern takes precedence over generic ones
    if mux.hosts {
        h, pattern = mux.match(host + path)
    }
    if h == nil {
        h, pattern = mux.match(path)
    }
    if h == nil {
        h, pattern = NotFoundHandler(), ""
    }
    return
}
```

原来他是**根据用户请求的URL和路由器里面存储的map去匹配的**，当**匹配到之后返回存储的handler，调用这个handler的ServeHTTP接口就可以执行到相应的函数了**。

### 4.http服务器代码的执行流程

通过对http包的分析之后，现在让我们来梳理一下整个的代码执行过程。

- **首先调用Http.HandleFunc**

 		往DefaultServeMux的map[string]muxEntry中增加对应的handler和路由规则，注册了一个handler函数sayhelloName，该handler函数对应客户端访问 "/" 时被调用

- **其次调用http.ListenAndServe(“:9090”, nil)**

  按顺序做了几件事情：

1. 实例化Server

2. 调用Server的ListenAndServe()

3. 调用net.Listen(“tcp”, addr)监听端口

4. 启动一个for循环，在循环体中Accept请求

5. 对每个请求实例化一个Conn，并且开启一个goroutine为这个请求进行服务go c.serve()

6. 读取每个请求的内容w, err := c.readRequest()

7. 判断handler是否为空，如果没有手动设置handler（这个例子就没有设置handler），handler就设置为DefaultServeMux

8. 调用handler的ServeHttp。在这个例子中，下面就进入到DefaultServeMux.ServeHttp

9. 根据request**(也就是根据http.HandleFunc的第一个参数)**选择handler，并且进入到这个handler的ServeHTTP

```
  mux.handler(r).ServeHTTP(w, r)
```

  	如何选择handler：

  A 判断是否有路由能满足这个request（循环遍历ServeMux的muxEntry）

  B 如果有路由满足，调用这个路由handler的ServeHTTP

  C 如果没有路由满足，调用NotFoundHandler的ServeHTTP

### 5. 自行设计http服务器的路由服务

如果 http.ListenAndServe()的第二个参数为nil，会默认使用DefaultServeMux路由，可以通过手动设置自己的路由服务

如下代码所示，我们自己实现了一个简易的路由器

```go
package main

import (
    "fmt"
    "net/http"
)

type MyMux struct {   //这个自行设计的路由类必须实现ServeHTTP方法，即实现Handler接口
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/" {
        sayhelloName(w, r)   //通过调用自行设计的handler函数sayhelloName，完成http处理
        return
    }
    http.NotFound(w, r)
    return
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello myroute!")
}

func main() {
    mux := &MyMux{}
    http.ListenAndServe(":9090", mux)   //第二个参数不再是nil，而是自行设计的路由对象
}
```

