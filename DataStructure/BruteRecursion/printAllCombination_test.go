package BruteRecursion

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrintAllCombination(t *testing.T) {
	count1, res1 := PrintAllCombination("hello")
	count2, res2 := PrintAllCombinationImproved("hello")
	fmt.Println(count1, count2)
	fmt.Println(reflect.DeepEqual(res1, res2))
}
