`RPC` 框架的一个基础能力是：像调用本地程序一样调用远程服务。那如何将程序映射为服务呢？

那么对 `Go` 来说，这个问题就变成了**如何将结构体的方法映射为服务**。

对 `net/rpc` 而言，一个函数需要能够被远程调用，需要满足如下五个条件：

- `the method’s type is exported` – 方法所属类型是导出的。
- `the method is exported` – 方式是导出的。
- `the method has two arguments, both exported (or builtin) types`  – 两个入参，均为导出或内置类型。
- `the method’s second argument is a pointer` – 第二个入参必须是一个指针。
- `the method has return type error` – 返回值为 error 类型。

更直观一些：

```go
func (t *T) MethodName(argType T1, replyType *T2) error
```

