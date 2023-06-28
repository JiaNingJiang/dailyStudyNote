package linkList

import "testing"

func TestLinkList_Add(t *testing.T) {
	ll := NewLinkList()

	//ll.TailAdd(3)
	//ll.TailAdd(4)
	//ll.TailAdd(5)
	ll.HeadAdd(3)
	ll.HeadAdd(4)
	ll.HeadAdd(5)
	ll.Print()

	//ll.Reverse1()
	//ll.Print()

	//ll.Reverse2()
	//ll.Print()

	//ll.Reverse3()
	//ll.Print()

	ll.Reverse4()
	ll.Print()
}
