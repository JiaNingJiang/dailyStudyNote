package lesson4

import (
	"fmt"
	"testing"
)

func TestDymMostFreq(t *testing.T) {
	dmf := NewDymMostFreq(4)

	dmf.AddWord("A").AddWord("B").AddWord("C").AddWord("D").AddWord("A")
	dmf.AddWord("E").AddWord("E").AddWord("E").AddWord("C").AddWord("C")

	fmt.Println(dmf.Pop())
	fmt.Println(dmf.Pop())
	fmt.Println(dmf.Pop())
	fmt.Println(dmf.Pop())
}
