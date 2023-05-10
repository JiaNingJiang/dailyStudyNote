package linkList

import (
	"fmt"
	"testing"
)

func TestIsHaveCircle(t *testing.T) {
	ll := NewLinkList()
	ll.TailAdd(8)
	ll.TailAdd(7)
	ll.TailAdd(5)
	ll.TailAdd(4)
	ll.TailAdd(3)

	ll.Print()

	AddCircleForLinkList(ll, 3)
	//ll.Print()
	flag, node := IsHaveCircle(ll)
	if flag {
		fmt.Printf("链表有环，入环点是：%d\n", node.Data)
	} else {
		fmt.Printf("链表无环\n")
	}
}
