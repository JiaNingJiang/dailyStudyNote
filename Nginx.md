### 一、Nginx的环境搭建与安装

#### 1. 环境搭建

需要安装的编译工具有：

```
gcc g++ autoconf  automake make
```

需要的第三方库文件有:

```
zlib  openssl  pcre  httpd-tools
```

```
sudo apt-get install libpcre3 libpcre3-dev
sudo apt-get install zlib1g zlib1g-dev

sudo apt-get install openssl 
sudo apt-get install libssl-dev
```

#### 2. 安装编译 Nginx 源码

Nginx开源版官网：nginx.org  或者安装淘宝的tengine ：http://tengine.taobao.org/

```
wget  http://nginx.org/download/nginx-1.23.1.tar.gz

wget  http://tengine.taobao.org/download/tengine-2.3.3.tar.gz
```

```
cd /usr/local
mkdir nginx
cd nginx
wget http://nginx.org/download/nginx-1.23.1.tar.gz
tar -xvf nginx-1.13.7.tar.gz 
```

编译：

```shell
# 进入nginx目录
/usr/local/nginx/nginx-1.13.7
# 执行命令
./configure    // ./configure --help 可以查看各种编译帮助信息(可以开启Nginx的特殊功能)
# 执行make命令
make
# 执行make install命令
make install
```

自定义configure：

```shell
./configure --prefix=/home/jiang/nginx --with-http_ssl_module --with-http_flv_module --with-http_gzip_static_module --with-http_stub_status_module --with-threads --with-file-aio
```

执行configure生成的makefile文件

```shell
make 
make install
```

#### 3. 查看安装后的Nginx目录

```shell
jiang@jiang-virtual-machine:~/nginx$ ls
conf  html  logs  sbin
```

```shell
conf : 存放nginx的配置文件,如 nginx.conf
html : 存放nginx的网页根目录文件,存放站点的静态文件数据
logs : 存放nginx的各种日志目录
sbin : 存放nginx的可执行文件 
```

#### 4. 将nginx可执行文件添加到PATH环境变量下

 若是直接使用 nginx 指令，默认是调用nginx程序

```shell
jiang@jiang-virtual-machine:~/nginx$ nginx
程序 'nginx' 已包含在下列软件包中：
 * nginx-core
 * nginx-extras
 * nginx-full
 * nginx-light
请尝试：sudo apt install <选定的软件包>
```

出现上述情况，表明nginx可执行文件没有添加到PATH环境变量下(但是可以通过绝对路径调用)。

配置PATH变量，将nginx的sbin目录添加到PATH中

```shell
vim /etc/profile
export PATH="$PATH:/usr/local/nginx/sbin"

source /etc/profile  ##重新加载配置文件信息
```

配置完成之后,可以使用nginx的各种指令了：

```shell
nginx  ##启动nginx
```

可以通过以下命令查看网络连接情况:

```shell
netstat -tunlp | grep 80  ##nginx是web服务器,因此会占用80端口
```

通过以下命令查看nginx的进程信息:

```shell
ps -aux | grep nginx
```

停止nginx服务

```shell
nginx -s stop
```

如果不想重启nginx就重新加载配置文件：

```shell
nginx -s reload  ##平滑重启
```

### 二、Nginx配置文件语法