### 一.切片

1. ##### 使用 make() 函数生成的切片一定发生了内存分配操作，但给定开始与结束位置（包括切片复位）的切片只是将新的切片结构指向已经分配好的内存区域，设定开始与结束位置，不会发生内存分配操作。

2.  ##### var numListEmpty1 = []int{}   与var numListEmpty2 = []int  是不一样的, 表现为前者≠nil而后者=nil(即声明但未使用的切片的默认值是 nil，将 numListEmpty1 声明为一个整型切片，本来会在`{}`中填充切片的初始化元素，这里没有填充，所以切片是空的，但是此时的 numListEmpty1 已经被分配了内存，只是还没有元素。)
3. ##### 从数组生成切片，代码如下：

```go
var a  = [3]int{1, 2, 3}
fmt.Println(a, a[1:2])
```

其中 a 是一个拥有 3 个整型元素的数组，被初始化为数值 1 到 3，使用 a[1:2] 可以生成一个新的切片，代码运行结果如下：

```go
[1 2 3]  [2]
```

取出元素不包含结束位置对应的索引，切片最后一个元素使用 slice[len(slice)] 获取；

4. ##### Go语言的内建函数 append() 可以为切片动态添加元素，代码如下所示：

```go
    var a []int
    a = append(a, 1) // 追加1个元素
    a = append(a, 1, 2, 3) // 追加多个元素, 手写解包方式
    a = append(a, []int{1,2,3}...) // 追加一个切片, 切片需要解包,注意有三个省略号
```

不过需要注意的是，在使用 append() 函数为切片动态添加元素时，如果空间不足以容纳足够多的元素，切片就会进行“扩容”，此时新切片的长度会发生改变。切片在扩容时，容量的扩展规律是按容量的 2 倍数进行扩充，例如 1、2、4、8、16……

5. ##### 除了在切片的尾部追加，我们还可以在切片的开头添加元素：

```go
    var a = []int{1,2,3}
    a = append([]int{0}, a...) // 在开头添加1个元素,注意省略号
    a = append([]int{-3,-2,-1}, a...) // 在开头添加1个切片,注意省略号
```

在切片开头添加元素一般都会导致==内存的重新分配==，而且会导致==已有元素全部被复制 1 次==，因此，从切片的开头添加元素的性能要比从尾部追加元素的性能差很多。



6. ##### 因为 append 函数返回新切片的特性，所以切片也支持链式操作，我们可以将多个 append 操作组合起来，实现在切片中间插入元素：

```go
    var a []int
    a = append(a[:i], append([]int{x}, a[i:]...)...) // 在第i个位置插入x
    a = append(a[:i], append([]int{1,2,3}, a[i:]...)...) // 在第i个位置插入切片
```

每个添加操作中的第二个 append 调用都会创建一个临时切片，并将 a[i:] 的内容复制到新创建的切片中，然后将临时创建的切片再追加到 a[:i] 中。

7. ##### 使用copy()内置函数完成切片拷贝

copy() 可以将一个数组切片复制到另一个数组切片中，如果加入的两个数组切片不一样大，就会==按照其中较小的那个数组切片的元素个数进行复制==。copy() 函数的使用格式如下：

```go
    slice1 := []int{1, 2, 3, 4, 5}
    slice2 := []int{5, 4, 3}
    copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
    copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
```

下面通过代码演示对切片的引用和复制操作后对切片元素的影响。

```go
    package main
    import "fmt"
    func main() {
        // 设置元素数量为1000
        const elementCount = 1000
        // 预分配足够多的元素切片
        srcData := make([]int, elementCount)
        // 将切片赋值
        for i := 0; i < elementCount; i++ {
            srcData[i] = i
        }
        // 引用切片数据
        refData := srcData
        // 预分配足够多的元素切片
        copyData := make([]int, elementCount)
        // 将数据复制到新的切片空间中
        copy(copyData, srcData)
        // 修改原始数据的第一个元素
        srcData[0] = 999
        // 打印引用切片的第一个元素
        fmt.Println(refData[0])
        // 打印复制切片的第一个和最后一个元素
        fmt.Println(copyData[0], copyData[elementCount-1])
        // 复制原始数据从4到6(不包含)
        copy(copyData, srcData[4:6])
        for i := 0; i < 5; i++ {
            fmt.Printf("%d ", copyData[i])
        }
    }
```

