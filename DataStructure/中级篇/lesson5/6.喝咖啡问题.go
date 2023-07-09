package lesson5

import (
	"DataStructure2/utils"
)

// 记录一台咖啡机的状态
type Coffer struct {
	Available int // 下一次可用的时间点
	Consume   int // 用一次需要耗费的时间
}

func LessCoffer(a, b interface{}) bool {
	aCoffer := a.(Coffer)
	bCoffer := b.(Coffer)

	aFinish := aCoffer.Consume + aCoffer.Available
	bFinish := bCoffer.Consume + bCoffer.Available
	if aFinish < bFinish {
		return true
	} else {
		return false
	}
}

// 记录喝咖啡问题的参数
type CoffeePro struct {
	Man              int   // 人数
	CoffeeConsume    []int // 各台咖啡机冲咖啡需要的时间
	WashConsume      int   // 洗一杯咖啡杯需要的时间
	EvaporateConsume int   // 挥发需要的时间

	DrinkSmallRoot []interface{} // 小根堆，用于计算完成Man人数喝咖啡问题的最早结束
	DrinkedTime    []int         // 记录N个人喝完咖啡的时间点
}

func NewCoffeePro(man int, coffeeConsume []int, washConsume int, evaporateConsume int) *CoffeePro {
	cfp := &CoffeePro{
		Man:              man,
		CoffeeConsume:    coffeeConsume,
		WashConsume:      washConsume,
		EvaporateConsume: evaporateConsume,
		DrinkSmallRoot:   utils.NewHeap(make([]interface{}, 0, len(coffeeConsume)), false, LessCoffer),
		DrinkedTime:      make([]int, 0, man),
	}
	// 将三个咖啡机的初始状态加入到小根堆中
	for i := 0; i < len(coffeeConsume); i++ {
		utils.HeapInsert(&cfp.DrinkSmallRoot, Coffer{0, coffeeConsume[i]}, false, LessCoffer)
	}
	return cfp
}

func (cfp *CoffeePro) MakeCoffee() {
	// 1.从小根堆中取出一个节点相当于某人用咖啡机做了一杯咖啡
	data := utils.PopAndheapify(&cfp.DrinkSmallRoot, false, LessCoffer)
	coffer := data.(Coffer)
	cfp.DrinkedTime = append(cfp.DrinkedTime, coffer.Consume+coffer.Available)

	// 2.当前人喝完以后，咖啡杯可以供其他人继续使用
	coffer.Available = coffer.Available + coffer.Consume
	utils.HeapInsert(&cfp.DrinkSmallRoot, coffer, false, LessCoffer)
}

// 让所有人完成喝咖啡，得到完整的drinkTime数组
func (cfp *CoffeePro) FinCoffeeDrink() {
	for i := 0; i < cfp.Man; i++ {
		cfp.MakeCoffee()
	}
}

// 计算让所有杯子变干净的最早时间
func (cfp *CoffeePro) MinTime() int {
	return minTime(cfp.DrinkedTime, cfp.WashConsume, cfp.EvaporateConsume, 0, 0)
}

func minTime(drinkedTime []int, washConsume int, evaporateConsume int, index int, washLine int) int {
	if index == len(drinkedTime)-1 {
		return getMin(getMax(washLine, drinkedTime[index])+washConsume, drinkedTime[index]+evaporateConsume)
	}

	// 1.计算当前咖啡杯采用机器洗的方法最终结束的时间
	wash := getMax(washLine, drinkedTime[index]) + washConsume
	next1 := minTime(drinkedTime, washConsume, evaporateConsume, index+1, wash)
	p1 := getMax(wash, next1)

	// 2.计算当前咖啡杯采用自然蒸发的方法最终结束的时间
	evaporate := drinkedTime[index] + evaporateConsume
	next2 := minTime(drinkedTime, washConsume, evaporateConsume, index+1, washLine)
	p2 := getMax(evaporate, next2)

	return getMin(p1, p2)
}
