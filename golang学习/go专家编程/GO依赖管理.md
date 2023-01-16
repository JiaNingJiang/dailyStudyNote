## 一、module

### 1. 基础概念

#### 1.1 Module的定义

首先，module是个新鲜又熟悉的概念。新鲜是指在以往GOPATH和vendor时代都没有提及，它是个新的词汇。
为什么说熟悉呢？因为它不是新的事物，事实上我们经常接触，这次只是官方给了一个统一的称呼而矣。

拿开源项目`https://github.com/blang/semver`举例，这个项目是一个语义化版本处理库，当你的项目需要时可以在你的项目中import，比如：

```golang
"github.com/blang/semver"
```

`https://github.com/blang/semver`项目中可以包含一个或多个package，不管有多少package，这些package都随项目一起发布，即**当我们说`github.com/blang/semver`某个版本时，说的是整个项目，而不是具体的package。此时项目`https://github.com/blang/semver`就是一个module。**

**官方给module的定义是：`A module is a collection of related Go packages that are versioned together as a single unit.`，**定义非常晰，一组package的集合，一起被标记版本，即是一个module。

通常而言，**一个仓库包含一个module（虽然也可以包含多个，但不推荐）**，所以仓库、module和package的关系如下：

- **一个仓库包含一个或多个Go module**；
- **每个Go module包含一个或多个Go package**；
- **每个package包含一个或多个Go源文件**；

此外，**一个module的版本号规则必须遵循语义化规范**（https://semver.org/），版本号必须使用格式`v(major).(minor).(patch)`，比如`v0.1.0`、`v1.2.3` 或`v1.5.0-rc.1`。

#### 1.2 语义化版本规范

语义化版本（Semantic Versioning）已成为事实上的标准，几乎知名的开源项目都遵循该规范，更详细的信息请前往https://semver.org/ 查看，在此只提炼一些要点，以便于后续的阅读。

版本格式`v(major).(minor).(patch)`中**major指的是大版本，minor指小版本，patch指补丁版本**。

- **major: 当发生不兼容的改动时才可以增加major版本**；比如`v2.x.y`与`v1.x.y`是不兼容的；
- **minor: 当有新增特性时才可以增加该版本**，比如`v1.17.0`是在`v1.16.0`基础上加了新的特性，同时兼容`v1.16.0`；
- **patch: 当有bug修复时才可以 增加该版本**，比如`v1.17.1`修复了`v1.17.0`上的bug，没有新特性增加；

语义化版本规范的好处是，用户通过版本号就能了解版本信息。

除了上面介绍的基础概念以外，还有描述依赖的`go.mod`和记录module的checksum的`go.sum`等内容，

### 2. 快速实践

#### 2.1 Go module到底是做什么的？

我们在前面的章节已介绍过，但还是想强调一下，**Go module实际上只是精准的记录项目的依赖情况，包括每个依赖的精确版本号，仅此而矣**。

那么，为什么需要记录这些依赖情况，或者记录这些依赖有什么好处呢？

试想一下，**在编译某个项目时，第三方包的版本往往是可以替换的，如果不能精确的控制所使用的第三方包的版本，最终构建出的可执行文件从本质上是不同的，这会给问题诊断带来极大的困扰。**

接下来，我们从一个Hello World项目开始，逐步介绍如何初始化module、如何记录依赖的版本信息。
项目托管在GitHub `https://github.com/renhongcai/gomodule`中，并使用版本号区别使用go module的阶段。

- v1.0.0 未引用任何第三方包，也未使用go module
- v1.1.0 未引用任何第三方包，已开始使用go module，但没有任何外部依赖
- v1.2.0 引用了第三方包，并更新了项目依赖

需要注意的是，下面的例子统一使用go 1.13版本，如果你使用go 1.11 或者go 1.12，运行效果可能略有不同。
本文最后部分我们尽量尝试记录一些版本间的差异，以供参考。

#### 2.2 初始化module

**一个项目若要使用Go module，那么其本身需要先成为一个module**，也即需要一个module名字。

在Go module机制下，项目的module名字以及其依赖信息记录在一个名为`go.mod`的文件中，该文件可以手动创建，也可以使用`go mod init`命令自动生成。推荐自动生成的方法，如下所示：

```go
[root@ecs-d8b6 gomodule]# go mod init github.com/renhongcai/gomodule
go: creating new go.mod: module github.com/renhongcai/gomodule
```

**完整的`go mod init`命令格式为`go mod init [module]`：其中`[module]`为module名字，如果不填，`go mod init`会尝试从版本控制系统或import的注释中猜测一个。**这里**推荐指定明确的module名字**，因为猜测有时需要一些额外的条件，比如 Go 1.13版本，只有项目位于GOPATH中才可以正确运行，而 Go 1.11版本则没有此要求。

上面的命令会自动创建一个`go.mod`文件，其中包括module名字，以及我们所使用的Go 版本：

```go
[root@ecs-d8b6 gomodule]# cat go.mod 
module github.com/renhongcai/gomodule

go 1.13
```

`go.mod`文件中的版本号`go 1.13`是在Go 1.12引入的，意思是开发此项目的Go语言版本，并不是编译该项目所限制的Go语言版本，但是如果项目中使用了Go 1.13的新特性，而你使用Go 1.11编译的话，编译失败时，编译器会提示你版本不匹配。

由于我们的项目还没有使用任何第三方包，所以`go.mod`中并没有记录依赖包的任何信息。我们把自动生成的`go.mod`提交，然后我们尝试引用一个第三方包。

#### 2.3 管理依赖

现在我们准备**引用一个第三方包`github.com/google/uuid`来生成一个UUID**，这样就会产生一个依赖，代码如下：

```go
package main

import (
    "fmt"

    "github.com/google/uuid"
)

func main() {
    id := uuid.New().String()
    fmt.Println("UUID: ", id)
}
```

在**开始编译以前**，我们先**使用`go get`来分析依赖情况，并会自动下载依赖**：

```go
[root@ecs-d8b6 gomodule]# go get 
go: finding github.com/google/uuid v1.1.1
go: downloading github.com/google/uuid v1.1.1
go: extracting github.com/google/uuid v1.1.1
```

从输出内容来看，**`go get`帮我们定位到可以使用`github.com/google/uuid`的v1.1.1版本，并下载再解压它们**。

注意：**`go get`总是获取依赖的最新版本**，如果`github.com/google/uuid`发布了新的版本，输出的版本信息会相应的变化。**关于Go Module机制中版本选择我们将在后续的章节详细介绍。**

`go get`命令会自动修改`go.mod`文件：

```go
[root@ecs-d8b6 gomodule]# cat go.mod 
module github.com/renhongcai/gomodule

go 1.13

require github.com/google/uuid v1.1.1
```

可以看到，现在**`go.mod`中增加了`require github.com/google/uuid v1.1.1`内容**，表示当前项目依赖`github.com/google/uuid`的`v1.1.1`版本，这就是我们所说的`go.mod`记录的依赖信息。

