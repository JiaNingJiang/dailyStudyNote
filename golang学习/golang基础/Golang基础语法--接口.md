### 一、多个类型可以实现共同的接口

一个接口的方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现。也就是说，使用者并不关心某个接口的方法是通过一个类型完全实现的，还是通过多个结构嵌入到一个结构体中拼凑起来共同实现的。

Service 接口定义了两个方法：一个是开启服务的方法（Start()），一个是输出日志的方法（Log()）。使用 GameService  结构体来实现 Service，GameService 自己的结构只能实现 Start() 方法，而 Service 接口中的 Log()  方法已经被一个能输出日志的日志器（Logger）实现了，无须再进行 GameService 封装，或者重新实现一遍。所以，选择将 Logger  嵌入到 GameService 能最大程度地避免代码冗余，简化代码结构。详细实现过程如下：

```go
    // 一个服务需要满足能够开启和写日志的功能
    type Service interface {
        Start()  // 开启服务
        Log(string)  // 日志输出
    }
    // 日志器
    type Logger struct {
    }
    // 实现Service的Log()方法
    func (g *Logger) Log(l string) {
    }
    // 游戏服务
    type GameService struct {
        Logger  // 嵌入日志器
    }
    // 实现Service的Start()方法
    func (g *GameService) Start() {
    }
```

此时，实例化 GameService，并将实例赋给 Service，代码如下：

```go
    var s Service = new(GameService)
    s.Start()
    s.Log(“hello”)
```

s 就可以使用 Start() 方法和 Log() 方法，其中，Start() 由 GameService 实现，Log() 方法由 Logger 实现。

### 二、类型断言

类型断言（Type Assertion）是一个使用在接口值上的操作，用于检查接口类型变量所持有的值是否实现了期望的接口或者具体的类型。

在Go语言中类型断言的语法格式如下：

```go
value, ok := x.(T)
```

其中，x 表示一个接口的类型，T 表示一个具体的类型（也可为接口类型）。

该断言表达式会返回 x 的值（也就是 value）和一个布尔值（也就是 ok），可根据该布尔值判断 x 是否为 T 类型：

- 如果 T 是具体某个类型，类型断言会检查 x 的动态类型是否等于具体类型 T。如果检查成功，类型断言返回的结果是 x 的动态值，其类型是 T。
- 如果 T 是接口类型，类型断言会检查 x 的动态类型是否满足 T。如果检查成功，x 的动态值不会被提取，返回值是一个类型为 T 的接口值。
- 无论 T 是什么类型，如果 x 是 nil 接口值，类型断言都会失败。

```go
    package main
    import (
        "fmt"
    )
    func main() {
        var x interface{}
        x = 10
        value, ok := x.(int)
        fmt.Print(value, ",", ok)
    }
```

运行结果如下：

```go
10,true
```

需要注意如果不接收第二个参数也就是上面代码中的 ok，断言失败时会直接造成一个 panic。如果 x 为 nil 同样也会 panic。

```go
    package main
    import (
        "fmt"
    )
    func main() {
        var x interface{}
        x = "Hello"
        value := x.(int)
        fmt.Println(value)
    }
```

运行结果如下：

```go
panic: interface conversion: interface {} is string, not int
```

### 三、排序（借助sort.Interface接口）

Go语言的 sort.Sort 函数不会对具体的序列和它的元素做任何假设。相反，它使用了一个接口类型 sort.Interface  来指定通用的排序算法和可能被排序到的序列类型之间的约定。这个接口的实现由序列的具体表示和它希望排序的元素决定，序列的表示经常是一个切片。

一个内置的排序算法需要知道三个东西：序列的长度，表示两个元素比较的结果，一种交换两个元素的方式；这就是 sort.Interface 的三个方法：

```go
    package sort
    type Interface interface {
        Len() int            // 获取元素数量
        Less(i, j int) bool // i，j是序列元素的指数。
        Swap(i, j int)        // 交换元素
    }
```

为了对序列进行排序，我们需要定义一个实现了这三个方法的类型，然后对这个类型的一个实例应用 sort.Sort  函数。思考对一个字符串切片进行排序，这可能是最简单的例子了。下面是这个新的类型 MyStringList 和它的 Len，Less 和  Swap 方法

