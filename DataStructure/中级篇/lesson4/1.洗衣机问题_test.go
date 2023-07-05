package lesson4

import (
	"fmt"
	"testing"
)

func TestWashMachine(t *testing.T) {
	//washer := []int{1, 2, 3, 4, 5}
	washer := []int{3, 3, 3, 3, 3}

	fmt.Printf("完成均分,至少需要(%d)轮\n", WashMachine(washer))
}
