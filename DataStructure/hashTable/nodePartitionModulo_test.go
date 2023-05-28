package hashTable

// import (
// 	"fmt"
// 	"testing"
// )

// func TestNodePartitionModulo(t *testing.T) {
// 	node1 := Node{ID: 1, IP: "192.168.0.1", Port: 8000}
// 	node2 := Node{ID: 2, IP: "192.168.0.2", Port: 8000}
// 	node3 := Node{ID: 3, IP: "192.168.0.3", Port: 8000}

// 	// Create the hash ring with the nodes and replication factor
// 	hashRing := NewHashRing([]*Node{&node1, &node2, &node3})

// 	// Perform some lookups
// 	key1 := "key1"
// 	node := hashRing.GetNode([]byte(key1))
// 	fmt.Printf("Key: %s, Node: %v\n", key1, *node)

// 	key2 := "key2"
// 	node = hashRing.GetNode([]byte(key2))
// 	fmt.Printf("Key: %s, Node: %v\n", key2, *node)

// 	// Add a new node to the hash ring
// 	node4 := &Node{ID: 4, IP: "192.168.0.4", Port: 8000}
// 	hashRing.AddNode(node4)

// 	// Perform lookup after adding the new node
// 	key3 := "key3"
// 	node = hashRing.GetNode([]byte(key3))
// 	fmt.Printf("Key: %s, Node: %v\n", key3, *node)

// }