```go
    type MyStringList  []string
    func (m MyStringList ) Len() int { return len(m) }
    func (m MyStringList ) Less(i, j int) bool { return m[i] < m[j] }
    func (m MyStringList ) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
```

只要对象实现了 sort.Interface 这个接口，就可以使用sort.Sort  函数完成排序

#### 常见类型的便捷排序

通过实现 sort.Interface  接口的排序过程具有很强的可定制性，可以根据被排序对象比较复杂的特性进行定制。例如，需要多种排序逻辑的需求就适合使用 sort.Interface 接口进行排序。但大部分情况中，只需要对字符串、整型等进行快速排序。==Go语言中提供了一些固定模式的封装以方便开发者迅速对内容进行排序。==

##### 1) 字符串切片的便捷排序

sort 包中有一个 StringSlice 类型，其源码实现如下：

```go
    type StringSlice []string
    func (p StringSlice) Len() int           { return len(p) }
    func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
    func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
    // Sort is a convenience method.
    func (p StringSlice) Sort() { Sort(p) }
```

sort 包中的 StringSlice 的代码与 MyStringList 的实现代码几乎一样。因此，==只需要使用 sort 包的 StringSlice 就可以更简单快速地进行字符串排序==。将代码1中的排序代码简化后如下所示：

```go
    names := sort.StringSlice{     //sort.StringSlice 类型
        "3. Triple Kill",
        "5. Penta Kill",
        "2. Double Kill",
        "4. Quadra Kill",
        "1. First Blood",
    }
    sort.Sort(names)
```

##### 2) 对整型切片进行排序

还可以使用 sort.IntSlice 进行整型切片的排序。sort.IntSlice 的定义如下：

```go
    type IntSlice []int
    func (p IntSlice) Len() int           { return len(p) }
    func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
    func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
    // Sort is a convenience method.
    func (p IntSlice) Sort() { Sort(p) }
```

sort 包在 sort.Interface 对各类型的封装上还有更进一步的简化，下面使用 sort.Strings 继续对代码1进行简化，代码如下：

```go
    names := []string{
        "3. Triple Kill",
        "5. Penta Kill",
        "2. Double Kill",
        "4. Quadra Kill",
        "1. First Blood",
    }
    sort.Strings(names)     //sort.Strings方法,完成对字符串数组的排序
    // 遍历打印结果
    for _, v := range names {
        fmt.Printf("%s\n", v)
    }
```

##### 3) sort包内建的类型排序接口一览

Go语言中的 sort 包中定义了一些常见类型的排序方法，如下表所示。

| 类  型                | 实现 sort.lnterface 的类型 | 直接排序方法               | 说  明            |
| --------------------- | -------------------------- | -------------------------- | ----------------- |
| 字符串（String）      | StringSlice                | sort.Strings(a [] string)  | 字符 ASCII 值升序 |
| 整型（int）           | IntSlice                   | sort.Ints(a []int)         | 数值升序          |
| 双精度浮点（float64） | Float64Slice               | sort.Float64s(a []float64) | 数值升序          |

编程中经常用到的 int32、int64、float32、bool 类型并没有由 sort 包实现，使用时依然需要开发者自己编写。

#### 对结构体数据进行排序

从 Go 1.8 开始，Go语言在 sort 包中提供了 sort.Slice() 函数进行更为简便的排序方法。sort.Slice()  函数只要求传入需要排序的数据，以及一个排序时对元素的回调函数，类型为 func(i,j int)bool，sort.Slice()  函数的定义如下：

```go
func Slice(slice interface{}, less func(i, j int) bool)
```

