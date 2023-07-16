package lesson9

import (
	"fmt"
	"testing"
)

func TestLogicalOpt(t *testing.T) {

	logicStr := "0|1&0|0"
	desire := false

	fmt.Println("运算方法数: ", LogicalOpt(logicStr, desire))

	fmt.Println("运算方法数: ", LogicalOptDP(logicStr, desire))
}