由于这是当前项目第一次引用外部依赖，`go get`命令还会生成一个`go.sum`文件，记录依赖包的hash值：

```go
[root@ecs-d8b6 gomodule]# cat go.sum 
github.com/google/uuid v1.1.1 h1:Gkbcsh/GbpXz7lPftLA3P6TYMwjCLYm83jiFQZF/3gY=
github.com/google/uuid v1.1.1/go.mod h1:TIyPZe4MgqvfeYDBFedMoGGpEw/LqOeaOT+nhxU+yHo=
```

**该文件通过记录每个依赖包的hash值，来确保依赖包没有被篡改。**关于此部分内容我们在此暂不展开介绍，留待后面的章节详细介绍。

经`go get`修改的`go.mod`和创建的`go.sum`都需要提交到代码库，这样别人获取到项目代码，编译时就会使用项目所要求的依赖版本。

至此，项目已经有一个依赖包，并且可以编译执行了，每次运行都会生成一个独一无二的UUID：

```go
[root@ecs-d8b6 gomodule]# go run main.go
UUID:  20047f5a-1a2a-4c00-bfcd-66af6c67bdfb
```

注：**如果你没有使用`go get`在执行之前下载依赖，而是直接使用`go build main.go`运行项目的话，依赖包也会被自动下载。**但是在`v1.13.4`中有个bug，即此时生成的`go.mod`中显示的依赖信息则会是`require github.com/google/uuid v1.1.1 // indirect`，注意行末的`indirect`表示间接依赖，这明显是错误的，因为我们直接`import`的。

### 3. replace指令

**`go.mod`文件中通过`指令`声明module信息**，用于控制命令行工具进行版本选择。一共有四个指令可供使用(**在go.mod文件中直接使用**)：

- module： 声明module名称；
- require： 声明依赖以及其版本号；
- replace： 替换require中声明的依赖，使用另外的依赖及其版本号；
- exclude： 禁用指定的依赖；

其中`module`和`require`我们前面已介绍过，`module`用于指定module的名字，如`module github.com/renhongcai/gomodule`，那么其他项目引用该module时其import路径需要指定`github.com/renhongcai/gomodule`。

**`require`用于指定依赖，如`require github.com/google/uuid v1.1.1`，该指令相当于告诉`go build`使用`github.com/google/uuid`的`v1.1.1`版本进行编译。**

#### 3.1 replace 工作机制

顾名思义，`replace`指替换，它**指示编译工具替换`require`指定中出现的包**，比如，我们在`require`中指定的依赖如下：

```go
module github.com/renhongcai/gomodule  

go 1.13  

require github.com/google/uuid v1.1.1
```

此时，我们可以**使用`go list -m all`命令查看依赖最终选定的版本**：

```go
[root@ecs-d8b6 gomodule]# go list -m all
github.com/renhongcai/gomodule
github.com/google/uuid v1.1.1
```

毫无意外，最终选定的uuid版本正是我们在require中指定的版本`v1.1.1`。



**如果我们想使用uuid的v1.1.0版本进行构建，可以修改require指定，还可以使用replace来指定。**
需要说明的是，**正常情况下不需要使用replace来修改版本，最直接的办法是修改require即可**，虽然replace也能够做到，但这不是replace的一般使用场景。
下面我们先通过一个简单的例子来说明replace的功能，随即介绍几种常见的使用场景。

比如，我们**修改`go.mod`，添加replace指令**：

```go
[root@ecs-d8b6 gomodule]# cat go.mod 
module github.com/renhongcai/gomodule

go 1.13

require github.com/google/uuid v1.1.1

replace github.com/google/uuid v1.1.1 => github.com/google/uuid v1.1.0
```

`replace github.com/google/uuid v1.1.1 => github.com/google/uuid v1.1.0`指定表示替换uuid v1.1.1版本为 v1.1.0，此时再次使用`go list -m all`命令查看最终选定的版本：

```shell
[root@ecs-d8b6 gomodule]# go list -m all 
github.com/renhongcai/gomodule
github.com/google/uuid v1.1.1 => github.com/google/uuid v1.1.0
```

可以看到其最终选择的uuid版本为 v1.1.0。如果你本地没有v1.1.0版本，你或许还会看到一条`go: finding github.com/google/uuid v1.1.0`信息，它表示在下载uuid v1.1.0包，也从侧面证明最终选择的版本为v1.1.0。

到此，我们可以看出`replace`的作用了，它用于替换`require`中出现的包，**它正常工作还需要满足两个条件：**

**第一，`replace`仅在当前module为`main module`时有效**，比如我们当前在编译`github.com/renhongcai/gomodule`，此时就是`main module`，**如果其他项目引用了`github.com/renhongcai/gomodule`，那么其他项目编译时，`replace`就会被自动忽略。**

**第二，`replace`指定中`=>`前面的包及其版本号必须出现在`require`中才有效，否则指令无效，也会被忽略。**
比如，上面的例子中，我们**指定`replace github.com/google/uuid => github.com/google/uuid v1.1.0`，或者指定`replace github.com/google/uuid v1.0.9 => github.com/google/uuid v1.1.0`，二者均都无效。**

#### 3.2 replace 使用场景

前面的例子中，我们使用`replace`替换`require`中的依赖，在实际项目中`replace`在项目中经常被使用，其中不乏一些精彩的用法。
但**不管应用在哪种场景，其本质都一样，都是替换`require`中的依赖。**

##### 3.2.1 替换无法下载的包

由于中国大陆网络问题，有些包无法顺利下载，比如`golang.org`组织下的包，值得庆幸的是这些包在GitHub都有镜像，此时
就可以使用GitHub上的包来替换。

比如，项目中使用了`golang.org/x/text`包：

```go
package main

import (
    "fmt"

    "github.com/google/uuid"
    "golang.org/x/text/language"
    "golang.org/x/text/message"
)

func main() {
    id := uuid.New().String()
    fmt.Println("UUID: ", id)

    p := message.NewPrinter(language.BritishEnglish)
    p.Printf("Number format: %v.\n", 1500)

    p = message.NewPrinter(language.Greek)
    p.Printf("Number format: %v.\n", 1500)
}
```

上面的简单例子，使用两种语言`language.BritishEnglish` 和`language.Greek`分别打印数字`1500`，来查看不同语言对数字格式的处理，一个是`1,500`，另一个是`1.500`。此时就会**分别引入`"golang.org/x/text/language"` 和`"golang.org/x/text/message"`。**

**执行`go get` 或`go build`命令时会就再次分析依赖情况，并更新`go.mod`文件**。网络正常情况下，`go.mod`文件将会变成下面的内容：

```go
module github.com/renhongcai/gomodule

go 1.13

require (
    github.com/google/uuid v1.1.1
    golang.org/x/text v0.3.2
)

replace github.com/google/uuid v1.1.1 => github.com/google/uuid v1.1.0
```

