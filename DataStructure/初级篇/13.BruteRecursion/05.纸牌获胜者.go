package BruteRecursion

import (
	"fmt"
	"math"
)

func SmartCard(card []int) {
	firstScore := first(card, 0, len(card)-1)
	secondScore := second(card, 0, len(card)-1)

	if firstScore > secondScore {
		fmt.Printf("比分为 %d:%d ,先手获胜\n", firstScore, secondScore)
	} else if firstScore == secondScore {
		fmt.Printf("比分为 %d:%d ,平手\n", firstScore, secondScore)
	} else {
		fmt.Printf("比分为 %d:%d ,后手获胜\n", firstScore, secondScore)
	}
}

// 先手函数
func first(card []int, start, end int) int {
	if start == end { // 只剩一张牌，那么作为先手就拿到这张牌
		return card[start]
	}
	// 作为先手，可以先选一张牌,之后要作为后手。math.Max表示先手者总是会做出对自己最有利的选择
	return int(math.Max(float64(card[start]+second(card, start+1, end)), float64(card[end]+second(card, start, end-1))))
}

// 后手函数
func second(card []int, start, end int) int {
	if start == end { // 只剩一张牌，那么作为后手就什么也拿不到
		return 0
	}
	// 作为后手，只能等待先手完成选牌后再选。math.Min表示后手者总是被迫接受对自己最不利的选择(因为对方会做出比自己好的选择)
	return int(math.Min(float64(first(card, start+1, end)), float64(first(card, start, end-1))))
}
