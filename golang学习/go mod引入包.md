### 一、引入本地自定义包

> 首先必须开启go mod模式：go env -w GO111MODULE=on

这里我们以实战进行展示：

1. 假设F:\数据结构\GO\SimpleStruct目录下有两个包： list和stack

​	![image-20220701090843178](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\image-20220701090843178.png)

2. stack包下的 linkedStack.go 需要引用list包
3. 在stack包和list包的共同上级目录SimpleStack文件路径下调用go mod init GO/SimpleStruct生成go.mod文件

![image-20220701091619791](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\image-20220701091619791.png)

![image-20220701091927025](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\image-20220701091927025.png)

4. 在被引入的包 list 包 目录路径下调用 go build 命令(将list包编译到缓存中，就可以被其他包引用了)
5. 接下来就可以在stack文件夹下的 linkedStack.go 中引入 list 包了：

![image-20220701091824073](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\image-20220701091824073.png)

### 二、引用网络包

 1.首先我们需要登录到 golang.org 网站上。点击右上栏的 package 搜索需要引用的包

 2.假设我们要使用 gin 包

![image-20220701092452608](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\image-20220701092452608.png)

 3.左侧README下的Installation会介绍如何对该包进行引用

![image-20220701092907335](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\image-20220701092907335.png)

4.Quick start 会介绍如何进行使用

5.下面以实战方式介绍如何引用 gin 包：

①假设SimpleStruct目录下的list包和stack包都需要导入gin包

②我们需要在 go mod所在目录，也就是SimpleStruct下cmd调用 go get -u github.com/gin-gonic/gin 命令

![image-20220701093244093](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\image-20220701093244093.png)

③之后将会自动下载所有所需包，同时自动更新 go.mod文件 以及加入新的go.sum文件

![image-20220701093358239](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\image-20220701093358239.png)

![image-20220701093416792](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\image-20220701093416792.png)

④假设还是 stack 包下的 linkedStack.go 需要引用 gin 包

![image-20220701093648759](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\image-20220701093648759.png)

⑤如果发现import时的路径存在问题，可以在 stack目录下调用 go mod tidy 命令(必须完成一次go get将gin包以及下载到本地了) (go mod tidy 负责从本地搜索import的包资源，更新 go.mod文件)

![image-20220701094039427](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\image-20220701094039427.png)