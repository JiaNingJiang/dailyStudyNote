### 一、包的基本概念

包可以定义在很深的目录中，包名的定义是不包括目录路径的，但是包在引用时一般使用全路径引用。比如在`GOPATH/src/a/b/ `下定义一个包 c。在包 c 的源码中只需声明为`package c`，而不是声明为`package a/b/c`，但是在导入 c 包时，需要带上路径，例如`import "a/b/c"`。

包的习惯用法：

- 包名一般是小写的，使用一个简短且有意义的名称。
- 包名一般要和所在的目录同名，也可以不同，包名中不能包含`- `等特殊符号。
- 包一般使用域名作为目录名称，这样能保证包名的唯一性，比如 GitHub 项目的包一般会放到`GOPATH/src/github.com/userName/projectName `目录下。
- 包名为 main 的包为应用程序的入口包，编译不包含 main 包的源码文件时不会得到可执行文件。
- 一个文件夹下的所有源码文件只能属于同一个包，同样属于同一个包的源码文件不能放在多个文件夹下。

### 二、包的导入

要在代码中引用其他包的内容，需要使用 import 关键字导入使用的包。具体语法如下：

```go
import "包的路径"
```

注意事项：

- import 导入语句通常放在源码文件开头包声明语句的下面；
- 导入的包名需要使用双引号包裹起来；
- 包名是从`GOPATH/src/ `后开始计算的（因此我们不需要手动添加GOPATH/src/），使用`/ `进行路径分隔。

### 三、包的导入路径

包的引用路径有两种写法，分别是全路径导入和相对路径导入。

#### 全路径导入

包的绝对路径就是`GOROOT/src/`或`GOPATH/src/`后面包的存放路径，如下所示：

```go
import "lab/test"
import "database/sql/driver"
import "database/sql"
```

上面代码的含义如下：

- test 包是自定义的包，其源码位于`GOPATH/src/lab/test `目录下；
- driver 包的源码位于`GOROOT/src/database/sql/driver `目录下；
- sql 包的源码位于`GOROOT/src/database/sql `目录下。

#### 相对路径导入

相对路径==只能用于导入`GOPATH `下的包==，标准包的导入只能使用全路径导入。

例如包 a 的所在路径是`GOPATH/src/lab/a`，包 b 的所在路径为`GOPATH/src/lab/b`，如果在包 b 中导入包 a ，则可以使用相对路径导入方式。示例如下：

```go
// 相对路径导入
import "../a"
```

当然了，也可以使用上面的全路径导入，如下所示：

```go
// 全路径导入
import "lab/a"
```

### 四、包的引用格式

包的引用有四种格式，下面以 fmt 包为例来分别演示一下这四种格式。

#### 1) 标准引用格式

```go
import "fmt"
```

#### 2) 自定义别名引用格式

在导入包的时候，我们还可以为导入的包设置别名，如下所示：

```go
import F "fmt"
```

其中 F 就是 fmt 包的别名，使用时我们可以使用`F.`来代替标准引用格式的`fmt.`来作为前缀使用 fmt 包中的方法。

#### 3) 省略引用格式

```go
import . "fmt"
```

这种格式相当于把 fmt 包直接合并到当前程序中，在使用 fmt 包内的方法是可以不用加前缀`fmt.`，直接引用

示例代码如下：

```go
    package main
    import . "fmt"
    func main() {
        //不需要加前缀 fmt.
        Println("C语言中文网")
    }
```

#### 4) 匿名引用格式

在引用某个包时，如果只是希望执行包初始化的 init 函数，而不使用包内部的数据时，可以使用匿名引用格式，如下所示：

```go
import _ "fmt"
```

匿名导入的包与其他方式导入的包一样都会被编译到可执行文件中。

使用标准格式引用包，但是代码中却没有使用包，编译器会报错。如果包中有 init 初始化函数，则通过`import _ "包的路径" `这种方式引用包，仅执行包的初始化函数，即使包没有 init 初始化函数，也不会引发编译器报错。

注意：

- 一个包可以有多个 init 函数，包加载时会执行全部的 init 函数，但并不能保证执行顺序，所以不建议在一个包中放入多个 init 函数，将需要初始化的逻辑放到一个 init 函数里面。
- 包不能出现环形引用（交叉引用）的情况，比如包 a 引用了包 b，包 b 引用了包 c，如果包 c 又引用了包 a，则编译不能通过。
- 包的重复引用是允许的，比如包 a 引用了包 b 和包 c，包 b 和包 c 都引用了包 d。这种场景相当于重复引用了 d，这种情况是允许的，并且 Go 编译器保证包 d 的 init 函数只会执行一次。

### 五、包加载

在执行 main 包的 mian 函数之前， Go 引导程序会先对整个程序的包进行初始化。整个执行的流程如下图所示。

![image-20220923205133653](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\image-20220923205133653.png)

Go语言包的初始化有如下特点：

- 包初始化程序从 main 函数引用的包开始，逐级查找包的引用，直到找到没有引用其他包的包，最终生成一个包引用的有向无环图。
- Go 编译器会将有向无环图转换为一棵树，然后从树的叶子节点开始逐层向上对包进行初始化。
- 单个包的初始化过程如上图所示，先初始化常量，然后是全局变量，最后执行包的 init 函数。

### 六、GOPATH详解（Go语言工作目录）

GOPATH 是 Go语言中使用的一个环境变量，它使用绝对路径提供项目的工作目录。

在 Go 1.8 版本之前，GOPATH 环境变量默认是空的。从 Go 1.8 版本开始，Go 开发包在安装完成后，将 GOPATH 赋予了一个默认的目录，参见下表。

| 平  台       | GOPATH 默认值    | 举 例              |
| ------------ | ---------------- | ------------------ |
| Windows 平台 | %USERPROFILE%/go | C:\Users\用户名\go |
| Unix 平台    | $HOME/go         | /home/用户名/go    |

#### 1.使用GOPATH的工程结构

在 GOPATH 指定的工作目录下，代码总是会保存在 $GOPATH/src 目录下。在工程经过 go build、go install 或  go get 等指令后，会将产生的二进制可执行文件放在 $GOPATH/bin 目录下，生成的中间缓存文件会被保存在 $GOPATH/pkg  下。

 如果需要将整个源码添加到版本管理工具（Version Control System，VCS）中时，只需要添加 $GOPATH/src 目录的源码即可。bin 和 pkg 目录的内容都可以由 src 目录生成。

#### 2.设置和使用GOPATH

以 Linux 为演示平台，演示使用 GOPATH 的方法。