代码说明如下：

- 第 19 行，将 refData 引用 srcData，切片不会因为等号操作进行元素的复制。
- 第 24 行，使用 copy() 函数将原始数据复制到 copyData 切片空间中。
- 第 30 行，引用数据的第一个元素将会发生变化。
- 第 33 行，打印复制数据的首位数据，由于数据是复制的，因此不会发生变化。
- 第 36 行，将 srcData 的局部数据复制到 copyData 中(将 srcData的第4和第5个元素复制到copyData的前两个元素上,copyData的其他元素保持不变)。
- 第 38～40 行，打印局部数据变化后的 copyData 的指定元素。

8. ##### 从切片中删除元素

Go语言并没有对删除切片元素提供专用的语法或者接口，需要使用切片本身的特性来删除元素，根据要删除元素的位置有三种情况，分别是从开头位置删除、从中间位置删除和从尾部删除，其中==删除切片尾部的元素速度最快==。

- 从开头位置删除

删除开头的元素可以直接移动数据指针：

```go
a = []int{1, 2, 3}
a = a[1:] // 删除开头1个元素
a = a[N:] // 删除开头N个元素
```

也可以不移动数据指针，但是将后面的数据向开头移动，可以用 append 原地完成（所谓原地完成是指在原有的切片数据对应的内存区间内完成，不会导致内存空间结构的变化）：

```go
    a = []int{1, 2, 3}
    a = append(a[:0], a[1:]...) // 删除开头1个元素
    a = append(a[:0], a[N:]...) // 删除开头N个元素
```

还可以用 copy() 函数来删除开头的元素：

```go
    a = []int{1, 2, 3}
    a = a[:copy(a, a[1:])] // 删除开头1个元素
    a = a[:copy(a, a[N:])] // 删除开头N个元素
```

- 从中间位置删除

对于删除中间的元素，需要对剩余的元素进行一次整体挪动，同样可以用 append 或 copy 原地完成：

```go
    a = []int{1, 2, 3, ...}
    a = append(a[:i], a[i+1:]...) // 删除中间1个元素
    a = append(a[:i], a[i+N:]...) // 删除中间N个元素
    a = a[:i+copy(a[i:], a[i+1:])] // 删除中间1个元素
    a = a[:i+copy(a[i:], a[i+N:])] // 删除中间N个元素
```

- 从尾部删除

```go
    a = []int{1, 2, 3}
    a = a[:len(a)-1] // 删除尾部1个元素
    a = a[:len(a)-N] // 删除尾部N个元素
```

Go语言中删除切片元素的本质是，以被删除元素为分界点，将前后两个部分的内存重新连接起来。连续容器的元素删除无论在任何语言中，都要将删除点前后的元素移动到新的位置，随着元素的增加，这个过程将会变得极为耗时，因此，当业务需要大量、频繁地从一个切片中删除元素时，如果对性能要求较高的话，就需要考虑更换其他的容器了（如双链表等能快速从删除点删除元素）。

9. ##### 利用ange关键字循环迭代切片

当迭代切片时，关键字 range 会返回两个值，第一个值是当前迭代到的索引位置，第二个值是该位置对应元素值的一份副本，如下图所示。

![image-20220922090015516](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\image-20220922090015516.png)

需要强调的是，==range 返回的是每个元素的副本，而不是直接返回对该元素的引用==，如下所示。

```go
    // 创建一个整型切片，并赋值
    slice := []int{10, 20, 30, 40}
    // 迭代每个元素，并显示值和地址
    for index, value := range slice {
        fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n", value, &value, &slice[index])
    }
```

输出结果为：

```go
Value: 10 Value-Addr: 10500168 ElemAddr: 1052E100
Value: 20 Value-Addr: 10500168 ElemAddr: 1052E104
Value: 30 Value-Addr: 10500168 ElemAddr: 1052E108
Value: 40 Value-Addr: 10500168 ElemAddr: 1052E10C
```

