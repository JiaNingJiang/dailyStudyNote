package BruteRecursion

import (
	"DataStructure/linkList"
	"fmt"
)

// n表示盘子的数量
func Hanoi(n int) {
	left := linkList.NewStack()
	middle := linkList.NewStack()
	right := linkList.NewStack()

	for i := n; i > 0; i-- {
		left.Push(i)
	}
	hanoi(n, left, middle, right)

	for {
		if left.Len == 0 {
			fmt.Println("---------")
			break
		}
		fmt.Printf("%d ", left.Pop())
	}
	for {
		if middle.Len == 0 {
			fmt.Println("---------")
			break
		}
		fmt.Printf("%d ", middle.Pop())
	}

	for {
		if right.Len == 0 {
			fmt.Println("---------")
			break
		}
		fmt.Printf("%d ", right.Pop())
	}
}

// n是剩余要移动的盘子的数量
// start是要移动的盘子的起点柱子
// end是要移动的盘子的终点柱子
// assist是辅助柱子
func hanoi(n int, start, assist, end *linkList.Stack) {
	if n == 0 { // 没有盘子需要移动
		return
	}
	// 1.先将上面的n-1个盘子移动到辅助柱子上（借助end柱子）
	hanoi(n-1, start, end, assist)
	// 2.再将最下面的第n个盘子移动到end柱子上(借助assist柱子)
	data := start.Pop()
	end.Push(data)
	// 3.最后将辅助柱子上的n-1个盘子移动到end柱子上(借助start柱子)
	hanoi(n-1, assist, start, end)
}
