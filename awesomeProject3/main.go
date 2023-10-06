package main

import (
	"fmt"
	"sort"
)

func main() {
	var set [3]struct{}
	sort.SearchInts()
	for i := range set {
		defer func() {
			fmt.Println(i)
		}()
	}
}
