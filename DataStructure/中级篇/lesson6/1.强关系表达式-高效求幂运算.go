package lesson6

func EffectivePow(x interface{}, y int64, selfMul func(interface{}, interface{}) interface{}) interface{} {
	// 1.将y拆解为二进制形式
	bitFormat := make([]byte, 0)
	for i := 0; 1<<i <= y; i++ {
		bitRes := ((1 << i) & y) >> i
		if bitRes == 1 {
			bitFormat = append(bitFormat, 1)
		} else {
			bitFormat = append(bitFormat, 0)
		}
	}
	var powRes interface{}

	for i := 0; i < len(bitFormat); i++ {
		if bitFormat[i] == 1 {
			powIndex := 1 << i
			for j := 0; j < powIndex; j++ {
				powRes = selfMul(powRes, x)
			}
		}
	}

	return powRes
}
