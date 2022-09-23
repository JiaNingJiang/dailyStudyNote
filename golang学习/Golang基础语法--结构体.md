### 一、结构体的成员变量

#### 1.结构体的嵌套

```go
type People struct {
    name  string
    child *People
}
```

People结构体中的child成员是一个嵌套结构体，被嵌套的结构体必须是结构体的指针类型，包含非指针类型会引起编译错误。

#### 2.匿名结构体

匿名结构体的初始化写法由结构体定义和键值对初始化两部分组成，结构体定义时没有结构体类型名，只有字段和类型定义，键值对初始化部分由可选的多个键值对组成，如下格式所示：

```go
ins := struct {
    // 匿名结构体字段定义
    字段1 字段类型1
    字段2 字段类型2
    …
}{
    // 字段值初始化
    初始化字段1: 字段1的值,
    初始化字段2: 字段2的值,
    …
}
```

```go
    // 实例化一个匿名结构体
    msg := &struct {  // 定义部分
        id   int
        data string
    }{  // 值初始化部分
        1024,
        "hello",
    }
```

匿名结构体的类型名是结构体包含字段成员的详细描述，匿名结构体在使用时需要重新定义，造成大量重复的代码，因此开发中较少使用。

### 二、类型内嵌和结构体内嵌

结构体可以包含一个或多个匿名（或内嵌）字段，即这些字段没有显式的名字，只有字段的类型是必须的，此时类型也就是字段的名字。匿名字段本身可以是一个结构体类型，即结构体可以包含内嵌结构体。可以粗略地将这个和面向对象语言中的继承概念相比较，随后将会看到它被用来模拟类似继承的行为。

```go
    package main
    import "fmt"
    type innerS struct {
        in1 int
        in2 int
    }
    type outerS struct {
        b int
        c float32
        int // anonymous field
        innerS //anonymous field
    }
    func main() {
        outer := new(outerS)
   
        outer.int = 60
        outer.in1 = 5
        outer.in2 = 10
    
        fmt.Printf("outer.int is: %d\n", outer.int)
        fmt.Printf("outer.in1 is: %d\n", outer.in1)
        fmt.Printf("outer.in2 is: %d\n", outer.in2)
        // 使用结构体字面量
        outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
        fmt.Printf("outer2 is:", outer2)
    }
```

运行结果如下所示：

```go
outer.int is: 60
outer.in1 is: 5
outer.in2 is: 10
outer2 is:{6 7.5 60 {5 10}}
```

通过类型 outer.int 的名字来获取存储在匿名字段中的数据，于是可以得出一个结论：在一个结构体中对于==每一种数据类型==只能有==一个匿名字段==。

#### 结构内嵌特性

##### 1) 内嵌的结构体可以直接访问其成员变量

嵌入结构体的成员，可以通过外部结构体的实例直接访问。如果结构体有多层嵌入结构体，结构体实例访问任意一级的嵌入结构体成员时都只用给出字段名，而无须像传统结构体字段一样，通过一层层的结构体字段访问到最终的字段。例如，==ins.a.b.c的访问可以简化为ins.c==。

##### 2) 内嵌结构体的字段名是它的类型名

内嵌结构体字段仍然可以使用详细的字段进行一层层访问，内嵌结构体的字段名就是它的类型名，代码如下：

```go
    var c Color
    c.BasicColor.R = 1
    c.BasicColor.G = 1
    c.BasicColor.B = 0
```

一个结构体只能嵌入一个同类型的成员，无须担心结构体重名和错误赋值的情况，编译器在发现可能的赋值歧义时会报错。

### 三、初始化内嵌结构体

```go
    package main
    import "fmt"
    // 车轮
    type Wheel struct {
        Size int
    }
    // 引擎
    type Engine struct {
        Power int    // 功率
        Type  string // 类型
    }
    // 车
    type Car struct {
        Wheel
        Engine
    }
    func main() {
        c := Car{
            // 初始化轮子
            Wheel: Wheel{
                Size: 18,
            },
            // 初始化引擎
            Engine: Engine{
                Type:  "1.4T",
                Power: 143,
            },
        }
        fmt.Printf("%+v\n", c)
    }
```

