package main

import (
	"fmt"
	"sort"
)

// 用户自定义的字符串数组要进行排序，必须让改自定义数组实现sort.Interface的三个接口
type MyStringList []string

func (m MyStringList) Len() int {
	return len(m)
}

// 如果返回 m[i] < m[j] 时返回true则为升序排序；如果 m[i] > m[j] 时返回true则为降序排序
func (m MyStringList) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m MyStringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

// 声明英雄的分类
type HeroKind int

// 定义HeroKind常量, 类似于枚举
const (
	None HeroKind = iota
	Tank
	Assassin
	Mage
)

type MyStruct struct {
	Name string
	Kind HeroKind
}

// 对结构体数组进行排序
type MyStructList []*MyStruct

func (m MyStructList) Len() int {
	return len(m)
}

func (m MyStructList) Less(i, j int) bool {
	if m[i].Kind != m[j].Kind {
		return m[i].Kind < m[j].Kind // 优先按照 MyStruct.Kind进行升序排序
	}
	return m[i].Name < m[j].Name
}

func (m MyStructList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func main1() {
	names := MyStringList{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	sort.Sort(names)

	for _, v := range names {
		fmt.Printf("%s\n", v)
	}

	heros := MyStructList{
		&MyStruct{"吕布", Tank},
		&MyStruct{"李白", Assassin},
		&MyStruct{"妲己", Mage},
		&MyStruct{"貂蝉", Assassin},
		&MyStruct{"关羽", Tank},
		&MyStruct{"诸葛亮", Mage},
	}

	sort.Sort(heros)
	// 遍历英雄列表打印排序结果
	for _, v := range heros {
		fmt.Printf("%+v\n", v)
	}
}
