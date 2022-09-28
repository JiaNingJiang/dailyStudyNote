### 一、简述

#### Goroutine 介绍

goroutine 是一种非常轻量级的实现，可在单个进程里执行成千上万的并发任务，它是Go语言并发设计的核心。

**说到底 goroutine 其实就是线程，但是它比线程更小，十几个 goroutine 可能体现在底层就是五六个线程，而且Go语言内部也实现了 goroutine 之间的内存共享。**

 使用 go 关键字就可以创建 goroutine，将 go 声明放到一个需调用的函数之前，在相同地址空间调用运行这个函数，这样该函数执行时便会作为一个独立的并发线程，这种线程在Go语言中则被称为 goroutine。

goroutine 虽然类似于线程概念，但是从调度性能上没有线程细致，而细致程度取决于 Go 程序的 goroutine 调度器的实现和运行环境。

#### channel介绍

channel 是Go语言在语言级别提供的 goroutine 间的通信方式。我们可以使用 channel 在两个或多个 goroutine 之间传递消息。

 channel 是进程内的通信方式，因此通过 channel  传递对象的过程和调用函数时的参数传递行为比较一致，比如也可以传递指针等。如果需要跨进程通信，我们建议用分布式系统的方法来解决，比如使用  Socket 或者 HTTP 等通信协议。Go语言对于网络方面也有非常完善的支持。

 channel 是类型相关的，也就是说，一个 channel 只能传递一种类型的值，这个类型需要在声明 channel 时指定。如果对 Unix 管道有所了解的话，就不难理解 channel，可以将其认为是一种类型安全的管道。

 定义一个 channel 时，也需要定义发送到 channel 的值的类型，注意，必须使用 make 创建 channel，代码如下所示：

```go
    ci := make(chan int)
    cs := make(chan string)
    cf := make(chan interface{})
```

回到在 Windows 和 Linux 出现之前的古老年代，在开发程序时并没有并发的概念，因为命令式程序设计语言是以串行为基础的，程序会顺序执行每一条指令，整个程序只有一个执行上下文，即一个调用栈，一个堆。

 并发则意味着程序在运行时有多个执行上下文，对应着多个调用栈。我们知道每一个进程在运行时，都有自己的调用栈和堆，有一个完整的上下文，而操作系统在调度进程的时候，会保存被调度进程的上下文环境，等该进程获得时间片后，再恢复该进程的上下文到系统中。

### 二、竞争状态简述

有并发，就有资源竞争，如果两个或者多个 goroutine 在没有相互同步的情况下，访问某个共享的资源，比如同时对该资源进行读写时，就会处于相互竞争的状态，这就是并发中的资源竞争。

下面的代码中就会出现竞争状态：

```go
    package main
    import (
        "fmt"
        "runtime"
        "sync"
    )
    var (
        count int32
        wg    sync.WaitGroup
    )
    func main() {
        wg.Add(2)
        go incCount()
        go incCount()
        wg.Wait()
        fmt.Println(count)
    }
    func incCount() {
        defer wg.Done()
        for i := 0; i < 2; i++ {
            value := count
            runtime.Gosched()
            value++
            count = value
        }
    }
```

这是一个资源竞争的例子，大家可以将程序多运行几次，会发现结果可能是 2，也可以是 3，还可能是 4。这是因为 count 变量没有任何同步保护，所以两个 goroutine 都会对其进行读写，会导致对已经计算好的结果被覆盖，以至于产生错误结果。

 **代码中的 runtime.Gosched() 是让当前 goroutine 暂停的意思，退回执行队列，让其他等待的 goroutine 运行**，目的是为了使资源竞争的结果更明显。

所以我们对于同一个资源的读写必须是**原子化**的，也就是说，**同一时间只能允许有一个 goroutine 对共享资源进行读写操作**。

 共享资源竞争的问题，非常复杂，并且难以察觉，好在 Go 为我们提供了一个工具帮助我们检查，这个就是`go build -race `命令。在项目目录下执行这个命令，生成一个可以执行文件，然后再运行这个可执行文件，就可以看到打印出的检测信息。

 **在`go build`命令中多加了一个`-race `标志，这样生成的可执行程序就自带了检测资源竞争的功能，运行生成的可执行文件**，效果如下所示：