##### 1) 设置当前目录为GOPATH

选择一个目录，在目录中的命令行中执行下面的指令：

```shell
export GOPATH=`pwd`
```

该指令中的 pwd 将输出当前的目录，使用反引号```将 pwd 指令括起来表示命令行替换，也就是说，使用``pwd``将获得 pwd 返回的当前目录的值。例如，假设你的当前目录是“/home/davy/go”，那么使用``pwd``将获得返回值“/home/davy/go”。

使用 export 指令可以将当前目录的值设置到环境变量 GOPATH中。

##### 2) 建立GOPATH中的源码目录

使用下面的指令创建 GOPATH 中的 src 目录，在 src 目录下还有一个 hello 目录，该目录用于保存源码。

```shell
mkdir -p src/hello
```

mkdir 指令的 -p 可以连续创建一个路径。

##### 3) 添加main.go源码文件

##### 4) 编译源码并运行

此时我们已经设定了 GOPATH，因此在 Go语言中可以通过 GOPATH 找到工程的位置。

 在命令行中执行如下指令编译源码：

```go
go install hello
```

编译完成的可执行文件会保存在 $GOPATH/bin 目录下。在 bin 目录中执行 ./hello，命令行输出如下：
 ```go
 hello world
 ```

#### 3.在多项目工程中使用GOPATH

在很多与 Go语言相关的书籍、文章中描述的 GOPATH 都是通过修改系统全局的环境变量来实现的。但这种设置全局 GOPATH 的方法可能会导致当前项目错误引用了其他目录的 Go 源码文件从而造成编译输出错误的版本或编译报出一些无法理解的错误提示。

比如说，将某项目代码保存在 /home/davy/projectA 目录下，将该目录设置为  GOPATH。随着开发进行，需要再次获取一份工程项目的源码，此时源码保存在 /home/davy/projectB 目录下，如果此时需要编译  projectB 目录的项目，但开发者忘记设置 GOPATH 而直接使用命令行编译，则当前的 GOPATH 指向的是  /home/davy/projectA 目录，而不是开发者编译时期望的 projectB  目录。编译完成后，开发者就会将错误的工程版本发布到外网。

因此，建议无论是使用命令行或者使用集成开发环境编译 Go 源码时，GOPATH 跟随项目设定。在 Jetbrains 公司的 GoLand 集成开发环境（IDE）中的 GOPATH 设置分为全局 GOPATH 和项目 GOPATH，如下图所示。

![image-20220926194348712](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\image-20220926194348712.png)

图中的 Global GOPATH 代表全局 GOPATH，一般来源于系统环境变量中的 GOPATH；Project GOPATH  代表项目所使用的 GOPATH，该设置会被保存在工作目录的 .idea 目录下，不会被设置到环境变量的 GOPATH  中，但会在编译时使用到这个目录。建议在开发时只填写项目 GOPATH，每一个项目尽量只设置一个 GOPATH，不使用多个 GOPATH 和全局的 GOPATH。

### 七、常用内置包简介

#### 1) fmt

fmt 包实现了格式化的标准输入输出，这与C语言中的 printf 和 scanf 类似。其中的 fmt.Printf() 和 fmt.Println() 是开发者使用最为频繁的函数。

 格式化短语派生于C语言，一些短语（%- 序列）是这样使用：

- %v：默认格式的值。当打印结构时，加号（%+v）会增加字段名；
- %#v：Go样式的值表达；
- %T：带有类型的 Go 样式的值表达。

#### 2) io

这个包提供了原始的 I/O 操作界面。它主要的任务是对 os 包这样的原始的 I/O 进行封装，增加一些其他相关，使其具有抽象功能用在公共的接口上。

#### 3) bufio

bufio 包通过对 io 包的封装，提供了数据缓冲功能，能够一定程度减少大块数据读写带来的开销。

 在 bufio 各个组件内部都维护了一个缓冲区，数据读写操作都直接通过缓存区进行。当发起一次读写操作时，会首先尝试从缓冲区获取数据，只有当缓冲区没有数据时，才会从数据源获取数据更新缓冲。

#### 4) sort

sort 包提供了用于对切片和用户定义的集合进行排序的功能。

#### 5) strconv

strconv 包提供了将字符串转换成基本数据类型，或者从基本数据类型转换为字符串的功能。

#### 6) os

os 包提供了不依赖平台的操作系统函数接口，设计像 Unix 风格，但错误处理是 go 风格，当 os 包使用时，如果失败后返回错误类型而不是错误数量。

#### 7) sync

sync 包实现多线程中锁机制以及其他同步互斥机制。

#### 8) flag

flag 包提供命令行参数的规则定义和传入参数解析的功能。绝大部分的命令行程序都需要用到这个包。

#### 9) encoding/json

JSON 目前广泛用做网络程序中的通信格式。encoding/json 包提供了对 JSON 的基本支持，比如从一个对象序列化为 JSON 字符串，或者从 JSON 字符串反序列化出一个具体的对象等。

#### 10) html/template

主要实现了 web 开发中生成 html 的 template 的一些函数。

#### 11) net/http

net/http 包提供 HTTP 相关服务，主要包括 http 请求、响应和 URL 的解析，以及基本的 http 客户端和扩展的 http 服务。

 通过 net/http 包，只需要数行代码，即可实现一个爬虫或者一个 Web 服务器，这在传统语言中是无法想象的。

#### 12) reflect

reflect 包实现了运行时反射，允许程序通过抽象类型操作对象。通常用于处理静态类型 interface{} 的值，并且通过 Typeof 解析出其动态类型信息，通常会返回一个有接口类型 Type 的对象。

#### 13) os/exec

os/exec 包提供了执行自定义 linux 命令的相关实现。

#### 14) strings

strings 包主要是处理字符串的一些函数集合，包括合并、查找、分割、比较、后缀检查、索引、大小写处理等等。

 strings 包与 bytes 包的函数接口功能基本一致。

#### 15) bytes

bytes 包提供了对字节切片进行读写操作的一系列函数。字节切片处理的函数比较多，分为基本处理函数、比较函数、后缀检查函数、索引函数、分割函数、大小写处理函数和子切片处理函数等。

#### 16) log

log 包主要用于在程序中输出日志。

 log 包中提供了三类日志输出接口，Print、Fatal 和 Panic。

- Print 是普通输出；
- Fatal 是在执行完 Print 后，执行 os.Exit(1)；
- Panic 是在执行完 Print 后调用 panic() 方法。

### 八、import导入包——在代码中使用其他的代码

#### 1.导入包后自定义引用的包名

如果我们想同时导入两个有着名字相同的包，例如 math/rand 包和 crypto/rand 包，那么导入声明必须至少为一个同名包指定一个新的包名以避免冲突。这叫做导入包的重命名。

```go
    import (
        "crypto/rand"
        mrand "math/rand" // 将名称替换为mrand避免冲突
    )
