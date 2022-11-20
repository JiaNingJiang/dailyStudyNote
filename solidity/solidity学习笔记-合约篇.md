### 一、合约介绍

Solidity 合约类似于面向对象语言中的类。合约中有用于数据持久化的状态变量，和可以修改状态变量的函数。 调用另一个合约实例的函数时，会执行一个 EVM 函数调用，这个操作会切换执行时的上下文，这样，前一个合约的状态变量就不能访问了。

### 二、创建合约

创建合约时， 合约的 [构造函数](https://learnblockchain.cn/docs/solidity/contracts.html#constructor)  (一个用关键字 `constructor` 声明的函数)会执行一次。 构造函数是可选的。只允许有一个构造函数，这意味着不支持重载。

**构造函数执行完毕后，合约的最终代码将部署到区块链上**。此代码包括所有公共和外部函数以及所有可以通过函数调用访问的函数。 部署的代码没有 包括构造函数代码或构造函数调用的**内部函数**。

#### 1. 可见性和 getter 函数

##### 状态变量可见性

状态变量有 3 种可见性：

- `public`

对于 public 状态变量会**自动生成一个  getter 函数**。 以便其他的合约读取他们的值。 当在用一个合约里使用时，外部方式访问 (如: `this.x`) 会调用getter 函数，而内部方式访问 (如: `x`) 会直接从存储中获取值。 **但是setter函数则不会被生成**，所以其他合约不能直接修改其值。

- `internal`

内部可见性状态变量只能在它们所定义的合约和派生合同中访问。 它们不能被外部访问。 这是状态变量的默认可见性。

- `private`

私有状态变量就像内部变量一样，但它们在派生合约中是不可见的。

```solidity
设置为 private或 internal，只能防止其他合约读取或修改信息，但它仍然可以在链外查看到。
```

##### 函数可见性

由于 Solidity 有两种函数调用：外部调用则会产生一个 EVM 调用，而内部调用不会， 更进一步， 函数可以确定其被内部及派生合约的可访问性，这里有 4 种可见性：

- `external`

外部可见性函数作为合约接口的一部分，意味着我们可以从其他合约和交易中调用。 一个外部函数 `f` 不能从内部调用（即 `f` 不起作用，但 `this.f()` 可以）。

- `public`

public 函数是合约接口的一部分，可以在内部或通过消息调用。

- `internal`

内部可见性函数访问可以在当前合约或派生的合约访问，不可以外部访问。 由于它们没有通过合约的ABI向外部公开，它们可以接受内部可见性类型的参数：比如映射或存储引用。

- `private`

private 函数和状态变量仅在当前定义它们的合约中使用，并且不能被派生合约使用。



可见性标识符的定义位置，对于状态变量来说是在类型后面，对于函数是在参数列表和返回关键字中间。

```solidity
pragma solidity  >=0.4.16 <0.9.0;

contract C {
    function f(uint a) private pure returns (uint b) { return a + 1; }
    function setData(uint a) internal { data = a; }
    uint public data;
}
```



##### Getter 函数

编译器自动为所有 **public** 状态变量创建 **getter 函数**。对于下面给出的合约，编译器会生成一个**名为 `data` 的函数**， 该函数没有参数，返回值是一个 `uint` 类型，即状态变量 `data` 的值。 状态变量的初始化可以在声明时完成。

```solidity
pragma solidity  >=0.4.16 <0.9.0;

contract C {
    uint public data = 42;
}

contract Caller {
    C c = new C();
    function f() public {
        uint local = c.data();   //存在c.data()函数
    }
}
```

getter 函数具有外部（external）可见性。如果在内部访问 getter（即没有 `this.` ），它被认为一个状态变量。 如果使用外部访问（即用 `this.` ），它被认作为一个函数。

```solidity
pragma solidity >=0.4.16 <0.9.0;

contract C {
    uint public data;
    function x() public {
        data = 3; // 内部访问方式
        uint val = this.data(); // 外部访问方式
    }
}
```

如果你有一个数组类型的 `public` 状态变量，那么你**只能通过生成的 getter 函数**访问数组的**单个元素**。 这个机制以避免返回整个数组时的高成本gas。 可以使用如 `myArray(0)` 用于指定参数要返回的单个元素。 如果要在一次调用中**返回整个数组**，则**需要写一个函数**，例如：

```solidity
pragma solidity >=0.4.0 <0.9.0;

contract arrayExample {
  // public state variable
  uint[] public myArray;

  // 指定生成的Getter 函数
  /*
  function myArray(uint i) public view returns (uint) {
      return myArray[i];
  }
  */

  // 返回整个数组
  function getArray() public view returns (uint[] memory) {
      return myArray;
  }
}
```

现在可以使用 `getArray()` 获得整个数组，而 `myArray(i)` 是返回单个元素。

#### 2. Constant 和 Immutable  状态变量

状态变量声明为 `constant` (常量)或者 `immutable` （不可变量），在这两种情况下，合约一旦部署之后，变量将不在修改。

对于 `constant` 常量, 他的值在编译器确定，而对于 `immutable`, 它的值在部署时确定。

编译器不会为这些变量预留存储位，它们的每次出现都会被替换为相应的常量表达式（它可能被优化器计算为实际的某个值）。

不是所有类型的状态变量都支持用 `constant` 或 `immutable` 来修饰，当前仅支持 [字符串](https://learnblockchain.cn/docs/solidity/types.html#strings) (仅常量) 和 [值类型](https://learnblockchain.cn/docs/solidity/types.html#value-types) .

##### Constant

如果状态变量声明为 `constant` (常量)。在这种情况下，只能使用那些在编译时有确定值的表达式来给它们赋值。 任何通过访问 storage，区块链数据（例如 `block.timestamp`, `address(this).balance` 或者 `block.number`）或执行数据（ `msg.value` 或 `gasleft()` ） 或对外部合约的调用来给它们赋值都是不允许的。

内建（built-in）函数 `keccak256` ， `sha256` ， `ripemd160` ， `ecrecover` ， `addmod` 和 `mulmod` 是允许的（即使他们确实会调用外部合约， `keccak256` 除外）。

##### Immutable

声明为不可变量(`immutable`)的变量的限制要比声明为常量(`constant`) 的变量的限制少：可以在合约的构造函数中或声明时为不可变的变量分配任意值。 不可变量只能赋值一次，并且在赋值之后才可以读取。

编译器生成的合约创建代码将在返回合约之前修改合约的运行时代码，方法是将对不可变量的所有引用替换为分配给它们的值。 如果要将编译器生成的运行时代码与实际存储在区块链中的代码进行比较，则这一点很重要。

```solidity
不可变量可以在声明时赋值，不过只有在合约的构造函数执行时才被视为视为初始化。 这意味着，你不能用一个依赖于不可变量的值在行内初始化另一个不可变量。 不过，你可以在合约的构造函数中这样做。

这是为了防止对状态变量初始化和构造函数顺序的不同解释，特别是继承时，出现问题。
```

```solidity
译者注：不可变量(Immutable) 是 Solidity 0.6.5 引入的，因此0.6.5 之前的版本不可用
```

#### 3. 函数

合约之外的函数（也称为“自由函数”）始终具有隐式的 `internal` [可见性](https://learnblockchain.cn/docs/solidity/contracts.html#visibility-and-getters)。 它们的代码包含在所有调用它们合约中，类似于内部库函数。

```solidity
// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.7.1 <0.9.0;

function sum(uint[] memory arr) pure returns (uint s) {  //只能由当前合约文件内被调用
    for (uint i = 0; i < arr.length; i++)
        s += arr[i];
}

contract ArrayExample {
    bool found;
    function f(uint[] memory arr) public {
        // This calls the free function internally.
        // The compiler will add its code to the contract.
        uint s = sum(arr);
        require(s >= 10);
        found = true;
    }
}
```

##### 函数参数（输入参数）

函数参数的声明方式与变量相同。不过未使用的参数可以省略参数名。

例如，如果我们希望合约接受有两个整数形参的函数的外部调用，可以像下面这样写：

```solidity
pragma solidity >=0.4.16 <0.9.0;

contract Simple {
    uint sum;
    function taker(uint a, uint b) public {
        sum = a + b;
    }
}
```

##### 返回变量

函数返回变量的声明方式在关键词 `returns` 之后，与参数的声明方式相同。

例如，如果我们需要返回两个结果：两个给定整数的和与积，我们应该写作：

```solidity
pragma solidity >=0.4.16 <0.9.0;

contract Simple {
    function arithmetic(uint a, uint b) public  pure returns (uint sum, uint product)
    {
        sum = a + b;
        product = a * b;
    }
}
```

##### View 视图函数

可以将函数声明为 `view` 函数类型，这种情况下函数保证不修改状态（但是可以读取状态变量）。

##### Pure 纯函数

在保证不读取或修改状态的情况下，函数可以被声明为 `pure` 函数。特别是，在编译时只给出函数输入和`msg.data` ，但又不知道当前[区块链](https://so.csdn.net/so/search?q=区块链&spm=1001.2101.3001.7020)状态的情况下，建议使用 `pure` 函数。

纯函数能够使用 `revert()` 和 `require()` 在 [发生错误](https://learnblockchain.cn/docs/solidity/control-structures.html#assert-and-require) 时去还原潜在状态更改。

#### 4. 特殊函数

##### receive 接收以太函数

一个合约最多有一个 `receive` 函数, 声明函数为： `receive() external payable { ... }`

不需要 `function` 关键字，也没有参数和返回值并且必须是　`external`　可见性和　`payable` 修饰． 它可以是 `virtual` 的，可以被重载也可以有 修改器（modifier） 。

在对合约没有任何附加数据调用（通常是对合约转账）是会执行 `receive` 函数．例如执行`.send()` 或者 `.transfer()` ，如果 `receive` 函数不存在，但是有payable的 `fallback` 回退函数 ,那么在进行纯以太转账时，`fallback `函数会调用．

如果两个函数都没有，这个合约就没法通过常规的转账交易接收以太（会抛出异常）．

##### Fallback 回退函数

合约可以最多有一个回退函数。函数声明为： `fallback () external [payable]` 或 `fallback (bytes calldata input) external [payable] returns (bytes memory output)`

没有　`function`　关键字。　必须是　`external`　可见性，它可以是 `virtual` 的，可以被重载也可以有 修改器（modifier） 。

如果在一个对合约调用中，没有其他函数与给定的函数标识符匹配，则fallback会被调用。或者在没有 receive 函数 时，fallback 函数也会被执行。

fallback函数始终会接收数据，但为了同时接收以太时，必须标记为　`payable` 。

如果使用了带参数的版本， `input` 将包含发送到合约的完整数据（等于 `msg.data` ），并且通过 `output` 返回数据。 返回数据不是 ABI 编码过的数据，相反，它返回不经过修改的数据。

```solidity
// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.6.2 <0.9.0;

contract Test {
    // 发送到这个合约的所有消息都会调用此函数（因为该合约没有其它函数）。
    // 向这个合约发送以太币会导致异常，因为 fallback 函数没有 `payable` 修饰符
    fallback() external { x = 1; }
    uint x;
}


// 这个合约会保留所有发送给它的以太币，没有办法返还。
contract TestPayable {
    uint x;
    uint y;

    // 除了纯转账外，所有的调用都会调用这个函数．
    // (因为除了 receive 函数外，没有其他的函数).
    // 任何对合约非空calldata 调用会执行回退函数(即使是调用函数附加以太).
    fallback() external payable { x = 1; y = msg.value; }

    // 纯转账调用这个函数，例如对每个空empty calldata的调用
    receive() external payable { x = 2; y = msg.value; }
}

contract Caller {
    function callTest(Test test) public returns (bool) {
        (bool success,) = address(test).call(abi.encodeWithSignature("nonExistingFunction()"));
        require(success);
        //  test.x 结果变成 == 1。

        // address(test) 不允许直接调用 ``send`` ,  因为 ``test`` 没有 payable 回退函数
        //  转化为 ``address payable`` 类型 , 然后才可以调用 ``send``
        address payable testPayable = payable(address(test));


        // 以下将不会编译，但如果有人向该合约发送以太币，交易将失败并拒绝以太币。
        // test.send(2 ether）;
    }

    function callTestPayable(TestPayable test) public returns (bool) {
        (bool success,) = address(test).call(abi.encodeWithSignature("nonExistingFunction()"));
        require(success);
        // 结果 test.x 为 1  test.y 为 0.
        (success,) = address(test).call{value: 1}(abi.encodeWithSignature("nonExistingFunction()"));
        require(success);
        // 结果test.x 为1 而 test.y 为 1.

        // 发送以太币, TestPayable 的 receive　函数被调用．

        // 因为函数有存储写入, 会比简单的使用 ``send`` or ``transfer``消耗更多的 gas。
        // 因此使用底层的call调用
        (success,) = address(test).call{value: 2 ether}("");
        require(success);

        // 结果 test.x 为 2 而 test.y 为 2 ether.

        return true;
    }

}
```

#### 5. 事件 Events

事件是以太坊虚拟机(EVM)日志基础设施提供的一个便利接口。当事件被发送（调用）时，会触发参数存储到交易的日志中（一种区块链上的特殊数据结构）。这些日志与合约的地址关联，并记录到区块链中。即是说：区块链是打包一系列交易的区块组成的链条，每一个交易“收据”会包含0到多个日志记录，日志代表着智能合约所触发的事件。

**在DAPP的应用中，如果监听了某事件，当事件发生时，会进行回调。** 不过要注意：日志和事件在合约内是无法被访问的，即使是创建日志的合约。

在Solidity 代码中，使用event 关键字来定义一个事件，如：

```solidity
event EventName(address bidder, uint amount);
```

这个用法和定义函数式一样的，并且事件在合约中同样可以被继承。触发一个事件使用emit(说明，之前的版本里并不需要使用emit)，如：

```solidity
emit EventName(msg.sender, msg.value);
```

触发事件可以在任何函数中调用，如：

```solidity
function testEvent() public {

    // 触发一个事件
     emit EventName(msg.sender, msg.value); 
}
```

##### 监听事件

##### 修改合约，定义事件及触发事件

合约代码：

```solidity
pragma solidity ^0.4.21;

contract InfoContract {
    
   string fName;
   uint age;
   
   function setInfo(string _fName, uint _age) public {
       fName = _fName;
       age = _age;
   }
   
   function getInfo() public constant returns (string, uint) {
       return (fName, age);
   }   
}
```

首先，需要定义一个事件：

```solidity
event Instructor(
       string name,
       uint age
    );
```

这个事件中，会接受两个参数：name 和 age , 也就是需要跟踪的两个信息。

然后，需要在setInfo函数中，触发Instructor事件，如：

```solidity
function setInfo(string _fName, uint _age) public {
       fName = _fName;
       age = _age;
       emit Instructor(_fName, _age);
   }
```

在[Web3与智能合约交互实战](https://link.zhihu.com/?target=https%3A//learnblockchain.cn/2018/04/15/web3-html/), 点击"Updata Info"按钮之后，会调用setInfo函数，函数时触发Instructor事件。

##### 使用Web3监听事件，刷新UI

现在需要使用Web3监听事件，刷新UI。 先回顾下之前的使用Web3和智能合约交互的代码：

```solidity
<script>
    if (typeof web3 !== 'undefined') {
        web3 = new Web3(web3.currentProvider);
    } else {
        // set the provider you want from Web3.providers
        web3 = new Web3(new Web3.providers.HttpProvider("http://localhost:7545"));
    }

    web3.eth.defaultAccount = web3.eth.accounts[0];

    var infoContract = web3.eth.contract(ABI INFO);

    var info = infoContract.at('CONTRACT ADDRESS');

    info.getInfo(function(error, result){
        if(!error)
            {
                $("#info").html(result[0]+' ('+result[1]+' years old)');
                console.log(result);
            }
        else
            console.error(error);
    });

    $("#button").click(function() {
        info.setInfo($("#name").val(), $("#age").val());
    });

</script>
```

现在可以不需要 info.getInfo()来获取信息，而改用监听事件获取信息，先定义一个变量引用事件：

```solidity
var instructorEvent = info.Instructor();
```

然后使用**.watch()**方法来添加一个回调函数：

```solidity
instructorEvent.watch(function(error, result) {
        if (!error)
            {
                $("#info").html(result.args.name + ' (' + result.args.age + ' years old)');
            } else {
                console.log(error);
            }
    });
```

