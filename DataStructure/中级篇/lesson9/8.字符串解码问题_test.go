package lesson9

import (
	"fmt"
	"testing"
)

func TestDesignatedStrIndex(t *testing.T) {

	fmt.Println("长度为1的字符串个数: ", fixedLen(1))
	fmt.Println("长度为2的字符串个数: ", fixedLen(2))

	fmt.Println("当前字符串位置: ", DesignatedStrIndex("abc"))
}
