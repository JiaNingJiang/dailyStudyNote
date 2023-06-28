package prefixTree

import (
	"fmt"
	"testing"
)

func TestPrefixTree(t *testing.T) {
	root := NewNode()

	root.Insert("nihao")
	root.Insert("nigao")
	root.Insert("niqao")
	fmt.Println(root.Search("nihao"))
	fmt.Println(root.Search("nigao"))
	fmt.Println(root.Search("niqao"))
	fmt.Println(root.SearchPre("ni"))

	fmt.Println()

	root.Delete("nihao")
	fmt.Println(root.Search("nihao"))
	fmt.Println(root.Search("nigao"))
	fmt.Println(root.Search("niqao"))
	fmt.Println(root.SearchPre("ni"))
}
