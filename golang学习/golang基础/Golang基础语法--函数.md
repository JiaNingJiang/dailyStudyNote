

### 一、函数的返回值

1. ##### 带有变量名的返回值

Go语言支持对返回值进行命名，这样返回值就和参数一样拥有参数变量名和类型。

下面代码中的函数拥有两个整型返回值，函数声明时将返回值命名为 a 和 b，因此可以在函数体中直接对函数返回值进行赋值，在命名的返回值方式的函数体中，在函数结束前需要显式地使用 return 语句进行返回，代码如下：

```go
    func namedRetValues() (a, b int) {
        a = 1
        b = 2
        return
    }
```

代码说明如下：

- 第 1 行，对两个整型返回值进行命名，分别为 a 和 b。
- 第 3 行和第 4 行，命名返回值的变量与这个函数的布局变量的效果一致，可以对返回值进行赋值和值获取。
- 第 6 行，当函数使用命名返回值时，可以在 return 中不填写返回值列表，如果填写也是可行的，下面代码的执行效果和上面代码的效果一样。

```go
    func namedRetValues() (a, b int) {
        a = 1
        return a, 2
    }
```

### 二、匿名函数

1. ##### 匿名函数做回调函数

下面的代码实现对切片的遍历操作，遍历中访问每个元素的操作使用匿名函数来实现，用户传入不同的匿名函数体可以实现对元素不同的遍历操作，代码如下：

```go
    package main
    import (
        "fmt"
    )
    // 遍历切片的每个元素, 通过给定函数进行元素访问
    func visit(list []int, f func(int)) {
        for _, v := range list {
            f(v)
        }
    }
    func main() {
        // 使用匿名函数打印切片内容
        visit([]int{1, 2, 3, 4}, func(v int) {
            fmt.Println(v)
        })
    }
```

代码说明如下：

- 第 8 行，使用 visit() 函数将整个遍历过程进行封装，当要获取遍历期间的切片值时，只需要给 visit() 传入一个回调参数即可。
- 第 18 行，准备一个整型切片 []int{1,2,3,4} 传入 visit() 函数作为遍历的数据。
- 第 19～20 行，定义了一个匿名函数，作用是将遍历的每个值打印出来。

2. ##### 使用匿名函数实现操作封装

下面这段代码将匿名函数作为 map 的value值，通过命令行参数动态调用匿名函数，代码如下：

```go
    package main
    import (
        "flag"
        "fmt"
    )
    var skillParam = flag.String("skill", "", "skill to perform")
    func main() {
        flag.Parse()
        var skill = map[string]func(){
            "fire": func() {
                fmt.Println("chicken fire")
            },
            "run": func() {
                fmt.Println("soldier run")
            },
            "fly": func() {
                fmt.Println("angel fly")
            },
        }
        if f, ok := skill[*skillParam]; ok {
            f()
        } else {
            fmt.Println("skill not found")
        }
    }
```

代码说明如下：

- 第 8 行，定义命令行参数 skill，从命令行输入 --skill 可以将` = `后的字符串传入 skillParam 指针变量。
- 第 12 行，解析命令行参数，解析完成后，skillParam 指针变量将指向命令行传入的值。
- 第 14 行，定义一个从字符串映射到 func() 的 map，然后填充这个 map。
- 第 15～23 行，初始化 map 的键值对，值为匿名函数。
- 第 26 行，skillParam 是一个 *string 类型的指针变量，使用 *skillParam 获取到命令行传过来的值，并在 map 中查找对应命令行参数指定的字符串的函数。
- 第 29 行，如果在 map 定义中存在这个参数就调用，否则打印“技能没有找到”。

运行代码，结果如下：

```go
PS D:\code> go run main.go --skill=fly
angel fly
PS D:\code> go run main.go --skill=run
soldier run 
```

### 三、函数类型实现接口

函数和其他类型一样都属于“一等公民”，其他类型能够实现接口，函数也可以，本节将对结构体与函数实现接口的过程进行对比。

假设有如下一个接口：

```go
    // 调用器接口
    type Invoker interface {
        // 需要实现一个Call()方法
        Call(interface{})
    }
```

这个接口需要实现 Call() 方法，调用时会传入一个 interface{} 类型的变量，这种类型的变量表示任意类型的值。

1. ##### 结构体实现接口

```go
    // 结构体类型
    type Struct struct {
    }
    // 实现Invoker的Call
    func (s *Struct) Call(p interface{}) {
        fmt.Println("from struct", p)
    }
```

将定义的 Struct 类型实例化，并传入接口中进行调用，代码如下：

```go
    // 声明接口变量
    var invoker Invoker
    // 实例化结构体
    s := new(Struct)
    // 将实例化的结构体赋值到接口
    invoker = s
    // 使用接口调用实例化结构体的方法Struct.Call
    invoker.Call("hello")
```



2. ##### 函数体实现接口

函数的声明不能直接实现接口，需要将函数定义为类型后，使用类型实现结构体，当类型方法被调用时，还需要调用函数本体。