```go
==================
WARNING: DATA RACE
Read at 0x000000619cbc by goroutine 8:
  main.incCount()
      D:/code/src/main.go:25 +0x80

Previous write at 0x000000619cbc by goroutine 7:
  main.incCount()
      D:/code/src/main.go:28 +0x9f

Goroutine 8 (running) created at:
  main.main()
      D:/code/src/main.go:17 +0x7e

Goroutine 7 (finished) created at:
  main.main()
      D:/code/src/main.go:16 +0x66
==================
4
Found 1 data race(s)
```

通过运行结果可以看出 goroutine 8 在代码 25 行读取共享资源` value := count`，而这时 goroutine 7 在代码 28 行修改共享资源` count = value`，而这两个 goroutine 都是从 main 函数的 16、17 行通过 go 关键字启动的。

#### 锁住共享资源

Go语言提供了传统的同步 goroutine 的机制，就是对共享资源加锁。atomic 和 sync 包里的一些函数就可以对共享的资源进行加锁操作。

##### 方式一：原子函数(atomic包)

原子函数能够以很底层的加锁机制来同步访问整型变量和指针，示例代码如下所示：

```go
    package main
    import (
        "fmt"
        "runtime"
        "sync"
        "sync/atomic"
    )
    var (
        counter int64
        wg      sync.WaitGroup
    )
    func main() {
        wg.Add(2)
        go incCounter(1)
        go incCounter(2)
        wg.Wait() //等待goroutine结束
        fmt.Println(counter)
    }
    func incCounter(id int) {
        defer wg.Done()
        for count := 0; count < 2; count++ {
            atomic.AddInt64(&counter, 1) //安全的对counter加1,强制同一时刻只能有一个 gorountie 运行并完成这个加法操作
            runtime.Gosched()
        }
    }
```

上述代码中使用了 **atmoic 包的 AddInt64 函数，这个函数会同步整型值的加法，方法是强制同一时刻只能有一个 gorountie  运行并完成这个加法操作**。当 goroutine 试图去调用任何原子函数时，这些 goroutine 都会自动根据所引用的变量做同步处理。

**另外两个有用的原子函数是 LoadInt64 和 StoreInt64**。这两个函数**提供了一种安全地读和写一个整型值的方式**。下面是代码就使用了  LoadInt64 和 StoreInt64 函数来创建一个同步标志，这个标志可以向程序里多个 goroutine 通知某个特殊状态。

```go
    package main
    import (
        "fmt"
        "sync"
        "sync/atomic"
        "time"
    )
    var (
        shutdown int64
        wg       sync.WaitGroup
    )
    func main() {
        wg.Add(2)
        go doWork("A")
        go doWork("B")
        time.Sleep(1 * time.Second)
        fmt.Println("Shutdown Now")
        atomic.StoreInt64(&shutdown, 1)    //原子写
        wg.Wait()
    }
    func doWork(name string) {
        defer wg.Done()
        for {
            fmt.Printf("Doing %s Work\n", name)
            time.Sleep(250 * time.Millisecond)
            if atomic.LoadInt64(&shutdown) == 1 {   //原子读
                fmt.Printf("Shutting %s Down\n", name)
                break
            }
        }
    }
```

上面代码中 main 函数使用 StoreInt64 函数来安全地修改 shutdown 变量的值。如果哪个 doWork goroutine  试图在 main 函数调用 StoreInt64 的同时调用 LoadInt64  函数，那么原子函数会将这些调用互相同步，保证这些操作都是安全的，不会进入竞争状态。

##### 方式二：互斥锁(sync包)

另一种同步访问共享资源的方式是使用互斥锁，互斥锁这个名字来自互斥的概念。互斥锁用于在代码上创建一个临界区，保证同一时间只有一个 goroutine 可以执行这个临界代码。

示例代码如下所示：

```go
    package main
    import (
        "fmt"
        "runtime"
        "sync"
    )
    var (
        counter int64
        wg      sync.WaitGroup
        mutex   sync.Mutex
    )
    func main() {
        wg.Add(2)
        go incCounter(1)
        go incCounter(2)
        wg.Wait()
        fmt.Println(counter)
    }
    func incCounter(id int) {
        defer wg.Done()
        for count := 0; count < 2; count++ {
            //同一时刻只允许一个goroutine进入这个临界区
            mutex.Lock()
            {
                value := counter
                runtime.Gosched()
                value++
                counter = value
            }
            mutex.Unlock() //释放锁，允许其他正在等待的goroutine进入临界区
        }
    }
```

同一时刻只有一个 goroutine 可以进入临界区。之后直到调用 Unlock 函数之后，其他 goroutine 才能进去临界区。当调用  runtime.Gosched 函数强制将当前 goroutine 退出当前线程后，调度器会再次分配这个 goroutine 继续运行。