```

导入包的重命名只影响当前的源文件。其它的源文件如果导入了相同的包，可以用导入包原本默认的名字或重命名为另一个完全不同的名字。

 导入包重命名是一个有用的特性，它不仅仅只是为了解决名字冲突。如果导入的一个包名很笨重，特别是在一些自动生成的代码中，这时候用一个简短名称会更方便。选择用简短名称重命名导入包时候最好统一，以避免包名混乱。选择另一个包名称还可以帮助避免和本地普通变量名产生冲突。例如，如果文件中已经有了一个名为 path 的变量，那么我们可以将"path"标准包重命名为 pathpkg。

 每个导入声明语句都明确指定了当前包和被导入包之间的依赖关系。如果遇到包循环导入的情况，Go语言的构建工具将报告错误。

#### 2.包在程序启动前的初始化入口：init

在某些需求的设计上需要在程序启动时统一调用程序引用到的所有包的初始化函数，如果需要通过开发者手动调用这些初始化函数，那么这个过程可能会发生错误或者遗漏。我们希望在被引用的包内部，由包的编写者获得代码启动的通知，在程序启动时做一些自己包内代码的初始化工作。

 例如，为了提高数学库计算三角函数的执行效率，可以在程序启动时，将三角函数的值提前在内存中建成索引表，外部程序通过查表的方式迅速获得三角函数的值。但是三角函数索引表的初始化函数的调用不希望由每一个外部使用三角函数的开发者调用，如果在三角函数的包内有一个机制可以告诉三角函数包程序何时启动，那么就可以解决初始化的问题。

 Go 语言为以上问题提供了一个非常方便的特性：init() 函数。

 init() 函数的特性如下：

- 每个源码可以使用 1 个 init() 函数。
- init() 函数会在程序执行前（main() 函数执行前）被自动调用。
- 调用顺序为 main() 中引用的包，以深度优先顺序初始化。


 例如，假设有这样的包引用关系：main→A→B→C，那么这些包的 init() 函数调用顺序为：

```go
C.init→B.init→A.init→main
```

说明：

- 同一个包中的多个 init() 函数的调用顺序不可预期。
- init() 函数不能被其他函数调用

### 九、sync包与锁：限制线程对变量的访问

#### 读写锁

读写锁有如下四个方法：

- 写操作的锁定和解锁分别是`func (*RWMutex) Lock`和`func (*RWMutex) Unlock`；

- 读操作的锁定和解锁分别是`func (*RWMutex) Rlock`和`func (*RWMutex) RUnlock`。

读写锁的区别在于：

- 当有一个 goroutine 获得写锁定，其它无论是读锁定还是写锁定都将阻塞直到写解锁；
- 当有一个 goroutine 获得读锁定，其它读锁定仍然可以继续；
- 当有一个或任意多个读锁定，写锁定将等待所有读锁定解锁之后才能够进行写锁定。



所以说这里的读锁定（RLock）目的其实是告诉写锁定，有很多协程或者进程正在读取数据，写操作需要等它们读（读解锁）完才能进行写（写锁定）。

 我们可以将其总结为如下三条：

- 同时只能有一个 goroutine 能够获得写锁定；
- 同时可以有任意多个 gorouinte 获得读锁定；
- 同时只能存在写锁定或读锁定（读和写互斥）。

### 十、big包：对整数的高精度计算

实际开发中，对于超出 int64 或者 uint64 类型的大数进行计算时，如果对精度没有要求，使用 float32 或者 float64 就可以胜任，但如果对精度有严格要求的时候，我们就不能使用浮点数了，因为浮点数在内存中只能被近似的表示。

Go语言中 math/big 包实现了大数字的多精度计算，支持 Int（有符号整数）、Rat（有理数）和 Float（浮点数）等数字类型。

这些类型可以实现任意位数的数字，只要内存足够大，但缺点是需要更大的内存和处理开销，这使得它们使用起来要比内置的数字类型慢很多。

#### 1.在 math/big 包中，Int 类型定义如下所示：

```go
    // An Int represents a signed multi-precision integer.
    // The zero value for an Int represents the value 0.
    type Int struct {
        neg bool // sign
        abs nat  // absolute value of the integer
    }
```

生成 Int 类型的方法为 NewInt()，如下所示：

```go
    // NewInt allocates and returns a new Int set to x.
    func NewInt(x int64) *Int {
        return new(Int).SetInt64(x)
    }
```

```go
注意：NewInt() 函数只对 int64 有效，其他类型必须先转成 int64 才行。
```

Go语言中还提供了许多 Set 函数，可以方便的把其他类型的整形存入 Int ，因此，我们可以先 new(int) 然后再调用 Set 函数，Set 函数有如下几种：

```go
    // SetInt64 函数将 z 转换为 x 并返回 z。
    func (z *Int) SetInt64(x int64) *Int {
        neg := false
        if x < 0 {
            neg = true
            x = -x
        }
        z.abs = z.abs.setUint64(uint64(x))
        z.neg = neg
        return z
    }
    
    // SetUint64 函数将 z 转换为 x 并返回 z。
    func (z *Int) SetUint64(x uint64) *Int {
        z.abs = z.abs.setUint64(x)
        z.neg = false
        return z
    }
    
    // Set 函数将 x 转换为Int格式的 z 并返回 z。
    func (z *Int) Set(x *Int) *Int {
        if z != x {
            z.abs = z.abs.set(x.abs)
            z.neg = x.neg
        }
        return z
    }
```

示例代码如下所示：

```go
    package main
    import (
        "fmt"
        "math/big"
    )
    func main() {
        big1 := new(big.Int).SetUint64(uint64(1000))
        fmt.Println("big1 is: ", big1)
        big2 := big1.Uint64()
        fmt.Println("big2 is: ", big2)
    }
