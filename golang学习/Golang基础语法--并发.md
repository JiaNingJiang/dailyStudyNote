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

### 六、无缓冲与有缓冲管道

Go语言中有缓冲的通道（buffered channel）是一种在被接收前能存储一个或者多个值的通道。这种类型的通道并不强制要求  goroutine  之间必须同时完成发送和接收。通道会阻塞发送和接收动作的条件也会不同。只有在通道中没有要接收的值时，接收动作才会阻塞。只有在通道没有可用缓冲区容纳被发送的值时，发送动作才会阻塞。

 这导致有缓冲的通道和无缓冲的通道之间的一个很大的不同：无缓冲的通道保证进行发送和接收的 goroutine 会在同一时间进行数据交换；有缓冲的通道没有这种保证。

 在无缓冲通道的基础上，为通道增加一个有限大小的存储空间形成带缓冲通道。带缓冲通道在发送时无需等待接收方接收即可完成发送过程，并且不会发生阻塞，只有当存储空间满时才会发生阻塞。同理，如果缓冲通道中有数据，接收时将不会发生阻塞，直到通道中没有数据可读时，通道将会再度阻塞。

 无缓冲通道保证收发过程同步。无缓冲收发过程类似于快递员给你电话让你下楼取快递，整个递交快递的过程是同步发生的，你和快递员不见不散。但这样做快递员就必须等待所有人下楼完成操作后才能完成所有投递工作。如果快递员将快递放入快递柜中，并通知用户来取，快递员和用户就成了异步收发过程，效率可以有明显的提升。带缓冲的通道就是这样的一个“快递柜”。

#### 有缓冲管道阻塞条件

带缓冲通道在很多特性上和无缓冲通道是类似的。无缓冲通道可以看作是长度永远为 0 的带缓冲通道。因此根据这个特性，带缓冲通道在下面列举的情况下依然会发生阻塞：

- 带缓冲通道被填满时，尝试再次发送数据时发生阻塞。
- 带缓冲通道为空时，尝试接收数据时发生阻塞。

##### 为什么Go语言对通道要限制长度而不提供无限长度的通道？

我们知道通道（channel）是在两个 goroutine 间通信的桥梁。使用 goroutine  的代码必然有一方提供数据，一方消费数据。当提供数据一方的数据供给速度大于消费方的数据处理速度时，如果通道不限制长度，那么内存将不断膨胀直到应用崩溃。因此，限制通道的长度有利于约束数据提供方的供给速度，供给数据量必须在消费方处理量+通道长度的范围内，才能正常地处理数据。

### 七、channel超时机制

select 的用法与 switch 语言非常类似，由 select 开始一个新的选择块，每个选择条件由 case 语句来描述。与 switch 语句相比，select 有比较多的限制，其中最大的一条限制就是每个 case 语句里必须是一个 IO 操作，大致的结构如下：

```go
select {
    case <-chan1:
    // 如果chan1成功读到数据，则进行该case处理语句
    case chan2 <- 1:
    // 如果成功向chan2写入数据，则进行该case处理语句
    default:
    // 如果上面都没有成功，则进入default处理流程
}
```

在一个 select 语句中，Go语言会按顺序**从头至尾**评估每一个发送和接收的语句。

 如果其中的任意一语句可以继续执行（即没有被阻塞），那么就从那些可以执行的语句中任意选择一条来使用。

 如果没有任意一条语句可以执行（即所有的通道都被阻塞），那么有如下两种可能的情况：

- 如果给出了 default 语句，那么就会执行 default 语句，同时程序的执行会从 select 语句后的语句中恢复；
- **如果没有 default 语句，那么 select 语句将被阻塞，直到至少有一个通信可以进行下去。**

示例代码如下所示：

```go
    package main
    import (
        "fmt"
        "time"
    )
    func main() {
        ch := make(chan int)
        quit := make(chan bool)
        //新开一个协程
        go func() {
            for {
                select {
                case num := <-ch:
                    fmt.Println("num = ", num)
                case <-time.After(3 * time.Second):
                    fmt.Println("超时")
                    quit <- true
                }
            }
        }() //别忘了()
        for i := 0; i < 5; i++ {
            ch <- i
            time.Sleep(time.Second)
        }
        <-quit
        fmt.Println("程序结束")
    }
```

运行结果如下：

```go
num =  0
num =  1
num =  2
num =  3
num =  4
超时
程序结束
```

### 八、等待组

Go语言中除了可以使用通道（channel）和互斥锁进行两个并发程序间的同步外，还可以使用等待组进行多个任务的同步，等待组可以保证在并发环境中完成指定数量的任务

 在 sync.WaitGroup（等待组）类型中，每个 sync.WaitGroup 值在内部维护着一个计数，此计数的初始默认值为零。

 等待组有下面几个方法可用，如下表所示。

