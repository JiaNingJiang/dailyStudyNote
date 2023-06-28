package hashTable

import "errors"

type BitMap struct {
	base []int32
}

// 新建一个对应bit位数的位图
func NewBitMap(bitSize uint64) *BitMap {
	num := bitSize / 32
	extra := bitSize % 32
	if extra != 0 { // 假如bitSize == 33 , 那么必定需要2个int32才能表示；假如bitSize == 31,那么需要1个int32
		num += 1
	}
	bitMap := &BitMap{base: make([]int32, num)}

	return bitMap
}

// 检查指定bit位的状态：是1则返回true; 是0则返回false
func (bm *BitMap) CheckBitStatus(index uint64) (bool, error) {
	if index > uint64(len(bm.base)*32) {
		return false, errors.New("位数不对")
	}

	numIndex := index / 32 // 找到i具体在哪一个int数上
	bitIndex := index % 32 // 找到i在这一个int数的哪一位上

	// 拿到第i位的具体状态（是1还是0）
	s := ((bm.base[numIndex] >> (bitIndex)) & 1)
	if s == 1 {
		return true, nil
	} else {
		return false, nil
	}
}

func (bm *BitMap) SetBitOne(index uint64) error {
	if index > uint64(len(bm.base)*32) {
		return errors.New("位数不对")
	}

	numIndex := index / 32 // 找到i具体在哪一个int数上
	bitIndex := index % 32 // 找到i在这一个int数的哪一位上

	// 把第i位的状态改成1
	bm.base[numIndex] = bm.base[numIndex] | (1 << (bitIndex))
	return nil
}

func (bm *BitMap) SetBitZero(index uint64) error {
	if index > uint64(len(bm.base)*32) {
		return errors.New("位数不对")
	}

	numIndex := index / 32 // 找到i具体在哪一个int数上
	bitIndex := index % 32 // 找到i在这一个int数的哪一位上

	// 把第i位的状态改成0
	bm.base[numIndex] = bm.base[numIndex] & (^(1 << bitIndex))

	return nil
}
