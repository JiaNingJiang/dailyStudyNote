package linkList

import "testing"

func TestPublicSection(t *testing.T) {
	l1 := NewLinkList()
	l2 := NewLinkList()

	l1.TailAdd(3)
	l1.TailAdd(4)
	l1.TailAdd(5)
	l1.TailAdd(6)
	l1.TailAdd(7)

	l2.TailAdd(3)
	l2.TailAdd(6)
	l2.TailAdd(9)

	l1.Print()
	l2.Print()

	PublicSection(l1, l2)
}
