package bruteRecursionPromote

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestChess(t *testing.T) {
	lineMin := 0
	colMin := 0
	lineMax := 8 // 0~8 共九行
	colMax := 9  // 0~9 共十列

	for i := 0; i < 10000; i++ {
		rand.Seed(time.Now().UnixNano())
		step := rand.Intn(6) + 1
		x := rand.Intn(lineMax)
		y := rand.Intn(colMax)

		res1 := Chess(x, y, step, lineMax, colMax, lineMin, colMin)
		res2 := ChessStrictTable(x, y, step, lineMax, colMax, lineMin, colMin)

		fmt.Printf("(0,0) -> (%d,%d) step:%d    res1:%d    res2:%d\n", x, y, step, res1, res2)
		if res1 != res2 {
			fmt.Println("run Faultily!")
			return
		}
	}
	fmt.Println("run Successfully!")
	//fmt.Println("走法: ", Chess(x, y, step, lineMax, colMax, lineMin, colMin))
	//fmt.Println("走法: ", ChessStrictTable(x, y, step, lineMax, colMax, lineMin, colMin))
}