#### 初始化内嵌匿名结构体

在前面描述车辆和引擎的例子中，有时考虑编写代码的便利性，会将结构体直接定义在嵌入的结构体中。也就是说，结构体的定义不会被外部引用到。在初始化这个被嵌入的结构体时，就需要再次声明结构才能赋予数据。具体请参考下面的代码。

```go
    package main
    import "fmt"
    // 车轮
    type Wheel struct {
        Size int
    }
    // 车
    type Car struct {
        Wheel
        // 引擎
        Engine struct {
            Power int    // 功率
            Type  string // 类型
        }
    }
    func main() {
        c := Car{
            // 初始化轮子
            Wheel: Wheel{
                Size: 18,
            },
            // 初始化引擎
            Engine: struct {
                Power int
                Type  string
            }{
                Type:  "1.4T",
                Power: 143,
            },
        }
        fmt.Printf("%+v\n", c)
    }
```

### 四、内嵌结构体成员名字冲突

嵌入结构体内部可能拥有相同的成员名，成员重名时会发生什么？下面通过例子来讲解。

```go
    package main
    import (
        "fmt"
    )
    type A struct {
        a int
    }
    type B struct {
        a int
    }
    type C struct {
        A
        B
    }
    func main() {
        c := &C{}
        c.A.a = 1
        fmt.Println(c)
    }
```

接着，将第 22 行修改为如下代码：

```go
    func main() {
        c := &C{}
        c.a = 1        //引起歧义
        fmt.Println(c)
    }
```

此时再编译运行，编译器报错：

```go
.\main.go:22:3: ambiguous selector c.a
```

编译器告知 C 的选择器 a 引起歧义，也就是说，编译器无法决定将 1 赋给 C 中的 A 还是 B 里的字段 a。在使用内嵌结构体时，Go语言的编译器会非常智能地提醒我们可能发生的歧义和错误。

### 五、垃圾回收和SetFinalizer

Go语言自带垃圾回收机制（GC）。GC 通过独立的进程执行，它会搜索不再使用的变量，并将其释放。需要注意的是，GC 在运行时会占用机器资源。

 GC 是自动进行的，如果要手动进行 GC，可以使用 runtime.GC() 函数，显式的执行 GC。显式的进行 GC  只在某些特殊的情况下才有用，比如当内存资源不足时调用 runtime.GC() ，这样会立即释放一大片内存，但是会造成程序短时间的性能下降。

 finalizer（终止器）是与对象关联的一个函数，通过 runtime.SetFinalizer 来设置，如果某个对象定义了 finalizer，当它被 GC 时候，这个 finalizer 就会被调用，以完成一些特定的任务，例如发信号或者写日志等。

在Go语言中 SetFinalizer 函数是这样定义的：

```go
func SetFinalizer(x, f interface{})
```

参数说明如下：

- 参数 x 必须是一个指向通过 new 申请的对象的指针，或者通过对复合字面值取址得到的指针。
- 参数 f 必须是一个函数，它接受单个可以直接用 x 类型值赋值的参数，也可以有任意个被忽略的返回值。

SetFinalizer 函数可以将 x 的终止器设置为 f，当垃圾收集器发现 x 不能再直接或间接访问时，它会清理 x 并调用 f(x)。

 另外，x 的终止器会在 x  不能直接或间接访问后的任意时间被调用执行，不保证终止器会在程序退出前执行，因此一般终止器只用于在长期运行的程序中释放关联到某对象的非内存资源。例如，当一个程序丢弃一个 os.File 对象时没有调用其 Close 方法，该 os.File 对象可以使用终止器去关闭对应的操作系统文件描述符。

 终止器会按依赖顺序执行：如果 A 指向 B，两者都有终止器，且 A 和 B 没有其它关联，那么只有 A 的终止器执行完成，并且 A 被释放后，B 的终止器才可以执行。

 如果 *x 的大小为 0 字节，也不保证终止器会执行。

 此外，我们也可以使用`SetFinalizer(x, nil)`来清理绑定到 x 上的终止器。

