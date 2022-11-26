> All transfer of value on the go channels happens with the copy of value.

就是说 **channel 的发送和接收操作本质上都是 “值的拷贝”**，无论是从 sender goroutine  的栈到 chan buf，还是从 chan buf 到 receiver goroutine，或者是直接从 sender goroutine 到 receiver goroutine。

举一个例子：

```go
type user struct {
    name string
    age int8
}

var u = user{name: "Ankur", age: 25}
var g = &u

func modifyUser(pu *user) {
    fmt.Println("modifyUser Received Vaule", pu)
    pu.name = "Anand"
}

func printUser(u <-chan *user) {
    time.Sleep(2 * time.Second)
    fmt.Println("printUser goRoutine called", <-u)
}

func main() {
    c := make(chan *user, 5)
    c <- g            //把一个地址传给管道，就是把一个地址拷贝到管道的底层数组
    fmt.Println(g)
    // modify g
    g = &user{name: "Ankur Anand", age: 100}
    go printUser(c)    //从管道(管道的底层数组)中取出上述的地址
    go modifyUser(g)
    time.Sleep(5 * time.Second)
    fmt.Println(g)
}
```

运行结果：

```go
&{Ankur 25}
modifyUser Received Vaule &{Ankur Anand 100}
printUser goRoutine called &{Ankur 25}
&{Anand 100}
```

这里就是一个很好的 `share memory by communicating` 的例子。

<img src="https://www.topgoer.cn/uploads/goquestions/images/m_d953dbc4cca421a8663ffe814dc73f70_r.png" alt="null" style="zoom: 80%;" />

一开始构造一个结构体 u，地址是 0x56420，图中地址上方就是它的内容。接着把 `&u` 赋值给指针 `g`，g 的地址是 0x565bb0，它的内容就是一个地址，指向 u。

main 程序里，先把 g 发送到 c，**根据 `copy value` 的本质，进入到 chan buf 里的就是 `0x56420`，它是指针 g 的值（不是它指向的内容）**，所以打印从 channel 接收到的元素时，它就是 `&{Ankur 25}`。因此，**这里并不是将指针 g “发送” 到了 channel 里，只是拷贝它的值而已。**