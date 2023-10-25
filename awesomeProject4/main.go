package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "12223"

	subStr := "4"

	index := strings.Index(str, subStr)

	fmt.Println(index)
}