```

运行结果如下：

```go
big1 is:  1000
big2 is:  1000
```

#### 2.除了上述的 Set 函数，math/big 包中还提供了一个 SetString() 函数，可以指定进制数，比如二进制、十进制或者十六进制等！

```go
    // SetString将z设置为s的值，以给定的基数进行解释，并返回z和表示成功的布尔值。
	//整个字符串（不仅仅是前缀）必须有效才能成功。
	//如果SetString失败，则z的值未定义，但返回的值为nil。
    //
    // The base argument must be 0 or a value between 2 and MaxBase. If the base
    // is 0, the string prefix determines the actual conversion base. A prefix of
    // ``0x'' or ``0X'' selects base 16; the ``0'' prefix selects base 8, and a
    // ``0b'' or ``0B'' prefix selects base 2. Otherwise the selected base is 10.
    //
    func (z *Int) SetString(s string, base int) (*Int, bool) {
        r := strings.NewReader(s)
        if _, _, err := z.scan(r, base); err != nil {
            return nil, false
        }
        // entire string must have been consumed
        if _, err := r.ReadByte(); err != io.EOF {
            return nil, false
        }
        return z, true // err == io.EOF => scan consumed all of s
    }
```

示例代码如下所示：

```go
    package main
    import (
        "fmt"
        "math/big"
    )
    func main() {
        big1, _ := new(big.Int).SetString("1000", 10)
        fmt.Println("big1 is: ", big1)
        big2 := big1.Uint64()
        fmt.Println("big2 is: ", big2)
    }
```

运行结果如下：

```go
big1 is:  1000
big2 is:  1000
```

因为Go语言不支持运算符重载，所以所有大数字类型都有像是 Add() 和 Mul() 这样的方法。Add 方法的定义如下所示：

```go
func (z *Int) Add(x, y *Int) *Int   //该方法会将 z 转换为 x + y 并返回 z。
```

【示例】计算第 1000 位的斐波那契数列。

```go
    package main
    import (
        "fmt"
        "math/big"
        "time"
    )
    const LIM = 1000 //求第1000位的斐波那契数列
    var fibs [LIM]*big.Int //使用数组保存计算出来的数列的指针
    func main() {
        result := big.NewInt(0)
        start := time.Now()
        for i := 0; i < LIM; i++ {
            result = fibonacci(i)
            fmt.Printf("数列第 %d 位: %d\n", i+1, result)
        }
        end := time.Now()
        delta := end.Sub(start)
        fmt.Printf("执行完成，所耗时间为: %s\n", delta)
    }
    func fibonacci(n int) (res *big.Int) {
        if n <= 1 {
            res = big.NewInt(1)
        } else {
            temp := new(big.Int)
            res = temp.Add(fibs[n-1], fibs[n-2])
        }
        fibs[n] = res
        return
    }
```

运行结果如下：

```go
数列第 1 位: 1
数列第 2 位: 1
数列第 3 位: 2
数列第 4 位: 3
数列第 5 位: 5
...
数列第 997 位: 10261062362033262336604926729245222132668558120602124277764622905699407982546711488272859468887457959
08773311924256407785074365766118082732679853917775891982813511440749936979646564952426675539110499009
9120377
数列第 998 位: 16602747662452097049541800472897701834948051198384828062358553091918573717701170201065510185595898605
10409473691887927846223301598102952299783631123261876053919903676539979992673143323971886037334508837
5054249
数列第 999 位: 26863810024485359386146727202142923967616609318986952340123175997617981700247881689338369654483356564
19182785616144335631297667364221035032463485041037768036733415117289916972319708276398561576445007847
4174626
数列第 1000 位: 4346655768693745643568852767504062580256466051737178040248172908953655541794905189040387984007925516
92959225930803226347752096896232398733224711616429964409065331879382989696499285160037044761377951668
49228875
执行完成，所耗时间为: 6.945ms
```

### 十一、time包：时间和日期

时间一般包含==时间值==和==时区==，可以从Go语言中 time 包的源码中看出：

```go
    type Time struct {
        // wall and ext encode the wall time seconds, wall time nanoseconds,
        // and optional monotonic clock reading in nanoseconds.
        //
        // From high to low bit position, wall encodes a 1-bit flag (hasMonotonic),
        // a 33-bit seconds field, and a 30-bit wall time nanoseconds field.
        // The nanoseconds field is in the range [0, 999999999].
        // If the hasMonotonic bit is 0, then the 33-bit field must be zero
        // and the full signed 64-bit wall seconds since Jan 1 year 1 is stored in ext.
        // If the hasMonotonic bit is 1, then the 33-bit field holds a 33-bit
        // unsigned wall seconds since Jan 1 year 1885, and ext holds a
        // signed 64-bit monotonic clock reading, nanoseconds since process start.
        wall uint64
        ext  int64
        // loc specifies the Location that should be used to
        // determine the minute, hour, month, day, and year
        // that correspond to this Time.
        // The nil location means UTC.
        // All UTC times are represented with loc==nil, never loc==&utcLoc.
        loc *Location
    }
```

上面代码中：

- wall：表示距离公元 1 年 1 月 1 日 00:00:00UTC 的秒数；
- ext：表示纳秒；
- loc：代表时区，主要处理偏移量，不同的时区，对应的时间不一样。

##### 如何正确表示时间呢？

公认最准确的计算应该是使用“原子震荡周期”所计算的物理时钟了（Atomic Clock, 也被称为原子钟），这也被定义为标准时间（International Atomic Time）。我们常常看见的 UTC（Universal Time Coordinated，世界协调时间）就是利用这种 Atomic Clock  为基准所定义出来的正确时间。UTC 标准时间是以 GMT（Greenwich Mean Time，格林尼治时间）这个时区为主，所以本地时间与  UTC 时间的时差就是本地时间与 GMT 时间的时差。

```go
UTC + 时区差 ＝ 本地时间
```

国内一般使用的是北京时间，与 UTC 的时间关系如下：

```go
UTC + 8 个小时 = 北京时间
```

在Go语言的 time 包里面有两个时区变量，如下：

- time.UTC：UTC 时间
- time.Local：本地时间

同时，Go语言还提供了 LoadLocation 方法和 FixedZone 方法来获取时区，如下：

```go
FixedZone(name string, offset int) *Location    //name 为时区名称，offset 是与 UTC 之前的时差。
```

```go
LoadLocation(name string) (*Location, error)  	//name 为时区的名字
```

#### 1.时间的获取

##### 1) 获取当前时间

通过 time.Now() 函数来获取当前的时间对象，然后通过事件对象来获取当前的时间信息。示例代码如下：

```go
    package main
    import (
        "fmt"
        "time"
    )
    func main() {
        now := time.Now() //获取当前时间
        fmt.Printf("current time:%v\n", now)
        year := now.Year()     //年
        month := now.Month()   //月
        day := now.Day()       //日
        hour := now.Hour()     //小时
        minute := now.Minute() //分钟
        second := now.Second() //秒
        fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
    }
