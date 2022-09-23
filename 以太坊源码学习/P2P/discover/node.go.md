### 1.Node结构体

```go
type Node struct {
	IP  net.IP // len 4 for IPv4 of 16 for IPv6
	UDP uint16 // UDP port number
	TCP uint16 // TCP port number
	ID  NodeID // the node's public key (ECC public key marshalled)
	sha common.Hash //NodeID的哈希值
}

```

### 2.函数

```go
//新建一个Node类型的节点
func newNode(id NodeID, ip net.IP, udpPort, tcpPort uint16) *Node
```

```go
//对字符串形式的url进行解析，将提取的信息进行组合构成Node结构体
//(取出各部分：NodeID , IP , UDP Port , TCP Port)
func ParseNode(rawurl string) (*Node, error)
```

```go
//底层调用了ParseNode()函数，也是对url进行解析。
//与ParseNode()不同的是增加了解析出错发出panic警报的功能
func MustParseNode(rawurl string) *Node
```

```go
//将一个表示十六进制数字的字符串转换为NodeID (字符串可以以0x开头)
func HexID(in string) (NodeID, error)
```

```go
//底层调用了HexID()，增加了转换失败时发出panic警报
func MustHexID(in string) NodeID
```

```go
//将传入的公钥采用椭圆曲线算法序列化后得到NodeID
func PubkeyID(pub *ecdsa.PublicKey) NodeID
```

```go
//通过数字签名与获取消息的哈希计算出签名者的公钥
func recoverNodeID(hash, sig []byte) (id NodeID, err error)
```

```go
//分别比较目标target与a和b的距离。如果a距离目标更近就返回-1  如果b距离目标更近就返回1  如果距离相等就返回0。
//距离的计算是通过异或比较哈希值，异或之后的结果越大说明距离target越远
func distcmp(target, a, b common.Hash) int

//需要注意：common.Hash是byte数组，因此每次异或比较都是对每一个元素byte进行比较(而不是以bit为单位进行比较)
//从比较的方式来看，common.Hash是小端字节序。common.Hash[0]存储最高位的哈希值
```

```go
//返回a与b的对数距离 --> log2(a ^ b) 。 这个距离是指异或后的前缀0的个数
func logdist(a, b common.Hash) int  

//返回值是a与b不相等的bit位数(异或后去除前缀0的剩余bit位数)
```

```go
//本函数返回一个满足 logdist(a, b) == n 的随机哈希值 b (哈希值a与距离n是已知的)
func hashAtDistance(a common.Hash, n int) (b common.Hash)
```

### 3.Node结构体的方法

```go
//以字符串形式返回NodeID
func (n NodeID) String() string
```

```go
//根据节点的ID值计算其哈希值，赋给n.sha
func (n *Node) Sha()
```

```go
//获取节点的UDP地址
func (n *Node) addr() *net.UDPAddr
```

```go
//以字符串形式返回节点n的url信息
func (n *Node) String() string

//注：url格式为 -->scheme://userinfo@host/path?query#fragment
//本函数省略了部分字段，返回格式为“dpnode://User@Host?discport=n.UDP
```

### 4.NodeID类型变量及其方法

```go
type NodeID [nodeIDBits / 8]byte
```

```go
//用NodeID创建一个椭圆曲线结构体公钥
func (id NodeID) Pubkey() (*ecdsa.PublicKey, error)
```

