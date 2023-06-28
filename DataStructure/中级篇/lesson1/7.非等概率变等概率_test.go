package lesson1

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestEqualP(t *testing.T) {
	count1 := 0
	count0 := 0

	round := 100
	for i := 1; i <= round; i++ {
		res := EqualP(NoEuqalP)
		fmt.Printf(" %d ", res)
		if i%10 == 0 {
			fmt.Println()
		}
		if res == 0 {
			count0++
		} else {
			count1++
		}
	}
	fmt.Println("0的概率: ", float32(count0)/float32(round))
	fmt.Println("1的概率: ", float32(count1)/float32(round))
}

// 1/3 概率返回 0 ， 2/3的概率返回1
func NoEuqalP() int {
	time.Sleep(time.Nanosecond)
	rand.Seed(time.Now().UnixNano())

	random := rand.Intn(8) + 1

	if random == 1 || random == 2 || random == 3 {
		return 0
	} else {
		return 1
	}

}
