package bitOperation

func IsTwoPower1(num uint32) bool {
	if num == 0 {
		return false
	}
	rightOne := num & (^num + 1)
	if num == rightOne {
		return true
	} else {
		return false
	}
}

func IsTwoPower2(num uint32) bool {
	if num == 0 {
		return false
	}
	if num&(num-1) == 0 {
		return true
	} else {
		return false
	}
}

func IsFourPower(num uint32) bool {
	if !IsTwoPower2(num) { // 4的幂次必须也同时是2的幂次(确保只有一位上为1，其他位都为0)
		return false
	}
	res := num & (0x55555555) // 为1的这一位是4的幂次
	if res != 0 {
		return true
	} else {
		return false
	}
}