> 提示：终止器只有在对象被 GC 时，才会被执行。其他情况下，都不会被执行，即使程序正常结束或者发生错误。

【示例】在函数 entry() 中定义局部变量并设置 finalizer，当函数 entry() 执行完成后，在 main 函数中手动触发 GC，查看 finalizer 的执行情况。

```go
    package main
    import (
        "log"
        "runtime"
        "time"
    )
    type Road int
    func findRoad(r *Road) {
        log.Println("road:", *r)
    }
    func entry() {
        var rd Road = Road(999)
        r := &rd
        runtime.SetFinalizer(r, findRoad)
    }
    func main() {
        entry()
        for i := 0; i < 10; i++ {
            time.Sleep(time.Second)
            runtime.GC()
        }
    }
```

运行结果如下：

```go
2019/11/28 15:32:16 road: 999
```

### 六、数据I/O对象及操作

Go语言标准库的 ==bufio 包==中，实现了对数据 I/O 接口的缓冲功能。这些功能封装于接口 io.ReadWriter、io.Reader 和 io.Writer 中，并对应创建了 ReadWriter、Reader 或 Writer 对象，在提供缓冲的同时实现了一些文本基本 I/O  操作功能。

#### 1.ReadWriter 对象

ReadWriter 对象可以对数据 I/O 接口 io.ReadWriter 进行输入输出缓冲操作，ReadWriter 结构定义如下：

```go
type ReadWriter struct {
    *Reader
    *Writer
}
```

可以使用 NewReadWriter() 函数创建 ReadWriter 对象，该函数的功能是根据指定的 Reader 和 Writer  创建一个 ReadWriter 对象，ReadWriter 对象将会向底层 io.ReadWriter 接口写入数据，或者从  io.ReadWriter 接口读取数据。该函数原型声明如下：

#### 2.Reader 对象

Reader 对象可以对数据 I/O 接口 io.Reader 进行输入缓冲操作，Reader 结构定义如下：

```go
type Reader struct {
    //contains filtered or unexported fields
)
```

#### 创建 Reader 对象

可以创建 Reader 对象的函数一共有两个，分别是 NewReader() 和 NewReaderSize()，下面分别介绍。

##### 1) NewReader() 函数

NewReader() 函数的功能是按照缓冲区默认长度创建 Reader 对象，Reader 对象会从底层 io.Reader 接口读取尽量多的数据进行缓存。该函数原型如下：

```go
func NewReader(rd io.Reader) *Reader	//其中，参数 rd 是 io.Reader 接口，Reader 对象将从该接口读取数据。
```

##### 2) NewReaderSize() 函数

NewReaderSize() 函数的功能是按照指定的缓冲区长度创建 Reader 对象，Reader 对象会从底层 io.Reader 接口读取尽量多的数据进行缓存。该函数原型如下：

```go
func NewReaderSize(rd io.Reader, size int) *Reader	//其中，参数 rd 是 io.Reader 接口，参数 size 是指定的缓冲区字节长度。
```

#### 操作 Reader 对象

操作 Reader 对象的方法共有 11 个，分别是  Read()、ReadByte()、ReadBytes()、ReadLine()、ReadRune  ()、ReadSlice()、ReadString()、UnreadByte()、UnreadRune()、Buffered()、Peek()，下面分别介绍。

##### 1) Read() 方法

Read() 方法的功能是读取数据，并存放到字节切片 p 中。Read() 执行结束会返回已读取的字节数，因为每次调用只调用底层的  io.Reader 一次，所以返回的 n 可能小于 len(p)，当字节流结束时，n 为 0，err 为 io. EOF。该方法原型如下：

```go
func (b *Reader) Read(p []byte) (n int, err error)
```

