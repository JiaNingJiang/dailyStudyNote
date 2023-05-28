package stringProblem

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestKMP(t *testing.T) {
	//str := "abbtabbkabbtabbz"
	//subStr := "abbtabbz"
	str := "222hellz world hello!!!"
	subStr := "hello"

	fmt.Println("匹配下标: ", KMP(str, subStr, 0))

	comparator(KMP, PlainMatch)
}

func comparator(f1, f2 func(string, string, int) int) bool {
	var testTime int = 5000 // 比较次数
	var maxLen int = 20     // 测试用输入主字符串的最大大小
	var succeed bool = true

	for i := 0; i < testTime; i++ {
		str, subStr := generateRandomStr(maxLen)

		index1 := f1(str, subStr, 0)
		index2 := f2(str, subStr, 0)
		if !reflect.DeepEqual(index1, index2) {
			fmt.Printf("主串:(%s) 子串:(%s)   index1:(%d) index2(%d) \n", str, subStr, index1, index2)
			succeed = false
			break
		} else {
			fmt.Printf("主串:(%s) 子串:(%s)   index1:(%d) index2(%d) \n", str, subStr, index1, index2)
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

func generateRandomStr(maxLen int) (string, string) {
	rand.Seed(time.Now().UnixNano())
	strLen := rand.Intn(maxLen) + 2 // 主串的长度
	subLen := 3                     // 子串的长度

	var str string
	// 生成主串
	for i := 0; i < strLen; i++ {
		str += fmt.Sprintf("%d", rand.Intn(2))
	}

	var subStr string
	// 生成子串
	for i := 0; i < subLen; i++ {
		subStr += fmt.Sprintf("%d", rand.Intn(2))
	}
	return str, subStr
}
