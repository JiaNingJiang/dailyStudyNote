package lesson3

import (
	"fmt"
	"testing"
)

func TestLongestPstr(t *testing.T) {
	pstr := "()()()()())"

	fmt.Println("最长完整字符串长度: ", LongestPstr(pstr))
}