| 方法名                          | 功能                                    |
| ------------------------------- | --------------------------------------- |
| (wg * WaitGroup) Add(delta int) | 等待组的计数器 +1                       |
| (wg * WaitGroup) Done()         | 等待组的计数器 -1                       |
| (wg * WaitGroup) Wait()         | 当等待组计数器不等于 0 时阻塞直到变 0。 |

对于一个可寻址的 sync.WaitGroup 值 wg：

- 我们可以使用方法调用 wg.Add(delta) 来改变值 wg 维护的计数。
- 方法调用 wg.Done() 和 wg.Add(-1) 是完全等价的。
- 如果一个 wg.Add(delta) 或者 wg.Done() 调用将 wg 维护的计数更改成一个负数，一个恐慌将产生。
- 当一个协程调用了 wg.Wait() 时，
  - 如果此时 wg 维护的计数为零，则此 wg.Wait() 此操作为一个空操作（noop）；
  - 否则（计数为一个正整数），此协程将进入**阻塞状态**。当以后其它某个协程将此计数更改至 0 时（一般通过调用 wg.Done()），此协程将重新进入运行状态（即 wg.Wait() 将返回）。

等待组内部拥有一个计数器，计数器的值可以通过方法调用实现计数器的增加和减少。当我们添加了 N 个并发任务进行工作时，就将等待组的计数器值增加  N。每个任务完成时，这个值减 1。同时，在另外一个 goroutine 中等待这个等待组的计数器值为 0 时，表示所有任务已经完成。

下面的代码演示了这一过程：

```go
package main
import (
    "fmt"
    "net/http"
    "sync"
)
func main() {
    // 声明一个等待组
    var wg sync.WaitGroup
    // 准备一系列的网站地址
    var urls = []string{
        "http://www.github.com/",
        "https://www.qiniu.com/",
        "https://www.golangtc.com/",
    }
    // 遍历这些地址
    for _, url := range urls {
        // 每一个任务开始时, 将等待组增加1
        wg.Add(1)
        // 开启一个并发
        go func(url string) {
            // 使用defer, 表示函数完成时将等待组值减1
            defer wg.Done()
            // 使用http访问提供的地址
            _, err := http.Get(url)
            // 访问完成后, 打印地址和可能发生的错误
            fmt.Println(url, err)
            // 通过参数传递url地址
        }(url)
    }
    // 等待所有的任务完成
    wg.Wait()
    fmt.Println("over")
```

### 九、死锁、活锁和饥饿概述

#### 1.死锁

死锁是指两个或两个以上的进程（或线程）在执行过程中，因争夺资源而造成的一种互相等待的现象，若无外力作用，它们都将无法推进下去。此时称系统处于死锁状态或系统产生了死锁，这些永远在互相等待的进程称为死锁进程。

 死锁发生的条件有如下几种：

##### 1) 互斥条件

线程对资源的访问是排他性的，如果一个线程对占用了某资源，那么其他线程必须处于等待状态，直到该资源被释放。

##### 2) 请求和保持条件

线程 T1 至少已经保持了一个资源 R1 占用，但又提出使用另一个资源 R2 请求，而此时，资源 R2 被其他线程 T2 占用，于是该线程 T1 也必须等待，但又对自己保持的资源 R1 不释放。

##### 3) 不剥夺条件

线程已获得的资源，在未使用完之前，不能被其他线程剥夺，只能在使用完以后由自己释放。

##### 4) 环路等待条件

在死锁发生时，必然存在一个“进程 - 资源环形链”，即：{p0,p1,p2,...pn}，进程 p0（或线程）等待 p1 占用的资源，p1 等待 p2 占用的资源，pn 等待 p0 占用的资源。

 最直观的理解是，p0 等待 p1 占用的资源，而 p1 而在等待 p0 占用的资源，于是两个进程就相互等待。



死锁解决办法：

- 如果并发查询多个表，约定访问顺序；
- 在同一个事务中，尽可能做到一次锁定获取所需要的资源；
- 对于容易产生死锁的业务场景，尝试升级锁颗粒度，使用表级锁；
- 采用分布式事务锁或者使用乐观锁。


 死锁程序是所有并发进程彼此等待的程序，在这种情况下，如果没有外界的干预，这个程序将永远无法恢复。

#### 2.活锁

