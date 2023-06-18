package bruteRecursionPromote

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestBeyondEdge(t *testing.T) {
	lineMin := 0
	colMin := 0
	lineMax := 8 // 0~8 共九行
	colMax := 9  // 0~9 共十列

	for i := 0; i < 10000; i++ {
		rand.Seed(time.Now().UnixNano())
		step := rand.Intn(6) + 1
		x := rand.Intn(lineMax)
		y := rand.Intn(colMax)

		res1 := BeyondEdge(x, y, step, lineMax, colMax, lineMin, colMin)
		res2 := BeyondEdgeStrictTable(x, y, step, lineMax, colMax, lineMin, colMin)

		fmt.Printf("(%d,%d) step:%d    res1:%f    res2:%f\n", x, y, step, res1, res2)
		if res1 != res2 {
			fmt.Println("run Faultily!")
			return
		}
	}
	fmt.Println("run Successfully!")
	//fmt.Println("存活概率： ", BeyondEdge(x, y, step, lineMax, colMax, lineMin, colMin))
	//fmt.Println("存活概率： ", BeyondEdgeStrictTable(x, y, step, lineMax, colMax, lineMin, colMin))
}
