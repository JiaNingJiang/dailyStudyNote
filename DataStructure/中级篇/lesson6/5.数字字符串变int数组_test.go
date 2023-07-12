package lesson6

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {

	//str := fmt.Sprintf("%d", 9223372036854775808)
	str := "9223372036854775807"

	fmt.Println("转换前: ", str)
	fmt.Println("转换结果: ", Convert(str))
}
