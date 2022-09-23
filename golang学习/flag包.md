1.如何使用自定义的flag对象

flag包提供了默认的flag会使用到的方法，但是我们还可以自定义“命令行参数解析规则”：

```go
var myFlagSetCmd = flag.NewFlagSet("myflagset", flag.ExitOnError)
```

上述语句创建了一个名为myFlagSetCmd的自定义flag对象，定义之后我们再使用下述语句：

```go
var stringFlag = myFlagSet.String("abc", "default value", "help mesage")
```

就为这个自定义的flag对象注册了一个 -abc的命令行参数，此参数默认值为"default value"，说明为"help mesage"