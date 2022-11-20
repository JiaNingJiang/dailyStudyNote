## 一、反射的基本概念

Go语言提供了一种机制在运行时更新和检查变量的值、调用变量的方法和变量支持的内在操作，但是在编译时并不知道这些变量的具体类型，这种机制被称为反射。反射也可以让我们将类型本身作为第一类的值类型处理。

 反射是指在程序运行期对程序本身进行访问和修改的能力，程序在编译时变量被转换为内存地址，变量名不会被编译器写入到可执行部分，在运行程序时程序无法获取自身的信息。

 支持反射的语言可以在程序编译期将变量的反射信息，如字段名称、类型信息、结构体信息等整合到可执行文件中，并给程序提供接口访问反射信息，这样就可以在程序运行期获取类型的反射信息，并且有能力修改它们。

 C/[C++](http://c.biancheng.net/cplus/)语言没有支持反射功能，只能通过 typeid 提供非常弱化的程序运行时类型信息；Java、[C#](http://c.biancheng.net/csharp/) 等语言都支持完整的反射功能；Lua、[JavaScript](http://c.biancheng.net/js/) 类动态语言，由于其本身的语法特性就可以让代码在运行期访问程序自身的值和类型信息，因此不需要反射系统。

 Go语言程序的反射系统无法获取到一个可执行文件空间中或者是一个包中的所有类型信息，需要配合使用标准库中对应的词法、语法解析器和抽象语法树（AST）对源码进行扫描后获得这些信息。

 Go语言提供了 reflect 包来访问程序的反射信息。

### 1. reflect 包

Go语言中的反射是由 reflect 包提供支持的，它定义了两个重要的类型 Type 和 Value。 **任意接口值在反射中都可以理解为由  reflect.Type 和 reflect.Value 两部分组成**，并且 **reflect 包提供了 reflect.TypeOf 和  reflect.ValueOf 两个函数来获取任意对象的 Value 和 Type。**

### 2. 反射的类型对象（reflect.Type）

在Go语言程序中，使用 reflect.TypeOf() 函数可以获得任意值的类型对象（reflect.Type），程序通过类型对象可以访问任意值的类型信息，下面通过示例来理解获取类型对象的过程：

```go
    package main
    import (
        "fmt"
        "reflect"
    )
    func main() {
        var a int
        typeOfA := reflect.TypeOf(a)
        fmt.Println(typeOfA.Name(), typeOfA.Kind())
    }
```

```go
运行结果如下：

int  int
```

代码说明如下：

- 第 9 行，定义一个 int 类型的变量。
- 第 10 行，通过 reflect.TypeOf() 取得变量 a 的类型对象 typeOfA，类型为 reflect.Type()。
- 第 11 行中，通过 typeOfA 类型对象的成员函数，可以分别获取到 typeOfA 变量的类型名为 int，种类（Kind）为 int。

### 3. 反射的类型（Type）与种类（Kind）

在使用反射时，需要首先理解**类型（Type）和种类（Kind）的区别**。编程中，使用最多的是类型(Type)，但在反射中，当需要区分一个大品种的类型时，就会用到种类（Kind）。例如需要统一判断类型中的指针时，使用种类（Kind）信息就较为方便。

#### 3.1 反射种类（Kind）的定义

Go语言程序中的类型（Type）指的是系统原生数据类型，如 int、string、bool、float32 等类型，以及使用 type  关键字定义的类型，这些类型的名称就是其类型本身的名称。例如使用 type A struct{} 定义结构体时，A 就是 struct{}  的类型（Type）。

 种类（Kind）指的是对象归属的品种，在 reflect 包中有如下定义：

```go
    type Kind uint
    const (
        Invalid Kind = iota  // 非法类型
        Bool                 // 布尔型
        Int                  // 有符号整型
        Int8                 // 有符号8位整型
        Int16                // 有符号16位整型
        Int32                // 有符号32位整型
        Int64                // 有符号64位整型
        Uint                 // 无符号整型
        Uint8                // 无符号8位整型
        Uint16               // 无符号16位整型
        Uint32               // 无符号32位整型
        Uint64               // 无符号64位整型
        Uintptr              // 指针
        Float32              // 单精度浮点数
        Float64              // 双精度浮点数
        Complex64            // 64位复数类型
        Complex128           // 128位复数类型
        Array                // 数组
        Chan                 // 通道
        Func                 // 函数
        Interface            // 接口
        Map                  // 映射
        Ptr                  // 指针
        Slice                // 切片
        String               // 字符串
        Struct               // 结构体
        UnsafePointer        // 底层指针
    )
```

Map、Slice、Chan 属于引用类型，使用起来类似于指针，但是在种类常量定义中仍然属于独立的种类，不属于 Ptr。type A struct{} 定义的结构体属于 Struct 种类，*A 属于 Ptr。

#### 3.2 从类型对象中获取类型名称和种类

Go语言中的类型**名称对应的反射获取方法是 reflect.Type 中的 Name() 方法**，返回表示类型名称的**字符串**；类型归属的**种类（Kind）使用的是 reflect.Type 中的 Kind() 方法**，返回 **reflect.Kind 类型的常量**。

下面的代码中会对常量和结构体进行类型信息获取。

```go
    package main
    import (
        "fmt"
        "reflect"
    )
    // 定义一个Enum类型
    type Enum int
    const (
        Zero Enum = 0
    )
    func main() {
        // 声明一个空结构体
        type cat struct {
        }
        // 获取结构体实例的反射类型对象
        typeOfCat := reflect.TypeOf(cat{})
        // 显示反射类型对象的名称和种类
        fmt.Println(typeOfCat.Name(), typeOfCat.Kind())
        // 获取Zero常量的反射类型对象
        typeOfA := reflect.TypeOf(Zero)
        // 显示反射类型对象的名称和种类
        fmt.Println(typeOfA.Name(), typeOfA.Kind())
    }
```

运行结果如下：

```go
cat struct
Enum int
```

代码说明如下：

- 第 17 行，声明结构体类型 cat。
- 第 20 行，将 cat 实例化，并且使用 reflect.TypeOf() 获取被实例化后的 cat 的反射类型对象。
- 第 22 行，输出 cat 的类型名称和种类，类型名称就是 cat，而 cat 属于一种结构体种类，因此种类为 struct。
- 第 24 行，Zero 是一个 Enum 类型的常量。这个 Enum 类型在第 9 行声明，第 12 行声明了常量。如没有常量也不能创建实例，通过 reflect.TypeOf() 直接获取反射类型对象。
- 第 26 行，输出 Zero 对应的类型对象的类型名和种类。

### 4. 指针与指针指向的元素

Go语言程序中**通过指针获取反射对象**时，可以通过 **reflect.Elem() 方法获取这个指针指向的元素类型**，这个获取过程被称为取元素，**等效于对指针类型变量做了一个`*`操作**，代码如下：

```go
    package main
    import (
        "fmt"
        "reflect"
    )
    func main() {
        // 声明一个空结构体
        type cat struct {
        }
        // 创建cat的实例
        ins := &cat{}
        // 获取结构体实例的反射类型对象
        typeOfCat := reflect.TypeOf(ins)
        // 显示反射类型对象的名称和种类
        fmt.Printf("name:'%v' kind:'%v'\n", typeOfCat.Name(), typeOfCat.Kind())
        // 取类型的元素
        typeOfCat = typeOfCat.Elem()
        // 显示反射类型对象的名称和种类
        fmt.Printf("element name: '%v', element kind: '%v'\n", typeOfCat.Name(), typeOfCat.Kind())
    }
```

运行结果如下：

```go
name:'' kind:'ptr'
element name: 'cat', element kind: 'struct'
```

代码说明如下：

- 第 13 行，创建了 cat 结构体的实例，ins 是一个 *cat 类型的指针变量。
- 第 15 行，对指针变量获取反射类型信息。
- 第 17 行，输出指针变量的类型名称和种类。Go语言的反射中对**所有指针变量的种类都是 Ptr**，但**需要注意的是，指针变量的类型名称是空，不是 *cat**。
- 第 19 行，取指针类型的元素类型，也就是 cat 类型。这个操作不可逆，不可以通过一个非指针类型获取它的指针类型。
- 第 21 行，输出指针变量指向元素的类型名称和种类，得到了 cat 的类型名称（cat）和种类（struct）。

### 5. 使用反射获取结构体的成员类型

任意值通过 reflect.TypeOf() 获得反射对象信息后，如果它的**类型是结构体**，可以**通过反射值对象 reflect.Type 的 NumField() 和 Field() 方法获得结构体成员的详细信息**。

 与**成员获取相关**的 reflect.Type 的方法如下表所示。

| 方法                                                        | 说明                                                         |
| ----------------------------------------------------------- | ------------------------------------------------------------ |
| Field(i int) StructField                                    | 根据索引返回索引对应的结构体字段的信息，当值不是结构体或索引超界时发生宕机 |
| NumField() int                                              | 返回结构体成员字段数量，当类型不是结构体或索引超界时发生宕机 |
| FieldByName(name string) (StructField, bool)                | 根据给定字符串返回字符串对应的结构体字段的信息，没有找到时 bool 返回 false，当类型不是结构体或索引超界时发生宕机 |
| FieldByIndex(index []int) StructField                       | 多层成员访问时，根据 []int 提供的每个结构体的字段索引，返回字段的信息，没有找到时返回零值。当类型不是结构体或索引超界时发生宕机 |
| FieldByNameFunc(match func(string) bool) (StructField,bool) | 根据匹配函数匹配需要的字段，当值不是结构体或索引超界时发生宕机 |

#### 5.1 结构体字段类型

reflect.Type 的 Field() 方法返回 **StructField  结构**，这个结构**描述结构体对应的成员信息**，通过这个信息可以获取成员与结构体的关系，如偏移、索引、是否为匿名字段、结构体标签（StructTag）等，而且还可以**通过 StructField 的 Type 字段进一步获取结构体成员的类型信息**。

StructField 的结构如下：

```go
    type StructField struct {
        Name string          // 字段名
        PkgPath string       // 字段路径
        Type      Type       // 字段反射类型对象
        Tag       StructTag  // 字段的结构体标签
        Offset    uintptr    // 字段在结构体中的相对偏移
        Index     []int      // Type.FieldByIndex中的返回的索引值
        Anonymous bool       // 是否为匿名字段
    }
```

字段说明如下：

- Name：为字段名称。
- PkgPath：字段在结构体中的路径。
- Type：字段本身的反射类型对象，类型为 reflect.Type，可以进一步获取字段的类型信息。
- Tag：结构体标签，为结构体字段标签的额外信息，可以单独提取。
- Index：FieldByIndex 中的索引顺序。
- Anonymous：表示该字段是否为匿名字段。

#### 5.2 获取成员反射信息

下面代码中，实例化一个结构体并遍历其结构体成员，再通过 reflect.Type 的 FieldByName() 方法查找结构体中指定名称的字段，直接获取其类型信息。

 反射访问结构体成员类型及信息：

```go
    package main
    import (
        "fmt"
        "reflect"
    )
    func main() {
        // 声明一个空结构体
        type cat struct {
            Name string
            // 带有结构体tag的字段
            Type int `json:"type" id:"100"`
        }
        // 创建cat的实例
        ins := cat{Name: "mimi", Type: 1}
        // 获取结构体实例的反射类型对象
        typeOfCat := reflect.TypeOf(ins)
        // 遍历结构体所有成员
        for i := 0; i < typeOfCat.NumField(); i++ {
            // 获取每个成员的结构体字段类型
            fieldType := typeOfCat.Field(i)
            // 输出成员名和tag
            fmt.Printf("name: %v  tag: '%v'\n", fieldType.Name, fieldType.Tag)
        }
        // 通过字段名, 找到字段类型信息
        if catType, ok := typeOfCat.FieldByName("Type"); ok {
            // 从tag中取出需要的tag
            fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
        }
    }
```

代码输出如下：

```go
name: Name  tag: ''
name: Type  tag: 'json:"type" id:"100"'
type 100
```

代码说明如下：

- 第 10 行，声明了带有两个成员的 cat 结构体。
- 第 13 行，Type 是 cat 的一个成员，这个成员类型后面带有一个以 ` 开始和结尾的字符串。这个字符串在Go语言中被称为 Tag（标签）。一般用于给字段添加自定义信息，方便其他模块根据信息进行不同功能的处理。
- 第 16 行，创建 cat 实例，并对两个字段赋值。结构体标签属于类型信息，无须且不能赋值。
- 第 18 行，获取实例的反射类型对象。
- 第 20 行，使用 reflect.Type 类型的 NumField() 方法获得一个结构体类型共有多少个字段。如果类型不是结构体，将会触发宕机错误。
- 第 22 行，**reflect.Type 中的 Field() 方法和 NumField 一般都是配对使用，用来实现结构体成员的遍历操作。**
- 第 24 行，使用 reflect.Type 的 Field() 方法返回的结构不再是 reflect.Type 而是 StructField 结构体。
- 第 27 行，使用 reflect.Type 的 FieldByName() 根据字段名查找结构体字段信息，catType 表示返回的结构体字段信息，类型为 StructField，ok 表示是否找到结构体字段的信息。
- 第 29 行中，**使用 StructField 中 Tag 的 Get() 方法，根据 Tag 中的名字进行信息获取。**

### 6. 结构体标签（Struct Tag）

通过 reflect.Type 获取结构体成员信息 reflect.StructField 结构中的 Tag 被称为结构体标签（StructTag）。结构体标签是对结构体字段的额外信息标签。

 JSON、BSON 等格式进行序列化及对象关系映射（Object Relational Mapping，简称  ORM）系统都会用到结构体标签，这些系统使用标签设定字段在处理时应该具备的特殊属性和可能发生的行为。**这些信息都是静态的，无须实例化结构体，可以通过反射获取到。**

#### 6.1 结构体标签的格式

Tag 在结构体字段后方书写的格式如下：

```go
`key1:"value1" key2:"value2"`
```

结构体标签由**一个或多个键值对组成**；键与值使用冒号分隔，**值用双引号括起来**；**键值对之间使用一个空格分隔**。

#### 6.2 从结构体标签中获取值

StructTag 拥有一些方法，可以进行 Tag 信息的解析和提取，如下所示：

- `func (tag StructTag) Get(key string) string`：根据 Tag 中的键获取对应的值，例如``key1:"value1" key2:"value2"``的 Tag 中，可以传入“key1”获得“value1”。
- `func (tag StructTag) Lookup(key string) (value string, ok bool)`：根据 Tag 中的键，查询值是否存在。

#### 6.3 结构体标签格式错误导致的问题

**编写 Tag 时，必须严格遵守键值对的规则**。**结构体标签的解析代码的容错能力很差**，一旦格式写错，编译和运行时都**不会提示任何错误**，示例代码如下：

```go
    package main
    import (
        "fmt"
        "reflect"
    )
    func main() {
        type cat struct {
            Name string
            Type int `json: "type" id:"100"`
        }
        typeOfCat := reflect.TypeOf(cat{})
        if catType, ok := typeOfCat.FieldByName("Type"); ok {
            fmt.Println(catType.Tag.Get("json"))
        }
    }
```

运行上面的代码会输出一个空字符串，并不会输出期望的 type。

代码第 11 行中，**在 json: 和 "type" 之间增加了一个空格，这种写法没有遵守结构体标签的规则，因此无法通过 Tag.Get  获取到正确的 json 对应的值**。这个错误在开发中非常容易被疏忽，造成难以察觉的错误。所以**将第 12 行代码修改为下面的样子**，则可以正常打印。

```go
    type cat struct {
        Name string
        Type int `json:"type" id:"100"`
    }
```

运行结果如下：

```go
type
```

## 二、反射规则浅析

反射是众多编程语言中的一个非常实用的功能，它是一种能够自描述、自控制的应用，Go语言也对反射提供了友好的支持。

 Go语言中使用反射可以在编译时不知道类型的情况下更新变量，在运行时查看值、调用方法以及直接对他们的布局进行操作。

 由于反射是建立在类型系统（type system）上的，所以我们先来复习一下Go语言中的类型。

### 1. Go语言中的类型

Go语言是一门静态类型的语言，每个变量都有一个静态类型，类型在编译的时候确定下来。

```go
type MyInt int

var i int
var j MyInt
```

变量 i 的类型是 int，变量 j 的类型是 MyInt，**虽然它们有着相同的基本类型，但静态类型却不一样**，在**没有类型转换的情况下，它们之间无法互相赋值**。

 接口是一个重要的类型，它意味着一个确定的方法集合，**一个接口变量可以存储任何实现了接口的方法的具体值**（除了接口本身），例如 io.Reader 和 io.Writer.

如果一个类型声明实现了 Reader（或 Writer）方法，那么它便实现了 io.Reader（或 io.Writer），这意味着**一个 io.Reader 的变量可以持有任何一个实现了 Read 方法的的类型的值**。

```go
var r io.Reader
r = os.Stdin
r = bufio.NewReader(r)
r = new(bytes.Buffer)
// and so on
```

必须要弄清楚的一点是，**不管变量 r 中的具体值是什么，r 的类型永远是 io.Reader**，由于**Go语言是静态类型的，r 的静态类型就是 io.Reader。**

在接口类型中有一个极为重要的例子——**空接口**：

```go
interface{}
```

它**表示了一个空的方法集，一切值都可以满足它，因为它们都有零值或方**法。

 有人说Go语言的接口是动态类型，这是错误的，它们都是**静态类型**，**虽然在运行时中，接口变量存储的值也许会变，但接口变量的类型是不会变的**。我们必须精确地了解这些，因为反射与接口是密切相关的。

 关于接口我们就介绍到这里，下面我们看看**Go语言的反射三定律**。

### 2. 反射第一定律：反射可以将“接口类型变量”转换为“反射类型对象”

> 注：这里反射类型指 reflect.Type 和 reflect.Value。

从使用方法上来讲，反射提供了一种机制，允许程序在运行时检查接口变量内部存储的 (value, type) 对。

 在最开始，我们先了解下 reflect 包的两种类型 Type 和  Value，这两种类型使访问接口内的数据成为可能，它们对应两个简单的方法，分别是 reflect.TypeOf 和  reflect.ValueOf，分别用来读取接口变量的 reflect.Type 和 reflect.Value 部分。

 当然，从 reflect.Value 也很容易获取到 reflect.Type，目前我们先将它们分开。

 首先，我们下看 reflect.TypeOf：

```go
    package main
    import (
        "fmt"
        "reflect"
    )
    func main() {
        var x float64 = 3.4
        fmt.Println("type:", reflect.TypeOf(x))
    }
```

运行结果如下：

```go
type: float64
```

大家可能会疑惑，为什么没看到接口？**这段代码看起来只是把一个 float64 类型的变量 x 传递给 reflect.TypeOf 并没有传递接口**。其实在 **reflect.TypeOf 的函数参数里包含一个空接口**：

```go
// TypeOf returns the reflection Type of the value in the interface{}.
func TypeOf(i interface{}) Type
```

我们调用 reflect.TypeOf(x) 时，**x 被存储在一个空接口变量中被传递过去**，然后 **reflect.TypeOf 对空接口变量进行拆解，恢复其类型信息**。

函数 reflect.ValueOf 也会**对底层的值进行恢复**：

```go
    package main
    import (
        "fmt"
        "reflect"
    )
    func main() {
        var x float64 = 3.4
        fmt.Println("value:", reflect.ValueOf(x))
    }
```

运行结果如下：

```go
value: 3.4
```

类型 reflect.Type 和 reflect.Value 都有很多方法，我们可以检查和使用它们，这里我们举几个例子。

1. 类型 **reflect.Value 有一个方法 Type()，它会返回一个 reflect.Type 类型的对象**。

2.  Type 和 Value **都有一个名为 Kind 的方法，它会返回一个常量，表示底层数据的类型**，常见值有：Uint、Float64、Slice 等。

3. **Value 类型**也有一些**类似于 Int、Float 的方法，用来提取底层的数据**：
   - Int 方法用来提取 int64
   - **Float 方法用来提取 float64（float类型里面最大的类型,因此不会是float32）**，示例代码如下：

```go
    package main
    import (
        "fmt"
        "reflect"
    )
    func main() {
        var x float64 = 3.4
        v := reflect.ValueOf(x)
        fmt.Println("type:", v.Type())
        fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
        fmt.Println("value:", v.Float())
    }
```

运行结果如下：

```go
type: float64
kind is float64: true
value: 3.4
```

反射库提供了很多值得列出来单独讨论的属性，下面就来介绍一下。

首先是介绍下 **Value 的 getter 和 setter 方法**，为了保证 API  的精简，这两个方法**操作的是某一组类型范围最大的那个**。比如，**处理任何含无符号整型数，都使用 uint64**，也就是说 **Value 类型的 uint  方法返回值为 uint64 类型**，SetInt 方法接收的参数类型也是 uint64 类型。实际使用时，可能需要转化为实际的类型：

```go
    package main
    import (
        "fmt"
        "reflect"
    )
    func main() {
        var x uint8 = 'x'
        v := reflect.ValueOf(x)
        fmt.Println("type:", v.Type())                            // uint8.
        fmt.Println("kind is uint8: ", v.Kind() == reflect.Uint8) // true.
        x = uint8(v.Uint())                                       // v.Uint returns a uint64.(必须强制类型转换)
    }
```

运行结果如下：

```go
type: uint8
kind is uint8: true
```

其次，**反射对象的 Kind 方法描述的是基础类型，而不是静态类型**。如果一个反射对象包含了用户定义类型的值，如下所示：

```go
type MyInt int
var x MyInt = 7
v := reflect.ValueOf(x)
```

上面的代码中，**虽然变量 v 的静态类型是 MyInt，而不是 int，但 Kind 方法仍然会返回 reflect.Int**。换句话说 **Kind 方法不会像 Type 方法一样区分 MyInt 和 int**。

### 3. 反射第二定律：反射可以将“反射类型对象”转换为“接口类型变量”

和物理学中的反射类似，Go语言中的反射也能创造自己反面类型的对象。

 根据一个 **reflect.Value 类型的变量**，我们可以**使用 Interface 方法恢复其接口类型的值**。事实上，这个方法**会把 type 和 value 信息打包并填充到一个接口变量中，然后返回**。

 其函数声明如下：

```go
// Interface returns v's value as an interface{}.
func (v Value) Interface() interface{}
```

然后，我们可以**通过断言，恢复底层的具体值**：

```go
y := v.Interface().(float64) // y will have type float64.
fmt.Println(y)
```

上面这段代码会打印出一个 float64 类型的值，也就是反射类型变量 v 所代表的值。

事实上，我们可以更好地利用这一特性，标准库中的 **fmt.Println 和 fmt.Printf 等函数都接收空接口变量作为参数**，**fmt  包内部会对接口变量进行拆包**，因此 fmt 包的打印函数在打印 reflect.Value 类型变量的数据时，**只需要把 Interface  方法的结果传给格式化打印程序**：

```go
fmt.Println(v.Interface())
```

为什么不直接使用 fmt.Println(v)？因为 v 的类型是 reflect.Value，我们需要的是它的具体值，**由于值的类型是 float64，我们也可以用浮点格式化打印它**：

```go
fmt.Printf("value is %7.1e\n", v.Interface())
```

运行结果如下：

```go
3.4e+00
```

同样，这次也**不需要对 v.Interface() 的结果进行类型断言，空接口值内部包含了具体值的类型信息，Printf 函数会恢复类型信息**。

简单来说 **Interface 方法和 ValueOf 函数作用恰好相反**，唯一一点是，**返回值的静态类型是 interface{}**。

Go的反射机制可以将“接口类型的变量”转换为“反射类型的对象”，然后再将“反射类型对象”转换过去。

### 4. 反射第三定律：如果要修改“反射类型对象”其值必须是“可写的”

这条定律很微妙，也很容易让人迷惑，但是如果从第一条定律开始看，应该比较容易理解。

下面这段代码虽然不能正常工作，但是非常值得研究：

```go
var x float64 = 3.4
v := reflect.ValueOf(x)
v.SetFloat(7.1) // Error: will panic
```

如果运行这段代码，它会抛出一个奇怪的异常：

```go
panic: reflect: reflect.flag.mustBeAssignable using unaddressable value
```

这里**问题不在于值`7.1 `不能被寻址**，而是**因为变量 v 是“不可写的”**，**“可写性”是反射类型变量的一个属性，但不是所有的反射类型变量都拥有这个属性**。

 我们可以**通过 CanSet 方法检查一个 reflect.Value 类型变量的“可写性”**，对于上面的例子，可以这样写：

```go
    package main
    import (
        "fmt"
        "reflect"
    )
    func main() {
        var x float64 = 3.4
        v := reflect.ValueOf(x)
        fmt.Println("settability of v:", v.CanSet())
    }
```

运行结果如下：

```go
settability of v: false
```

**对于一个不具有“可写性”的 Value 类型变量，调用 Set 方法会返回false**。

首先我们要**弄清楚什么是“可写性”**，**“可写性”有些类似于寻址能力，但是更严格，它是反射类型变量的一种属性，赋予该变量修改底层存储数据的能力**。“可写性”最终**是由一个反射对象是否存储了原始值而决定的**。

示例代码如下：

```go
var x float64 = 3.4
v := reflect.ValueOf(x)
```

这里我们**传递给 reflect.ValueOf 函数的是变量 x 的一个拷贝，而非 x 本身**，想象一下如果下面这行代码能够成功执行：

```go
v.SetFloat(7.1)
```

**如果这行代码能够成功执行，它不会更新 x，虽然看起来变量 v 是根据 x 创建的，相反它会更新 x 存在于反射对象 v 内部的一个拷贝，而变量 x 本身完全不受影响。这会造成迷惑，并且没有任何意义，所以是不合法的。**“可写性”就是为了避免这个问题而设计的。

这看起来很诡异，事实上并非如此，而且类似的情况很常见。考虑下面这行代码：

```go
f(x)
```

代码中，我们把变量 x 的一个拷贝传递给函数，因此不期望它会改变 x 的值。**如果期望函数 f 能够修改变量 x，我们必须传递 x 的地址（即指向 x 的指针）给函数 f，**如下所示：

```go
f(&x)
```

反射的工作机制与此相同，**如果想通过反射修改变量 x，就要把想要修改的变量的指针传递给反射库**。

首先，像通常一样初始化变量 x，然后**创建一个指向它的反射对象**，命名为 p：

```go
    package main
    import (
        "fmt"
        "reflect"
    )
    func main() {
        var x float64 = 3.4
        p := reflect.ValueOf(&x) // Note: take the address of x.
        fmt.Println("type of p:", p.Type())
        fmt.Println("settability of p:", p.CanSet())
    }
```

运行结果如下：

```go
type of p: *float64
settability of p: false
```

反射对象 p 是不可写的，但是我们也不像修改 p，**事实上我们要修改的是 *p**。**为了得到 p 指向的数据，可以调用 Value 类型的 Elem 方法**。Elem 方法能够对指针进行“解引用”，然后将结果存储到反射 Value 类型对象 v 中

```go
    package main
    import (
        "fmt"
        "reflect"
    )
    func main() {
        var x float64 = 3.4
        p := reflect.ValueOf(&x) // Note: take the address of x.
        v := p.Elem()
        fmt.Println("settability of v:", v.CanSet())
    }
```

运行结果如下：

```go
settability of v: true
```

由于**变量 v 代表 x， 因此我们可以使用 v.SetFloat 修改 x 的值**：

```go
    package main
    import (
        "fmt"
        "reflect"
    )
    func main() {
        var x float64 = 3.4
        p := reflect.ValueOf(&x) // Note: take the address of x.
        v := p.Elem()
        v.SetFloat(7.1)
        fmt.Println(v.Interface())
        fmt.Println(x)
    }
```

运行结果如下：

```go
7.1
7.1
```

> 只需要记住：只要反射对象要修改它们表示的对象，就必须获取它们表示的对象的地址。

### 5. 结构体

我们一般使用反射修改结构体的字段，**只要有结构体的指针，我们就可以修改它的字段**。

 下面是一个解析结构体变量 t 的例子，**用结构体的地址创建反射变量，再修改它**。然后我们**对它的类型设置了 typeOfT，并用调用简单的方法迭代字段**。

```go
    package main
    import (
        "fmt"
        "reflect"
    )
    func main() {
        type T struct {
            A int
            B string
        }
        t := T{23, "skidoo"}
        s := reflect.ValueOf(&t).Elem()
        typeOfT := s.Type()
        for i := 0; i < s.NumField(); i++ {
            f := s.Field(i)
            fmt.Printf("%d: %s %s = %v\n", i,
                typeOfT.Field(i).Name, f.Type(), f.Interface())
        }
    }
```

运行结果如下：

```go
0: A int = 23
1: B string = skidoo
```

T 的**各个字段名之所以大写，是因为结构体中只有可导出的字段是“可设置”的**。

**因为 s 包含了一个可设置的反射对象，我们可以修改结构体字段**：

```go
    package main
    import (
        "fmt"
        "reflect"
    )
    func main() {
        type T struct {
            A int
            B string
        }
        t := T{23, "skidoo"}
        s := reflect.ValueOf(&t).Elem()
        s.Field(0).SetInt(77)
        s.Field(1).SetString("Sunset Strip")
        fmt.Println("t is now", t)
    }
```

运行结果如下：

```go
t is now {77 Sunset Strip}
```

## 三、通过反射获取值信息

### 1. 使用反射值对象包装任意值

Go语言中，使用 reflect.ValueOf() 函数获得值的反射值对象（reflect.Value）。书写格式如下：

```go
value := reflect.ValueOf(rawValue)
```

reflect.ValueOf 返回 reflect.Value 类型，包含有 rawValue 的值信息。**reflect.Value 与原值间可以通过值包装和值获取互相转化**。reflect.Value 是一些反射操作的重要类型，如反射调用函数。

### 2. 从反射值对象获取被包装的值

Go语言中可以通过 reflect.Value 重新获得原始值。

#### 2.1  从反射值对象（reflect.Value）中获取值的方法

可以通过下面几种方法从反射值对象 reflect.Value 中获取原值，如下表所示。

| 方法名                   | 说  明                                                       |
| ------------------------ | ------------------------------------------------------------ |
| Interface() interface {} | 将值以 interface{} 类型返回，可以通过类型断言转换为指定类型  |
| Int() int64              | 将值以 int 类型返回，所有有符号整型均可以此方式返回          |
| Uint() uint64            | 将值以 uint 类型返回，所有无符号整型均可以此方式返回         |
| Float() float64          | 将值以双精度（float64）类型返回，所有浮点数（float32、float64）均可以此方式返回 |
| Bool() bool              | 将值以 bool 类型返回                                         |
| Bytes() []bytes          | 将值以字节数组 []bytes 类型返回                              |
| String() string          | 将值以字符串类型返回                                         |

下面代码中，将整型变量中的值使用 reflect.Value 获取反射值对象（reflect.Value）。再通过 reflect.Value 的 Interface() 方法获得 interface{} 类型的原值，通过 int 类型对应的 reflect.Value 的 Int()  方法获得整型值。

```go
    package main
    import (
        "fmt"
        "reflect"
    )
    func main() {
        // 声明整型变量a并赋初值
        var a int = 1024
        // 获取变量a的反射值对象
        valueOfA := reflect.ValueOf(a)
        // 获取interface{}类型的值, 通过类型断言转换
        var getA int = valueOfA.Interface().(int)
        // 获取int64位的值, 强制类型转换为int类型
        var getA2 int = int(valueOfA.Int())
        fmt.Println(getA, getA2)
    }
```

代码输出如下：

```go
1024 1024
```

## 四、通过反射访问结构体成员的值

和反射的类型对象(reflect.Type)一样，反射值对象（reflect.Value）也提供**对结构体访问的方法**，通过这些方法可以**完成对结构体任意值的访问**，如下表所示。

| 方  法                                         | 备  注                                                       |
| ---------------------------------------------- | ------------------------------------------------------------ |
| Field(i int) Value                             | 根据索引，返回索引对应的结构体成员字段的反射值对象。当值不是结构体或索引超界时发生宕机 |
| NumField() int                                 | 返回结构体成员字段数量。当值不是结构体或索引超界时发生宕机   |
| FieldByName(name string) Value                 | 根据给定字符串返回字符串对应的结构体字段。没有找到时返回零值，当值不是结构体或索引超界时发生宕机 |
| FieldByIndex(index []int) Value                | 多层成员访问时，根据 []int 提供的每个结构体的字段索引，返回字段的值。 没有找到时返回零值，当值不是结构体或索引超界时发生宕机 |
| FieldByNameFunc(match func(string) bool) Value | 根据匹配函数匹配需要的字段。找到时返回零值，当值不是结构体或索引超界时发生宕机 |

下面代码构造一个结构体包含不同类型的成员。通过 reflect.Value 提供的成员访问函数，可以获得结构体值的各种数据。

反射访问结构体成员值：

```go
    package main
    import (
        "fmt"
        "reflect"
    )
    // 定义结构体
    type dummy struct {
        a int
        b string
        // 嵌入字段
        float32
        bool
        next *dummy
    }
    func main() {
        // 值包装结构体
        d := reflect.ValueOf(dummy{
                next: &dummy{},
        })
        // 获取字段数量
        fmt.Println("NumField", d.NumField())
        // 获取索引为2的字段(float32字段)
        floatField := d.Field(2)
        // 输出字段类型
        fmt.Println("Field", floatField.Type())
        // 根据名字查找字段
        fmt.Println("FieldByName(\"b\").Type", d.FieldByName("b").Type())
        // 根据索引查找值中, next字段的int字段的值
        fmt.Println("FieldByIndex([]int{4, 0}).Type()", d.FieldByIndex([]int{4, 0}).Type())
    }
```

代码说明如下：

- 第 9 行，定义结构体，结构体的每个字段的类型都不一样。
- 第 24 行，实例化结构体并包装为 reflect.Value 类型，成员中包含一个 *dummy 的实例。
- 第 29 行，获取结构体的字段数量。
- 第 32 和 35 行，获取索引为2的字段值（float32 字段），并且打印类型。
- 第 38 行，根据`b`字符串，查找到 b 字段的类型。
- 第 41 行，**[]int{4,0} 中的 4 表示，在 dummy 结构中索引值为 4 的成员，也就是 next。next 的类型为  dummy，也是一个结构体，因此使用 []int{4,0} 中的 0 继续在 next 值的基础上索引，结构为 dummy 中索引值为 0 的 a 字段，类型为 int。**

代码输出如下：

```go
NumField 5
Field float32
FieldByName("b").Type string
FieldByIndex([]int{4, 0}).Type() int
```

## 五、判断反射值的空和有效性

反射值对象（reflect.Value）提供一系列方法进行零值和空判定，如下表所示。

| 方 法          | 说 明                                                        |
| -------------- | ------------------------------------------------------------ |
| IsNil() bool   | 返回值是否为 nil。如果值类型不是通道（channel）、函数、接口、map、指针或 切片时发生 panic，类似于语言层的`v== nil`操作 |
| IsValid() bool | 判断值是否有效。 当值本身非法时，返回 false，例如 reflect Value不包含任何值，值为 nil 等。 |

下面的例子将会对各种方式的空指针进行 IsNil() 和 IsValid() 的返回值判定检测。**同时对结构体成员及方法查找 map 键值对的返回值进行 IsValid() 判定**，参考下面的代码。

**IsNil() 常被用于判断指针是否为空；IsValid() 常被用于判定返回值是否有效。**

反射值对象的零值和有效性判断：

```go
    package main
    import (
        "fmt"
        "reflect"
    )
    func main() {
        // *int的空指针
        var a *int
        fmt.Println("var a *int:", reflect.ValueOf(a).IsNil())
        // nil值
        fmt.Println("nil:", reflect.ValueOf(nil).IsValid())
        // *int类型的空指针
        fmt.Println("(*int)(nil):", reflect.ValueOf((*int)(nil)).Elem().IsValid())
        // 实例化一个结构体
        s := struct{}{}
        // 尝试从结构体中查找一个不存在的字段
        fmt.Println("不存在的结构体成员:", reflect.ValueOf(s).FieldByName("").IsValid())
        // 尝试从结构体中查找一个不存在的方法
        fmt.Println("不存在的结构体方法:", reflect.ValueOf(s).MethodByName("").IsValid())  //结构体方法也可以查询
        // 实例化一个map
        m := map[int]int{}   //此方法与make等效
        // 尝试从map中查找一个不存在的键
        fmt.Println("不存在的键：", reflect.ValueOf(m).MapIndex(reflect.ValueOf(3)).IsValid())//map可以查询键值对
    }
```

代码输出如下：

```go
var a *int: true
nil: false
(*int)(nil): false
不存在的结构体成员: false
不存在的结构体方法: false
不存在的键： false
```

代码说明如下：

- 第 11 行，声明一个 *int 类型的指针，初始值为 nil。
- 第 12 行，将变量 a 包装为 reflect.Value 并且判断是否为空，此时变量 a 为空指针，因此返回 true。
- 第 15 行，对 nil 进行 IsValid() 判定（有效性判定），返回 false。
- 第 18 行，(*int)(nil) 的含义是将 nil 转换为 * int，也就是* int 类型的空指针。此行将 nil 转换为 * int  类型，并取指针指向元素。由于 nil 不指向任何元素，* int 类型的 nil 也不能指向任何元素，值不是有效的。因此这个反射值使用  Isvalid() 判断时返回 false。
- 第 21 行，实例化一个结构体。
- 第 24 行，通过 FieldByName 查找 s 结构体中一个空字符串的成员，如成员不存在，IsValid() 返回 false。
- 第 27 行，通过 MethodByName 查找 s 结构体中一个空字符串的方法，如方法不存在，IsValid() 返回 false。
- **第 30 行，实例化一个 map，这种写法与 make 方式创建的 map 等效。**
- 第 33 行，**MapIndex() 方法能根据给定的 reflect.Value 类型的值查找 map，并且返回查找到的结果。**

## 六、通过反射修改变量的值

Go语言中类似 x、x.f[1] 和 *p 形式的表达式都可以表示变量，但是其它如 x + 1 和 f(2) 则不是变量。一个变量就是一个可寻址的内存空间，里面存储了一个值，并且存储的值可以通过内存地址来更新。

对于 reflect.Values 也有类似的区别。有一些 reflect.Values 是可取地址的；其它一些则不可以。考虑以下的声明语句：

```go
x := 2 // value type variable?
a := reflect.ValueOf(2) // 2 int no
b := reflect.ValueOf(x) // 2 int no
c := reflect.ValueOf(&x) // &x *int no
d := c.Elem() // 2 int yes (x)
```

其中 a 对应的变量则不可取地址。因为 a 中的值仅仅是整数 2 的拷贝副本。b 中的值也同样不可取地址。c  中的值还是不可取地址，它只是一个指针 &x 的拷贝。**实际上，所有通过 reflect.ValueOf(x) 返回的  reflect.Value 都是不可取地址的。**

但是对于 d，它是 c 的解引用方式生成的，指向另一个变量，因此是可取地址的。我们**可以通过调用  reflect.ValueOf(&x).Elem()，来获取任意变量x对应的可取地址的 Value。**

我们**可以通过调用 reflect.Value 的 CanAddr 方法来判断其是否可以被取地址**：

```go
fmt.Println(a.CanAddr()) // "false"
fmt.Println(b.CanAddr()) // "false"
fmt.Println(c.CanAddr()) // "false"
fmt.Println(d.CanAddr()) // "true"
```

> 下面这段话未理解

每当我们**通过指针间接地获取的 reflect.Value 都是可取地址的**，**即使开始的是一个不可取地址的  Value**。在反射机制中，所有关于是否支持取地址的规则都是类似的。**例如，slice 的索引表达式  e[i]将隐式地包含一个指针，它就是可取地址的，即使开始的e表达式不支持也没有关系。**

 以此类推，**reflect.ValueOf(e).Index(i) 对于的值也是可取地址的，即使原始的 reflect.ValueOf(e) 不支持也没有关系。**

### 1. 判定及获取元素的相关方法

使用 reflect.Value 取元素、取地址及修改值的属性方法请参考下表。

| 方法名         | 备  注                                                       |
| -------------- | ------------------------------------------------------------ |
| Elem() Value   | 取值指向的元素值，类似于语言层`*`操作。当值类型不是指针或接口时发生宕 机，空指针时返回 nil 的 Value |
| Addr() Value   | 对可寻址的值返回其地址，类似于语言层`&`操作。当值不可寻址时发生宕机 |
| CanAddr() bool | 表示值是否可寻址                                             |
| CanSet() bool  | 返回值能否被修改。要求值可寻址且是可导出的字段               |

### 2. 值修改相关方法

使用 reflect.Value 修改值的相关方法如下表所示。

| Set(x Value)        | 将值设置为传入的反射值对象的值                               |
| ------------------- | ------------------------------------------------------------ |
| Setlnt(x int64)     | 使用 int64 设置值。当值的类型不是 int、int8、int16、 int32、int64 时会发生宕机 |
| SetUint(x uint64)   | 使用 uint64 设置值。当值的类型不是 uint、uint8、uint16、uint32、uint64 时会发生宕机 |
| SetFloat(x float64) | 使用 float64 设置值。当值的类型不是 float32、float64 时会发生宕机 |
| SetBool(x bool)     | 使用 bool 设置值。当值的类型不是 bod 时会发生宕机            |
| SetBytes(x []byte)  | 设置字节数组 []bytes值。当值的类型不是 []byte 时会发生宕机   |
| SetString(x string) | 设置字符串值。当值的类型不是 string 时会发生宕机             |

以上方法，**在 reflect.Value 的 CanSet 返回 false 仍然修改值时会发生宕机**。

在**已知值的类型时，应尽量使用值对应类型的反射设置值。**

### 3. 值可修改条件之一：可被寻址

通过反射修改变量值的前提条件之一：这个值必须可以被寻址。简单地说就是这个变量必须能被修改。示例代码如下：

```go
    package main
    import (
        "reflect"
    )
    func main() {
        // 声明整型变量a并赋初值
        var a int = 1024
        // 获取变量a的反射值对象
        valueOfA := reflect.ValueOf(a)
        // 尝试将a修改为1(此处会发生崩溃)
        valueOfA.SetInt(1)
    }
```

程序运行崩溃，打印错误：

```go
panic: reflect: reflect.Value.SetInt using unaddressable value
```

报错意思是：SetInt 正在使用一个不能被寻址的值。**从 reflect.ValueOf 传入的是 a 的值，而不是 a 的地址，这个 reflect.Value 当然是不能被寻址的**。将代码修改一下，重新运行：

```go
    package main
    import (
        "fmt"
        "reflect"
    )
    func main() {
        // 声明整型变量a并赋初值
        var a int = 1024
        // 获取变量a的反射值对象(a的地址)
        valueOfA := reflect.ValueOf(&a)
        // 取出a地址的元素(a的值)
        valueOfA = valueOfA.Elem()
        // 修改a的值为1
        valueOfA.SetInt(1)
        // 打印a的值
        fmt.Println(valueOfA.Int())
    }
```

代码输出如下：

```go
1
```

### 4. 值可修改条件之二：被导出

结构体成员中，**如果字段没有被导出**，即便不使用反射也**可以被访问**，**但不能通过反射修改**，代码如下：

```go
package main
import (
	"reflect"
)
func main() {
	type dog struct {
		legCount int
	}
	// 获取dog实例的反射值对象
	valueOfDog := reflect.ValueOf(&dog{}).Elem()
	// 获取legCount字段的值
	vLegCount := valueOfDog.FieldByName("legCount")
	// 尝试设置legCount的值(这里会发生崩溃)
	vLegCount.SetInt(4)
}

```

程序发生崩溃，报错：

```go
panic: reflect: reflect.Value.SetInt using value obtained using unexported field
```

报错的意思是：SetInt() 使用的值来自于一个未导出的字段。

为了能修改这个值，需要将该字段导出。将 dog 中的 legCount 的成员首字母大写，导出 LegCount 让反射可以访问，修改后的代码如下：

```go
    type dog struct {
        LegCount int
    }
```

然后根据字段名获取字段的值时，将字符串的字段首字母大写，修改后的代码如下：

```go
vLegCount := valueOfDog.FieldByName("LegCount")
```

## 七、通过类型信息创建实例

当**已知 reflect.Type 时，可以动态地创建这个类型的实例，实例的类型为指针**。例如 reflect.Type 的类型为 int 时，**创建 int 的指针**，即`*int`，代码如下：

```go
    package main
    import (
        "fmt"
        "reflect"
    )
    func main() {
        var a int
        // 取变量a的反射类型对象
        typeOfA := reflect.TypeOf(a)
        // 根据反射类型对象创建类型实例
        aIns := reflect.New(typeOfA)
        // 输出Value的类型和种类
        fmt.Println(aIns.Type(), aIns.Kind())
    }
```

代码输出如下：

```go
*int ptr
```

代码说明如下：

- 第 13 行，获取变量 a 的反射类型对象。
- 第 16 行，使用 reflect.New() 函数传入变量 a 的反射类型对象，创建这个类型的实例值，值以 reflect.Value 类型返回。**这步操作等效于：new(int)，因此返回的是 *int 类型的实例。**
- 第 19 行，打印 aIns 的类型为 *int，种类为指针。

## 八、通过反射调用函数

如果**反射值对象（reflect.Value）中值的类型为函数**时，**可以通过 reflect.Value  调用该函数**。**使用反射调用函数时，需要将参数使用反射值对象的切片 []reflect.Value 构造后传入 Call()  方法中，调用完成时，函数的返回值通过 []reflect.Value 返回。**

下面的代码声明一个加法函数，传入两个整型值，返回两个整型值的和。将函数保存到反射值对象（reflect.Value）中，然后将两个整型值构造为反射值对象的切片（[]reflect.Value），使用 Call() 方法进行调用。

```go
    package main
    import (
        "fmt"
        "reflect"
    )
    // 普通函数
    func add(a, b int) int {
        return a + b
    }
    func main() {
        // 将函数包装为反射值对象
        funcValue := reflect.ValueOf(add)
        // 构造函数参数, 传入两个整型值
        paramList := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}
        // 反射调用函数
        retList := funcValue.Call(paramList)
        // 获取第一个返回值, 取整数值
        fmt.Println(retList[0].Int())
    }
```

### 1. 反射调用函数的使用提示

反射调用函数的过程需要构造大量的 reflect.Value  和中间变量，对函数参数值进行逐一检查，还需要将调用参数复制到调用函数的参数内存中。调用完毕后，还需要将返回值转换为  reflect.Value，用户还需要从中取出调用值。因此，**反射调用函数的性能问题尤为突出，不建议大量使用反射函数调用。**