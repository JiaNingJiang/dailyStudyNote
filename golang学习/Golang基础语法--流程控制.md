### 一、switch case语句

1. ##### 基本写法

Go语言改进了 switch 的语法设计，case 与 case 之间是独立的代码块，不需要通过 break 语句跳出当前 case 代码块以避免执行到下一行，示例代码如下：

```go
    var a = "hello"
    switch a {
    case "hello":
        fmt.Println(1)
    case "world":
        fmt.Println(2)
    default:
        fmt.Println(0)
    }
```

case 后不仅仅只是常量，还可以和 if 一样添加表达式，代码如下：

```go
    var r int = 11
    switch {
    case r > 10 && r < 20:
        fmt.Println(r)
    }
```

注意，这种情况的 switch 后面不再需要跟判断变量。

2. ##### 跨越 case 的 fallthrough——兼容C语言的 case 设计

在Go语言中 case 是一个独立的代码块，执行完毕后不会像C语言那样紧接着执行下一个 case，但是为了兼容一些移植代码，依然加入了 fallthrough 关键字来实现这一功能，代码如下：

```go
    var s = "hello"
    switch {
    case s == "hello":
        fmt.Println("hello")
        fallthrough     //跨越case
    case s != "world":
        fmt.Println("world")
    }
```

代码输出如下：

```go
hello
world
```

新编写的代码，不建议使用 fallthrough。

### 二、goto语句

1. ##### 使用 goto 退出多层循环

下面这段代码在满足条件时，需要连续退出两层循环，使用==传统的编码方式==如下：

```go
    package main
    import "fmt"
    func main() {
        var breakAgain bool
        // 外循环
        for x := 0; x < 10; x++ {
            // 内循环
            for y := 0; y < 10; y++ {
                // 满足某个条件时, 退出循环
                if y == 2 {
                    // 设置退出标记
                    breakAgain = true
                    // 退出本次循环
                    break
                }
            }
            // 根据标记, 还需要退出一次循环
            if breakAgain {
                    break
            }
        }
        fmt.Println("done")
    }
```

将上面的代码使用Go语言的 goto 语句进行优化：

```go
    package main
    import "fmt"
    func main() {
        for x := 0; x < 10; x++ {
            for y := 0; y < 10; y++ {
                if y == 2 {
                    // 跳转到标签
                    goto breakHere
                }
            }
        }
        // 手动返回, 避免执行进入标签
        return
        // 标签
    breakHere:
        fmt.Println("done")
    }
```

2. ##### 使用 goto 集中处理错误

多处错误处理存在代码重复时是非常棘手的，例如：

```go
    err := firstCheckError()
    if err != nil {
        fmt.Println(err)
        exitProcess()
        return
    }
    err = secondCheckError()
    if err != nil {
        fmt.Println(err)
        exitProcess()
        return
    }
    fmt.Println("done")
```

代码说明如下：

- 第 1 行，执行某逻辑，返回错误。
- 第 2～6 行，如果发生错误，打印错误退出进程。
- 第 8 行，执行某逻辑，返回错误。
- 第 10～14 行，发生错误后退出流程。
- 第 16 行，没有任何错误，打印完成。

在上面代码中，有一部分都是重复的错误处理代码，如果后期在这些代码中添加更多的判断，就需要在这些雷同的代码中依次修改，极易造成疏忽和错误。

使用 goto 语句来实现同样的逻辑：

```go
        err := firstCheckError()
        if err != nil {
            goto onExit
        }
        err = secondCheckError()
        if err != nil {
            goto onExit
        }
        fmt.Println("done")
        return
    onExit:           //汇总所有流程进行错误打印并退出进程。
        fmt.Println(err)
        exitProcess()
```

### 三、break跳出循环

Go语言中 break 语句可以结束 for、switch 和 select 的代码块，另外 break 语句还可以在语句后面添加标签，表示退出某个标签对应的代码块，标签要求必须定义在对应的 for、switch 和 select 的代码块上。

跳出指定循环：

```go
    package main
    import "fmt"
    func main() {
    OuterLoop:
        for i := 0; i < 2; i++ {
            for j := 0; j < 5; j++ {
                switch j {
                case 2:
                    fmt.Println(i, j)
                    break OuterLoop
                case 3:
                    fmt.Println(i, j)
                    break OuterLoop
                }
            }
        }
    }
```

代码输出如下：

```go
0  2
```

代码说明如下：

- 第 7 行，外层循环的标签。
- 第 8 行和第 9 行，双层循环。
- 第 10 行，使用 switch 进行数值分支判断。
- 第 13 和第 16 行，退出 OuterLoop 对应的循环之外，也就是==跳转到第 20 行==。

### 四、continue（继续下一次循环）

Go语言中 continue 语句可以结束当前循环，开始下一次的循环迭代过程，仅限在 for 循环内使用，在 continue 语句后添加标签时，表示开始标签对应的循环，例如：

```go
    package main
    import "fmt"
    func main() {
    OuterLoop:
        for i := 0; i < 2; i++ {
            for j := 0; j < 5; j++ {
                switch j {
                case 2:
                    fmt.Println(i, j)
                    continue OuterLoop
                }
            }
        }
    }
```

代码输出结果如下：

```go
0 2
1 2
```

代码说明：第 14 行将结束当前循环，开启下一次的外层循环，而不是第 10 行的循环。
