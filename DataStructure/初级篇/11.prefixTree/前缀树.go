package prefixTree

type Node struct {
	pass int       // 经过此节点的路径数
	end  int       // 以此节点为终点的路径数
	next [26]*Node // 当前节点的下一节点(有26种可能)
}

func NewNode() *Node {
	return &Node{
		pass: 0,
		end:  0,
		next: [26]*Node{},
	}
}

func (node *Node) Insert(str string) {
	if str == "" {
		return
	}
	charSet := []byte(str) // 转化为字符串数组
	node.pass++            // 所有字符串都必须起源于根
	current := node
	for _, char := range charSet {
		offset := char - 'a'
		if current.next[offset] == nil { // 重要：新的路径,需要新建(旧的路径则直接复用)
			current.next[offset] = NewNode()
		}

		current.next[offset].pass++
		current = current.next[offset]
	}
	current.end++
}

// 返回某字符串出现的次数
func (node *Node) Search(str string) int {
	if str == "" {
		return 0
	}
	charSet := []byte(str) // 转化为字符串数组
	current := node
	for _, char := range charSet {
		offset := char - 'a'
		if current.next[offset] == nil { // 搜索过程中中断，则说明此字符串不存在
			return 0
		}
		current = current.next[offset]
	}
	return current.end
}

func (node *Node) SearchPre(pre string) int {
	if pre == "" {
		return 0
	}

	charSet := []byte(pre) // 转化为字符串数组
	current := node
	for _, char := range charSet {
		offset := char - 'a'
		if current.next[offset] == nil { // 搜索过程中中断，则说明此字符串不存在
			return 0
		}
		current = current.next[offset]
	}
	return current.pass
}

func (node *Node) Delete(word string) {
	if node.Search(word) == 0 { // 删除之前，必须要确保字符串存在于前缀树中
		return
	}

	node.pass--
	charSet := []byte(word) // 转化为字符串数组
	current := node

	for _, char := range charSet {
		offset := char - 'a'
		current.next[offset].pass--
		if current.next[offset].pass == 0 { // 之后路径上的的全部剩余节点清理掉(GC清理)
			current.next[offset] = nil
			return
		}
		current = current.next[offset]
	}
	current.end--
}
