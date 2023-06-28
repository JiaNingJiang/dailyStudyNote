package mergeTable

import "DataStructure/linkList"

type Node struct {
	Data    interface{}
	Feature *Node // 指向所在集合的特征节点
}

type UnionFindSet struct {
	ElementMap map[interface{}]*Node // 值 --> 节点
	FeatureMap map[*Node]*Node       // 节点 -- > 特征节点
	SizeMap    map[*Node]int         // 特征节点 -- > 集合的节点个数
}

// 根据给定的数据集合生成一个并查集(每个数据属于一个集合，并且是该集合的唯一节点和特征节点)
func NewUnionFindSet(dataSet []interface{}) *UnionFindSet {

	ufs := new(UnionFindSet)
	ufs.ElementMap = make(map[interface{}]*Node)
	ufs.FeatureMap = make(map[*Node]*Node)
	ufs.SizeMap = make(map[*Node]int)

	for _, data := range dataSet {
		newNode := new(Node)
		ufs.ElementMap[data] = newNode
		ufs.FeatureMap[newNode] = newNode
		ufs.SizeMap[newNode]++
	}
	return ufs
}

// 返回节点所在集合的特征节点，同时重新调整集合内节点指向的特征节点(在进行集合合并时，会导致一部分节点不能指向新的特征节点)
func (ufs *UnionFindSet) FindFeature(node *Node) *Node {
	stack := linkList.NewStack()
	current := node
	for {
		if ufs.FeatureMap[current] != current {
			stack.Push(current) // 将所有不是特征节点的节点加入到栈中
			current = ufs.FeatureMap[current]
		} else { // current 目前指向集合最新的特征节点
			break
		}
	}
	for {
		if stack.Len == 0 {
			return current
		}
		commonNode := stack.Pop().(*Node)
		ufs.FeatureMap[commonNode] = current // 更新节点指向的特征节点为最新的
	}
}

// 判断两个节点是否在同一个集合中
func (ufs *UnionFindSet) IsSameSet(a, b interface{}) bool {
	if ufs.ElementMap[a] != nil && ufs.ElementMap[b] != nil {
		if ufs.FindFeature(ufs.ElementMap[a]) == ufs.FindFeature(ufs.ElementMap[b]) {
			return true
		}
	}
	return false
}

// 将两个节点所在的集合进行合并
func (ufs *UnionFindSet) Union(a, b interface{}) {
	aNode := ufs.ElementMap[a]
	bNode := ufs.ElementMap[b]
	if aNode == nil || bNode == nil {
		return
	}
	if ufs.IsSameSet(a, b) {
		return
	}
	aNodeFeature := ufs.FindFeature(aNode)
	bNodeFeature := ufs.FindFeature(bNode)

	aNodeSetCap := ufs.SizeMap[aNodeFeature]
	bNodeSetCap := ufs.SizeMap[bNodeFeature]

	if aNodeSetCap >= bNodeSetCap { // 两个集合节点个数相等或者A集合节点数更多，都合并到集合a中
		ufs.FeatureMap[bNodeFeature] = aNodeFeature
		ufs.SizeMap[aNodeFeature] += ufs.SizeMap[bNodeFeature]
		delete(ufs.SizeMap, bNodeFeature)
	} else if aNodeSetCap < bNodeSetCap { // B集合节点数更多，都合并到集合B中
		ufs.FeatureMap[aNodeFeature] = bNodeFeature
		ufs.SizeMap[bNodeFeature] += ufs.SizeMap[aNodeFeature]
		delete(ufs.SizeMap, aNodeFeature)
	}
}
