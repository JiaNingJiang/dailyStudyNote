package bruteRecursionPromote

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestRobotWalk(t *testing.T) {
	totalStep := 3
	end := 2
	fmt.Printf("总步数:%d  目标地点:%d   走法:%d\n", totalStep, end,
		RobotWalk(totalStep, end, 1, totalStep))

	fmt.Printf("总步数:%d  目标地点:%d   走法:%d\n", totalStep, end,
		RobotWalkMemoryCache(totalStep, end, 1, totalStep))

	fmt.Printf("总步数:%d  目标地点:%d   走法:%d\n", totalStep, end,
		RobotWalk3(totalStep, end, 1, totalStep))

	Comparator(RobotWalk, RobotWalkMemoryCache)
	fmt.Println()
	Comparator(RobotWalk, RobotWalk3)
}

func Comparator(f1, f2 func(int, int, int, int) int) bool {
	var testTime int = 5000 // 比较次数
	var left int = 1
	var maxStep int = 20

	var succeed bool = true

	for i := 0; i < testTime; i++ {
		arr := generateRandomArray(maxStep, left) // 生成一个长度随机，元素值也随机的数组

		res1 := f1(arr[0], arr[1], 1, arr[2]) // 在 0~arr[2]上行动arr[0]步，返回最终走到arr[1]的走法
		res2 := f2(arr[0], arr[1], 1, arr[2])
		//fmt.Printf("totalStep: %v , end:%v , res1:%v , res2:%v\n", arr[0], arr[1], res1, res2)
		if !reflect.DeepEqual(res1, res2) {
			succeed = false
			fmt.Printf("totalStep: %v , end:%v , res1:%v , res2:%v\n", arr[0], arr[1], res1, res2)
			break
		}
	}

	if succeed {
		fmt.Printf("insertionSort run successfully!")
		return true
	} else {
		fmt.Printf("insertionSort run Faultily!")
		return false
	}
}

func generateRandomArray(maxStep, left int) []int {
	rand.Seed(time.Now().UnixNano())

	arr := make([]int, 3)
	arr[0] = rand.Intn(maxStep)            // 最大步数
	arr[2] = left + rand.Intn(maxStep) + 1 // 右边界(至少比左边界大1)
	arr[1] = left + rand.Intn(arr[2]-left) // 目标位置(必须在左右边界之内)

	return arr
}