```

运行结果如下：

```go
current time:2019-12-12 12:33:19.4712277 +0800 CST m=+0.006980401
2019-12-12 12:33:19
```

##### 2) 获取时间戳

时间戳是自 1970 年 1 月 1 日（08:00:00GMT）至当前时间的总毫秒数，它也被称为 Unix 时间戳（UnixTimestamp）。

 基于时间对象获取时间戳的示例代码如下：

```go
    package main
    import (
        "fmt"
        "time"
    )
    func main() {
        now := time.Now()            //获取当前时间
        timestamp1 := now.Unix()     //时间戳
        timestamp2 := now.UnixNano() //纳秒时间戳
        fmt.Printf("现在的时间戳：%v\n", timestamp1)
        fmt.Printf("现在的纳秒时间戳：%v\n", timestamp2)
    }
```

运行结果如下：

```go
现在的时间戳：1576127858
现在的纳秒时间戳：1576127858829900100
```

使用 time.Unix() 函数可以将时间戳转为时间格式，示例代码如下：

```go
    package main
    import (
        "fmt"
        "time"
    )
    func main() {
        now := time.Now()                  //获取当前时间
        timestamp := now.Unix()            //时间戳
        timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
        fmt.Println(timeObj)
        year := timeObj.Year()     //年
        month := timeObj.Month()   //月
        day := timeObj.Day()       //日
        hour := timeObj.Hour()     //小时
        minute := timeObj.Minute() //分钟
        second := timeObj.Second() //秒
        fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
    }
```

运行结果如下：

```go
2019-12-12 13:24:09 +0800 CST
2019-12-12 13:24:09
```

##### 3) 获取当前是星期几

time 包中的 Weekday 函数能够返回某个时间点所对应是一周中的周几，示例代码如下：

```go
    package main
    import (
        "fmt"
        "time"
    )
    func main() {
        //时间戳
        t := time.Now()
        fmt.Println(t.Weekday().String())
    }
```

运行结果如下：

```go
Thursday
```

#### 2.时间操作函数`

##### 1) Add

在日常的开发过程中可能会遇到要求某个时间 + 时间间隔之类的需求，Go语言中的 Add 方法如下：

```go
func (t Time) Add(d Duration) Time
```

【示例】求一个小时之后的时间：

```go
    package main
    import (
        "fmt"
        "time"
    )
    func main() {
        now := time.Now()
        later := now.Add(time.Hour) // 当前时间加1小时后的时间
        fmt.Println(later)
    }
```

##### 2) Sub

```go
func (t Time) Sub(u Time) Duration    //求两个时间之间的差值
```

返回一个时间段 t - u 的值。如果结果超出了 Duration 可以表示的最大值或最小值，将返回最大值或最小值，要获取时间点 t - d（d 为 Duration），可以使用 t.Add(-d)。

##### 3) Equal

```go
func (t Time) Equal(u Time) bool    //判断两个时间是否相同
```

Equal 函数会考虑时区的影响，因此不同时区标准的时间也可以正确比较，Equal 方法和用 t==u 不同，Equal 方法还会比较地点和时区信息。

##### 4) Before

判断一个时间点是否在另一个时间点之前：

```go
func (t Time) Before(u Time) bool    //如果 t 代表的时间点在 u 之前，则返回真，否则返回假。
```

##### 5) After

判断一个时间点是否在另一个时间点之后：

```go
func (t Time) After(u Time) bool    //如果 t 代表的时间点在 u 之后，则返回真，否则返回假。
```

#### 3.定时器

使用 time.Tick(时间间隔) 可以设置定时器，定时器的本质上是一个通道（channel），示例代码如下：

```go
    package main
    import (
        "fmt"
        "time"
    )
    func main() {
        ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
        for i := range ticker {      //管道也能range
            fmt.Println(i) //每秒都会执行的任务
        }
    }
```

运行结果如下：

```go
2019-12-12 15:14:26.4158067 +0800 CST m=+16.007460701
2019-12-12 15:14:27.4159467 +0800 CST m=+17.007600701
2019-12-12 15:14:28.4144689 +0800 CST m=+18.006122901
2019-12-12 15:14:29.4159581 +0800 CST m=+19.007612101
2019-12-12 15:14:30.4144337 +0800 CST m=+20.006087701
...
```

#### 4.时间格式化

时间类型有一个自带的 Format 方法进行格式化，需要注意的是Go语言中格式化时间模板不是常见的`Y-m-d H:M:S `而是使用Go语言的诞生时间 2006 年 1 月 2 号 15 点 04 分 05 秒。

> 提示：如果想将时间格式化为 12 小时格式，需指定 PM。

```go
    package main
    import (
        "fmt"
        "time"
    )
    func main() {
        now := time.Now()
        // 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
        // 24小时制
        fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
        // 12小时制
        fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
        fmt.Println(now.Format("2006/01/02 15:04"))
        fmt.Println(now.Format("15:04 2006/01/02"))
        fmt.Println(now.Format("2006/01/02"))
    }
```

运行结果如下：

```go
2019-12-12 15:20:52.037 Thu Dec
2019-12-12 03:20:52.037 PM Thu Dec
2019/12/12 15:20
15:20 2019/12/12
2019/12/12
```

#### 5.解析字符串格式的时间

Parse 函数可以解析一个格式化的时间字符串并返回它代表的时间。

```go
func Parse(layout, value string) (Time, error)   //layout为时间格式，value为要解析的时间字符串
```

与 Parse 函数类似的还有 ParseInLocation 函数。

```go
func ParseInLocation(layout, value string, loc *Location) (Time, error)
```

ParseInLocation 与 Parse 函数类似，但有两个重要的不同之处：

- 第一，当缺少时区信息时，Parse 将时间解释为 UTC 时间，而 ParseInLocation 将返回值(Time类对象)的 Location 设置为 loc；
- 第二，当时间字符串提供了时区偏移量信息时，Parse 会尝试去匹配本地时区，而 ParseInLocation 会去匹配 loc

示例代码如下：

```go
    package main
    import (
        "fmt"
        "time"
    )
    func main() {
        var layout string = "2006-01-02 15:04:05"    //格式
        var timeStr string = "2019-12-12 15:22:12"   //待解析的时间字符串
        timeObj1, _ := time.Parse(layout, timeStr)   //解析获得的Time对象
        fmt.Println(timeObj1)
        timeObj2, _ := time.ParseInLocation(layout, timeStr, time.Local)
        fmt.Println(timeObj2)
    }
```

运行结果如下：

```go
2019-12-12 15:22:12 +0000 UTC
2019-12-12 15:22:12 +0800 CST
```

### 十二、os包用法简述