```go
    package main
    import (
        "fmt"
        "sort"
    )
    type HeroKind int
    const (
        None = iota
        Tank
        Assassin
        Mage
    )
    type Hero struct {
        Name string
        Kind HeroKind
    }
    func main() {
        heros := []*Hero{
            {"吕布", Tank},
            {"李白", Assassin},
            {"妲己", Mage},
            {"貂蝉", Assassin},
            {"关羽", Tank},
            {"诸葛亮", Mage},
        }
        sort.Slice(heros, func(i, j int) bool {    //完成对结构体的排序
            if heros[i].Kind != heros[j].Kind {
                return heros[i].Kind < heros[j].Kind   // 如果英雄的分类不一致时, 优先对分类进行排序
            }
            return heros[i].Name < heros[j].Name
        })
        for _, v := range heros {
            fmt.Printf("%+v\n", v)
        }
    }
```

使用 sort.Slice() 不仅可以完成结构体切片排序，还可以对各种切片类型进行自定义排序。

### 四、接口和类型之间的转换

#### 1.类型断言的格式

类型断言的基本格式如下：

```go
t := i.(T)
```

这里有两种可能。第一种，如果断言的类型 T 是一个具体类型(实体类)，然后类型断言检查 i 的动态类型(接口)是否和 T  相同。如果这个检查成功了，类型断言的结果是 i 的动态值，当然它的类型是  T。换句话说，具体类型的类型断言从它的操作对象中获得具体的值。如果检查失败，接下来这个操作会抛出 panic。例如：

```go
    var w io.Writer
    w = os.Stdout
    f := w.(*os.File) // 成功: f == os.Stdout
    c := w.(*bytes.Buffer) // 死机：接口保存*os.file(os.Stdout实体类对应的接口)，而不是*bytes.buffer
```

第二种，如果相反断言的类型 T 是一个接口类型，然后类型断言检查是否 i 的动态类型(对应接口)满足  T。如果这个检查成功了，动态值没有获取到；这个结果仍然是一个有相同类型和值部分的接口值，但是结果有类型  T。换句话说，对一个接口类型的类型断言改变了类型的表述方式，改变了可以获取的方法集合（通常更大），但是它保护了接口值内部的动态类型和值的部分。

在下面的第一个类型断言后，w 和 rw 都持有 os.Stdout 因此它们每个有一个动态类型 *os.File，但是变量 w 是一个 io.Writer 类型只对外公开出文件的 Write 方法，然而 rw 变量也只公开它的 Read 方法。

```go
    var w io.Writer
    w = os.Stdout
    rw := w.(io.ReadWriter) // 成功：*os.file具有读写功能
    w = new(ByteCounter)
    rw = w.(io.ReadWriter) // 死机：*字节计数器没有读取方法
```

如果断言操作的对象是一个 nil 接口值，那么不论被断言的类型是什么这个类型断言都会失败。几乎不需要对一个更少限制性的接口类型（更少的方法集合）做断言，因为它表现的就像赋值操作一样，除了对于 nil 接口值的情况。



如果 i 没有完全实现 T 接口的方法，这个语句将会触发宕机。触发宕机不是很友好，因此上面的语句还有一种写法：

```go
t,ok := i.(T)
```

这种写法下，如果接口未实现，将会把 ok 置为 false，t 置为 T 类型的 0 值。正常实现时，ok 为 true。这里 ok 可以被认为是：i 接口是否实现 T 类型的结果。

#### 2.将接口转换为其他类型

可以实现将接口转换为普通的指针类型。例如将 Walker 接口转换为 *pig 类型，请参考下面的代码：

```go
package main
import "fmt"
// 定义飞行动物接口
type Flyer interface {
    Fly()
}
// 定义行走动物接口
type Walker interface {
    Walk()
}
// 定义鸟类
type bird struct {
}
// 实现飞行动物接口
func (b *bird) Fly() {
    fmt.Println("bird: fly")
}
// 为鸟添加Walk()方法, 实现行走动物接口
func (b *bird) Walk() {
    fmt.Println("bird: walk")
}
// 定义猪
type pig struct {
}
// 为猪添加Walk()方法, 实现行走动物接口
func (p *pig) Walk() {
    fmt.Println("pig: walk")
}

func main() {
    p1 := new(pig)
    var a Walker = p1
    p2 := a.(*pig)
    fmt.Printf("p1=%p p2=%p", p1, p2)
}
```

对代码的说明如下：

