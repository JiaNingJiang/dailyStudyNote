package main

import "fmt"

func main() {
	var set [3]struct{}

	for i := range set {
		defer func() {
			fmt.Println(i)
		}()
	}
}
