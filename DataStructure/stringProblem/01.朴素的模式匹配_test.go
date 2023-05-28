package stringProblem

import (
	"fmt"
	"testing"
)

func TestPlainMatch(t *testing.T) {
	str := "1001000001101111011"
	subStr := "010"

	fmt.Println("匹配下标: ", PlainMatch(str, subStr, 0))
}
