package lesson8

import (
	"fmt"
	"testing"
)

func TestMinLight(t *testing.T) {
	area := "..x...xx.x..x."

	fmt.Println("需要的最少路灯数: ", MinLight(area))
}
