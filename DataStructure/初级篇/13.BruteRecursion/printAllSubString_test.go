package BruteRecursion

import (
	"fmt"
	"testing"
)

func TestPrintAllSubString(t *testing.T) {
	count1 := PrintAllSubString("hello!")
	count2 := PrintAllSubStrImproved("hello!")

	fmt.Println(count1 == count2)
}
