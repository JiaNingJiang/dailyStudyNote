package lesson5

import (
	"fmt"
	"testing"
)

func TestRotationStr(t *testing.T) {
	str1 := "abcde"
	str2 := "cdeab"

	if res := RotationStr(str1, str2); res {
		fmt.Printf("(%s)与(%s)互为旋转词\n", str1, str2)
	} else {
		fmt.Printf("(%s)与(%s)没有旋转关系\n", str1, str2)
	}
}
