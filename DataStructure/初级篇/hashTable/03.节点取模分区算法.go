package hashTable

// import (
// 	"hash/fnv"
// 	"reflect"
// )

// // 存储数据的节点
// type Node struct {
// 	ID   int
// 	IP   string
// 	Port int
// }

// type HashRing struct {
// 	Nodes []*Node
// }

// func NewHashRing(nodes []*Node) *HashRing {
// 	hr := new(HashRing)
// 	hr.Nodes = append(hr.Nodes, nodes...)
// 	return hr
// }

// func (hr *HashRing) AddNode(node *Node) {
// 	hr.Nodes = append(hr.Nodes, node)
// }

// func (hr *HashRing) DelNode(node *Node) {
// 	index := 0
// 	for i, n := range hr.Nodes {
// 		if reflect.DeepEqual(*n, *node) {
// 			index = i
// 		}
// 	}
// 	hr.Nodes = append(hr.Nodes[:index], hr.Nodes[index+1:]...)
// }

// // 查看数据存储在哪一个节点上(存储的时候查看存储目标，查询的时候获知数据位置)
// func (hr *HashRing) GetNode(data []byte) *Node {
// 	hash := hashData(data)
// 	index := hash % uint32(len(hr.Nodes))
// 	return hr.Nodes[index]
// }

// // 计算数据的哈希值
// func hashData(data []byte) uint32 {
// 	h := fnv.New32a()
// 	h.Write(data)
// 	return h.Sum32()
// }