#### 1. os包的常用函数

##### 1) Hostname

函数定义:

```go
func Hostname() (name string, err error)   //Hostname 函数会返回内核提供的主机名。
```

##### 2) Environ

函数定义：

```go
func Environ() []string    //Environ 函数会返回所有的环境变量，返回值格式为“key=value”的字符串的切片拷贝。
```

##### 3) Getenv
函数定义：

```go
func Getenv(key string) string   //Getenv 函数会检索并返回名为 key 的环境变量的值。如果不存在该环境变量则会返回空字符串。
```

##### 4) Setenv
函数定义：

```go
func Setenv(key, value string) error  //Setenv 函数可以设置名为 key 的环境变量，如果出错会返回该错误。
```

##### 5) Exit
函数定义：

```go
func Exit(code int)   //Exit 函数可以让当前程序以给出的状态码 code 退出。一般来说，状态码 0 表示成功，非 0 表示出错。程序会立刻终止，并且 defer 的函数不会被执行。
```

##### 6) Getuid
函数定义：

```go
func Getuid() int   //Getuid 函数可以返回调用者的用户 ID。
```

##### 7) Getgid
函数定义：

```go
func Getgid() int   //Getgid 函数可以返回调用者的组 ID。
```

##### 8) Getpid
函数定义：

```go
func Getpid() int  //Getpid 函数可以返回调用者所在进程的进程 ID。
```

##### 9) Getwd
函数定义：

```go
func Getwd() (dir string, err error)   //返回一个对应当前工作目录的根路径。如果当前目录可以经过多条路径抵达（因为硬链接），Getwd 会返回其中一个。
```

##### 10) Mkdir
函数定义：

```go
func Mkdir(name string, perm FileMode) error   //使用指定的权限和名称创建一个目录。如果出错，会返回 *PathError 底层类型的错误。
```

##### 11) Exit
函数定义：

```go
func Exit(code int)   //Exit 函数可以让当前程序以给出的状态码 code 退出。一般来说，状态码 0 表示成功，非 0 表示出错。程序会立刻终止，并且 defer 的函数不会被执行。
```

##### 12) Remove
函数定义：

```go
func Remove(name string) error //Remove 函数会删除 name 指定的文件或目录。如果出错，会返回 *PathError 底层类型的错误。RemoveAll 函数跟 Remove 用法一样，区别是会递归的删除所有子目录和文件。
```

#### 2. os/exec 执行外部命令

exec 包可以执行外部命令，它包装了 os.StartProcess 函数以便更容易的修正输入和输出，使用管道连接 I/O，以及作其它的一些调整。

```go
func LookPath(file string) (string, error)   //在环境变量 PATH 指定的目录中搜索可执行文件，如果 file 中有斜杠，则只在当前目录搜索。返回完整路径或者相对于当前目录的一个相对路径。
```

示例代码如下：

```go
    package main
    import (
        "fmt"
        "os/exec"
    )
    func main() {
        f, err := exec.LookPath("main")
        if err != nil {
            fmt.Println(err)
        }
        fmt.Println(f)
    }
```

运行结果如下：

```fo
main.exe
```

#### 3. os/user 获取当前用户信息

可以通过 os/user 包中的 Current() 函数来获取当前用户信息，该函数会返回一个 User 结构体，结构体中的 Username、Uid、HomeDir、Gid 分别表示当前用户的名称、用户 id、用户主目录和用户所属组 id，函数原型如下：

```go
func Current() (*User, error)
```

示例代码如下：

```go
    package main
    import (
        "log"
        "os/user"
    )
    func main() {
        u, _ := user.Current()
        log.Println("用户名：", u.Username)
        log.Println("用户id", u.Uid)
        log.Println("用户主目录：", u.HomeDir)
        log.Println("主组id：", u.Gid)
        // 用户所在的所有的组的id
        s, _ := u.GroupIds()
        log.Println("用户所在的所有组：", s)
    }
```

运行结果：

```go
2019/12/13 15:12:14 用户名： LENOVO-PC\Administrator
2019/12/13 15:12:14 用户id S-1-5-21-711400000-2334436127-1750000211-000
2019/12/13 15:12:14 用户主目录： C:\Users\Administrator
2019/12/13 15:12:14 主组id： S-1-5-22-766000000-2300000100-1050000262-000
2019/12/13 15:12:14 用户所在的所有组： [S-1-5-32-544 S-1-5-22-000 S-1-5-21-777400999-2344436111-1750000262-003]
```

#### 4. os/signal 信号处理

一个运行良好的程序在退出（正常退出或者强制退出，如 Ctrl+C，kill  等）时是可以执行一段清理代码的，将收尾工作做完后再真正退出。一般采用系统 Signal 来通知系统退出，如 kill  pid，在程序中针对一些系统信号设置了处理函数，当收到信号后，会执行相关清理程序或通知各个子进程做自清理。

Go语言中对信号的处理主要使用 os/signal 包中的两个方法，一个是 Notify 方法用来监听收到的信号，一个是 Stop 方法用来取消监听。

```go
func Notify(c chan<- os.Signal, sig ...os.Signal)  //第一个参数表示接收信号的 channel，第二个及后面的参数表示设置要监听的信号，如果不设置表示监听所有的信号。
```

【示例 1】使用 Notify 方法来监听收到的信号：

```go
    package main
    import (
        "fmt"
        "os"
        "os/signal"
    )
    func main() {
        c := make(chan os.Signal, 0)
        signal.Notify(c)
        // Block until a signal is received.
        s := <-c
        fmt.Println("Got signal:", s)
    }
```

运行该程序，然后在 CMD 窗口中通过 Ctrl+C 来结束该程序，便会得到输出结果：

```go
Got signal: interrupt
```

【示例 2】使用 stop 方法来取消监听：

```go
    package main
    import (
        "fmt"
        "os"
        "os/signal"
    )
    func main() {
        c := make(chan os.Signal, 0)
        signal.Notify(c)
        signal.Stop(c) //不允许继续往c中存入内容
        s := <-c       //c无内容，此处阻塞，所以不会执行下面的语句，也就没有输出
        fmt.Println("Got signal:", s)
    }
```

因为使用 Stop 方法取消了 Notify 方法的监听，所以运行程序没有输出结果。

### 十三、flag包：命令行参数解析

#### 1. flag参数类型

flag 包支持的命令行参数类型有 bool、int、int64、uint、uint64、float、float64、string、duration，如下表所示：

