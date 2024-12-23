package hashTable

import (
	//	"fmt"
	"testing"
)

const (
	node1 = "192.168.1.1"
	node2 = "192.168.1.2"
	node3 = "192.168.1.3"
)

// 返回3个真实节点(node1、node2、node3)分别对应的虚拟节点的个数
func getNodesCount(nodes nodesArray) (int, int, int) {
	node1Count := 0
	node2Count := 0
	node3Count := 0

	for _, node := range nodes {
		if node.nodeKey == node1 {
			node1Count += 1
		}
		if node.nodeKey == node2 {
			node2Count += 1

		}
		if node.nodeKey == node3 {
			node3Count += 1

		}
	}
	return node1Count, node2Count, node3Count
}

func TestHash(t *testing.T) {

	nodeWeight := make(map[string]int)
	nodeWeight[node1] = 2
	nodeWeight[node2] = 2
	nodeWeight[node3] = 3
	vitualSpots := 100

	hash := NewHashRing(vitualSpots) // 新建一个哈希环

	hash.AddNodes(nodeWeight)
	if hash.GetNode("1") != node3 {
		t.Fatalf("expetcd %v got %v", node3, hash.GetNode("1"))
	}
	if hash.GetNode("2") != node3 {
		t.Fatalf("expetcd %v got %v", node3, hash.GetNode("2"))
	}
	if hash.GetNode("3") != node2 {
		t.Fatalf("expetcd %v got %v", node2, hash.GetNode("3"))
	}
	c1, c2, c3 := getNodesCount(hash.nodes) // 计算三个真实节点在哈希环上分别对应的虚拟节点的个数
	t.Logf("len of nodes is %v after AddNodes node1:%v, node2:%v, node3:%v", len(hash.nodes), c1, c2, c3)

	hash.RemoveNode(node3) // 移除掉真实节点node3,重新再生成哈希环
	if hash.GetNode("1") != node1 {
		t.Fatalf("expetcd %v got %v", node1, hash.GetNode("1"))
	}
	if hash.GetNode("2") != node2 {
		t.Fatalf("expetcd %v got %v", node1, hash.GetNode("2"))
	}
	if hash.GetNode("3") != node2 {
		t.Fatalf("expetcd %v got %v", node2, hash.GetNode("3"))
	}
	c1, c2, c3 = getNodesCount(hash.nodes)
	t.Logf("len of nodes is %v after RemoveNode node1:%v, node2:%v, node3:%v", len(hash.nodes), c1, c2, c3)

	hash.AddNode(node3, 3) // 重新添加真实节点node3,重新再生成哈希环
	if hash.GetNode("1") != node3 {
		t.Fatalf("expetcd %v got %v", node3, hash.GetNode("1"))
	}
	if hash.GetNode("2") != node3 {
		t.Fatalf("expetcd %v got %v", node3, hash.GetNode("2"))
	}
	if hash.GetNode("3") != node2 {
		t.Fatalf("expetcd %v got %v", node2, hash.GetNode("3"))
	}
	c1, c2, c3 = getNodesCount(hash.nodes)
	t.Logf("len of nodes is %v after AddNode node1:%v, node2:%v, node3:%v", len(hash.nodes), c1, c2, c3)

}
