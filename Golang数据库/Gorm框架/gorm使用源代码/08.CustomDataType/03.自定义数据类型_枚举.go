package CustomDataType

import (
	"encoding/json"
	"fmt"
	"time"
)

type Weekday int

const (
	Sunday    Weekday = iota + 1 // EnumIndex = 1
	Monday                       // EnumIndex = 2
	Tuesday                      // EnumIndex = 3
	Wednesday                    // EnumIndex = 4
	Thursday                     // EnumIndex = 5
	Friday                       // EnumIndex = 6
	Saturday                     // EnumIndex = 7
)

var WeekStringList = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
var WeekTypeList = []Weekday{Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday}

// String 转字符串
func (w Weekday) String() string { //当以字符串格式输出Weekday类型数据时，会调用此方法
	return WeekStringList[w-1]
}

// MarshalJSON 自定义类型转换为json
func (w Weekday) MarshalJSON() ([]byte, error) {
	return json.Marshal(w.String()) //先将weekday类型转化为字符串格式，再进行json格式化
}

// EnumIndex 自定义类型转原始类型
func (w Weekday) EnumIndex() int { //将weekday类型转化为int类型
	return int(w)
}

// ParseWeekDay 字符串转自定义类型
func ParseWeekDay(week string) Weekday { //根据传入的字符串，返回对应的Weekday类型数据
	for i, i2 := range WeekStringList {
		if week == i2 {
			return WeekTypeList[i]
		}
	}
	return Monday
}

// ParseIntWeekDay 数字转自定义类型
func ParseIntWeekDay(week int) Weekday { //将int类型转化为weekday类型
	return Weekday(week)
}

type DayInfo struct {
	Weekday Weekday   `json:"weekday"` //weekday自定义数据类型
	Date    time.Time `json:"date"`
}

func MeiJu() {
	w := Sunday //weekday类型
	fmt.Println(w)
	dayInfo := DayInfo{Weekday: Sunday, Date: time.Now()}
	data, err := json.Marshal(dayInfo)
	fmt.Println(string(data), err)
	week := ParseWeekDay("Sunday")
	fmt.Println(week)
	week = ParseIntWeekDay(2)
	fmt.Println(week)
}
