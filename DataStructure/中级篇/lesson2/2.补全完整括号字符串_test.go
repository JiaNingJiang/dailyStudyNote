package lesson2

import (
	"fmt"
	"testing"
)

func TestCompleteParentheses(t *testing.T) {
	//str1 := "()()()))("
	str2 := "(()())()"
	fmt.Println("缺少的括号个数: ", CompleteParentheses(str2))
}