### 三、GOMAXPROCS（调整并发的运行性能）

在 Go语言程序运行时（runtime）实现了一个小型的任务调度器。这套调度器的工作原理类似于操作系统调度线程，Go 程序调度器可以高效地将  CPU 资源分配给每一个任务。传统逻辑中，开发者需要维护线程池中线程与 CPU 核心数量的对应关系。同样的，Go 地中也可以通过  runtime.GOMAXPROCS() 函数做到，格式为：

```go
runtime.GOMAXPROCS(逻辑CPU数量)
```

这里的逻辑CPU数量可以有如下几种数值：

- <1：不修改任何数值。
- =1：单核心执行。
- \>1：多核并发执行。

一般情况下，可以使用 runtime.NumCPU() 查询 CPU 数量，并使用 runtime.GOMAXPROCS() 函数进行设置，例如：

```go
runtime.GOMAXPROCS(runtime.NumCPU())
```

Go 1.5 版本之前，默认使用的是单核心执行。从 Go 1.5 版本开始，默认执行上面语句以便让代码并发执行，最大效率地利用 CPU。GOMAXPROCS 同时也是一个环境变量，在应用程序启动前设置环境变量也可以起到相同的作用。

### 四、通道（chan）——goroutine之间通信的管道

Go语言中的通道（channel）是一种特殊的类型。**在任何时候，同时只能有一个 goroutine 访问通道进行发送和获取数据。**goroutine 间通过通道就可以通信。

通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。

```go
    ch1 := make(chan int)                 // 创建一个整型类型的通道
    ch2 := make(chan interface{})         // 创建一个空接口类型的通道, 可以存放任意格式
    type Equip struct{ /* 一些字段 */ }
    ch2 := make(chan *Equip)             // 创建Equip指针类型的通道, 可以存放*Equip
```

#### 1.发送将持续阻塞直到数据被接收

把数据往通道中发送时，如果接收方一直都没有接收，那么发送操作将持续阻塞。同时，Go 程序运行时能**智能地**发现一些**永远无法发送成功的语句**并做出提示，代码如下：

```go
    package main
    func main() {
        // 创建一个整型通道
        ch := make(chan int)
        // 尝试将0通过通道发送
        ch <- 0
    }
```

运行代码，报错：

```go
fatal error: all goroutines are asleep - deadlock!
```

报错的意思是：运行时发现所有的 goroutine（包括main）都处于等待 goroutine。也就是说所有 goroutine 中的 channel 并**没有形成发送和接收对应的代码**。

#### 2.使用通道接收数据

通道接收同样使用`<-`操作符，通道接收有如下特性：

① 通道的收发操作在**两个不同的 goroutine 间**进行。(如果在一个 goroutine 内同时进行管道的读写，会提示死锁deadlock)

② 接收将持续阻塞直到发送方发送数据。

③ 每次接收一个元素。

通道的数据接收一共有以下 4 种写法。

##### 1) 阻塞接收数据

阻塞模式接收数据时，将接收变量作为`<-`操作符的左值，格式如下：

```go
data := <-ch    //执行该语句时将会阻塞，直到接收到数据并赋值给 data 变量。 
```

##### 2) 非阻塞接收数据

使用非阻塞方式从通道接收数据时，**语句不会发生阻塞**，格式如下：

```go
data, ok := <-ch
```

- data：表示接收到的数据。**未接收到数据时，data 为通道类型的零值**。
- ok：表示是否接收到数据。

非阻塞的通道接收方法**可能造成高的 CPU 占用，因此使用非常少**。如果需要实现接收超时检测，可以配合 select 和计时器 channel 进行.

##### 3) 接收任意数据，忽略接收的数据

阻塞接收数据后，忽略从通道返回的数据，格式如下：

```go
<-ch      //执行该语句时将会发生阻塞，直到接收到数据，但接收到的数据会被忽略。这个方式实际上只是通过通道在 goroutine 间阻塞收发实现并发同步。
```

##### 4) 循环接收

**通道的数据接收可以借用 for range 语句进行多个元素的接收操作**，格式如下：

```go
    for data := range ch {
    }
```

通道 ch 是可以进行遍历的，遍历的结果就是接收到的数据。数据类型就是通道的数据类型。通过 for 遍历获得的变量只有一个，即上面例子中的 data。

遍历通道数据的例子请参考下面的代码。

