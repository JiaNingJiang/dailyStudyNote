package stringProblem

import (
	"fmt"
	"testing"
)

func TestManacherFormat(t *testing.T) {
	str := "aabaa"
	sep := "#"

	fmt.Println(manacherFormat(str, sep))
}

func TestManacher(t *testing.T) {
	str := "aacac"

	fmt.Println("最大回文长度:", Manacher(str))
}