活锁是另一种形式的活跃性问题，该问题尽管不会阻塞线程，但也不能继续执行，因为线程将不断重复同样的操作，而且总会失败。

 例如线程 1 可以使用资源，但它很礼貌，让其他线程先使用资源，线程 2 也可以使用资源，但它同样很绅士，也让其他线程先使用资源。就这样你让我，我让你，最后两个线程都无法使用资源。

 活锁通常发生在处理事务消息中，如果不能成功处理某个消息，那么消息处理机制将回滚事务，并将它重新放到队列的开头。这样，错误的事务被一直回滚重复执行，这种形式的活锁通常是由过度的错误恢复代码造成的，因为它错误地将不可修复的错误认为是可修复的错误。

 当多个相互协作的线程都对彼此进行相应而修改自己的状态，并使得任何一个线程都无法继续执行时，就导致了活锁。这就像两个过于礼貌的人在路上相遇，他们彼此让路，然后在另一条路上相遇，然后他们就一直这样避让下去。

 要解决这种活锁问题，需要在重试机制中引入随机性。例如在网络上发送数据包，如果检测到冲突，都要停止并在一段时间后重发，如果都在 1 秒后重发，还是会冲突，所以引入随机性可以解决该类问题。

#### 3.饥饿

饥饿是指一个可运行的进程尽管能继续执行，但被调度器无限期地忽视，而不能被调度执行的情况。

 与死锁不同的是，饥饿锁在一段时间内，优先级低的线程最终还是会执行的，比如高优先级的线程执行完之后释放了资源。

 活锁与饥饿是无关的，因为在活锁中，所有并发进程都是相同的，并且没有完成工作。更广泛地说，饥饿通常意味着有一个或多个贪婪的并发进程，它们不公平地阻止一个或多个并发进程，以尽可能有效地完成工作，或者阻止全部并发进程。

 下面的示例程序中包含了一个贪婪的 goroutine 和一个平和的 goroutine：

```go
    package main
    import (
        "fmt"
        "runtime"
        "sync"
        "time"
    )
    func main() {
        runtime.GOMAXPROCS(3)
        var wg sync.WaitGroup
        const runtime = 1 * time.Second
        var sharedLock sync.Mutex
        greedyWorker := func() {
            defer wg.Done()
            var count int
            for begin := time.Now(); time.Since(begin) <= runtime; {
                sharedLock.Lock()
                time.Sleep(3 * time.Nanosecond)
                sharedLock.Unlock()
                count++
            }
            fmt.Printf("Greedy worker was able to execute %v work loops\n", count)
        }
        politeWorker := func() {
            defer wg.Done()
            var count int
            for begin := time.Now(); time.Since(begin) <= runtime; {
                sharedLock.Lock()
                time.Sleep(1 * time.Nanosecond)
                sharedLock.Unlock()
                sharedLock.Lock()
                time.Sleep(1 * time.Nanosecond)
                sharedLock.Unlock()
                sharedLock.Lock()
                time.Sleep(1 * time.Nanosecond)
                sharedLock.Unlock()
                count++
            }
            fmt.Printf("Polite worker was able to execute %v work loops\n", count)
        }
        wg.Add(2)
        go greedyWorker()
        go politeWorker()
        wg.Wait()
    }
```

输出如下：

```go
Greedy worker was able to execute 276 work loops
Polite worker was able to execute 92 work loops
```

贪婪的 worker 会贪婪地抢占共享锁，以完成整个工作循环，而平和的 worker 则试图只在需要时锁定。两种 worker  都做同样多的模拟工作（sleeping 时间为 3ns），可以看到，在同样的时间里，贪婪的 worker 工作量几乎是平和的 worker  工作量的两倍！

 假设两种 worker 都有同样大小的临界区，而不是认为贪婪的 worker 的算法更有效（或调用 Lock 和 Unlock  的时候，它们也不是缓慢的），我们得出这样的结论，贪婪的 worker 不必要地扩大其持有共享锁上的临界区，井阻止（通过饥饿）平和的 worker 的 goroutine 高效工作。

### 总结

不恰当的用锁肯定会出问题。如果用了，虽然解了前面的问题，但是又出现了更多的新问题。

- 死锁：是因为错误的使用了锁，导致异常；
- 活锁：是饥饿的一种特殊情况，逻辑上感觉对，程序也一直在正常的跑，但就是效率低，逻辑上进行不下去；
- 饥饿：与锁使用的粒度有关，通过计数取样，可以判断进程的工作效率。

### 十、CSP：通信顺序进程

