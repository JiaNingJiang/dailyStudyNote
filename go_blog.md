## 一、项目架构

项目主要分为以下几个模块：

```
go-gin-example/
   ├── conf
   ├── middleware
   ├── models
   ├── pkg
   ├── routers
   └── runtime
```

- conf：用于存储配置文件
- middleware：应用中间件
- models：应用数据库模型
- pkg：第三方包
- routers 路由逻辑处理
- runtime：应用运行时数据

功能：

用户可以基于 Web 接口，对数据库中的**文章表**和**标签表**进行增删改查。

文章和标签的关系：每个文章有自己的标签，因此在执行查询的时候，能够达到`Article`、`Tag`关联查询的功能

## 二、配置文件读取模块

整个系统的配置文件被存储在 `/conf/app.ini` 文件下。包括 `MySQL` 数据库相关信息、`Redis` 数据库相关信息、`Web` 服务相关信息、其他功能模块（日志模块、图片存储查询模块、`JWT` 模块等等）的相关信息。

采用第三方包 `go-ini` 完成对配置文件`app.ini` 的读取，因为是导入的第三方库，因此整个程序逻辑放置在 `pkg/setting` 目录下。

从配置文件中读取到的配置信息以**全局对象**的形式存在，包括`setting/Server`（包括 `Web` 服务相关配置项），`setting/Database`（包括`MySQL` 相关配置项），`setting/Redis`（包括 `Redis`  相关配置项），`setting/App`（包括其他功能模块的相关配置项）

## 三、编写 `Web` 接口的错误码包

用户在访问 `Web` 服务的时候，可能遇到各种各种的错误，当前包将所有可能出现的错误与错误码相关联，同时将冗长的错误信息用简短的字符常量表示。

- `pkg/e/msg.go` ：将冗长的错误信息用简短的字符常量表示。
- `pkg/e/code.go`：将错误与错误码相关联。

当服务器在返回错误的使用，同时向客户端返回错误码和具体的错误内容。

## 四、`Model` 层设计

### 4.1 连接并配置数据库

从 `setting/DatabaseSetting` 中完成对 `MySQL` 配置项的获取。

基于 `gorm` 库完成对 `MySQL` 数据库的连接和配置，生成一个数据库对象 `db *gorm.DB`

### 4.2 编写数据库对象的 `Hook` 函数

每当我们新生成或者更新了一条 `SQL` 记录，就要记录下当前的产生时间或者修改时间。

因此对于每个 `SQL` 对象都要创建两个字段：`CreatedOn` 字段和 `ModifiedOn` 字段。创建 `SQL` 记录的时候这两个字段都需要修改，如果只是更新操作则只需要更新后者，为了避免手动更新，将此操作注册为 `Hook` 函数：

- `updateTimeStampForCreateCallback`：插入 `SQL` 时的 `Hook`
- `updateTimeStampForUpdateCallback`：更新 `SQL` 时的 `Hook`

这两个 `Hook` 函数将会在实际向数据库插入或更新数据之前被调用。 

### 4.2 编写标签类的 `Models` 逻辑

新建 `Tag` 类对象与 `blog_tag` 表中的记录进行关联。

```go
type Tag struct { // 无论是在查询还是插入时，Tag类的对象都会直接在blog_tag表中完成映射(前缀blog_是自行指定的)
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}
```

`Tag` 对象的方法完成对 `blog_tag` 表中记录的"增删改查"

### 4.3 编写文章类的 `Models` 逻辑

新建 `Article` 类对象与 `blog_article` 表中的记录进行关联。

```go
type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`  // 外键
	Tag   Tag `json:"tag"`   // 与标签表进行关联查询

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}
```

- `Article`有一个结构体成员是**`TagID`，就是外键**。`gorm`会通过**类名+ID** 的方式去找到这**两个类之间的关联关系**
- `Article`有一个结构体成员是`Tag`，就是我们**嵌套在`Article`里的`Tag`结构体**，我们可以**通过`Related`进行关联查询**

`Article` 对象的方法完成对 `blog_article` 表中记录的"增删改查"

如何进行关联查询？

- 针对一条数据的关联查询：

```go
func GetArticle(id int) (article Article) {
    db.Where("id = ?", id).First(&article)
    db.Model(&article).Related(&article.Tag)  // 获取与之关联的Tag的SQL记录，存储在 Tag 字段中
    return
}
```

- 针对多条数据同时关联查询

```go
func GetArticles(pageNum int, pageSize int, maps interface {}) (articles []Article) {
    db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

    return
}
```

### 4.4 编写用户类的 `Models` 逻辑

存在的目的是进行登录验证，每一个注册用户都有自己的用户名和密码，在进行 `JWT` 验证的时候，用户名存在且密码正确才能登录成功，然后采取后续行为。

```go
type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
```



## 五、`Web` 层路由设计

### 5.1 读取 `Web` 配置并定义接口

从 `setting/ServerSetting` 中完成对 `Web` 服务配置项的获取。

基于 `Gin` 框架完成对接口服务的定义和注册：

分为**标签类**的接口和**文章类**的接口

```
获取标签列表：GET(“/tags”)
新建标签：POST(“/tags”)
更新指定标签：PUT(“/tags/:id”)
删除指定标签：DELETE(“/tags/:id”)

获取文章列表：GET(“/articles”)
获取指定文章：POST(“/articles/:id”)
新建文章：POST(“/articles”)
更新指定文章：PUT(“/articles/:id”)
删除指定文章：DELETE(“/articles/:id”)
```

### 5.2 编写标签类的路由处理函数

针对定义的接口，实现具体的路由处理函数：

解析用户的 URL 以及 POST 表单，调用 `Model` 层提供的方法，从对应的表中获取记录，返回给用户。

```
//获取多个文章标签
func GetTags(c *gin.Context) {
}

//新增文章标签
func AddTag(c *gin.Context) {
}

//修改文章标签
func EditTag(c *gin.Context) {
}

//删除文章标签
func DeleteTag(c *gin.Context) {
}
```

### 5.3 编写文章类的路由处理函数

## 六、使用 `JWT` 完成身份认证

### 6.1 `JWT` 生成

客户端实现会在服务端上进行注册，其用户名和密码会在服务端留存。

客户端在**第一次登录的时候**，需要先依据自己的用户名和密码，申请获取一个 `token`，如果服务器对客户端发送来的用户名和密码完成验证，则执行以下操作：

服务端依据客户端的**用户名、密码**，通过**指定的Hash计算方法**对公共的 `jwt` 密码进行**签名**获取**单向性的token**，然后将其返回给用户

### 6.2 `JWT` 认证

后续客户端在每次向服务端发送 HTTP 请求的时候，都需要将自身获取的 `token` 一并发送，当且仅当服务器验证了 `token` 的有效性，才会对客户端的后续请求进行响应。

为了实现上述效果，需要将 `JWT` 验证服务作为一个 `gin` 的中间件进行注册：

```go
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT()) // 添加中间件,验证 token 的有效性
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag) // id 是动态params
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag) // id 是动态params

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)

	}
```

