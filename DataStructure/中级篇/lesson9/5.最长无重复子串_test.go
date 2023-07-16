package lesson9

import (
	"fmt"
	"testing"
)

func TestMaxUnique(t *testing.T) {
	//str := "abcabcbb"
	//str := "bbbbb"
	str := "pwwkew"
	fmt.Println("最长无重复子串长度: ", MaxUnique(str))
}
