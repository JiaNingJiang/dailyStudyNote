package bruteRecursionPromote

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestCardWin(t *testing.T) {
	card := generateRandomArray2(7, 8)

	fmt.Println(card)

	fmt.Println("最高分数: ", CardWin(card))

	fmt.Println("最高分数: ", CardWinStrictTable(card))
}

func generateRandomArray2(maxCardNumber, maxCardValue int) []int {
	rand.Seed(time.Now().UnixNano())
	cardNum := rand.Intn(maxCardNumber) + 2

	cardSet := make([]int, cardNum)
	for i := 0; i < cardNum; i++ {
		cardSet[i] = rand.Intn(maxCardValue) + 1
	}

	return cardSet
}