Go实现了两种并发形式，第一种是大家普遍认知的多线程共享内存，其实就是 [Java](http://c.biancheng.net/java/) 或 [C++](http://c.biancheng.net/cplus/) 等语言中的多线程开发；另外一种是Go语言特有的，也是Go语言推荐的 CSP（communicating sequential processes）并发模型。

在并发编程中，对共享资源的正确访问需要精确地控制，在目前的绝大多数语言中，都是通过加锁等线程同步方案来解决这一困难问题，而Go语言却另辟蹊径，它将共享的值通过通道传递（实际上多个独立执行的线程很少主动共享资源）。

1. 并发编程的核心概念是同步通信，但是同步的方式却有多种。先以大家熟悉的互斥量 sync.Mutex 来实现同步通信，示例代码如下所示：

```go
    package main
    import (
        "fmt"
        "sync"
    )
    func main() {
        var mu sync.Mutex
        go func() {
            fmt.Println("C语言中文网")
            mu.Lock()
        }()
        mu.Unlock()
    }
```

由于 mu.Lock() 和 mu.Unlock() 并不在同一个 Goroutine 中，所以也就不满足顺序一致性内存模型。同时它们也没有其他的同步事件可以参考，也就是说这两件事是可以并发的。由于可能是并发的事件，所以 **main() 函数中的 mu.Unlock() 很有可能先发生，而这个时刻 mu 互斥对象还处于未加锁的状态，因而会导致运行时异常**。

2. 下面是修复后的代码：

```go
    package main
    import (
        "fmt"
        "sync"
    )
    func main() {
        var mu sync.Mutex
        mu.Lock()
        go func() {
            fmt.Println("C语言中文网")
            mu.Unlock()
        }()
        mu.Lock()
    }
```

修复的方式是在 main() 函数所在线程中执行两次 mu.Lock()，当第二次加锁时会因为锁已经被占用（不是递归锁）而阻塞，main() 函数的阻塞状态驱动后台线程继续向前执行。

 **当后台线程执行到 mu.Unlock() 时解锁，此时打印工作已经完成了，解锁会导致 main() 函数中的第二个 mu.Lock()  阻塞状态取消**，此时后台线程和主线程再没有其他的同步事件参考，**它们退出的事件将是并发的**，在 main()  函数退出导致程序退出时，后台线程可能已经退出了，也可能没有退出。**虽然无法确定两个线程退出的时间，但是打印工作是可以正确完成的**。

3. 使用 sync.Mutex 互斥锁同步是比较低级的做法，我们现在**改用无缓存通道来实现同步**：

```go
    package main
    import (
        "fmt"
    )
    func main() {
        done := make(chan int)
        go func() {
            fmt.Println("C语言中文网")
            <-done
        }()
        done <- 1
    }
```

根据Go语言内存模型规范，对于从**无缓存通道**进行的接收，发生在对该通道进行的发送完成之前。因此，后台线程`<-done `接收操作完成之后，main 线程的`done <- 1 `发送操作才可能完成（从而退出 main、退出程序），而此时打印工作已经完成了。

4. 上面的代码虽然可以正确同步，但是**对通道的缓存大小太敏感**，如果通道有缓存，就无法保证 main() 函数退出之前后台线程能正常打印了，**更好的做法是将通道的发送和接收方向调换一下，这样可以避免同步事件受通道缓存大小的影响**：

```go
    package main
    import (
        "fmt"
    )
    func main() {
        done := make(chan int, 1) // 带缓存通道
        go func() {
            fmt.Println("C语言中文网")
            done <- 1
        }()
        <-done
    }
```

5. 基于**带缓存通道**，我们可以很容易将打印线程扩展到 N 个，下面的示例是开启 10 个后台线程分别打印：

```go
    package main
    import (
        "fmt"
    )
    func main() {
        done := make(chan int, 10) // 带10个缓存
        // 开N个后台打印线程
        for i := 0; i < cap(done); i++ {
            go func() {
                fmt.Println("C语言中文网")
                done <- 1
            }()
        }
        // 等待N个后台线程完成
        for i := 0; i < cap(done); i++ {
            <-done
        }
    }
```

6. 对于这种要等待 N 个线程完成后再进行下一步的同步操作有一个简单的做法，就是使用 sync.WaitGroup 来等待一组事件：

```go
    package main
    import (
        "fmt"
        "sync"
    )
    func main() {
        var wg sync.WaitGroup
        // 开N个后台打印线程
        for i := 0; i < 10; i++ {
            wg.Add(1)
            go func() {
                fmt.Println("C语言中文网")
                wg.Done()
            }()
        }
        // 等待N个后台线程完成
        wg.Wait()
    }
```

其中 wg.Add(1)  用于增加等待事件的个数，必须确保在后台线程启动之前执行（如果放到后台线程之中执行则不能保证被正常执行到）。当后台线程完成打印工作之后，调用  wg.Done() 表示完成一个事件，main() 函数的 wg.Wait() 是等待全部的事件完成。
