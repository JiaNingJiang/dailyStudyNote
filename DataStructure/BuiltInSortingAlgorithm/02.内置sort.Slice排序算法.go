package main

import (
	"fmt"
	"sort"
)

// sort.Slice(x any, less func(i int, j int) bool)
// 第一个参数是要排序的数组arr，第二个参数是比较函数：如果arr[i]<arr[j]时返回true则为升序，如果arr[i]>arr[j]时返回true则为降序
// 具体的大小关系，也需要在less函数中指定
func main() {
	heros := []*MyStruct{
		{"吕布", Tank},
		{"李白", Assassin},
		{"妲己", Mage},
		{"貂蝉", Assassin},
		{"关羽", Tank},
		{"诸葛亮", Mage},
	}
	sort.Slice(heros, func(i, j int) bool {
		if heros[i].Kind != heros[j].Kind {
			return heros[i].Kind < heros[j].Kind // 优先按照MyStruct.Kind进行升序排序
		}
		return heros[i].Name < heros[j].Name // 按照MyStruct.Name进行升序排序
	})
	for _, v := range heros {
		fmt.Printf("%+v\n", v)
	}
}
