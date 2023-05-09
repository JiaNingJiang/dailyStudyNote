package linkList

import "testing"

func TestSLinkList(t *testing.T) {
	sl := NewSLinkList()

	sl.TailAdd(1)
	sl.TailAdd(2)
	sl.TailAdd(3)

	sl.RandomAdd(1, 3)
	sl.RandomAdd(2, 1)
	sl.RandomAdd(3, sl.Len+1)

	sl.Print()

	copysl := CopyLinkListWithRandomP(sl)

	copysl.Print()
}