```go
    package main
    import (
        "fmt"
        "time"
    )
    func main() {
        // 构建一个通道
        ch := make(chan int)
        // 开启一个并发匿名函数
        go func() {
            // 从3循环到0
            for i := 3; i >= 0; i-- {
                // 发送3到0之间的数值
                ch <- i
                // 每次发送完时等待
                time.Sleep(time.Second)
            }
        }()
        // 遍历接收通道数据
        for data := range ch {
            // 打印通道数据
            fmt.Println(data)
            // 当遇到数据0时, 退出接收循环
            if data == 0 {
                    break
            }
        }
    }
```

执行代码，输出如下：

```go
3
2
1
0
```

#### 3.并发打印（借助通道实现）

下面通过一个并发打印的例子，将 goroutine 和 channel 放在一起展示它们的用法。

```go
package main
import (
    "fmt"
)
func printer(c chan int) {
    // 开始无限循环等待数据
    for {
        // 从channel中获取一个数据
        data := <-c
        // 将0视为数据结束
        if data == 0 {
            break
        }
        // 打印数据
        fmt.Println(data)
    }
    // 通知main已经结束循环(我搞定了!)
    c <- 0
}
func main() {
    // 创建一个channel
    c := make(chan int)
    // 并发执行printer, 传入channel
    go printer(c)
    for i := 1; i <= 10; i++ {
        // 将数据通过channel投送给printer
        c <- i
    }
    // 通知并发的printer结束循环(没数据啦!)
    c <- 0
    // 等待printer结束(搞定喊我!)
    <-c

```

运行代码，输出如下：

```go
1
2
3
4
5
6
7
8
9
10
```

**本例的[设计模式](http://c.biancheng.net/design_pattern/)就是典型的生产者和消费者**。生产者是第 37 行的循环，而消费者是 printer() 函数。整个例子使用了两个 goroutine，一个是 main()，一个是通过第 35 行  printer() 函数创建的 goroutine。两个 goroutine 通过第 32 行创建的通道进行通信。**这个通道有下面两重功能**。

- **数据传送**：第 40 行中发送数据和第 13 行接收数据。
- **控制指令**：类似于信号量的功能。同步 goroutine 的操作。功能简单描述为：
  - 第 44 行：“没数据啦！”
  - 第 25 行：“我搞定了！”
  - 第 47 行：“搞定喊我！”

### 五、单向通道

Go语言的类型系统提供了单方向的 channel 类型，顾名思义，单向 channel 就是只能用于写入或者只能用于读取数据。当然 channel 本身必然是同时支持读写的，否则根本没法用。

假如一个 channel 真的只能读取数据，那么它肯定只会是空的，因为你没机会往里面写数据。同理，如果一个 channel  只允许写入数据，即使写进去了，也没有丝毫意义，因为没有办法读取到里面的数据。所谓的单向 channel 概念，其实只是对 channel  的一种使用限制。

#### 1.单向通道的使用例子

示例代码如下：

```go
    ch := make(chan int)
    // 声明一个只能写入数据的通道类型, 并赋值为ch
    var chSendOnly chan<- int = ch
    //声明一个只能读取数据的通道类型, 并赋值为ch
    var chRecvOnly <-chan int = ch
```

上面的例子中，chSendOnly 只能写入数据，如果尝试读取数据，将会出现如下报错：

```go
invalid operation: <-chSendOnly (receive from send-only type chan<- int)
```

同理，chRecvOnly 也是不能写入数据的。

当然，使用 make 创建通道时，也可以创建一个只写入或只读取的通道：

```go
    ch := make(<-chan int)
    var chReadOnly <-chan int = ch
    <-chReadOnly
```

上面代码编译正常，运行也是正确的。但是，**一个不能写入数据只能读取的通道是毫无意义的**。

其次，**我们在将一个 channel 变量传递到一个函数时，可以通过将其指定为单向 channel 变量，从而限制该函数中可以对此 channel 的操作，比如只能往这个 channel 中写入数据，或者只能从这个 channel 读取数据。**

#### 2.关闭 channel

关闭 channel 非常简单，直接使用Go语言内置的 close() 函数即可：

```go
close(ch)
```

在介绍了如何关闭 channel 之后，我们就多了一个问题：如何判断一个 channel 是否已经被关闭？我们可以在读取的时候使用多重返回值的方式：

```go
x, ok := <-ch
```

这个用法与 map 中的按键获取 value 的过程比较类似，只需要看第二个 bool 返回值即可，如果返回值是 false 则表示 ch 已经被关闭。
