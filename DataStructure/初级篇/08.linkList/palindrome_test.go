package linkList

import (
	"fmt"
	"testing"
)

func TestIsNotPalindrome(t *testing.T) {
	ll := NewLinkList()
	//ll.TailAdd(3)
	//ll.TailAdd(4)
	//ll.TailAdd(5)
	//ll.TailAdd(4)
	//ll.TailAdd(3)

	ll.TailAdd(1)
	ll.TailAdd(2)
	ll.TailAdd(2)
	ll.TailAdd(1)

	ll.Print()

	flag := IsNotPalindrome(ll)
	if flag {
		fmt.Println("链表是回文结构")
	} else {
		fmt.Println("链表不是回文结构")
	}

	flag1 := IsNotPalindromeImproved(ll)
	if flag1 {
		fmt.Println("链表是回文结构")
	} else {
		fmt.Println("链表不是回文结构")
	}
}
