### 一、基础使用

在线编译运行平台：https://remix.ethereum.org

#### 1. SPDX 版权许可标识

每个源文件都应该以注释开始以说明其版权许可证。

若是开源:

```solidity
// SPDX-License-Identifier: MIT
```

若是不开源

```solidity
// UNLICENSED
```



#### 2. 标注使用的solidity版本号

```solidity
pragma solidity >=0.4.16 <0.9.0;     //表示版本号大于等于0.4.16 ,小于 0.9.0

pragma solidity ^0.4.16;     //表示版本号不兼容0.4.16一下的版本和0.5.0之上的版本
```

关键字 `pragma` 版本标识指令，用来启用某些编译器检查， 版本 标识（pragma） 指令通常只对本文件有效，所以我们需要把这个版本 标识（pragma） 添加到项目中所有的源文件。 如果使用了 [import 导入](https://learnblockchain.cn/docs/solidity/layout-of-source-files.html#import) 其他的文件, 标识（pragma） 并不会从被导入的文件加入到导入的文件中。

使用版本标准不会改变编译器的版本，它不会启用或关闭任何编译器的功能。 他仅仅是告知编译器去检查版本是否匹配， 如果不匹配，编译器就会提示一个错误。

#### 3. 使用contract关键字生成合约

```solidity
contract HelloWorld {   //每一条语句都必须以分号“;”结尾
    uint256 storedData;   // 状态变量,是永久地存储在合约存储中的值。

	// 函数是代码的可执行单元。函数通常在合约内部定义，但也可以在合约外定义。
    function set(uint256 x) public {   //public关键字表示该合约函数可以被任何人调用
        storedData = x;
    }

    function get() public view returns (uint256) {   //view关键字表示调用该合约函数不消耗任何gas
        return storedData;
    }

    function pureTest(uint64 name) pure public returns(uint64){  //pure关键字也表示调用该函数不消耗任何gas
        return name;
    }

    function pureTest1(string memory name) pure public returns(string memory){ //字符串型变量后面必须用memory修饰
        return name;  //输入string类型变量时必须用双引号""格式
    }

}
```

#### 4. 导入其他源文件

在全局层面上，可使用如下格式的导入语句：

```solidity
import "filename";  
```

`filename` 部分称为导入路径（ *import path*）。 此语句将从 “filename” 中**导入所有的全局符号到当前全局作用域**中



### 二、数据类型

#### 1. bool 类型

```solidity
bool a;   //bool型变量默认值为false

!a   //翻转bool类型变量的值

&&  ||  ! //关系运算符 与  或  非
```

#### 2. 整型

`int` / `uint` ：分别表示有符号和无符号的不同位数的整型变量。 支持关键字 `uint8` 到 `uint256` （无符号，从 8 位到 256 位）以及 `int8` 到 `int256`，**以 `8` 位为步长递增**。 **`uint` 和 `int` 分别是 `uint256` 和 `int256` 的别名**。

运算符：

- 比较运算符： `<=` ， `<` ， `==` ， `!=` ， `>=` ， `>` （返回布尔值）
- 位运算符： `&` ， `|` ， `^` （异或）， `~` （位取反）
- 移位运算符： `<<` （左移位） ， `>>` （右移位）
- 算数运算符： `+` ， `-` ， 一元运算负 `-` （仅针对有符号整型）， `*` ， `/` ， `%` （取余或叫模运算） ， `**` （幂）

对于整形 `X`，可以使用 `type(X).min` 和 `type(X).max` 去获取**这个类型**的最小值与最大值。

> Solidity中的整数是有取值范围的。 例如 `uint32` 类型的取值范围是 `0` 到 `2 ** 32-1` 。 0.8.0 开始，算术运算有两个计算模式：一个是 “wrapping”（截断）模式或称 “unchecked”（不检查）模式，一个是”checked” （检查）模式。 默认情况下，算术运算在 “checked” 模式下，即都会进行溢出检查，如果结果落在取值范围之外，调用会通过 [失败异常](https://learnblockchain.cn/docs/solidity/control-structures.html#assert-and-require) 回退。 你也可以通过 `unchecked { ... }` 切换到 “unchecked”模式
>
> 溢出的检查功能是在 0.8.0 版本加入的，在此版本之前，请使用 OpenZepplin SafeMath 库。

#### 3. 定长浮点型

**Solidity 还没有完全支持定长浮点型。可以声明定长浮点型的变量，但不能给它们赋值或把它们赋值给其他变量。**

`fixed` / `ufixed`：表示各种大小的有符号和无符号的**定长浮点型**。 在关键字 `ufixedMxN` 和 `fixedMxN` 中，`M` 表示该类型占用的位数，`N` 表示可用的小数位数。 `M` 必须能整除 8，即 8 到 256 位。 `N` 则可以是从 0 到 80 之间的任意数。 `ufixed` 和 `fixed` 分别是 `ufixed128x19` 和 `fixed128x19` 的别名。

运算符：

- 比较运算符：`<=`， `<`， `==`， `!=`， `>=`， `>` （返回值是布尔型）
- 算术运算符：`+`， `-`， 一元运算 `-`， 一元运算 `+`， `*`， `/`， `%` （取余数）

注意：**solidity完全不支持浮点型数据。**

#### 4. 地址类型 Address

地址类型有两种形式，他们大致相同：

> - `address`：保存一个20字节的值（以太坊地址的大小）。
> - `address payable` ：可支付地址，与 `address` 相同，不过有成员函数 `transfer` 和 `send` 。

这种区别背后的思想是 `address payable` 可以向其他地址发送以太币，而不能让一个普通的 `address` 发送以太币，例如，它可能是一个智能合约地址，并且不支持接收以太币。

类型转换:

允许从 `address payable` 到 `address` 的**隐式转换**，而从 `address` 到 `address payable` 必须**显示的转换**, 通过 `payable(<address>)` 进行转换。

> 在0.5版本,执行这种显式转换的唯一方法是使用中间类型，先转换为 `uint160` 如,  address payable ap = address(uint160(addr));

`address` 允许和 `uint160`、 整型字面常量、`bytes20` 及合约类型相互转换。



只能通过 `payable(...)` 表达式把 `address` 类型和合约类型转换为 `address payable`。 只有能接收以太币的合约类型，才能够进行此转换。例如合约要么有  [receive](https://learnblockchain.cn/docs/solidity/contracts.html#receive-ether-function) 或可支付的回退函数。 注意 `payable(0)` 是有效的，这是此规则的例外。

> `address` 和 `address payable` 之间的区别是从0.5.0版本引入的。 同样从该版本开始，**合约**不能**隐式地**转换为 `address` 类型，但仍然可以**显式地**转换为 `address` 或者 `address payable` ，**如果它们有一个receive 或 payable的回退函数的话**。

运算符:

- `<=`, `<`, `==`, `!=`, `>=` and `>`

##### 1. 地址类型成员变量

- `balance` 和 `transfer` 成员

可以使用 `balance` 属性来查询一个地址的余额， 也可以使用 `transfer` 函数向一个**可支付地址（payable address）**发送 以太币（Ether） （以 wei 为单位）：

```solidity
address x = 0x123;
address myAddress = this;
if (x.balance < 10 && myAddress.balance >= 10) x.transfer(10);
```

如果当前合约的余额不够多，则 `transfer` 函数会执行失败，或者如果以太转移被接收帐户拒绝， `transfer` 函数同样会失败而进行回退。

```solidity
如果 x 是一个合约地址，它的代码（更具体来说是, 如果有receive函数, 执行 receive 接收以太, 或者存在fallback函数,执行 Fallback 回退函数）会跟 transfer 函数调用一起执行（这是 EVM 的一个特性，无法阻止）。 如果在执行过程中用光了 gas 或者因为任何原因执行失败，以太币 交易会被打回，当前的合约也会在终止的同时抛出异常。
```

- `send` 成员

`send` 是 `transfer` 的低级版本。如果执行失败，当前的合约不会因为异常而终止，但 `send` 会返回 `false`。

```solidity
在使用 send 的时候会有些风险：如果调用栈深度是 1024 会导致发送失败（这总是可以被调用者强制），如果接收者用光了 gas 也会导致发送失败。 所以为了保证 以太币 发送的安全，一定要检查 send 的返回值，可以使用 transfer 或者更好的办法： 保证使用接收者自己能够取回资金的模式。
```

- `call`， `delegatecall` 和 `staticcall`

TODO : 待补充学习

- `code` 和 `codehash` 成员

你可以查询任何智能合约的部署代码。使用 `.code` 来获取EVM的字节码，其返回 `bytes memory` ，值可能是空。 使用 `.codehash` 获得该代码的 Keccak-256哈希值 (为 `bytes32` )。注意， `addr.codehash` 比使用 `keccak256(addr.code)` 更便宜。

```solidity
所有合约都可以转换为 address 类型，因此可以使用 address(this).balance 查询当前合约的余额。
```

关于this指针的说明，比如下面的代码中 this 指的就是 当前代码中 Score合约。

```solidity
contract Score {
    mapping(address => uint256) score;
    address teacher;
 
    constructor() {
        teacher = address(new Teacher(address(this)));   //所以说，this通常指向当前合约contract
    }
}
```

#### 5. 合约类型

每一个 [contract](https://learnblockchain.cn/docs/solidity/contracts.html#contracts) 定义都有他自己的类型。

可以隐式地将合约转换为继承它的子类合约。 合约可以显式转换为 `address` 类型( 使用 `address(x)` 执行 )。

只有当合约具有 接收receive函数 或 payable 回退函数时，才能显式和 `address payable` 类型相互转换。

- 合约不支持任何运算符。
- 合约和 `address` 的数据表示是相同的
- 可以通过 `new` 关键字实例化合约（即新创建一个合约对象）
- 合约类型的成员是合约的外部函数及 public 的 状态变量
- 对于合约  `C` 可以使用 `type(C)` 获取合约的类型信息

#### 6. 定长字节数组

关键字有：`bytes1`， `bytes2`， `bytes3`， …， `bytes32`。

运算符：

- 比较运算符： `<=`， `<`， `==`， `!=`， `>=`， `>` （返回布尔型）
- 位运算符： `&`， `|`， `^` （按位异或）， `~` （按位取反）
- 移位运算符： `<<` （左移位）， `>>` （右移位）
- 索引访问：如果 `x` 是 `bytesI` 类型，那么 `x[k]` （其中 `0 <= k < I`）返回该字节数组x第 `k` 个字节（只读）。

成员变量：

- `.length` 表示这个字节数组的长度（只读）.

> 在 0.8.0 之前, `byte` 用作为 `bytes1` 的别名。只有一个字节的字节数组就是一个字节

#### 7. 变长字节数组

- `bytes`:

变长字节数组，参见 [数组](https://learnblockchain.cn/docs/solidity/types.html#arrays)。它并不是值类型，而是引用类型

- `string`:

变长 UTF-8 编码字符串类型，参见 [数组](https://learnblockchain.cn/docs/solidity/types.html#arrays)。并不是值类型，是引用类型

#### 8. 地址字面常量

比如像 `0xdCad3a6d3569DF655070DEd06cb7A1b2Ccd1D3AF` 这样的通过了地址校验和测试的十六进制字面常量会作为 `address` 类型。 而没有通过校验测试, 长度在 39 到 41 个数字之间的十六进制字面常量，会产生一个错误。对于整数类型可以通过在前面添加0来解决整个问题，对于bytesNN 类型则需要通过在后面添加0来解决这个问题。

