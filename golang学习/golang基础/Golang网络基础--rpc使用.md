## 一、服务器端

```go
package main

import (
  "fmt"
  "net"
  "net/http"
  "net/rpc"
)

type Server struct {
}
type Req struct {
  Num1 int
  Num2 int
}
type Res struct {
  Num int
}

func (s Server) Add(req Req, res *Res) error {
  res.Num = req.Num1 + req.Num2
  return nil
}

func main() {
  // 注册rpc服务
  rpc.Register(new(Server))
  rpc.HandleHTTP()
  listen, err := net.Listen("tcp", ":8080")
  if err != nil {
    fmt.Println(err)
    return
  }
  http.Serve(listen, nil)
}
```

## 二、客户端

```go
package main

import (
  "fmt"
  "net/rpc"
)

type Req struct {
  Num1 int
  Num2 int
}
type Res struct {
  Num int
}

func main() {
  req := Req{1, 2}
  client, err := rpc.DialHTTP("tcp", ":8080")   //获取rpc客户端
  if err != nil {
    fmt.Println(err)
    return
  }
  var res Res
  client.Call("Server.Add", req, &res)   //远程调用rpc服务端服务对象Server的Add方法
  fmt.Println(res)
}
```

