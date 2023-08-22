package linkList

import "testing"

func TestSeparateLinkList(t *testing.T) {
	ll := NewLinkList()
	ll.TailAdd(8)
	ll.TailAdd(7)
	ll.TailAdd(5)
	ll.TailAdd(4)
	ll.TailAdd(3)

	ll.Print()

	//SeparateLinkList(ll, 5)
	SeparateLinkListImproved(ll, 9)

	ll.Print()
}