在方法 Read() 中，参数 p 是用于存放读取数据的字节切片。示例代码如下：

```go
    package main
    import (
        "bufio"
        "bytes"
        "fmt"
    )
    func main() {
        data := []byte("C语言中文网")
        rd := bytes.NewReader(data)
        r := bufio.NewReader(rd)
        var buf [128]byte
        n, err := r.Read(buf[:])
        fmt.Println(string(buf[:n]), n, err)
    }
```

##### 2) ReadByte() 方法

ReadByte() 方法的功能是读取并返回一个字节，如果没有字节可读，则返回错误信息。该方法原型如下：

```go
func (b *Reader) ReadByte() (c byte,err error)
```

示例代码如下：

```go
    package main
    import (
        "bufio"
        "bytes"
        "fmt"
    )
    func main() {
        data := []byte("Go语言入门教程")
        rd := bytes.NewReader(data)
        r := bufio.NewReader(rd)
        c, err := r.ReadByte()
        fmt.Println(string(c), err)
    }
```

##### 3) ReadBytes() 方法

ReadBytes() 方法的功能是读取数据直到遇到第一个分隔符“delim”，并返回读取的字节序列（包括“delim”）。如果  ReadBytes 在读到第一个“delim”之前出错，它返回已读取的数据和那个错误（通常是  io.EOF）。只有当返回的数据不以“delim”结尾时，返回的 err 才不为空值。该方法原型如下：

```go
func (b *Reader) ReadBytes(delim byte) (line []byte, err error)
```

```go
    package main
    import (
        "bufio"
        "bytes"
        "fmt"
    )
    func main() {
        data := []byte("C语言中文网, Go语言入门教程")
        rd := bytes.NewReader(data)
        r := bufio.NewReader(rd)
        var delim byte = ','
        line, err := r.ReadBytes(delim)
        fmt.Println(string(line), err)
    }
```

##### 4) ReadLine() 方法

ReadLine() 是一个低级的用于读取一行数据的方法，大多数调用者应该使用 ReadBytes('\n') 或者  ReadString('\n')。ReadLine 返回一行，不包括结尾的回车字符，如果一行太长（超过缓冲区长度），参数 isPrefix  会设置为 true 并且只返回前面的数据，剩余的数据会在以后的调用中返回。

当返回最后一行数据时，参数 isPrefix 会置为 false。返回的字节切片只在下一次调用 ReadLine 前有效。ReadLine 会返回一个非空的字节切片或一个错误，方法原型如下：

```go
func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
```

```go
    package main
    import (
        "bufio"
        "bytes"
        "fmt"
    )
    func main() {
        data := []byte("Golang is a beautiful language. \r\n I like it!")
        rd := bytes.NewReader(data)
        r := bufio.NewReader(rd)
        line, prefix, err := r.ReadLine()
        fmt.Println(string(line), prefix, err)
    }
```

运行结果如下：

```go
Golang is a beautiful language.  false <nil>
```

##### 5) ReadRune() 方法

ReadRune() 方法的功能是读取一个 UTF-8 编码的字符，并返回其 Unicode 编码和字节数。如果编码错误，ReadRune 只读取一个字节并返回 unicode.ReplacementChar(U+FFFD) 和长度 1。该方法原型如下：

```go
func (b *Reader) ReadRune() (r rune, size int, err error)
```

```go
    package main
    import (
        "bufio"
        "bytes"
        "fmt"
    )
    func main() {
        data := []byte("C语言中文网")
        rd := bytes.NewReader(data)
        r := bufio.NewReader(rd)
        ch, size, err := r.ReadRune()   //第一次
        fmt.Println(string(ch), size, err)
        
        ch, size, err := r.ReadRune()  //第二次
        fmt.Println(string(ch), size, err)
    }
```

运行结果如下：

```go
C 1 <nil>
语 3 <nil>
```

##### 6) ReadSlice() 方法