**没有合适的网络代理情况下，`golang.org/x/text` 很可能无法下载**。那么此时，就可以**使用`replace`来让我们的项目使用GitHub上相应的镜像包。**我们可以添加一条新的`replace`条目，如下所示：

```go
replace (
    github.com/google/uuid v1.1.1 => github.com/google/uuid v1.1.0
    golang.org/x/text v0.3.2 => github.com/golang/text v0.3.2
)
```

此时，**项目编译时就会从GitHub下载包**。我们**源代码中import路径 `golang.org/x/text/xxx`不需要改变。**

也许有读者会问，**是否可以将import路径由`golang.org/x/text/xxx`改成`github.com/golang/text/xxx`？**这样一来，就不需要使用replace来替换包了。

**遗憾的是，不可以。因为`github.com/golang/text`只是镜像仓库，其`go.mod`文件中定义的module还是`module golang.org/x/text`，这个module名字直接决定了你的import的路径。**

##### 3.2.2 调试依赖包

有时我们需要调试依赖包，此时就可以使用`replace`来修改依赖，如下所示：

```go
replace (
	github.com/google/uuid v1.1.1 => ../uuid
	golang.org/x/text v0.3.2 => github.com/golang/text v0.3.2
)
```

语句`github.com/google/uuid v1.1.1 => ../uuid`**使用本地的uuid来替换依赖包**，此时，我们**可以任意地修改`../uuid`目录的内容来进行调试。**

除了使用相对路径，还可以使用绝对路径，甚至还可以使用自已的fork仓库。

##### 3.2.3 禁止被依赖