| flag 参数      | 有效值                                                       |
| -------------- | ------------------------------------------------------------ |
| 字符串 flag    | 合法字符串                                                   |
| 整数 flag      | 1234、0664、0x1234 等类型，也可以是负数                      |
| 浮点数 flag    | 合法浮点数                                                   |
| bool 类型 flag | 1、0、t、f、T、F、true、false、TRUE、FALSE、True、False      |
| 时间段 flag    | 任何合法的时间段字符串，如“300ms”、“-1.5h”、“2h45m”，  合法的单位有“ns”、“us”、“µs”、“ms”、“s”、“m”、“h” |

#### 2. flag 包基本使用

有以下两种常用的定义命令行 flag 参数的方法：

##### 1) flag.Type()

基本格式如下：

```go
flag.Type(flag 名, 默认值, 帮助信息) *Type
```

Type 可以是 Int、String、Bool 等，返回值为一个相应类型的指针，例如我们要定义姓名、年龄、婚否三个命令行参数，我们可以按如下方式定义：

```go
    name := flag.String("name", "张三", "姓名")
    age := flag.Int("age", 18, "年龄")
    married := flag.Bool("married", false, "婚否")
    delay := flag.Duration("d", 0, "时间间隔")
```

需要注意的是，此时 name、age、married、delay 均为对应类型的==指针==。

##### 2) flag.TypeVar()

```go
flag.TypeVar(Type 指针, flag 名, 默认值, 帮助信息)
```

TypeVar 可以是 IntVar、StringVar、BoolVar 等，其功能为将 flag 绑定到一个变量上，例如我们要定义姓名、年龄、婚否三个命令行参数，我们可以按如下方式定义：

```go
    var name string
    var age int
    var married bool
    var delay time.Duration
    flag.StringVar(&name, "name", "张三", "姓名")
    flag.IntVar(&age, "age", 18, "年龄")
    flag.BoolVar(&married, "married", false, "婚否")
    flag.DurationVar(&delay, "d", 0, "时间间隔")
```

#### flag.Parse()

通过以上两种方法定义好命令行 flag 参数后，需要通过调用 flag.Parse() 来对命令行参数进行解析。

支持的命令行参数格式有以下几种：

- -flag：只支持 bool 类型；
- -flag=x；
- -flag x：只支持非 bool 类型。

其中，布尔类型的参数必须使用等号的方式指定。

flag 包的其他函数：

```go
flag.Args()  //返回命令行参数后的其他参数，以 []string 类型
flag.NArg()  //返回命令行参数后的其他参数个数
flag.NFlag() //返回使用的命令行参数个数
```

一个实例，代码如下：

```go
    package main
    import (
        "flag"
        "fmt"
    )
    var Input_pstrName = flag.String("name", "gerry", "input ur name")
    var Input_piAge = flag.Int("age", 20, "input ur age")
    var Input_flagvar int
    func Init() {
        flag.IntVar(&Input_flagvar, "flagname", 1234, "help message for flagname")
    }
    func main() {
        Init()
        flag.Parse()
        // After parsing, the arguments after the flag are available as the slice flag.Args() or individually as flag.Arg(i). The arguments are indexed from 0 through flag.NArg()-1
        // Args returns the non-flag command-line arguments
        // NArg is the number of arguments remaining after flags have been processed
        fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
        for i := 0; i != flag.NArg(); i++ {
            fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
        }
        fmt.Println("name=", *Input_pstrName)
        fmt.Println("age=", *Input_piAge)
        fmt.Println("flagname=", Input_flagvar)
    }
```

运行结果如下：

```go
go run main.go -name "aaa" -age=123 -flagname=999
args=[], num=0
name= aaa
age= 123
flagname= 999
```

#### 自定义 Value

另外，我们还可以创建自定义 flag，只要实现 flag.Value 接口即可（要求 receiver 是指针类型），这时候可以通过如下方式定义该 flag：

```go
flag.Var(&flagVal, "name", "help message for flagname")
```

【示例】解析喜欢的编程语言，并直接解析到 slice 中，我们可以定义如下 sliceValue 类型，然后实现 Value 接口：

```go
    package main
    import (
        "flag"
        "fmt"
        "strings"
    )
    //定义一个类型，用于增加该类型方法
    type sliceValue []string
    //new一个存放命令行参数值的slice
    func newSliceValue(vals []string, p *[]string) *sliceValue {
        *p = vals
        return (*sliceValue)(p)
    }
    /*
    flag包中的Value接口：
    type Value interface {
        String() string
        Set(string) error
    }
    实现flag包中的Value接口，将命令行接收到的值用,分隔存到slice里
    */
    func (s *sliceValue) Set(val string) error {
        *s = sliceValue(strings.Split(val, ","))
        return nil
    }
    //flag为slice的默认值default is me,和return返回值没有关系
    func (s *sliceValue) String() string {
        *s = sliceValue(strings.Split("default is me", ","))
        return "It's none of my business"
    }
    /*
    可执行文件名 -slice="java,go"  最后将输出[java,go]
    可执行文件名 最后将输出[default is me]
    */
    func main(){
        var languages []string
        flag.Var(newSliceValue([]string{}, &languages), "slice", "I like programming `languages`")
        flag.Parse()
        //打印结果slice接收到的值
        fmt.Println(languages)
    }
```


通过-slice go,php 这样的形式传递参数，languages 得到的就是 [go, php]，如果不加-slice 参数则打印默认值[default is me]，如下所示：

```go
go run main.go -slice go,php,java
[go php java]
```

### 十四、go mod包依赖管理工具使用详解

#### 1.常用的go mod命令

| 命令            | 作用                                           |
| --------------- | ---------------------------------------------- |
| go mod download | 下载依赖包到本地（默认为 GOPATH/pkg/mod 目录） |
| go mod edit     | 编辑 go.mod 文件                               |
| go mod graph    | 打印模块依赖图                                 |
| go mod init     | 初始化当前文件夹，创建 go.mod 文件             |
| go mod tidy     | 增加缺少的包，删除无用的包                     |
| go mod vendor   | 将依赖复制到 vendor 目录下                     |
| go mod verify   | 校验依赖                                       |
| go mod why      | 解释为什么需要依赖                             |

##### 使用go get命令下载指定版本的依赖包

执行`go get `命令，在下载依赖包的同时还可以指定依赖包的版本。

- 运行`go get -u`命令会将项目中的包升级到最新的次要版本或者修订版本；
- 运行`go get -u=patch`命令会将项目中的包升级到最新的修订版本；
- 运行`go get [包名]@[版本号]`命令会下载对应包的指定版本或者将对应包升级到指定的版本。