ReadSlice() 方法的功能是读取数据直到分隔符“delim”处，并返回读取数据的字节切片，下次读取数据时返回的切片会失效。如果 ReadSlice 在查找到“delim”之前遇到错误，它返回读取的所有数据和那个错误（通常是 io.EOF）。

 如果缓冲区满时也没有查找到“delim”，则返回 ErrBufferFull 错误。ReadSlice 返回的数据会在下次 I/O  操作时被覆盖，大多数调用者应该使用 ReadBytes 或者 ReadString。只有当 line  不以“delim”结尾时，ReadSlice 才会返回非空 err。该方法原型如下：

```go
func (b *Reader) ReadSlice(delim byte) (line []byte, err error)
```

```go
    package main
    import (
        "bufio"
        "bytes"
        "fmt"
    )
    func main() {
        data := []byte("C语言中文网, Go语言入门教程")
        rd := bytes.NewReader(data)
        r := bufio.NewReader(rd)
        var delim byte = ','
        line, err := r.ReadSlice(delim)
        fmt.Println(string(line), err)
        line, err = r.ReadSlice(delim)
        fmt.Println(string(line), err)
        line, err = r.ReadSlice(delim)
        fmt.Println(string(line), err)
    }
```

运行结果如下：

```go
C语言中文网, <nil>
Go语言入门教程 EOF
EOF
```

##### 7) ReadString() 方法

ReadString() 方法的功能是读取数据直到分隔符“delim”第一次出现，并返回一个包含“delim”的字符串。如果  ReadString 在读取到“delim”前遇到错误，它返回已读字符串和那个错误（通常是  io.EOF）。只有当返回的字符串不以“delim”结尾时，ReadString 才返回非空 err。该方法原型如下：

```go
func (b *Reader) ReadString(delim byte) (line string, err error)
```

```go
    package main
    import (
        "bufio"
        "bytes"
        "fmt"
    )
    func main() {
        data := []byte("C语言中文网, Go语言入门教程")
        rd := bytes.NewReader(data)
        r := bufio.NewReader(rd)
        var delim byte = ','
        line, err := r.ReadString(delim)
        fmt.Println(line, err)
    }
```

##### 8) UnreadByte() 方法

UnreadByte() 方法的功能是取消已读取的最后一个字节（即把字节重新放回读取缓冲区的前部）。只有最近一次读取的单个字节才能取消读取。该方法原型如下：

```go
func (b *Reader) UnreadByte() error
```

##### 9) UnreadRune() 方法

UnreadRune() 方法的功能是取消读取最后一次读取的 Unicode 字符。如果最后一次读取操作不是  ReadRune，UnreadRune 会返回一个错误（在这方面它比 UnreadByte 更严格，因为 UnreadByte  会取消上次任意读操作的最后一个字节）。该方法原型如下：

```go
func (b *Reader) UnreadRune() error
```

##### 10) Buffered() 方法

Buffered() 方法的功能是返回剩余可从缓冲区读出数据的字节数, 示例代码如下：

```go
    package main
    import (
        "bufio"
        "bytes"
        "fmt"
    )
    func main() {
        data := []byte("Go语言入门教程")
        rd := bytes.NewReader(data)
        r := bufio.NewReader(rd)
        var buf [14]byte
        n, err := r.Read(buf[:])
        fmt.Println(string(buf[:n]), n, err)
        rn := r.Buffered()
        fmt.Println(rn)
        n, err = r.Read(buf[:])
        fmt.Println(string(buf[:n]), n, err)
        rn = r.Buffered()
        fmt.Println(rn)
    }
```

运行结果如下：

```go
Go语言入门 14 <nil>
6
教程 6 <nil>
0
```

##### 11) Peek() 方法

Peek() 方法的功能是读取指定字节数的数据，这些被读取的数据不会从缓冲区中清除。在下次读取之后，本次返回的字节切片会失效。如果 Peek  返回的字节数不足 n 字节，则会同时返回一个错误说明原因，如果 n 比缓冲区要大，则错误为 ErrBufferFull。该方法原型如下:

```go
func (b *Reader) Peek(n int) ([]byte, error)
```

在方法 Peek() 中，参数 n 是希望读取的字节数。示例代码如下：

```go
    package main
    import (
        "bufio"
        "bytes"
        "fmt"
    )
    func main() {
        data := []byte("Go语言入门教程")
        rd := bytes.NewReader(data)
        r := bufio.NewReader(rd)
        bl, err := r.Peek(8)
        fmt.Println(string(bl), err)
        bl, err = r.Peek(14)
        fmt.Println(string(bl), err)
        bl, err = r.Peek(20)
        fmt.Println(string(bl), err)
    }
```

运行结果如下：

```go
Go语言 <nil>
Go语言入门 <nil>
Go语言入门教程 <nil>
```





#### 3.Writer 对象

Writer 对象可以对数据 I/O 接口 io.Writer 进行输出缓冲操作，Writer 结构定义如下：

```go
type Writer struct {
    //contains filtered or unexported fields
}
```

默认情况下 Writer 对象没有定义初始值，如果输出缓冲过程中发生错误，则数据写入操作立刻被终止，后续的写操作都会返回写入异常错误。

#### 创建 Writer 对象

创建 Writer 对象的函数共有两个分别是 NewWriter() 和 NewWriterSize()，下面分别介绍一下。

##### 1) NewWriter() 函数

NewWriter() 函数的功能是按照默认缓冲区长度创建 Writer 对象，Writer 对象会将缓存的数据批量写入底层 io.Writer 接口。该函数原型如下：

```go
func NewWriter(wr io.Writer) *Writer
```

其中，参数 wr 是 io.Writer 接口，Writer 对象会将数据写入该接口。

##### 2) NewWriterSize() 函数

NewWriterSize() 函数的功能是按照指定的缓冲区长度创建 Writer 对象，Writer 对象会将缓存的数据批量写入底层 io.Writer 接口。该函数原型如下：

```go
func NewWriterSize(wr io.Writer, size int) *Writer
```

其中，参数 wr 是 io.Writer 接口，参数 size 是指定的缓冲区字节长度。

#### 操作 Writer 对象

操作 Writer 对象的方法共有 7 个，分别是 Available()、Buffered()、Flush()、Write()、WriteByte()、WriteRune() 和 WriteString() 方法，下面分别介绍。

##### 1) Available() 方法

Available() 方法的功能是返回缓冲区中未使用的字节数，该方法原型如下：

```go
func (b *Writer) Available() int
```

```go
    package main
    import (
        "bufio"
        "bytes"
        "fmt"
    )
    func main() {
        wr := bytes.NewBuffer(nil)
        w := bufio.NewWriter(wr)      //按照默认缓冲区长度创建 Writer 对象
        p := []byte("C语言中文网")
        fmt.Println("写入前未使用的缓冲区为：", w.Available())
        w.Write(p)
        fmt.Printf("写入%q后，未使用的缓冲区为：%d\n", string(p), w.Available())
    }
```

运行结果如下：

```go
写入前未使用的缓冲区为： 4096
写入"C语言中文网"后，未使用的缓冲区为：4080
```

##### 2) Buffered() 方法

Buffered() 方法的功能是返回已写入当前缓冲区中的字节数，该方法原型如下：

```go
func (b *Writer) Buffered() int
```

```go
    package main
    import (
        "bufio"
        "bytes"
        "fmt"
    )
    func main() {
        wr := bytes.NewBuffer(nil)
        w := bufio.NewWriter(wr)
        p := []byte("C语言中文网")
        fmt.Println("写入前未使用的缓冲区为：", w.Buffered())
        w.Write(p)
        fmt.Printf("写入%q后，未使用的缓冲区为：%d\n", string(p), w.Buffered())
        w.Flush()
        fmt.Println("执行 Flush 方法后，写入的字节数为：", w.Buffered())
    }
```

该例测试结果为：

