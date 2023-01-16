## 一、`websocket`包

websocket是[socket](https://so.csdn.net/so/search?q=socket&spm=1001.2101.3001.7020)连接和http协议的结合体，可以实现网页和服务端的长连接

```go
go get github.com/gorilla/websocket
```

## 二、服务端实现

### 2.1 服务器：一端口一客户端

```go
package main

import (
  "fmt"
  "github.com/gorilla/websocket"
  "net/http"
)

var UP = websocket.Upgrader{  //此对象用于设置从http升级到wesocket的配置信息(读缓冲、写缓冲大小)
  ReadBufferSize:  1024,
  WriteBufferSize: 1024,
}

func handler(res http.ResponseWriter, req *http.Request) {
  // 服务升级
  conn, err := UP.Upgrade(res, req, nil)   //将http升级成websocket
  if err != nil {
    fmt.Println(err)
    return
  }
  for {
    // 消息类型，消息，错误
    t, p, err := conn.ReadMessage()  //循环读取websocket conn的接收数据
    if err != nil {
      break
    }
    // 将数据处理后回复给客户端：回复消息类型码 + 回复消息实体
    conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("你说的是：%s吗？", string(p))))
    fmt.Println(t, string(p))
  }
  defer conn.Close()
  fmt.Println("服务关闭")
}

// 起始，与客户端通过http通信，在handler函数中将http协议升级为websocket协议
func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}
```

### 2.2 服务器：广播式

服务器将接收到的数据广播给所有客户端

```go
package main

import (
  "fmt"
  "github.com/gorilla/websocket"
  "net/http"
)

var UP = websocket.Upgrader{
  ReadBufferSize:  1024,
  WriteBufferSize: 1024,
}

var connLis []*websocket.Conn   //新增数据类型，用于保存所有客户端的websocket conn

func handler(res http.ResponseWriter, req *http.Request) {
  // 服务升级
  conn, err := UP.Upgrade(res, req, nil)
  if err != nil {
    fmt.Println(err)
    return
  }
  connLis = append(connLis, conn)   //每与一个客户端建立websocket连接，就将其存入connLis切片
  for {
    // 消息类型，消息，错误
    t, p, err := conn.ReadMessage()
    if err != nil {
      break
    }
    for index := range connLis {   //将接收的客户端消息广播给所有websocket conn
      connLis[index].WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("你说的是：%s吗？", string(p))))
    }
    fmt.Println(t, string(p))
  }
  defer conn.Close()
  fmt.Println("服务关闭")
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}
```



## 三、客户端实现

```go
package main

import (
  "bufio"
  "fmt"
  "github.com/gorilla/websocket"
  "os"
)

func main() {
  dl := websocket.Dialer{}   // websocket连接器
  conn, _, err := dl.Dial("ws://127.0.0.1:8080", nil)  //用连接器与服务器建立连接，注意协议名为：ws
  if err != nil {
    fmt.Println(err)
    return
  }
  go send(conn)   //发送子协程
  for {   //主协程负责接收服务器消息
    t, p, err := conn.ReadMessage()
    if err != nil {
      break
    }
    fmt.Println(t, string(p))
  }
}

func send(conn *websocket.Conn) {
  for {
    reader := bufio.NewReader(os.Stdin)   //以标准输入设备创建reader
    l, _, _ := reader.ReadLine()   //每次读取用户的一行输入
    conn.WriteMessage(websocket.TextMessage, l)
  }
}

```

## 四、使用PostMan作为测试客户端

![image-20230116192221294](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\image-20230116192221294.png)

![image-20230116192227358](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\image-20230116192227358.png)