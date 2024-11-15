1. 视图切换中的改进：通过心跳包确保被选举的`Leader`总是能正常通信工作的。可能还需要引入区块数作为一个选举的标准？（按常理应该是选择有效区块数最多的作为新的`Leader`，有效区块数最多应该可以按照投票的方式决定：①区块数量作为一个投票标准②最近的若干个区块相同作为一个标准）

2. 交易打包，提高一轮共识中的承载量，提高共识效率，可以具体测试大概累计多少笔交易才能让共识效率最高？根据当前TPS动态实时调整？
3. 通过让每一个节点即担任共识节点，又担任客户端的角色，这样的好处大概有以下几点：
   1. 每个节点都有相同的交易(p2p广播获得，而且参与共识时交易顺序和数量应该是一样的)，这样在启动一轮新的view-change时，新的`Leader`节点可以天然知道旧`Leader`哪些交易没有完成共识，而不需要从日志中搜集获取
   2. view-change可以有每一个共识节点主动发起，而不是需要由外界`client`提议才能进行。这样的好处是：网络总是会调整自己保持一个正常的状态，在网络TPS不是太高的情况下，可以最大程序降低不必要的view-change消耗。

4. 舍弃了检查点机制，原因：因为当选Leader的总是具有最为正确的区块数量，没有必要从其他节点处进行区块同步。而对于`follower`，它们通过区块同步模块主动同步共识结果(区块)，而不是当`Leader`完成一轮共识后去广播。这样网络中所有节点的区块链总是保持同步的，检查点就是最新区块。
5. 



待添加的：

1. 明显非法的交易就不应该参与共识（进行预处理：重复性验证 --> 布隆过滤器 和 双花验证？），这样可以降低系统开销
2. 一段时间内确实没有搜集到交易，主节点需要发心跳包给其他跟随节点表名自己的活性，防止不必要的view-change启动。---->  当节点兼具共识节点和客户端时，优势就体现出来了：跟随节点可以验证当前时间内是否如主节点表示的一样，确实没有任何交易产生，否则对于主节点的这项声明是不好验证的。
3. 消息的有效性认证是否必须对完整消息求取哈希值？是否能用一种相对更低耗的方式代替？
4. 在区块链系统中，交易的顺序是很重要的，有些交易存在先后依赖关系，是否能通过某种方式解决？（交易人为引入优先级和相应的合约---> 不同的合约的用不同的pbft round）
5. PBFT算法可扩展性不行，为此可以引入一个引导节点，专门负责将新节点加入到pbft共识群组（所有节点都需要与引导节点连接，因此它能获取其他所有节点的网络地址信息），新加入的节点通过区块同步模块快速同步已有区块。



文献：

1. [1]赵振龙. 带有主动恢复的拜占庭容错算法在区块链中的应用[D].浙江大学,2018.