```go
写入前未使用的缓冲区为： 0
写入"C语言中文网"后，未使用的缓冲区为：16
执行 Flush 方法后，写入的字节数为： 0
```

##### 3) Flush() 方法

Flush() 方法的功能是把缓冲区中的数据写入底层的 io.Writer，并返回错误信息。如果成功写入，error 返回 nil，否则 error 返回错误原因。该方法原型如下：

```go
func (b *Writer) Flush() error
```

```go
    package main
    import (
        "bufio"
        "bytes"
        "fmt"
    )
    func main() {
        wr := bytes.NewBuffer(nil)
        w := bufio.NewWriter(wr)
        p := []byte("C语言中文网")
        w.Write(p)
        fmt.Printf("未执行 Flush 缓冲区输出 %q\n", string(wr.Bytes()))
        w.Flush()
        fmt.Printf("执行 Flush 后缓冲区输出 %q\n", string(wr.Bytes()))
    }
```

运行结果如下：

```go
未执行 Flush 缓冲区输出 ""
执行 Flush 后缓冲区输出 "C语言中文网"
```

##### 4) Write() 方法

Write() 方法的功能是把字节切片 p 写入==缓冲区==，返回已写入的字节数 nn。如果 nn 小于 len(p)，则同时返回一个错误原因。该方法原型如下：

```go
func (b *Writer) Write(p []byte) (nn int, err error)
```

其中，参数 p 是要写入的字节切片。示例代码如下：

```go
    package main
    import (
        "bufio"
        "bytes"
        "fmt"
    )
    func main() {
        wr := bytes.NewBuffer(nil)
        w := bufio.NewWriter(wr)
        p := []byte("C语言中文网")
        n, err := w.Write(p)
        w.Flush()     //仍然需要调用Flush()，将数据写入底层io.Writer
        fmt.Println(string(wr.Bytes()), n, err)
    }
```

运行结果如下：

```go
C语言中文网 16 <nil>
```

##### 5) WriteByte() 方法

WriteByte() 方法的功能是向缓冲区写入一个字节，如果成功写入，error 返回 nil，否则 error 返回错误原因。该方法原型如下：

```go
func (b *Writer) WriteByte(c byte) error
```

其中，参数 c 是要写入的字节数据，比如 ASCII 字符。示例代码如下：

```go
    package main
    import (
        "bufio"
        "bytes"
        "fmt"
    )
    func main() {
        wr := bytes.NewBuffer(nil)
        w := bufio.NewWriter(wr)
        var c byte = 'G'
        err := w.WriteByte(c)
        w.Flush()
        fmt.Println(string(wr.Bytes()), err)
    }
```

运行结果如下：

```go
G <nil>
```

##### 6) WriteRune() 方法

WriteRune() 方法的功能是以 UTF-8 编码写入一个 Unicode 字符，返回写入的字节数和错误信息。该方法原型如下：

```go
func (b *Writer) WriteRune(r rune) (size int,err error)
```

其中，参数 r 是要写入的 Unicode 字符。示例代码如下：

```go
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	wr := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wr)
	var r rune = '我'
	size, err := w.WriteRune(r)
	w.Flush()
	fmt.Println(string(wr.Bytes()), size, err)
}
```

该例测试结果为：

```go
我 3 <nil>
```

##### 7) WriteString() 方法

WriteString() 方法的功能是写入一个字符串，并返回写入的字节数和错误信息。如果返回的字节数小于 len(s)，则同时返回一个错误说明原因。该方法原型如下：

```go
func (b *Writer) WriteString(s string) (int, error)
```

其中，参数 s 是要写入的字符串。示例代码如下：

```go
    package main
    import (
        "bufio"
        "bytes"
        "fmt"
    )
    func main() {
        wr := bytes.NewBuffer(nil)
        w := bufio.NewWriter(wr)
        s := "C语言中文网"
        n, err := w.WriteString(s)
        w.Flush()
        fmt.Println(string(wr.Bytes()), n, err)
    }
```

运行结果如下：

```go
C语言中文网 16 <nil>
```