> 提示：`go get [包名]@[版本号]`命令中版本号可以是 x.y.z 的形式，例如 go get foo@v1.2.3，也可以是 git 上的分支或 tag，例如 go get foo@master，还可以是 git 提交时的哈希值，例如 go get foo@e3702bed2。

#### 2.如何在项目中使用module

【示例 1】创建一个新项目：

1) 在 GOPATH 目录之外新建一个目录，并使用`go mod init`初始化生成 go.mod 文件。

```go
go mod init hello
go: creating new go.mod: module hello
```

go.mod 文件一旦创建后，它的内容将会被 go toolchain 全面掌控，go toolchain 会在各类命令执行时，比如`go get`、`go build`、`go mod`等修改和维护 go.mod 文件。

go.mod 提供了 module、require、replace 和 exclude 四个命令：

- module 语句指定包的名字（路径）；
- require 语句指定的依赖项模块；
- replace 语句可以替换依赖项模块；
- exclude 语句可以忽略依赖项模块。

初始化生成的 go.mod 文件如下所示：

```go
module hello

go 1.13
```

2. 添加依赖

新建一个 main.go 文件，写入以下代码：

```go
    package main
    import (
        "net/http"
        "github.com/labstack/echo"
    )
    func main() {
        e := echo.New()
        e.GET("/", func(c echo.Context) error {
            return c.String(http.StatusOK, "Hello, World!")
        })
        e.Logger.Fatal(e.Start(":1323"))
    }
```

执行`go run main.go`运行代码会发现 go mod 会自动查找依赖自动下载：

```go
go run main.go
go: finding github.com/labstack/echo v3.3.10+incompatible
go: downloading github.com/labstack/echo v3.3.10+incompatible
go: extracting github.com/labstack/echo v3.3.10+incompatible
go: finding github.com/labstack/gommon v0.3.0
......
go: finding golang.org/x/text v0.3.0

   ____    __
  / __/___/ /  ___
/ _// __/ _ \/ _ \
/___/\__/_//_/\___/ v3.3.10-dev
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                                      O\
⇨ http server started on [::]:1323
```

现在查看 go.mod 内容：

```go
module hello

go 1.13

require (
    github.com/labstack/echo v3.3.10+incompatible // indirect
    github.com/labstack/gommon v0.3.0 // indirect
    golang.org/x/crypto v0.0.0-20191206172530-e9b2fee46413 // indirect
)
```

go module 安装 package 的原则是先拉取最新的 release tag，若无 tag 则拉取最新的 commit，详见[ Modules 官方](https://github.com/golang/go/wiki/Modules)介绍。

go 会自动生成一个 go.sum 文件来记录 dependency tree：

```go
github.com/davecgh/go-spew v1.1.0/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/labstack/echo v3.3.10+incompatible h1:pGRcYk231ExFAyoAjAfD85kQzRJCRI8bbnE7CX5OEgg=
github.com/labstack/echo v3.3.10+incompatible/go.mod h1:0INS7j/VjnFxD4E2wkz67b8cVwCLbBmJyDaka6Cmk1s=
github.com/labstack/gommon v0.3.0 h1:JEeO0bvc78PKdyHxloTKiF8BD5iGrH8T6MSeGvSgob0=
github.com/labstack/gommon v0.3.0/go.mod h1:MULnywXg0yavhxWKc+lOruYdAhDwPK9wf0OL7NoOu+k=
github.com/mattn/go-colorable v0.1.2 h1:/bC9yWikZXAL9uJdulbSfyVNIR3n3trXl+v8+1sx8mU=
... 省略很多行
```

再次执行脚本`go run main.go`发现跳过了检查并安装依赖的步骤。

可以使用命令`go list -m -u all`来检查可以升级的 package，使用`go get -u need-upgrade-package`升级后会将新的依赖版本更新到 go.mod * 也可以使用`go get -u`升级所有依赖。

【示例 2】改造现有项目。

项目目录结构为：

```go
├─ main.go
│
└─ api
      └─ apis.go
```

main.go 源码为：

```go
    package main
    import (
        api "./api"  // 这里使用的是相对路径
        "github.com/labstack/echo"
    )
    func main() {
        e := echo.New()
        e.GET("/", api.HelloWorld)
        e.Logger.Fatal(e.Start(":1323"))
    }
```

api/apis.go 源码为：

```go
    package api
    import (
        "net/http"
        "github.com/labstack/echo"
    )
    func HelloWorld(c echo.Context) error {
        return c.JSON(http.StatusOK, "hello world")
    }
```

1) **使用 `go mod init ***` 初始化 go.mod。**

```go
go mod init hello
go: creating new go.mod: module hello
```

2) **运行`go run main.go`。**

```go
go run main.go
go: finding golang.org/x/crypto latest
build _/D_/code/src/api: cannot find module for path _/D_/code/src/api
```

首先还是会查找并下载安装依赖，然后运行脚本 main.go，这里会抛出一个错误：

```go
build _/D_/code/src/api: cannot find module for path _/D_/code/src/api
```

但是 go.mod 已经更新：

```go
module hello

go 1.13

require (
    github.com/labstack/echo v3.3.10+incompatible // indirect
    github.com/labstack/gommon v0.3.0 // indirect
    golang.org/x/crypto v0.0.0-20191206172530-e9b2fee46413 // indirect
)
```

那为什么会抛出这个错误呢？

这是因为 main.go 中使用 internal package 的方法跟以前已经不同了，由于 go.mod 会扫描同工作目录下所有  package 并且变更引入方法，必须将 hello 当成路径的前缀，也就是需要写成 import hello/api，以往  GOPATH/dep 模式允许的 import ./api 已经失效。

3)  **更新旧的 package import 方式。**

所以 main.go 需要改写成：

```go
    package main
    import (
        api "hello/api" // 修改路径，将 hello 当成路径的前缀
        "github.com/labstack/echo"
    )
    func main() {
        e := echo.New()
        e.GET("/", api.HelloWorld)
        e.Logger.Fatal(e.Start(":1323"))
    }
```

#### 3.使用 replace 替换无法直接获取的 package

由于某些已知的原因，并不是所有的 package 都能成功下载，比如：golang.org 下的包。

modules 可以通过在 go.mod 文件中使用 replace 指令替换成 github 上对应的库，比如：

```go
replace (
    golang.org/x/crypto v0.0.0-20190313024323-a1f597ede03a => github.com/golang/crypto v0.0.0-20190313024323-a1f597ede03a
)
```

或者

```go
replace golang.org/x/crypto v0.0.0-20190313024323-a1f597ede03a => github.com/golang/crypto v0.0.0-20190313024323-a1f597ede03a
```