- 第 3 行，由于 pig 类的对象 p1 实现了 Walker 接口，因此可以被隐式转换为 Walker 接口类型保存于 a 中。
- 第 4 行，由于 a 中保存的本来就是 *pig 本体，因此可以转换为 *pig 类型。
- 第 6 行，对比发现，p1 和 p2 指针是相同的。

如果尝试将上面这段代码中的 Walker 类型的 a 转换为 *bird 类型，将会发出运行时错误，请参考下面的代码：

```go
    p1 := new(pig)
    var a Walker = p1
    p2 := a.(*bird)
```

运行时报错：

```go
panic: interface conversion: main.Walker is *main.pig, not *main.bird
```

报错意思是：接口转换时，main.Walker 接口的内部保存的是 *main.pig，而不是 *main.bird。

 因此，接口在转换为其他类型时，接口内保存的实例对应的类型指针，必须是要转换的对应的类型指针。

### 五、空接口类型（interface{}）

空接口的内部实现保存了对象的类型和指针。使用空接口保存一个数据的过程会比直接用数据对应类型的变量保存稍慢。因此在开发中，应在需要的地方使用空接口，而不是在所有地方使用空接口。

#### 1.将值保存到空接口

```go
    var any interface{}
    any = 1
    fmt.Println(any)
    any = "hello"
    fmt.Println(any)
    any = false
    fmt.Println(any)
```

代码输出如下：

```go
1
hello
false
```

#### 2.从空接口获取值

保存到空接口的值，如果直接取出指定类型的值时，会发生编译错误，代码如下：

```go
    // 声明a变量, 类型int, 初始值为1
    var a int = 1
    // 声明i变量, 类型为interface{}, 初始值为a, 此时i的值变为1
    var i interface{} = a
    // 声明b变量, 尝试赋值i
    var b int = i
```

第8行代码编译报错：

```go
cannot use i (type interface {}) as type int in assignment: need type assertion
```

编译器告诉我们，不能将i变量视为int类型赋值给b。

在代码第 5 行中，将 a 的值赋值给 i 时，虽然 i 在赋值完成后的内部值为 int，但 i 还是一个 interface{} 类型的变量。类似于无论集装箱装的是茶叶还是烟草，集装箱依然是金属做的，不会因为所装物的类型改变而改变。

 为了让第 8 行的操作能够完成，编译器提示我们得使用 type assertion，意思就是类型断言。

使用类型断言修改第 8 行代码如下：

```go
var b int = i.(int)
```

修改后，代码可以编译通过，并且 b 可以获得 i 变量保存的 a 变量的值：1。

#### 3.空接口的值比较

空接口在保存不同的值后，可以和其他变量值一样使用`==`进行比较操作。空接口的比较有以下几种特性。

##### 1) 类型不同的空接口间的比较结果不相同

保存有类型不同的值的空接口进行比较时，Go语言会优先比较值的类型。因此类型不同，比较结果也是不相同的，代码如下：

```go
    // a保存整型
    var a interface{} = 100
    // b保存字符串
    var b interface{} = "hi"
    // 两个空接口不相等
    fmt.Println(a == b)
```

代码输出如下：

```go
false
```

##### 2) 不能比较空接口中的动态值

当接口中保存有动态类型的值时，运行时将触发错误，代码如下：

```go
    // c保存包含10的整型切片
    var c interface{} = []int{10}
    // d保存包含20的整型切片
    var d interface{} = []int{20}
    // 这里会发生崩溃
    fmt.Println(c == d)
```

代码运行到第8行时发生崩溃：

```go
panic: runtime error: comparing uncomparable type []int
```

这是一个运行时错误，提示 []int 是不可比较的类型。下表中列举出了类型及比较的几种情况。

| map             | 宕机错误，不可比较                                           |
| --------------- | ------------------------------------------------------------ |
| 切片（[]T）     | 宕机错误，不可比较                                           |
| 通道（channel） | 可比较，必须由同一个 make 生成，也就是同一个通道才会是 true，否则为 false |
| 数组（[容量]T） | 可比较，编译期知道两个数组是否一致                           |
| 结构体          | 可比较，可以逐个比较结构体的值                               |
| 函数            | 可比较                                                       |