因为迭代返回的变量是一个在迭代过程中根据切片依次赋值的新变量，所以 ==value 的地址总是相同的==，要想获取每个元素的地址，需要使用切片变量和索引值（例如上面代码中的 &slice[index]）。

10. ##### 多维切片

```go
    // 声明一个二维整型切片并赋值
    slice := [][]int{{10}, {100, 200}}
```

上面的代码中展示了一个包含两个元素的外层切片，同时每个元素包又含一个内层的整型切片，切片 slice 的值如下图所示。

![image-20220922090608999](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\image-20220922090608999.png)

通过上图可以看到外层的切片包括两个元素，每个元素都是一个切片，第一个元素中的切片使用单个整数 10 来初始化，第二个元素中的切片包括两个整数，即 100 和 200。内置函数[ append() ](http://c.biancheng.net/view/28.html)的规则也可以应用到组合后的切片上，如下所示。

```go
    // 声明一个二维整型切片并赋值
    slice := [][]int{{10}, {100, 200}}
    // 为第一个切片追加值为 20 的元素
    slice[0] = append(slice[0], 20)
```

![image-20220922090937283](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\image-20220922090937283.png)

### 二.map

1. ##### map的容量问题

和数组不同，map 可以根据新增的 key-value 动态的伸缩，因此它不存在固定长度或者最大限制，但是也可以选择标明 map 的初始容量 capacity，格式如下：

```go
make(map[keytype]valuetype, cap)  例如  map2 := make(map[string]float, 100)
```

==当 map 增长到容量上限的时候，如果再增加新的 key-value，map 的大小会自动加 1==，所以出于性能的考虑，对于大的 map 或者会快速扩张的 map，即使只是大概知道容量，也最好先标明。

这里有一个 map 的具体例子，即将音阶和对应的音频映射起来：

```go
    noteFrequency := map[string]float32 {
    "C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83,
    "G0": 24.50, "A0": 27.50, "B0": 30.87, "A4": 440}
```

注意：可以使用 make()，但不能使用 new() 来构造 map，如果错误的使用 new() 分配了一个引用对象，会获得一个空引用的指针，相当于声明了一个未初始化的变量并且取了它的地址：

```go
mapCreated := new(map[string]float)
```

接下来当我们调用`mapCreated["key1"] = 4.5`的时候，编译器会报错：

```go
invalid operation: mapCreated["key1"] (index of type *map[string]float).
```

2. ##### 用切片作为 map 的值

既然一个 key 只能对应一个 value，而 value 又是一个原始类型，那么==如果一个 key 要对应多个值怎么办？==例如，当我们要处理  unix 机器上的所有进程，以父进程（pid 为整形）作为 key，所有的子进程（以所有子进程的 pid 组成的切片）作为 value。通过将  value 定义为 []int 类型或者其他类型的切片，就可以优雅的解决这个问题，示例代码如下所示：

```go
    mp1 := make(map[int][]int)
    mp2 := make(map[int]*[]int)
```

3. ##### 利用 for range 循环完成map遍历

遍历对于Go语言的很多对象来说都是差不多的，直接使用 for range 语法即可，遍历时，可以同时获得键和值，如只遍历值，可以使用下面的形式：

```go
for _, v := range scene {   //将不需要的键使用_改为匿名变量形式。
```

只遍历键时，使用下面的形式：

```go
for k := range scene {   //无须将值改为匿名变量形式，忽略值即可。
```

4. ##### 清空map数据

使用 delete() 内建函数从 map 中删除一组键值对，delete() 函数的格式如下：

```go
delete(map, 键)
```

如果要清空 map 中的所有元素，Go语言中并没有为 map 提供任何清空所有元素的函数、方法，==清空 map 的唯一办法就是重新 make 一个新的 map==，不用担心垃圾回收的效率，Go语言中的并行垃圾回收效率比写一个清空函数要高效的多。

5. ##### sync.Map（在并发环境中使用的map）

Go语言中的 map 在并发情况下，只读是线程安全的，==同时读写==是线程不安全的。

```go
    // 创建一个int到int的映射
    m := make(map[int]int)
    // 开启一段并发代码
    go func() {
        // 不停地对map进行写入
        for {
            m[1] = 1
        }
    }()
    // 开启一段并发代码
    go func() {
        // 不停地对map进行读取
        for {
            _ = m[1]
        }
    }()
    // 无限循环, 让并发程序在后台执行
    for {
    }
```

运行代码会报错，输出如下：

```go
fatal error: concurrent map read and map write
```

错误信息显示，并发的 map 读和 map 写，也就是说使用了==两个并发函数不断地对 map 进行读和写而发生了竞态问题==，map 内部会对这种并发操作进行检查并提前发现。

==需要并发读写时，一般的做法是加锁，但这样性能并不高==，==Go语言在 1.9 版本中提供了一种效率较高的并发安全的 sync.Map==，sync.Map 和 map 不同，不是以语言原生形态提供，而是在 sync 包下的特殊结构。

sync.Map 有以下特性：

- 无须初始化，直接声明即可。
- sync.Map 不能使用 map 的方式进行取值和设置等操作，而是使用 sync.Map 的方法进行调用，Store 表示存储，Load 表示获取，Delete 表示删除。
- 使用 Range 配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值，Range 参数中回调函数的返回值在需要继续迭代遍历时，返回 true，终止迭代遍历时，返回 false。

并发安全的 sync.Map 演示代码如下：

```go
    package main
    import (
          "fmt"
          "sync"
    )
    func main() {
        var scene sync.Map
        // 将键值对保存到sync.Map
        scene.Store("greece", 97)
        scene.Store("london", 100)
        scene.Store("egypt", 200)
        // 从sync.Map中根据键取值
        fmt.Println(scene.Load("london"))
        // 根据键删除对应的键值对
        scene.Delete("london")
        // 遍历所有sync.Map中的键值对
        scene.Range(func(k, v interface{}) bool {
            fmt.Println("iterate:", k, v)
            return true
        })
    }
```

代码输出如下：

```go
100 true
iterate: egypt 200
iterate: greece 97
```

代码说明如下：

- 第 10 行，声明 scene，类型为 sync.Map，注意，sync.Map 不能使用 make 创建。
- 第 13～15 行，将一系列键值对保存到 sync.Map 中，sync.Map ==将键和值以 interface{} 类型进行保存==。
- 第 18 行，提供一个 sync.Map 的键给 scene.Load() 方法后将查询到键对应的值返回。
- 第 21 行，sync.Map 的 Delete 可以使用指定的键将对应的键值对删除。
- 第 24 行，Range() 方法可以遍历 sync.Map，遍历需要提供一个匿名函数，==参数为 k、v，类型为 interface{}==，==每次 Range() 在遍历一个元素时，都会调用这个匿名函数把结果返回。==

sync.Map 没有提供获取 map 数量的方法，替代方法是在获取 sync.Map 时遍历自行计算数量，sync.Map 为了保证并发安全有一些性能损失，因此在非并发情况下，使用 map 相比使用 sync.Map 会有更好的性能。

### 三、list列表

列表是一种非连续的存储容器，由多个节点组成，节点通过一些变量记录彼此之间的关系，列表有多种实现方法，如单链表、双链表等。==在Go语言中，列表使用 container/list 包来实现==，==内部的实现原理是双链表==，列表能够高效地进行任意位置的元素插入和删除操作。

1. ##### 初始化列表

- 通过 container/list 包的 New() 函数初始化 list

```go
变量名 := list.New()
```

- 通过 var 关键字声明初始化 list

```go
var 变量名 list.List
```

列表与切片和 map 不同的是，列表并没有具体元素类型的限制，因此，列表的元素可以是任意类型，这既带来了便利，也引来一些问题，例如==给列表中放入了一个 interface{} 类型的值，取出值后，如果要将 interface{} 转换为其他类型将会发生宕机==。

2. ##### 在列表中插入元素

双链表支持从队列前方或后方插入元素，分别对应的方法是 PushFront 和 PushBack。

==这两个方法都会返回一个 *list.Element 结构==，如果在以后的使用中==需要删除插入的元素，则只能通过 *list.Element 配合 Remove() 方法进行删除==，这种方法可以让删除更加效率化，同时也是双链表特性之一。

```go
    l := list.New()
    l.PushBack("fist")   //尾插法,插入一个字符串
    l.PushFront(67)		//前插法，插入一个整形数值
```

列表插入元素的其他方法如下表所示。

| 方  法                                                | 功  能                                            |
| ----------------------------------------------------- | ------------------------------------------------- |
| InsertAfter(v interface {}, mark * Element) * Element | 在 mark 点之后插入元素，mark 点由其他插入函数提供 |
| InsertBefore(v interface {}, mark * Element) *Element | 在 mark 点之前插入元素，mark 点由其他插入函数提供 |
| PushBackList(other *List)                             | 添加 other 列表元素到尾部                         |
| PushFrontList(other *List)                            | 添加 other 列表元素到头部                         |

3. ##### 从列表中删除元素

列表插入函数的返回值会提供一个 *list.Element 结构，这个结构记录着列表元素的值以及与其他节点之间的关系等信息，从列表中删除元素时，需要用到这个结构进行快速删除。

```go
    package main
    import "container/list"
    func main() {
        l := list.New()
        // 尾部添加
        l.PushBack("canon")
        // 头部添加
        l.PushFront(67)
        // 尾部添加后保存元素句柄
        element := l.PushBack("fist")
        // 在fist之后添加high
        l.InsertAfter("high", element)
        // 在fist之前添加noon
        l.InsertBefore("noon", element)
        // 使用
        l.Remove(element)
    }
```

代码说明如下：
 第 6 行，创建列表实例。
 第 9 行，将字符串 canon 插入到列表的尾部。
 第 12 行，将数值 67 添加到列表的头部。
 第 15 行，将字符串 fist 插入到列表的尾部，并将这个元素的内部结构保存到 element 变量中。
 第 18 行，使用 element 变量，在 element 的位置后面插入 high 字符串。
 第 21 行，使用 element 变量，在 element 的位置前面插入 noon 字符串。
 第 24 行，移除 element 变量对应的元素。

下表中展示了每次操作后列表的实际元素情况。

| 操作内容                        | 列表元素                    |
| ------------------------------- | --------------------------- |
| l.PushBack("canon")             | canon                       |
| l.PushFront(67)                 | 67, canon                   |
| element := l.PushBack("fist")   | 67, canon, fist             |
| l.InsertAfter("high", element)  | 67, canon, fist, high       |
| l.InsertBefore("noon", element) | 67, canon, noon, fist, high |
| l.Remove(element)               | 67, canon, noon, high       |

4. ##### 遍历列表——访问列表的每一个元素

遍历双链表需要配合 Front() 函数获取头元素，遍历时只要元素不为空就可以继续进行，每一次遍历都会调用元素的 Next() 函数，代码如下所示。

```go
    l := list.New()
    // 尾部添加
    l.PushBack("canon")
    // 头部添加
    l.PushFront(67)
    for i := l.Front(); i != nil; i = i.Next() {
        fmt.Println(i.Value)   //使用遍历返回的 *list.Element 的 Value 成员取得放入列表时的原值。
    }
```

代码输出如下：

```go
67
canon
```

### 四、nil：空值/零值

1. ##### nil 标识符是不能比较的

```go
    package main
    import (
        "fmt"
    )
    func main() {
        fmt.Println(nil==nil)
    }
```

运行结果如下所示：

```go
PS D:\code> go run .\main.go
# command-line-arguments
.\main.go:8:21: invalid operation: nil == nil (operator == not defined on nil)
```

从上面的运行结果不难看出，`==`对于 nil 来说是一种未定义的操作。

2. ##### nil 不是关键字或保留字

nil 并不是Go语言的关键字或者保留字，也就是说我们可以定义一个名称为 nil 的变量，比如下面这样：

```go
var nil = errors.New("my god")
```

虽然上面的声明语句可以通过编译，但是并不提倡这么做。

3. ##### nil 没有默认类型

```go
    package main
    import (
        "fmt"
    )
    func main() {
        fmt.Printf("%T", nil)
        print(nil)
    }
```

运行结果如下所示：

```go
PS D:\code> go run .\main.go
# command-line-arguments
.\main.go:9:10: use of untyped nil
```

4. ##### 不同类型 nil 的指针是一样的

```go
    package main
    import (
        "fmt"
    )
    func main() {
        var arr []int
        var num *int
        fmt.Printf("%p\n", arr)
        fmt.Printf("%p", num)
    }
```

运行结果如下所示：

```go
PS D:\code> go run .\main.go
0x0
0x0
```

通过运行结果可以看出 arr 和 num 的指针都是 0x0。

5. ##### 不同类型的 nil 是不能比较的

```go
    package main
    import (
        "fmt"
    )
    func main() {
        var m map[int]string
        var ptr *int
        fmt.Printf(m == ptr)
    }
```

运行结果如下所示：

```go
PS D:\code> go run .\main.go
# command-line-arguments
.\main.go:10:20: invalid operation: arr == ptr (mismatched types []int and *int)
```

6. ##### 两个相同类型的 nil 值也可能无法比较

在Go语言中 map、slice 和 function 类型的 nil 值不能比较，比较两个无法比较类型的值是非法的，下面的语句无法编译。

```go
    package main
    import (
        "fmt"
    )
    func main() {
        var s1 []int
        var s2 []int
        fmt.Printf(s1 == s2)
    }
```

运行结果如下所示：

```go
PS D:\code> go run .\main.go
# command-line-arguments
.\main.go:10:19: invalid operation: s1 == s2 (slice can only be compared to nil)
```

通过上面的错误提示可以看出，==能够将上述不可比较类型的空值直接与 nil 标识符进行比较==，如下所示：

```go
    package main
    import (
        "fmt"
    )
    func main() {
        var s1 []int
        fmt.Println(s1 == nil)
    }
```

运行结果如下所示：

```go
PS D:\code> go run .\main.go
true
```

7. ##### nil 是 map、slice、pointer、channel、func、interface 的零值

```go
    package main
    import (
        "fmt"
    )
    func main() {
        var m map[int]string
        var ptr *int
        var c chan int
        var sl []int
        var f func()
        var i interface{}
        fmt.Printf("%#v\n", m)
        fmt.Printf("%#v\n", ptr)
        fmt.Printf("%#v\n", c)
        fmt.Printf("%#v\n", sl)
        fmt.Printf("%#v\n", f)
        fmt.Printf("%#v\n", i)
    }
```

运行结果如下所示：

```go
PS D:\code> go run .\main.go
map[int]string(nil)
(*int)(nil)
(chan int)(nil)
[]int(nil)
(func())(nil)
<nil>
```

零值是Go语言中变量在声明之后但是未初始化被赋予的该类型的一个默认值。

8. ##### 不同类型的 nil 值占用的内存大小可能是不一样的

一个类型的所有的值的内存布局都是一样的，nil 也不例外，nil 的大小与同类型中的非 nil 类型的大小是一样的。但是不同类型的 nil 值的大小可能不同。

```go
    package main
    import (
        "fmt"
        "unsafe"
    )
    func main() {
        var p *struct{}
        fmt.Println( unsafe.Sizeof( p ) ) // 8
        var s []int
        fmt.Println( unsafe.Sizeof( s ) ) // 24
        var m map[int]bool
        fmt.Println( unsafe.Sizeof( m ) ) // 8
        var c chan string
        fmt.Println( unsafe.Sizeof( c ) ) // 8
        var f func()
        fmt.Println( unsafe.Sizeof( f ) ) // 8
        var i interface{}
        fmt.Println( unsafe.Sizeof( i ) ) // 16
    }
```

运行结果如下所示：

```go
PS D:\code> go run .\main.go
8
24
8
8
8
16
```

具体的大小取决于编译器和架构，上面打印的结果是在 64 位架构和标准编译器下完成的，对应 32 位的架构的，打印的大小将减半。

