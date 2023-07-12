package lesson7

import (
	"DataStructure2/utils"
	"fmt"
	"testing"
)

func TestBSTtoLinkList(t *testing.T) {
	arr := []int{5, 3, 7, 2, 4, 6, 8, 1}
	tree := utils.NewTreeWithArr(arr)

	utils.InOrderTraversal(tree)
	fmt.Println()

	head, tail := BSTtoLinkList(tree.Root)

	cur := head
	for {
		if cur == tail {
			fmt.Printf("%d ", cur.Data)
			break
		}
		fmt.Printf("%d ", cur.Data)
		cur = cur.Right
	}
	fmt.Println()

	cur = tail
	for {
		if cur == head {
			fmt.Printf("%d ", cur.Data)
			break
		}
		fmt.Printf("%d ", cur.Data)
		cur = cur.Left
	}
	fmt.Println()
}