```go
    // 函数定义为类型
    type FuncCaller func(interface{})
    // 实现Invoker的Call
    func (f FuncCaller) Call(p interface{}) {
        // 调用f()函数本体
        f(p)
    }
```

上面代码只是定义了函数类型，需要函数本身进行逻辑处理，FuncCaller 无须被实例化，只需要将函数转换为 FuncCaller 类型即可，函数来源可以是命名函数、匿名函数或闭包，参见下面代码：

```go
    // 声明接口变量
    var invoker Invoker
    // 将匿名函数转为FuncCaller类型, 再赋值给接口
    invoker = FuncCaller(func(v interface{}) {   //强制类型转换
        fmt.Println("from function", v)
    })
    // 使用接口调用FuncCaller.Call, 内部会调用函数本体
    invoker.Call("hello")
```

### 四、闭包

Go语言中闭包是引用了自由变量的函数，被引用的自由变量和函数一同存在，即使已经离开了自由变量的环境也不会被释放或者删除，在闭包中可以继续使用这个自由变量，因此，简单的说：

> 函数 + 引用环境 = 闭包

同一个函数与不同引用环境组合，可以形成不同的实例，如下图所示。

![image-20220922185716516](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\image-20220922185716516.png)

一个函数类型就像结构体一样，可以被实例化，函数本身不存储任何信息，只有与引用环境结合后形成的闭包才具有“记忆性”，函数是编译期静态的概念，而闭包是运行期动态的概念。(闭包（Closure）在某些编程语言中也被称为 Lambda 表达式。)

##### 1.在闭包内部修改引用的变量

```go
    // 准备一个字符串
    str := "hello world"
    // 创建一个匿名函数
    foo := func() {
       
        // 匿名函数中访问str
        str = "hello dude"
    }
    // 调用匿名函数
    foo()
```

代码说明如下：

- 第 2 行，准备一个字符串用于修改。
- 第 5 行，创建一个匿名函数。
- 第 8 行，在匿名函数中并没有定义 str，str 的定义在匿名函数之前，此时，str 就被引用到了匿名函数中形成了闭包。
- 第 12 行，执行闭包，此时 str 发生修改，变为 hello dude。

##### 2.示例：闭包的记忆效应

被捕获到闭包中的变量让闭包本身拥有了记忆效应，闭包中的逻辑可以修改闭包捕获的变量，变量会跟随闭包生命期一直存在，闭包本身就如同变量一样拥有了记忆效应。

> 累加器的实现：

```go
    package main
    import (
        "fmt"
    )
    // 提供一个值, 每次调用函数会指定对值进行累加
    func Accumulate(value int) func() int {
        // 返回一个闭包
        return func() int {
            // 累加
            value++
            // 返回一个累加值
            return value
        }
    }
    func main() {
        // 创建一个累加器, 初始值为1
        accumulator := Accumulate(1)
        // 累加1并打印
        fmt.Println(accumulator())
        fmt.Println(accumulator())
        // 打印累加器的函数地址
        fmt.Printf("%p\n", &accumulator)
        // 创建一个累加器, 初始值为1
        accumulator2 := Accumulate(10)
        // 累加1并打印
        fmt.Println(accumulator2())
        // 打印累加器的函数地址
        fmt.Printf("%p\n", &accumulator2)
    }
```

输出结果：

```go
2
3           
0xc000006028
11          
0xc000006038
```

代码说明如下：

- 第 8 行，累加器生成函数，这个函数输出一个初始值，调用时返回一个为初始值创建的闭包函数。
- 第 11 行，返回一个闭包函数，每次返回会创建一个新的函数实例。
- 第 14 行，对引用的 Accumulate 参数变量进行累加，注意 value 不是第 11 行匿名函数定义的，但是被这个匿名函数引用，所以形成闭包。
- 第 17 行，将修改后的值通过闭包的返回值返回。
- 第 24 行，创建一个累加器，初始值为 1，返回的 accumulator 是类型为 func()int 的函数变量。
- 第 27 行，调用 accumulator() 时，代码从 11 行开始执行匿名函数逻辑，直到第 17 行返回。
- 第 32 行，打印累加器的函数地址。

==对比输出的日志发现 accumulator 与 accumulator2 输出的函数地址不同，因此它们是两个不同的闭包实例。==

 ==每调用一次 accumulator 都会自动对引用的变量进行累加。==

##### 3.示例：闭包实现生成器

