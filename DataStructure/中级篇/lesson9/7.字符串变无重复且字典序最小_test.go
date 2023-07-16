package lesson9

import (
	"fmt"
	"testing"
)

func TestOperation(t *testing.T) {
	str := "dbcacbca"
	fmt.Println("结果: ", Operation(str))
}