另一种使用`replace`的场景是你的module不希望被直接引用，比如开源软件[kubernetes](https://github.com/kubernetes/kubernetes)，在它的**`go.mod`中`require`部分有大量的`v0.0.0`依赖**，比如：

```go
module k8s.io/kubernetes

require (
    ...
    k8s.io/api v0.0.0
    k8s.io/apiextensions-apiserver v0.0.0
    k8s.io/apimachinery v0.0.0
    k8s.io/apiserver v0.0.0
    k8s.io/cli-runtime v0.0.0
    k8s.io/client-go v0.0.0
    k8s.io/cloud-provider v0.0.0
    ...
)
```

由于**上面的依赖都不存在v0.0.0版本**，所以**其他项目直接依赖`k8s.io/kubernetes`时会因无法找到版本而无法使用**。
因为**Kubernetes不希望作为module被直接使用，其他项目可以使用kubernetes其他子组件**。

**kubernetes 对外隐藏了依赖版本号，其真实的依赖通过`replace`指定：**

```go
replace (
    k8s.io/api => ./staging/src/k8s.io/api
    k8s.io/apiextensions-apiserver => ./staging/src/k8s.io/apiextensions-apiserver
    k8s.io/apimachinery => ./staging/src/k8s.io/apimachinery
    k8s.io/apiserver => ./staging/src/k8s.io/apiserver
    k8s.io/cli-runtime => ./staging/src/k8s.io/cli-runtime
    k8s.io/client-go => ./staging/src/k8s.io/client-go
    k8s.io/cloud-provider => ./staging/src/k8s.io/cloud-provider
)
```

前面我们说过，**`replace`指令在当前模块不是`main module`时会被自动忽略的，Kubernetes正是利用了这一特性来实现对外隐藏依赖版本号来实现禁止直接引用的目的。**

### 4. exclude指令

`go.mod`文件中的**`exclude`指令用于排除某个包的特定版本**，其与`replace`类似，也**仅在当前module为`main module`时有效，其他项目引用当前项目时，`exclude`指令会被忽略。**

`exclude`指令在实际的项目中很少被使用，因为很少会显式地排除某个包的某个版本，除非我们知道某个版本有严重bug。
比如**指令`exclude github.com/google/uuid v1.1.0`，表示不使用v1.1.0 版本**。

下面我们还是使用`github.com/renhongcai/gomodule`来举例说明。

#### 4.1 排除指定版本

在 `github.com/renhongcai/gomodule`的v 1.3.0 版本中，我们的`go.mod`文件如下：

```go
module github.com/renhongcai/gomodule  

go 1.13  

require (  
  github.com/google/uuid v1.0.0  
  golang.org/x/text v0.3.2  
)  

replace golang.org/x/text v0.3.2 => github.com/golang/text v0.3.2
```

`github.com/google/uuid v1.0.0`说明我们期望使用 uuid包的`v1.0.0`版本。

假如，当前uuid仅有`v1.0.0` 、`v1.1.0`和`v1.1.1`三个版本可用，而且我们假定`v1.1.0`版本有严重bug。
此时**可以使用`exclude`指令将uuid的`v1.1.0`版本排除在外，即在`go.mod`文件添加如下内容：**

```go
exclude github.com/google/uuid v1.1.0
```

虽然我们暂时没有使用uuid的`v1.1.0`版本，但**如果将来引用了其他包，正好其他包引用了uuid的`v1.1.0`版本的话，此时添加的`exclude`指令就会跳过`v1.1.0`版本。**

下面我们创建`github.com/renhongcai/exclude`包来验证该问题。

#### 4.2 创建依赖包

为了进一步说明`exclude`用法，我们创建了一个仓库`github.com/renhongcai/exclude`，并在其中创建了一个module`github.com/renhongcai/exclude`，其中`go.mod`文件（v1.0.0版本）如下：

```go
module github.com/renhongcai/exclude

go 1.13

require github.com/google/uuid v1.1.0
```

可以看出其依赖`github.com/google/uuid` 的 `v1.1.0` 版本。创建`github.com/renhongcai/exclude`的目的是供`github.com/renhongcai/gomodule`使用的。

#### 4.3 使用依赖包

由于`github.com/renhongcai/exclude`也引用了uuid包且引用了更新版本的uuid，那么在`github.com/renhongcai/gomodule`引用`github.com/renhongcai/exclude`时，会**被动的提升uuid的版本**。

在**没有添加`exclude`之前，编译时`github.com/renhongcai/gomodule`依赖的uuid版本会提升到`v1.1.0`，与`github.com/renhongcai/exclude`保持一致，相应的`go.mod`也会被自动修改**，如下所示：

```go
module github.com/renhongcai/gomodule

go 1.13

require (
    github.com/google/uuid v1.1.0
    github.com/renhongcai/exclude v1.0.0
    golang.org/x/text v0.3.2
)

replace golang.org/x/text v0.3.2 => github.com/golang/text v0.3.2
```

但如果**添加了`exclude github.com/google/uuid v1.1.0` 指令**后，**编译时`github.com/renhongcai/gomodule`依赖的uuid版本会自动跳过`v1.1.0`，即选择`v1.1.1`版本**，相应的`go.mod`文件如下所示：

```go
module github.com/renhongcai/gomodule

go 1.13

require (
    github.com/google/uuid v1.1.1
    github.com/renhongcai/exclude v1.0.0
    golang.org/x/text v0.3.2
)

replace golang.org/x/text v0.3.2 => github.com/golang/text v0.3.2

exclude github.com/google/uuid v1.1.0
```

在本例中，在选择版本时，**跳过uuid `v1.1.0`版本后还有`v1.1.1`版本可用，Go 命令行工具可以自动选择`v1.1.1`版本，但如果没有更新的版本时将会报错而无法编译。**

### 5. indirect含义

在使用 Go module 过程中，随着引入的依赖增多，也许你会发现`go.mod`文件中部分依赖包后面会出现一个**`// indirect`的标识**。这个标识**总是出现在`require`指令中**，其中**`//`与代码的行注释一样表示注释的开始**，**`indirect`表示间接的依赖**。

比如开源软件 Kubernetes（v1.17.0版本）的 go.mod 文件中就有数十个依赖包被标记为`indirect`：

```go
require (
    github.com/Rican7/retry v0.1.0 // indirect
    github.com/auth0/go-jwt-middleware v0.0.0-20170425171159-5493cabe49f7 // indirect
    github.com/boltdb/bolt v1.3.1 // indirect
    github.com/checkpoint-restore/go-criu v0.0.0-20190109184317-bdb7599cd87b // indirect
    github.com/codegangsta/negroni v1.0.0 // indirect
    ...
)
```

在**执行命令`go mod tidy`时，Go module 会自动整理`go.mod 文件`**，如果**有必要会在部分依赖包的后面增加`// indirect`注释**。一般而言，**被添加注释的包肯定是间接依赖的包**，而**没有添加`// indirect`注释的包则是直接依赖的包，即明确的出现在某个`import`语句中**。

然而，这里需要着重强调的是：**并不是所有的间接依赖都会出现在 `go.mod`文件中**。

**间接依赖出现在`go.mod`文件的情况，可能符合下面所列场景的一种或多种**：

- 直接依赖的项目未启用 Go module
- 直接依赖的项目其go.mod 文件中缺失部分依赖

#### 5.1 直接依赖的项目未启用 Go module

如下图所示，**Module A 依赖 B，但是 B 还未切换成 Module，也即没有`go.mod`文件**，此时，当使用`go mod tidy`命令更新A的`go.mod`文件时，**B的两个依赖B1和B2将会被添加到A的`go.mod`文件中（前提是A之前没有依赖B1和B2）**，并且**B1 和B2还会被添加`// indirect`的注释**。

![null](https://www.topgoer.cn/uploads/gozhuanjia/images/m_82b92ed511cfce12ce51b43624e5563d_r.png)

此时Module A的`go.mod`文件中require部分将会变成：

```go
require (
    B vx.x.x
    B1 vx.x.x // indirect
    B2 vx.x.x // indirect
)
```

依赖B及B的依赖B1和B2都会出现在`go.mod`文件中。

#### 5.2 直接依赖的项目其 go.mod 文件不完整

如上面所述，如果依赖B没有`go.mod`文件，则Module A 将会把B的所有依赖记录到A 的`go.mod`文件中。**即便B拥有`go.mod`，如果`go.mod`文件不完整的话，Module A依然会记录部分B的依赖到`go.mod`文件中。**

如下图所示，**Module B虽然提供了`go.mod`文件中，但`go.mod`文件中只添加了依赖B1，那么此时A在引用B时，则会在A的`go.mod`文件中添加B2作为间接依赖，B1则不会出现在A的`go.mod`文件中。**

![null](https://www.topgoer.cn/uploads/gozhuanjia/images/m_f6e6b0f7d488c98682354d6e704aa965_r.png)

此时Module A的`go.mod`文件中require部分将会变成：

```go
require (
    B vx.x.x
    B2 vx.x.x // indirect
)
```

由于**B1已经包含进B的`go.mod`文件中，A的`go.mod`文件则不必再记录，只会记录缺失的B2。**

#### 5.3 总结

##### 5.3.1 为什么要记录间接依赖

在上面的例子中，如果某个依赖B 没有`go.mod`文件，在A 的`go.mod`文件中已经记录了依赖B及其版本号，为什么还要增加间接依赖呢？

我们知道Go module需要精确地记录软件的依赖情况，虽然此处记录了依赖B的版本号，但B的依赖情况没有记录下来，所以如果B的`go.mod`文件缺失了（或没有）这个信息，则需要在A的`go.mod`文件中记录下来。此时间接依赖的版本号将会根据Go module的版本选择机制确定一个最优版本。

##### 5.3.2 如何处理间接依赖

综上所述**间接依赖出现在`go.mod`中，可以一定程度上说明依赖有瑕疵，要么是其不支持Go module，要么是其`go.mod`文件不完整**。

由于Go 语言从v1.11版本才推出module的特性，众多开源软件迁移到go module还需要一段时间，在过渡期必然会出现间接依赖，但随着时间的推进，在`go.mod`中出现`// indirect`的机率会越来越低。

**出现间接依赖可能意味着你在使用过时的软件，如果有精力的话还是推荐尽快消除间接依赖。可以通过使用依赖的新版本或者替换依赖的方式消除间接依赖。**

##### 5.3.3 如何查找间接依赖来源

**Go module提供了`go mod why` 命令来解释为什么会依赖某个软件包**，若要查看`go.mod`中某个间接依赖是被哪个依赖引入的，可以**使用命令`go mod why -m <pkg>`来查看。**

比如，我们有如下的`go.mod`文件片断：

```go
require (
    github.com/Rican7/retry v0.1.0 // indirect
    github.com/google/uuid v1.0.0
    github.com/renhongcai/indirect v1.0.0
    github.com/spf13/pflag v1.0.5 // indirect
    golang.org/x/text v0.3.2
)
```

我们**希望确定间接依赖`github.com/Rican7/retry v0.1.0 // indirect`是被哪个依赖引入的**，则可以使用命令`go mod why`来查看：

```go
[root@ecs-d8b6 gomodule]# go mod why -m github.com/Rican7/retry
# github.com/Rican7/retry
github.com/renhongcai/gomodule
github.com/renhongcai/indirect
github.com/Rican7/retry
```

上面的打印信息中**`# github.com/Rican7/retry` 表示当前正在分析的依赖**，**后面几行则表示依赖链（上面的包依赖于下面的包）**。`github.com/renhongcai/gomodule` 依赖`github.com/renhongcai/indirect`，而`github.com/renhongcai/indirect`依赖`github.com/Rican7/retry`。由此我们就可以判断出**间接依赖`github.com/Rican7/retry`是被`github.com/renhongcai/indirect`引入的。**

另外，**命令`go mod why -m all`则可以分析所有依赖的依赖链**。

### 6. 版本选择机制

在前面的章节中，我们使用过`go get <pkg>`来获取某个依赖，如果没有特别指定依赖的版本号，`go get`会自动选择一个最优版本，并且如果本地有`go.mod`文件的话，还会自动更新`go.mod`文件。

事实上**除了`go get`，`go build`和`go mod tidy`也会自动帮我们选择依赖的版本**。**这些命令选择依赖版本时都遵循一些规则**，本节我们就开始介绍Go module涉及到的版本选择机制。

#### 6.1 依赖包版本约定

关于如何管理依赖包的版本，Go语言提供了一个规范，并且Go语言的演进过程中也一直遵循这个规范。

这个非强制性的规范主要围绕包的兼容性展开。对于如何处理依赖包的兼容性，根据是否支持Go module分别有不同的建议。

##### 6.1.1 Go module 之前版本兼容性

在Go v1.11（开始引入Go module的版本）之前，**Go 语言建议依赖包需要保持向后兼容**，这包括可导出的函数、变量、类型、常量等不可以随便删除。以函数为例，如果需要修改函数的入参，可以增加新的函数而不是直接修改原有的函数。

如果确实需要做一些打破兼容性的修改，建议创建新的包。

比如**仓库`github.com/RainbowMango/xxx`中包含一个package A，此时该仓库只有一个package**：

- `github.com/RainbowMango/xxx/A`

那么其他项目引用该依赖时的import 路径为：

```
import "github.com/RainbowMango/xxx/A"
```

**如果该依赖包需要引入一个不兼容的特性，可以在该仓库中增加一个新的package A1**，此时该仓库包含两个包：

- `github.com/RainbowMango/xxx/A`
- `github.com/RainbowMango/xxx/A1`

**那么其他项目在升级依赖包版本后不需要修改原有的代码可以继续使用package A，如果需要使用新的package A1，只需要将import 路径修改为`import "github.com/RainbowMango/xxx/A1"`并做相应的适配即可**。

##### 6.1.2 Go module 之后版本兼容性

从Go v1.11版本开始，随着Go module特性的引入，依赖包的兼容性要求有了进一步的延伸，Go module开始关心依赖包版本管理系统（如Git）中的版本号。尽管如此，兼容性要求的核心内容没有改变：

- 如果新package 和旧的package拥有相同的import 路径，那么新package必须兼容旧的package; 
- **如果新的package不能兼容旧的package，那么新的package需要更换import路径**；

在前面的介绍中，我们知道Go module 的`go.mod`中记录的module名字决定了import路径。例如，要引用module `module github.com/renhongcai/indirect`中的内容时，其import路径需要为`import github.com/renhongcai/indirect`。

**在Go module时代，module版本号要遵循语义化版本规范**，即版本号格式为`v<major>.<minor>.<patch>`，如v1.2.3。**当有不兼容的改变时，需要增加`major`版本号**，如v2.1.0。

Go module规定，**如果`major`版本号大于`1`**，**则`major`版本号需要显式地标记在module名字中，如`module github.com/my/mod/v2`**。这样做的好处是**Go module 会把`module github.com/my/mod/v2` 和 `module github.com/my/mod`视做两个module，他们甚至可以被同时引用**。

另外，**如果module的版本为`v0.x.x`或`v1.x.x`则都不需要在module名字中体现版本号**。

#### 6.2 版本选择机制

Go 的多个命令行工具都有自动选择依赖版本的能力，如`go build` 和`go mod tidy`，**当在源代码中增加了新的import，这些命令将会自动选择一个最优的版本，并更新`go.mod`文件。**

需要特别说明的是，**如果`go.mod`文件中已标记了某个依赖包的版本号，则这些命令不会主动更新`go.mod`中的版本号**。所谓**自动更新版本号只在`go.mod`中缺失某些依赖或者依赖不匹配时才会发生。**

##### 6.2.1 最新版本选择

当在源代码中**新增加了一个import**，比如：

```
import "github.com/RainbowMango/M"
```

如果**`go.mod`的require指令中并没有包含`github.com/RainbowMango/M`这个依赖，那么`go build` 或`go test`命令则会去`github.com/RainbowMango/M`仓库寻找最新的符合语义化版本规范的版本，比如v1.2.3，并在`go.mod`文件中增加一条require依赖：**

```
require github.com/RainbowMango/M v1.2.3
```

这里，**由于import路径里没有类似于`v2`或更高的版本号，所以版本选择时只会选择v1.x.x的版本，不会去选择v2.x.x或更高的版本**。

##### 6.2.2 最小版本选择

创建这样一种情况: 几个模块(A、B 和 C)依赖于同一个模块(D) ，但每个模块需要不同的版本。

```text
$ go list -m -versions github.com/sirupsen/logrus

github.com/sirupsen/logrus v0.1.0 v0.1.1 v0.2.0
v0.3.0 v0.4.0 v0.4.1 v0.5.0 v0.5.1 v0.6.0 v0.6.1
v0.6.2 v0.6.3 v0.6.4 v0.6.5 v0.6.6 v0.7.0 v0.7.1 
v0.7.2 v0.7.3 v0.8.0 v0.8.1 v0.8.2 v0.8.3 v0.8.4
v0.8.5 v0.8.6 v0.8.7 v0.9.0 v0.10.0 v0.11.0 v0.11.1
v0.11.2 v0.11.3 v0.11.4 v0.11.5 v1.0.0 v1.0.1 v1.0.3
v1.0.4 v1.0.5 v1.0.6 v1.1.0 v1.1.1 v1.2.0 v1.3.0
v1.4.0 v1.4.1 v1.4.2
```

清单 1 显示了模块 D 的所有版本，其中显示了“最新最大”的版本为 v1.4.2。

<img src="https://pic3.zhimg.com/80/v2-94218871b6d83adea5a79be4abc60752_720w.jpg" alt="img" style="zoom:50%;" />

图 1 显示了模块 A、B 和 C 各自独立地需要模块 D，并且**每个模块都需要不同版本的模块**。

如果项目现在只用到了A模块，应该为项目选择哪个版本的模块 D？实际上有两种选择。第一个选择是选择“最新最大”的版本(在这一行的主要版本 1 版本中) ，它将是版本 1.4.2。第二个选择是选择模块 A 需要的版本，即 v1.0.6 版本。

**Go 将尊重模块 A 的要求，选择 v1.0.6 版本**。**Go 为项目中需要该模块的所有依赖项选择当前在所需版本集中的“最小”版本**。换句话说，现在只有模块 A 需要模块 D，而模块 A 已经指定它需要版本 v1.0.6，因此这将作为模块 D 的版本。

如果我引入新的代码，要求项目导入模块 B，会怎么样？**一旦模块 B 被导入到项目中，Go 将该项目的模块 D 的版本从 v1.0.6 升级到  v1.2.0**。**再次为项目中需要模块 D 的所有依赖项(模块 A 和模块 B)选择模块 D 的“最小”版本**，该版本目前位于所需版本集(v1.0.6 和 v1.2.0 )中。

如果我**再次引入需要项目导入模块 C 的新代码**会怎么样？然后 **Go 将从所需的版本集(v1.0.6、 v1.2.0、  v1.3.2)中选择最新版本(v1.3.2)**。请注意，v1.3.2 版本仍然是“最小”版本，而不是模块 D (v1.4.2)的“最新最大”版本。

最后，**如果我删除刚刚为模块 C 添加的代码**会怎样？**Go 将把该项目锁定到模块 D 的版本 v1.3.2 中**，**而不是降级回到版本 v1.2.0**  。因为 Go 知道版本 v1.3.2 工作正常且稳定，因此版本 v1.3.2 仍然是该项目模块 D  的“最新非最大”或“最小”版本。另外，模块文件只维护一个快照，而不是日志。没有关于历史撤销或降级的信息。

### 7. incompatible

在前面的章节中，我们介绍了Go module的版本选择机制，其中介绍了一个**Module的版本号需要遵循`v<major>.<minor>.<patch>`的格式**，此外，**如果`major`版本号大于1时，其版本号还需要体现在Module名字中**。

**比如Module `github.com/RainbowMango/m`，如果其版本号增长到`v2.x.x`时，其Module名字也需要相应的改变为：**
**`github.com/RainbowMango/m/v2`。即，如果`major`版本号大于1时，需要在Module名字中体现版本。**

那么**如果Module的`major`版本号虽然变成了`v2.x.x`，但Module名字仍保持原样会怎么样呢？ 其他项目是否还可以引用呢？其他项目引用时有没有风险呢？**这就是今天要讨论的内容。

#### 7.1 能否引用不兼容的包

我们还是以Module `github.com/RainbowMango/m` 为例，**假如其当前版本为`v3.6.0`，因为其Module名字未遵循Golang所推荐的风格**，即**Module名中未附带版本信息**，我们**称这个Module为不规范的Module**。

**不规范的Module还是可以引用的，但跟引用规范的Module略有差别。**

如果我们在项目A中引用了该module，**使用命令`go mod tidy`**，go 命令会**自动查找Module m的最新版本**，即`v3.6.0`。
由于Module为不规范的Module，**为了加以区分，go 命令会在`go.mod`中增加`+incompatible` 标识。**

```
require (
    github.com/RainbowMango/m v3.6.0+incompatible
)
```

**除了增加`+incompatible`（不兼容）标识外，在其使用上没有区别。**

#### 7.2 如何处理incompatible

`go.mod`文件中出现`+incompatible`，说明你引用了一个不规范的Module，正常情况下，只能说明这个Module版本未遵循版本化语义规范。但**引用这个规范的Module还是有些困扰，可能还会有一定的风险**。

比如，我们拿某开源Module `github.com/blang/semver`为例，编写本文时，该Module最新版本为`v3.6.0`，但其`go.mod`中记录的Module却是：

```
module github.com/blang/semver
```

Module `github.com/blang/semver` 在另一个著名的开源软件`Kubernetes`（github.com/kubernetes/kubernetes）中被引用，那么`Kubernetes`的`go.mod`文件则会标记这个Module为`+incompatible`：

```go
require (
    ...
    github.com/blang/semver v3.5.0+incompatible
    ...
）
```

站在`Kubernetes`的角度，此处的困扰在于，**如果将来 `github.com/blang/semver`发布了新版本`v4.0.0`，但不幸的是Module名字仍然为`github.com/blang/semver`。那么，升级这个Module的版本将会变得困难。因为`v3.6.0`到`v4.0.0`跨越了大版本，按照语义化版本规范来解释说明发生了不兼容的改变，即然不兼容，项目维护者有必须对升级持谨慎态度，甚至放弃升级。**

站在`github.com/blang/semver`的角度，如果迟迟不能将自身变得”规范”，那么其他项目有可能放弃本Module，转而使用其他更规范的Module来替代，开源项目如果没有使用者，也就走到了尽头。

### 8. 伪版本

**在`go.mod`中通常使用语义化版本来标记依赖，比如`v1.2.3`、`v0.1.5`等**。因为`go.mod`文件通常是`go`命令自动生成并修改的，所以实际上是`go`命令习惯使用语义化版本。

诸如`v1.2.3`和`v0.1.5`这样的**语义化版本，实际是某个commit ID的标记，真正的版本还是commit ID**。比如**`github.com/renhongcai/gomodule`项目的`v1.5.0`对应的真实版本为`20e9757b072283e5f57be41405fe7aaf867db220`。**

由于**语义化版本比`commit ID`更直观（方便交流与比较版本大小），所以一般情况下使用语义化版本。**

#### 8.1 什么是伪版本

在实际项目中，**有时不得不直接使用一个`commit ID`**，比如**某项目发布了`v1.5.0`版本，但随即又修复了一个bug（引入一个新的commit ID），而且没有发布新的版本。此时，如果我们希望使用最新的版本，就需要直接引用最新的`commit ID`，而不是之前的语义化版本`v1.5.0`。**
**使用`commit ID`的版本在Go语言中称为`pseudo-version`，可译为”伪版本”**。

**伪版本的版本号通常会使用`vx.y.z-yyyymmddhhmmss-abcdefabcdef`格式**，其中**`vx.y.z`看上去像是一个真实的语义化版本，但通常并不存在该版本**，所以称为伪版本。另外**`abcdefabcdef`表示某个commit ID的前12位**，而**`yyyymmddhhmmss`则表示该commit的提交时间**，方便做版本比较。

使用伪版本的`go.mod`举例如下：

```go
...
require (
    go.etcd.io/etcd v0.0.0-20191023171146-3cf2f69b5738
)
...
```

#### 8.2 伪版本风格

伪版本格式都为`vx.y.z-yyyymmddhhmmss-abcdefabcdef`，但**`vx.y.z`部分在不同情况下略有区别，有时可能是`vx.y.z-pre.0`或者`vx.y.z-0`，甚至`vx.y.z-dev.2.0`等**。

**`vx.y.z`的具体格式取决于所引用`commit ID`之前的版本号**，如果所引用`commit ID`之前的最新的tag版本为`v1.5.0`，那么伪版本号则在其基础上增加一个标记，即`v1.5.1-0`，看上去像是下一个版本一样。

**实际使用中`go`命令会帮我们自动生成伪版本，不需要手动计算**，所以此处我们仅做基本说明。

#### 8.3 如何获取伪版本

我们使用具体的例子还演示如何使用伪版本。在仓库`github.com/renhongcai/gomodule`中存在`v1.5.0` tag 版本，在`v1.5.0`之后又提交了一个commit，并没有发布新的版本。其版本示意图如下：

![null](https://www.topgoer.cn/uploads/gozhuanjia/images/m_7eb056cd2a4e4f96cfb5a405fa40c227_r.png)

为了方便描述，我们**把`v1.5.0`对应的commit 称为`commit-A`**，而其**随后的commit称为`commit-B`**。

如果我们**要使用commit-A，即`v1.5.0`，可使用`go get github.com/renhongcai/gomodule@v1.5.0`命令**：

```go
[root@ecs-d8b6 ~]# go get github.com/renhongcai/gomodule@v1.5.0
go: finding github.com/renhongcai/gomodule v1.5.0
go: downloading github.com/renhongcai/gomodule v1.5.0
go: extracting github.com/renhongcai/gomodule v1.5.0
go: finding github.com/renhongcai/indirect v1.0.1
```

此时，如果存在`go.mod`文件，`github.com/renhongcai/gomodule`体现在`go.mod`文件的版本为`v1.5.0`。

**如果我们要使用commit-B，可使用`go get github.com/renhongcai/gomodule@6eb27062747a458a27fb05fceff6e3175e5eca95`命令（可以使用完整的commit id，也可以只使用前12位）**：

```go
[root@ecs-d8b6 ~]# go get github.com/renhongcai/gomodule@6eb27062747a458a27fb05fceff6e3175e5eca95
go: finding github.com 6eb27062747a458a27fb05fceff6e3175e5eca95
go: finding github.com/renhongcai/gomodule 6eb27062747a458a27fb05fceff6e3175e5eca95
go: finding github.com/renhongcai 6eb27062747a458a27fb05fceff6e3175e5eca95
go: downloading github.com/renhongcai/gomodule v1.5.1-0.20200203082525-6eb27062747a   //伪版本
go: extracting github.com/renhongcai/gomodule v1.5.1-0.20200203082525-6eb27062747a
go: finding github.com/renhongcai/indirect v1.0.2
```

此时，可以看到生成的伪版本号为`v1.5.1-0.20200203082525-6eb27062747a`，当前最新版本为`v1.5.0`，`go`命令生成伪版本号时自动增加了版本。此时，如果存在`go.mod`文件的话，`github.com/renhongcai/gomodule`体现在`go.mod`文件中的版本则为该伪版本号。

### 9. 依赖包存储(获取的依赖包的存储方式和位置)

在前面介绍`GOPATH`的章节中，我们提到`GOPATH`模式下不方便使用同一个依赖包的多个版本。在`GOMODULE`模式下这个问题得到了很好的解决。

**`GOPATH`模式下，依赖包存储在`$GOPATH/src`，该目录下只保存特定依赖包的一个版本**，而在**`GOMODULE`模式下，依赖包存储在`$GOPATH/pkg/mod`，该目录中可以存储特定依赖包的多个版本**。

需要注意的是`$GOPATH/pkg/mod`目录下有个`cache`目录，它用来存储依赖包的缓存，简单说，`go`命令每次下载新的依赖包都会在该`cache`目录中保存一份。关于该目录的工作机制我们留到`GOPROXY`章节时再详细介绍。

接下来，我们使用开源项目`github.com/google/uuid`为例分别说明`GOPATH`模式和`GOMODULE`模式下特定依赖包存储机制。在下面的操作中，我们会使用`GO111MODULE`环境变量控制具体的模式：

- `export GO111MODULE=off`切换到`GOPATH`模式
- `export GO111MODULE=on`切换到`GOMODULE`模式。

#### 9.1 GOPATH 依赖包存储

为了实验`GOPATH`模式下依赖包的存储方式，我们可以使用以下命令来获取`github.com/google/uuid`：

```
# export GO111MODULE=off
# go get -v github.com/google/uuid
```

在`GOPATH`模式下，**`go get`命令会将依赖包下载到`$GOPATH/src/google`目录中**。

该命令**等同于在`$GOPATH/src/google`目录下执行`git clone https://github.com/google/uuid.git`**，也就是**`$GOPATH/src/google/uuid`目录中存储的是完整的仓库**。

#### 9.2 GOMODULE 依赖包存储

为了实验`GOMODULE`模式下依赖的存储方式，我们使用以下命令来获取`github.com/google/uuid`：

```
# export GO111MODULE=on
# go get -v github.com/google/uuid
# go get -v github.com/google/uuid@v1.0.0
# go get -v github.com/google/uuid@v1.1.0
# go get -v github.com/google/uuid@v1.1.1
```

在`GOMODULE`模式下，`go get`命令会**将依赖包下载到`$GOPATH/pkg/mod`目录下**，并且**按照依赖包的版本分别存放**。（注：`go get`命令**不指定特定版本时，默认会下载最新版本**，即v1.1.1，如软件包有新版本发布，实验结果将有所不同。）

此时`$GOPATH/pkg/mod`目录结构如下：

```
${GOPATH}/pkg/mod/github.com/google
├── uuid@v1.0.0
├── uuid@v1.1.0
├── uuid@v1.1.1
```

相较于`GOPATH`模式，`GOMODULE`有两处不同点：

- 一是**依赖包的目录中包含了版本号，每个版本占用一个目录**；
- 二是依赖包的特定版本目录中**只包含依赖包文件，不包含`.git`目录**；

由于依赖包的每个版本都有一个唯一的目录，所以在**多项目场景中需要使用同一个依赖包的多版本时才不会产生冲突**。另外，由于依赖包的**每个版本都有唯一的目录**，也**表示该目录内容不会发生改变**，也就**不必再存储其位于版本管理系统(如git)中的信息**。

#### 9.3 包名大小写敏感问题

有时我们使用的**包名中会包含大写字母**，比如`github.com/Azure/azure-sdk-for-go`，**`GOMODULE`模式下，在存储时会将包名做大小写编码处理，即每个大写字母将变与`!`+相应的小写字母**，比如`github.com/Azure`包在存储时将会被放置在`$GOPATH/pkg/mod/github.com/!azure`目录中。

需要注意的是，`GOMODULE`模式下，我们使用`go get`命令时，**如果不小心将某个包名大小写搞错**，比如`github.com/google/uuid`写成`github.com/google/UUID`时，在**存储依赖包时会严格按照`go get`命令指示的包名进行存储**。

如下所示，使用大写的`UUID`:

```go
[root@ecs-d8b6 uuid]# go get -v github.com/google/UUID@v1.0.0
go: finding github.com v1.0.0
go: finding github.com/google v1.0.0
go: finding github.com/google/UUID v1.0.0
go: downloading github.com/google/UUID v1.0.0
go: extracting github.com/google/UUID v1.0.0
github.com/google/UUID
```

由于`github.com/google/uuid`**域名不区分大小写**，所以使用`github.com/google/UUID`**下载包时仍然可以下载**，但在**存储时将会严格区分大小写**，此时`$GOPATH/pkg/mod/google/`目录下将会多出一个[d@v1.0.0](mailto:`!u!u!i!"">`!u!u!i!d@v1.0.0`目录：

```go
${GOPATH}/pkg/mod/github.com/google
├── uuid@v1.0.0
├── uuid@v1.1.0
├── uuid@v1.1.1
├── !u!u!i!d@v1.0.0    //错误的大小写存储了错误的包名
```

在`go get`中使用错误的包名，除了会增加额外的不必要存储外，还可能会影响`go`命令解析依赖，还可能将错误的包名使用到`import`指令中，所以在实际使用时应该尽量避免。

### 10. go.sum文件

为了确保一致性构建，Go引入了`go.mod`文件来标记每个依赖包的版本，在构建过程中`go`命令会下载`go.mod`中的依赖包，下载的依赖包会缓存在本地，以便下次构建。
**考虑到下载的依赖包有可能是被黑客恶意篡改的，以及缓存在本地的依赖包也有被篡改的可能，单单一个`go.mod`文件并不能保证一致性构建。**

为了解决Go module的这一安全隐患，Go开发团队在引入`go.mod`的同时也引入了**`go.sum`文件**，用于**记录每个依赖包的哈希值**，在构建时，**如果本地的依赖包hash值与`go.sum`文件中记录得不一致，则会拒绝构建**。

本节暂不对模块校验细节展开介绍，只从日常应用层面介绍：

- go.sum 文件记录含义
- go.sum文件内容是如何生成的
- go.sum是如何保证一致性构建的

#### 10.1 go.sum文件记录

`go.sum`文件中每行记录由`module`名、版本和哈希组成，并由空格分开：

```
<module> <version>[/go.mod] <hash>
```

比如，某个`go.sum`文件中记录了`github.com/google/uuid` 这个依赖包的`v1.1.1`版本的哈希值：

```
github.com/google/uuid v1.1.1 h1:Gkbcsh/GbpXz7lPftLA3P6TYMwjCLYm83jiFQZF/3gY=  
github.com/google/uuid v1.1.1/go.mod h1:TIyPZe4MgqvfeYDBFedMoGGpEw/LqOeaOT+nhxU+yHo=
```

在Go module机制下，我们需要**同时使用依赖包的名称和版本**才可以准确的描述一个依赖，为了方便叙述，**下面我们使用`依赖包版本`来指代依赖包名称和版本**。

正常情况下，**每个`依赖包版本`会包含两条记录**，**第一条记录为该`依赖包版本`整体（所有文件）的哈希值**，**第二条记录仅表示该`依赖包版本`中`go.mod`文件的哈希值**，如果该`依赖包版本`没有`go.mod`文件，则只有第一条记录。**如上面的例子中，`v1.1.1`表示该`依赖包版本`整体，而`v1.1.1/go.mod`表示该`依赖包版本`中`go.mod`文件。**

**`依赖包版本`中任何一个文件（包括`go.mod`）改动，都会改变其整体哈希值**，此处再额外记录`依赖包版本`的`go.mod`文件主要用于计算依赖树时不必下载完整的`依赖包版本`，只根据`go.mod`即可计算依赖树。

每条记录中的哈希值前均有一个表示哈希算法的`h1:`，表示后面的哈希值是由算法`SHA-256`计算出来的，自Go module从v1.11版本初次实验性引入，直至v1.14 ，只有这一个算法。

此外，细心的读者或许会发现**`go.sum`文件中记录的`依赖包版本`数量往往比`go.mod`文件中要多**，这是因为二者记录的粒度不同导致的。`go.mod`只需要记录**直接依赖的**`依赖包版本`，只在**`依赖包版本`不包含`go.mod`文件时候**才会记录**间接`依赖包版本`**，而**`go.sum`则是要记录构建用到的所有`依赖包版本`**。

#### 10.2 生成

假设我们在开发某个项目，当我们在GOMODULE模式下引入一个新的依赖时，通常会使用`go get`命令获取该依赖，比如：

```
go get github.com/google/uuid@v1.0.0
```

**`go get`命令首先会将该依赖包下载到本地缓存目录`$GOPATH/pkg/mod/cache/download`，该依赖包为一个后缀为`.zip`的压缩包，如`v1.0.0.zip`。`go get`下载完成后会对该`.zip`包做哈希运算，并将结果存放在后缀为`.ziphash`的文件中，如`v1.0.0.ziphash`。如果在项目的根目录中执行`go get`命令的话，`go get`会同步更新`go.mod`和`go.sum`文件，`go.mod`中记录的是依赖名及其版本**，如：

```
require (
    github.com/google/uuid v1.0.0
)
```

`go.sum`文件中则会记录依赖包的哈希值（同时还有依赖包中go.mod的哈希值），如：

```
github.com/google/uuid v1.0.0 h1:b4Gk+7WdP/d3HZH8EJsZpvV7EtDOgaZLtnaNGIu1adA=
github.com/google/uuid v1.0.0/go.mod h1:TIyPZe4MgqvfeYDBFedMoGGpEw/LqOeaOT+nhxU+yHo=
```

值得一提的是，**在更新`go.sum`之前，为了确保下载的依赖包是真实可靠的，`go`命令在下载完依赖包后还会查询`GOSUMDB`环境变量所指示的服务器，以得到一个权威的`依赖包版本`哈希值。如果`go`命令计算出的`依赖包版本`哈希值与`GOSUMDB`服务器给出的哈希值不一致，`go`命令将拒绝向下执行，也不会更新`go.sum`文件。**

`go.sum`存在的意义在于，我们希望别人或者在别的环境中构建当前项目时所使用依赖包跟`go.sum`中记录的是完全一致的，从而达到一致构建的目的。

#### 10.3 校验

假设我们**拿到某项目的源代码并尝试在本地构建**，`go`命令会**从本地缓存中查找所有`go.mod`中记录的依赖包，并计算本地依赖包的哈希值，然后与`go.sum`中的记录进行对比**，即**检测本地缓存中使用的`依赖包版本`是否满足项目`go.sum`文件的期望**。

如果校验失败，说明本地缓存目录中`依赖包版本`的哈希值和项目中`go.sum`中记录的哈希值不一致，`go`命令将拒绝构建。
这就是`go.sum`存在的意义，即如果不使用我期望的版本，就不能构建。

当校验失败时，**有必要确认到底是本地缓存错了，还是`go.sum`记录错了**。
需要说明的是，二者都可能出错，本地缓存目录中的`依赖包版本`有可能被有意或无意地修改过，项目中`go.sum`中记录的哈希值也可能被篡改过。

当校验失败时，**`go`命令倾向于相信`go.sum`**，因为一个新的`依赖包版本`在被添加到`go.sum`前是经过`GOSUMDB`（校验和数据库）验证过的。此时即便系统中配置了`GOSUMDB`（校验和数据库），`go`命令也不会查询该数据库。

#### 10.4 校验和数据库

**环境变量`GOSUMDB`标识一个`checksum database`，即校验和数据库**，实际上是一个web服务器，**该服务器提供查询`依赖包版本`哈希值的服务**。

该数据库中记录了很多`依赖包版本`的哈希值，比如Google官方的`sum.golang.org`则记录了所有的可公开获得的`依赖包版本`。除了使用官方的数据库，还**可以指定自行搭建的数据库，甚至干脆禁用它（`export GOSUMDB=off`）**。

如果系统配置了`GOSUMDB`，在`依赖包版本`被写入`go.sum`之前会向该数据库查询该`依赖包版本`的哈希值进行二次校验，校验无误后再写入`go.sum`。

**如果系统禁用了`GOSUMDB`，在`依赖包版本`被写入`go.sum`之前则不会进行二次校验，`go`命令会相信所有下载到的依赖包，并把其哈希值记录到`go.sum`中**。