闭包的记忆效应被用于实现类似于[设计模式](http://c.biancheng.net/design_pattern/)中工厂模式的生成器，下面的例子展示了创建一个玩家生成器的过程。

> 玩家生成器的实现：

```go
    package main
    import (
        "fmt"
    )
    // 创建一个玩家生成器, 输入名称, 输出生成器
    func playerGen(name string) func() (string, int) {
        // 血量一直为150
        hp := 150
        // 返回创建的闭包
        return func() (string, int) {
            // 将变量引用到闭包中
            return name, hp
        }
    }
    func main() {
        // 创建一个玩家生成器
        generator := playerGen("high noon")
        // 返回玩家的名字和血量
        name, hp := generator()
        // 打印值
        fmt.Println(name, hp)
    }
```

代码输出如下：

```go
high noon 150
```

代码说明如下：

- 第 8 行，playerGen() 需要提供一个名字来创建一个玩家的生成函数。
- 第 11 行，声明并设定 hp 变量为 150。
- 第 14～18 行，将 hp 和 name 变量引用到匿名函数中形成闭包。
- 第 24 行中，通过 playerGen 传入参数调用后获得玩家生成器。
- 第 27 行，调用这个玩家生成器函数，可以获得玩家的名称和血量。

### 五、可变参数

##### 1.可变参数类型

可变参数是指函数传入的参数个数是可变的，为了做到这点，首先需要将函数定义为可以接受可变参数的类型：

```go
    func myfunc(args ...int) {
        for _, arg := range args {
            fmt.Println(arg)
        }
    }
```

上面这段代码的意思是，函数 myfunc() 接受不定数量的参数，这些参数的类型全部是 int，所以它可以用如下方式调用：

```go
myfunc(2, 3, 4)
myfunc(1, 3, 7, 13)
```

形如`...type`格式的类型只能作为函数的参数类型存在，并且必须是最后一个参数。从内部实现机理上来说，类型`...type`本质上是一个数组切片，也就是`[]type`，这也是为什么上面的参数 args 可以用 for 循环来获得每个传入的参数。

##### 2.任意类型的可变参数

之前的例子中将可变参数类型约束为 int，如果你希望传任意类型，可以指定类型为 interface{}，下面是Go语言标准库中 fmt.Printf() 的函数原型：

```go
func Printf(format string, args ...interface{}) {
    // ...
}
```

用 interface{} 传递任意类型数据是Go语言的惯例用法，使用 interface{} 仍然是类型安全的，这和 C/[C++](http://c.biancheng.net/cplus/) 不太一样，下面通过示例来了解一下如何分配传入 interface{} 类型的数据。

```go
    package main
    import "fmt"
    func MyPrintf(args ...interface{}) {
        for _, arg := range args {
            switch arg.(type) {
                case int:
                    fmt.Println(arg, "is an int value.")
                case string:
                    fmt.Println(arg, "is a string value.")
                case int64:
                    fmt.Println(arg, "is an int64 value.")
                default:
                    fmt.Println(arg, "is an unknown type.")
            }
        }
    }
    func main() {
        var v1 int = 1
        var v2 int64 = 234
        var v3 string = "hello"
        var v4 float32 = 1.234
        MyPrintf(v1, v2, v3, v4)
    }
```

该程序的输出结果为：

```go
1 is an int value.
234 is an int64 value.
hello is a string value.
1.234 is an unknown type.
```

##### 3.遍历可变参数列表——获取每一个参数的值

可变参数列表的数量不固定，传入的参数是一个切片，如果需要获得每一个参数的具体值时，可以对可变参数变量进行遍历。如果要获取可变参数的数量，可以使用 len() 函数对可变参数变量对应的切片进行求长度操作，以获得可变参数数量。

##### 4.在多个可变参数函数中传递参数

可变参数变量是一个包含所有参数的切片，如果要将这个含有可变参数的变量传递给下一个可变参数函数，可以==在传递时给可变参数变量后面添加`...`==，这样就可以==将切片中的元素进行传递，而不是传递可变参数变量本身==。

下面的例子模拟 print() 函数及实际调用的 rawPrint() 函数，两个函数都拥有可变参数，需要将参数从 print 传递到 rawPrint 中。

```go
    package main
    import "fmt"
    // 实际打印的函数
    func rawPrint(rawList ...interface{}) {
        // 遍历可变参数切片
        for _, a := range rawList {
            // 打印参数
            fmt.Println(a)
        }
    }
    // 打印函数封装
    func print(slist ...interface{}) {
        // 将slist可变参数切片完整传递给下一个函数
        rawPrint(slist...)
    }
    func main() {
        print(1, 2, 3)
    }
```

代码输出如下：

```go
1
2
3
```

如果尝试将第 20 行修改为：

```go
rawPrint(slist)
```

再次执行代码，将输出：

```go
[1 2 3]
```

此时，slist（类型为 []interface{}）将被作为一个整体传入 rawPrint()，rawPrint() 函数中遍历的变量也就是 slist 的切片值。

### 六、defer（延迟执行语句）

Go语言的 defer 语句会将其后面跟随的语句进行延迟处理，在 defer 归属的函数即将返回时，将==延迟处理的语句按 defer 的逆序进行执行==，也就是说，==先被 defer 的语句最后被执行，最后被 defer 的语句，最先被执行==。（延迟调用是在 defer 所在函数结束时进行，==函数结束可以是正常返回时，也可以是发生宕机时==。）

当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出），下面的代码是将一系列的数值打印语句按顺序延迟处理，如下所示：

```go
    package main
    import (
        "fmt"
    )
    func main() {
        fmt.Println("defer begin")
        // 将defer放入延迟调用栈
        defer fmt.Println(1)
        defer fmt.Println(2)
        // 最后一个放入, 位于栈顶, 最先调用
        defer fmt.Println(3)
        fmt.Println("defer end")
    }
```

代码输出如下：

```go
defer begin
defer end
3
2
1
```

### 七、处理运行时的错误

Go语言的错误处理思想及设计包含以下特征：

- 一个可能造成错误的函数，需要返回值中返回一个错误接口（error），如果调用是成功的，错误接口将返回 nil，否则返回错误。
- 在函数调用后需要检查错误，如果发生错误，则进行必要的错误处理。

##### 1.错误接口的定义格式

error 是 ==Go 系统==声明的接口类型，代码如下：

```go
    type error interface {
        Error() string
    }
```

所有符合 Error()string 格式的方法，都能实现错误接口，Error() 方法返回错误的具体描述，使用者可以通过这个字符串知道发生了什么错误。

##### 2.自定义一个错误

返回错误前，需要定义会产生哪些可能的错误，在Go语言中，使用 errors 包进行错误的定义，格式如下：`

```go
var err = errors.New("this is an error")   //err为实现了error类型接口的变量
```

Go语言的 errors包中对 New 的定义非常简单，代码如下：

```go
    // 创建错误对象
	func New(text string) error {   //errors.New()方法
        return &errorString{text}
    }
    // 错误字符串
    type errorString struct {
        s string
    }
    // 返回发生何种错误
    func (e *errorString) Error() string {
        return e.s
    }
```

##### 3.示例：在解析中使用自定义错误

使用 errors.New 定义的错误字符串的错误类型是无法提供丰富的错误信息的，那么，如果需要携带错误信息返回，就需要==借助自定义结构体实现错误接口==。

下面代码将实现一个解析错误（ParseError），这种错误包含两个内容，分别是文件名和行号，解析错误的结构还实现了 error 接口的 Error() 方法，返回错误描述时，就需要将文件名和行号返回。

```go
package main
import (
    "fmt"
)
// 声明一个解析错误
type ParseError struct {
    Filename string // 文件名
    Line     int    // 行号
}
// 实现error接口，返回错误描述
func (e *ParseError) Error() string {
    return fmt.Sprintf("%s:%d", e.Filename, e.Line)
}
// 创建一些解析错误
func newParseError(filename string, line int) error {
    return &ParseError{filename, line}
}
func main() {
    var e error
    // 创建一个错误实例，包含文件名和行号
    e = newParseError("main.go", 1)
    // 通过error接口查看错误描述
    fmt.Println(e.Error())
    // 根据错误接口具体的类型，进一步获取详细错误信息
    switch detail := e.(type) {
    case *ParseError: // 这是一个解析错误
        fmt.Printf("Filename: %s Line: %d\n", detail.Filename, detail.Line)
    default: // 其他类型的错误
        fmt.Println("other error")
    }
```

代码输出如下：

```go
main.go:1
Filename: main.go Line: 1
```

错误对象都要实现 error 接口的 Error() 方法，这样，所有的错误都可以获得字符串的描述，如果想进一步知道错误的详细信息，可以通过类型断言，将错误对象转为具体的错误类型进行错误详细信息的获取。

### 八、Golang中的宕机 panic

Go语言的类型系统会在编译时捕获很多错误，但有些错误只能在运行时检查，如数组访问越界、空指针引用等，这些运行时错误会引起宕机panic。一般而言，当宕机发生时，程序会中断运行，并立即执行在该 goroutine（可以先理解成线程）中被延迟的函数（defer  机制），随后，程序崩溃并输出日志信息，日志信息包括 panic value 和函数调用的堆栈跟踪信息，panic value  通常是某种错误信息。

虽然Go语言的 panic 机制类似于其他语言的异常，但 panic 的适用场景有一些不同，由于 panic 会引起程序的崩溃，因此 panic  一般用于严重错误，如程序内部的逻辑不一致。任何崩溃都表明了我们的代码中可能存在漏洞，所以对于大部分漏洞，我们应该使用Go语言提供的错误机制，而不是 panic。

##### 1.手动触发宕机

Go语言程序在宕机时，会将堆栈和 goroutine 信息输出到控制台，所以宕机也可以方便地知晓发生错误的位置，那么我们要如何触发宕机呢，示例代码如下所示：

```go
    package main
    func main() {
        panic("crash")
    }
```

代码运行崩溃并输出如下：

```go
panic: crash

goroutine 1 [running]:
main.main()
    D:/code/main.go:4 +0x40
exit status 2
```

##### 2.在宕机时触发延迟执行语句

当 panic() 触发的宕机发生时，panic() 后面的代码将不会被运行，但是在 panic() 函数前面已经运行过的 defer 语句依然会在宕机发生时发生作用，参考下面代码：

```go
    package main
    import "fmt"
    func main() {
        defer fmt.Println("宕机后要做的事情1")
        defer fmt.Println("宕机后要做的事情2")
        panic("宕机")
    }
```

代码输出如下：

```go
宕机后要做的事情2
宕机后要做的事情1
panic: 宕机

goroutine 1 [running]:
main.main()
    D:/code/main.go:8 +0xf8
exit status 2
```

==宕机前，defer 语句会被优先执行==，由于第 7 行的 defer 后执行，因此会在宕机前，这个 defer 会优先处理，随后才是第 6 行的 defer 对应的语句，这个特性可以用来在宕机发生前进行宕机信息处理。

### 九、宕机恢复（recover）——防止程序崩溃

Recover 是一个Go语言的内建函数，可以让进入宕机流程中的 goroutine 恢复过来，==recover 仅在延迟函数 defer  中有效==，在正常的执行过程中，调用 recover 会返回 nil 并且没有其他任何效果，==如果当前的 goroutine 陷入恐慌，调用  recover 可以捕获到 panic 的输入值，并且恢复正常的执行==。

通常来说，不应该对进入 panic 宕机的程序做任何处理，但有时，需要我们可以从宕机中恢复，至少我们可以在程序崩溃前，做一些操作，举个例子，当  web 服务器遇到不可预料的严重问题时，在崩溃前应该将所有的连接关闭，如果不做任何处理，会使得客户端一直处于等待状态，如果 web  服务器还在开发阶段，服务器甚至可以将异常信息反馈到客户端，帮助调试。

##### 1.让程序在崩溃时继续执行

下面的代码实现了 ProtectRun() 函数，该函数传入一个匿名函数或闭包后的执行函数，当传入函数以任何形式发生 panic 崩溃后，可以将崩溃发生的错误打印出来，同时允许后面的代码继续运行，不会造成整个进程的崩溃。

保护运行函数：

```go	
    package main
    import (
        "fmt"
        "runtime"
    )
    // 崩溃时需要传递的上下文信息
    type panicContext struct {
        function string // 所在函数
    }
    // 保护方式允许一个函数
    func ProtectRun(entry func()) {
        // 延迟处理的函数
        defer func() {
            // 发生宕机时，获取panic传递的上下文并打印
            err := recover()
            switch err.(type) {
            case runtime.Error: // 运行时错误
                fmt.Println("runtime error:", err)
            default: // 非运行时错误
                fmt.Println("error:", err)
            }
        }()
        entry()
    }
    func main() {
        fmt.Println("运行前")
        // 允许一段手动触发的错误
        ProtectRun(func() {
            fmt.Println("手动宕机前")
            // 使用panic传递上下文
            panic(&panicContext{
                "手动触发panic",
            })
            fmt.Println("手动宕机后")
        })
        // 故意造成空指针访问错误
        ProtectRun(func() {
            fmt.Println("赋值宕机前")
            var a *int
            *a = 1
            fmt.Println("赋值宕机后")
        })
        fmt.Println("运行后")
    }
```

代码输出结果：

```go
运行前
手动宕机前
error: &{手动触发panic}
赋值宕机前
runtime error: runtime error: invalid memory address or nil pointer dereference
运行后
```

对代码的说明：

- 第 9 行声明描述错误的结构体，保存执行错误的函数。
- 第 17 行使用 defer 将闭包延迟执行，当 panic 触发崩溃时，ProtectRun() 函数将结束运行，此时 defer 后的闭包将会发生调用。
- 第 20 行，recover() 获取到 panic 传入的参数。
- 第 22 行，使用 switch 对 err 变量进行类型断言。
- 第 23 行，如果错误是有 Runtime 层抛出的运行时错误，如空指针访问、除数为 0 等情况，打印运行时错误。
- 第 25 行，其他错误，打印传递过来的错误数据。
- 第 44 行，使用 panic 手动触发一个错误，并将一个结构体附带信息传递过去，此时，recover 就会获取到这个结构体信息，并打印出来。
- 第 57 行，模拟代码中空指针赋值造成的错误，此时会由 Runtime 层抛出错误，被 ProtectRun() 函数的 recover() 函数捕获到。

##### 2.panic 和 recover 的关系

panic 和 recover 的组合有如下特性：

- 有 panic 没 recover，程序宕机。
- 有 panic 也有 recover，程序不会宕机，执行完对应的 defer 后，从宕机点退出当前函数后继续执行。

#### 提示

虽然 panic/recover 能模拟其他语言的异常机制，但并不建议在编写普通函数时也经常性使用这种特性。

 ==在 panic 触发的 defer 函数内，可以继续调用 panic，进一步将错误外抛，直到程序整体崩溃。==

### 十、计算函数执行时间

在Go语言中我们可以使用 time 包中的 Since() 函数来获取函数的运行时间，Go语言官方文档中对 Since() 函数的介绍是这样的。

```go	
func Since(t Time) Duration
```

Since() 函数返回从 t 到现在经过的时间，等价于`time.Now().Sub(t)`。

【示例】使用 Since() 函数获取函数的运行时间。

```go
    package main
    import (
        "fmt"
        "time"
    )
    func test() {
        start := time.Now() // 获取当前时间
        sum := 0
        for i := 0; i < 100000000; i++ {
            sum++
        }
        elapsed := time.Since(start)
        fmt.Println("该函数执行完成耗时：", elapsed)
    }
    func main() {
        test()
    }
```

运行结果如下所示：

```go
该函数执行完成耗时： 39.8933ms
```

上面我们提到了 time.Now().Sub() 的功能类似于 Since() 函数，想要使用 time.Now().Sub() 获取函数的运行时间只需要把我们上面代码的第 14 行简单修改一下就行。

【示例 2】使用 time.Now().Sub() 获取函数的运行时间。

``` go
    package main
    import (
        "fmt"
        "time"
    )
    func test() {
        start := time.Now() // 获取当前时间
        sum := 0
        for i := 0; i < 100000000; i++ {
            sum++
        }
        elapsed := time.Now().Sub(start)
        fmt.Println("该函数执行完成耗时：", elapsed)
    }
    func main() {
        test()
    }
```

运行结果如下所示：

```go
该函数执行完成耗时： 36.8769ms
```

### 十一、Test功能测试函数

#### 1.测试规则

要开始一个单元测试，需要准备一个 go 源码文件，在命名文件时文件名必须以`_test.go`结尾，单元测试源码文件可以由多个测试用例（可以理解为函数）组成，每个测试用例的名称需要以 Test 为前缀，例如：

```go
func TestXxx( t *testing.T ){
    //......
}
```

编写测试用例有以下几点需要注意：

- 测试用例文件不会参与正常源码的编译，不会被包含到可执行文件中；
- 测试用例的文件名必须以`_test.go`结尾；
- 需要使用 import 导入 testing 包；
- 测试函数的名称要以`Test`或`Benchmark`开头，后面可以跟任意字母组成的字符串，但第一个字母必须大写，例如 TestAbc()，一个测试用例文件中可以包含多个测试函数；
- 单元测试则以`(t *testing.T)`作为参数，性能测试以`(t *testing.B)`做为参数；
- 测试用例文件使用` go test `命令来执行，源码中不需要 main() 函数作为入口，所有以`_test.go`结尾的源码文件内以`Test`开头的函数都会自动执行。

Go语言的 testing 包提供了三种测试方式，分别是==单元（功能）测试==、==性能（压力）测试==和==覆盖率测试==。

#### 2.单元（功能）测试

#### demo.go：

```go
    package demo
    // 根据长宽获取面积
    func GetArea(weight int, height int) int {
        return weight * height
    }
```

#### demo_test.go：

```go
    package demo
    import "testing"
    func TestGetArea(t *testing.T) {
        area := GetArea(40, 50)
        if area != 2000 {
            t.Error("测试失败")
        }
    }
```

执行测试命令，运行结果如下所示：

```go
PS D:\code> go test -v   // -v 指令
=== RUN   TestGetArea
--- PASS: TestGetArea (0.00s)
PASS
ok      _/D_/code       0.435s
```

#### 3.性能（压力）测试

#### demo_test.go：

```go
    package demo
    import "testing"
    func BenchmarkGetArea(t *testing.B) {    //测试函数名前缀为Benchmark
        for i := 0; i < t.N; i++ {
            GetArea(40, 50)    //反复执行被测试函数t.N次
        }
    }
```

执行测试命令，运行结果如下所示：

```go
PS D:\code> go test -bench="."  // -bench="."  指令
goos: windows
goarch: amd64
BenchmarkGetArea-4      2000000000               0.35 ns/op
PASS
ok      _/D_/code       1.166s
```

上面信息显示了程序执行 2000000000 次，平均每次耗时 0.35 纳秒

#### 4.覆盖率测试

覆盖率测试能知道测试程序总共覆盖了多少业务代码（也就是 demo_test.go 中测试了多少 demo.go 中的代码），可以的话最好是覆盖100%。

#### demo_test.go：

```go
    package demo
    import "testing"
    func TestGetArea(t *testing.T) {
        area := GetArea(40, 50)
        if area != 2000 {
            t.Error("测试失败")
        }
    }
    func BenchmarkGetArea(t *testing.B) {
        for i := 0; i < t.N; i++ {
            GetArea(40, 50)
        }
    }
```

执行测试命令，运行结果如下所示：

```go
PS D:\code> go test -cover   //  -cover指令
PASS
coverage: 100.0% of statements
ok      _/D_/code       0.437s
```

#### 5.示例测试

源代码文件`example.go`中包含`SayHello()`、`SayGoodbye()`和`PrintNames()`三个方法，如下所示：

```go
package gotest

import "fmt"

// SayHello 打印一行字符串
func SayHello() {
    fmt.Println("Hello World")
}

// SayGoodbye 打印两行字符串
func SayGoodbye() {
    fmt.Println("Hello,")
    fmt.Println("goodbye")
}

// PrintNames 打印学生姓名
func PrintNames() {
    students := make(map[int]string, 4)
    students[1] = "Jim"
    students[2] = "Bob"
    students[3] = "Tom"
    students[4] = "Sue"
    for _, value := range students {
        fmt.Println(value)
    }
}
```

测试文件`example_test.go`中包含3个测试方法，于源代码文件中的3个方法一一对应，测试文件如下所示：

```go
package gotest_test

import "gotest"

// 检测单行输出
func ExampleSayHello() {     //命名规则为ExampleXxx
    gotest.SayHello()
    // OutPut: Hello World 
}

// 检测多行输出
func ExampleSayGoodbye() {
    gotest.SayGoodbye()
    // OutPut:
    // Hello,
    // goodbye
}

// 检测乱序输出
func ExamplePrintNames() {
    gotest.PrintNames()
    // Unordered output:
    // Jim
    // Bob
    // Tom
    // Sue
}
```

例子测试函数命名规则为”ExampleXxx”，其中”Xxx”为自定义的标识，通常为待测函数名称。

这三个测试函数分别代表三种场景：(==每个例子测试函数中的注释都是必要的,是ExampleXxx函数用来对比函数输出值的期望值==)

- ExampleSayHello()： 待测试函数只有一行输出，使用”// OutPut: “检测。
- ExampleSayGoodbye()：待测试函数有多行输出，使用”// OutPut: “检测，其中期望值也是多行。
- ExamplePrintNames()：待测试函数有多行输出，但输出次序不确定，使用”// Unordered output:”检测。

注：字符串比较时会忽略前后的空白字符。

命令行下，使用`go test`或`go test example_test.go`命令即可启动测试，运行结果如下所示：

```go
有正确注释：(执行ExampleSayHello())
=== RUN   ExampleSayHello
--- PASS: ExampleSayHello (0.00s)
PASS

无正确注释:(执行ExampleSayHello())
=== RUN   ExampleSayHello
--- FAIL: ExampleSayHello (0.00s)
got:
Hello World
want:
Hello World！！！

FAIL
```

#### 6.子测试

简单的说，子测试提供一种在一个测试函数中执行多个测试的能力，比如原来有TestA、TestB和TestC三个测试函数，每个测试函数执行开始都需要做些相同的初始化工作，那么可以利用子测试将这三个测试合并到一个测试中，这样初始化工作只需要做一次。

##### 6.1 简单例子：

```go
package gotest_test

import (
    "testing"
    "gotest"
)

// sub1 为子测试，只做加法测试
func sub1(t *testing.T) {
    var a = 1
    var b = 2
    var expected = 3

    actual := gotest.Add(a, b)
    if actual != expected {
        t.Errorf("Add(%d, %d) = %d; expected: %d", a, b, actual, expected)
    }
}

// sub2 为子测试，只做加法测试
func sub2(t *testing.T) {
    var a = 1
    var b = 2
    var expected = 3

    actual := gotest.Add(a, b)
    if actual != expected {
        t.Errorf("Add(%d, %d) = %d; expected: %d", a, b, actual, expected)
    }
}

// sub3 为子测试，只做加法测试
func sub3(t *testing.T) {
    var a = 1
    var b = 2
    var expected = 3

    actual := gotest.Add(a, b)
    if actual != expected {
        t.Errorf("Add(%d, %d) = %d; expected: %d", a, b, actual, expected)
    }
}

// TestSub 内部调用sub1、sub2和sub3三个子测试
func TestSub(t *testing.T) {
    // setup code

    t.Run("A=1", sub1)
    t.Run("A=2", sub2)
    t.Run("B=1", sub3)

    // tear-down code
}
```

本例中`TestSub()`通过`t.Run()`依次执行三个子测试。`t.Run()`函数声明如下：

```go
func (t *T) Run(name string, f func(t *T)) bool
```

`name`参数为子测试的名字，`f`为子测试函数，本例中`Run()`一直阻塞到`f`执行结束后才返回，返回值为f的执行结果。
`Run()`会启动新的协程来执行`f`，并阻塞等待`f`执行结束才返回，除非`f`中使用`t.Parallel()`设置子测试为并发。

本例中`TestSub()`把三个子测试合并起来，可以共享setup和tear-down部分的代码。

我们在命令行下，使用`-v`参数执行测试：

```go
E:\OpenSource\GitHub\RainbowMango\GoExpertProgrammingSourceCode\GoExpert\src\gotest>go test subunit_test.go -v
=== RUN   TestSub
=== RUN   TestSub/A=1
=== RUN   TestSub/A=2
=== RUN   TestSub/B=1
--- PASS: TestSub (0.00s)
    --- PASS: TestSub/A=1 (0.00s)
    --- PASS: TestSub/A=2 (0.00s)
    --- PASS: TestSub/B=1 (0.00s)
PASS
ok      command-line-arguments  0.354s
```

从输出中可以看出，三个子测试都被执行到了，而且执行次序与调用次序一致。

##### 6.2  子测试命名规则

通过上面的例子我们知道`Run()`方法第一个参数为子测试的名字，而==实际上子测试的内部命名规则为：”*<父测试名字>*/*<传递给Run的名字>*”==。比如，==传递给`Run()`的名字是“A=1”，那么子测试名字为“TestSub/A=1”==。这个在上面的命令行输出中也可以看出。

##### 6.3  过滤筛选

通过测试的名字，可以在执行中过滤掉一部分测试。

比如，只执行上例中“A=*”的子测试，那么执行时使用`-run Sub/A=`参数即可：

```go
E:\OpenSource\GitHub\RainbowMango\GoExpertProgrammingSourceCode\GoExpert\src\gotest>go test subunit_test.go -v -run Sub/A=
=== RUN   TestSub
=== RUN   TestSub/A=1
=== RUN   TestSub/A=2
--- PASS: TestSub (0.00s)
    --- PASS: TestSub/A=1 (0.00s)
    --- PASS: TestSub/A=2 (0.00s)
PASS
ok      command-line-arguments  0.340s
```

上例中，使用参数`-run Sub/A=`则只会执行`TestSub/A=1`和`TestSub/A=2`两个子测试。

对于子性能测试则使用`-bench`参数来筛选，此处不再赘述。

注意：此处的筛选不是严格的正则匹配，而是包含匹配（字符串子串包含匹配）。比如，`-run A=`那么所有测试（含子测试）的名字中如果包含“A=”则会被选中执行。

##### 6.3  子测试并发

前面提到的多个子测试共享setup和teardown有一个前提是子测试没有并发，如果子测试使用`t.Parallel()`指定并发，那么就没办法共享teardown了，因为执行顺序很可能是setup->子测试1->teardown->子测试2…。

==如果子测试可能并发，则可以把子测试通过`Run()`再嵌套一层，`Run()`可以保证其下的所有子测试执行结束后再返回。==

为便于说明，我们创建文件`subparallel_test.go`用于说明：

```go
package gotest_test

import (
    "testing"
    "time"
)

// 并发子测试，无实际测试工作，仅用于演示
func parallelTest1(t *testing.T) {
    t.Parallel()               //启用并发
    time.Sleep(3 * time.Second)
    // do some testing
}

// 并发子测试，无实际测试工作，仅用于演示
func parallelTest2(t *testing.T) {
    t.Parallel()               //启用并发
    time.Sleep(2 * time.Second)
    // do some testing
}

// 并发子测试，无实际测试工作，仅用于演示
func parallelTest3(t *testing.T) {
    t.Parallel()               //启用并发
    time.Sleep(1 * time.Second)
    // do some testing
}

// TestSubParallel 通过把多个子测试放到一个组中并发执行，同时多个子测试可以共享setup和tear-down
func TestSubParallel(t *testing.T) {
    // setup
    t.Logf("Setup")

    t.Run("group", func(t *testing.T) {
        t.Run("Test1", parallelTest1)
        t.Run("Test2", parallelTest2)
        t.Run("Test3", parallelTest3)
    })

    // tear down
    t.Logf("teardown")
}
```

上面三个子测试中分别sleep了3s、2s、1s用于观察并发执行顺序。通过`Run()`将多个子测试“封装”到一个组中，可以保证所有子测试全部执行结束后再执行tear-down。

命令行下的输出如下：

```go
E:\OpenSource\GitHub\RainbowMango\GoExpertProgrammingSourceCode\GoExpert\src\gotest>go test subparallel_test.go -v -run SubParallel
=== RUN   TestSubParallel
=== RUN   TestSubParallel/group
=== RUN   TestSubParallel/group/Test1
=== RUN   TestSubParallel/group/Test2
=== RUN   TestSubParallel/group/Test3
--- PASS: TestSubParallel (3.01s)
        subparallel_test.go:25: Setup
    --- PASS: TestSubParallel/group (0.00s)
        --- PASS: TestSubParallel/group/Test3 (1.00s)
        --- PASS: TestSubParallel/group/Test2 (2.01s)
        --- PASS: TestSubParallel/group/Test1 (3.01s)
        subparallel_test.go:34: teardown
PASS
ok      command-line-arguments  3.353s
```

通过该输出可以看出：

1. 子测试是并发执行的（Test1最先被执行却最后结束）
2. tear-down在所有子测试结束后才执行

##### 6.4  Main测试

我们知道子测试的一个方便之处在于可以让多个测试共享Setup和Tear-down。但这种程度的共享有时并不满足需求，有时希望在整个测试程序做一些全局的setup和Tear-down，这时就需要Main测试了。

所谓Main测试，即==声明一个`func TestMain(m *testing.M)`==，它是名字比较特殊的测试，==参数类型为`testing.M`指针==。如果声明了这样一个函数，==当前测试程序将不是直接执行各项测试，而是将测试交给TestMain调度==。

```go
// TestMain 用于主动执行各种测试，可以测试前后做setup和tear-down操作
func TestMain(m *testing.M) {
    println("TestMain setup.")

    retCode := m.Run() // 执行测试，包括单元测试、性能测试和示例测试

    println("TestMain tear-down.")

    os.Exit(retCode)
}
```

 TestMain测试函数无法手动调用，而是会在运行其他测试函数时自动执行。
