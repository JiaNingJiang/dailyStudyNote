package bitOperation

import (
	"fmt"
	"testing"
)

func TestIsTwoPower1(t *testing.T) {
	fmt.Printf("%d是2的幂次?%v\n", 0, IsTwoPower1(0))
	fmt.Printf("%d是2的幂次?%v\n", 1, IsTwoPower1(1))
	fmt.Printf("%d是2的幂次?%v\n", 2, IsTwoPower1(2))
	fmt.Printf("%d是2的幂次?%v\n", 3, IsTwoPower1(3))
	fmt.Printf("%d是2的幂次?%v\n", 4, IsTwoPower1(4))
	fmt.Printf("%d是2的幂次?%v\n", 5, IsTwoPower1(5))
	fmt.Printf("%d是2的幂次?%v\n", 6, IsTwoPower1(6))
	fmt.Printf("%d是2的幂次?%v\n", 7, IsTwoPower1(7))
	fmt.Printf("%d是2的幂次?%v\n", 8, IsTwoPower1(8))
}

func TestIsTwoPower2(t *testing.T) {
	fmt.Printf("%d是2的幂次?%v\n", 0, IsTwoPower2(0))
	fmt.Printf("%d是2的幂次?%v\n", 1, IsTwoPower2(1))
	fmt.Printf("%d是2的幂次?%v\n", 2, IsTwoPower2(2))
	fmt.Printf("%d是2的幂次?%v\n", 3, IsTwoPower2(3))
	fmt.Printf("%d是2的幂次?%v\n", 4, IsTwoPower2(4))
	fmt.Printf("%d是2的幂次?%v\n", 5, IsTwoPower2(5))
	fmt.Printf("%d是2的幂次?%v\n", 6, IsTwoPower2(6))
	fmt.Printf("%d是2的幂次?%v\n", 7, IsTwoPower2(7))
	fmt.Printf("%d是2的幂次?%v\n", 8, IsTwoPower2(8))
}

func TestIsFourPower(t *testing.T) {
	fmt.Printf("%d是4的幂次?%v\n", 0, IsFourPower(0))
	fmt.Printf("%d是4的幂次?%v\n", 1, IsFourPower(1))
	fmt.Printf("%d是4的幂次?%v\n", 2, IsFourPower(2))
	fmt.Printf("%d是4的幂次?%v\n", 3, IsFourPower(3))
	fmt.Printf("%d是4的幂次?%v\n", 4, IsFourPower(4))
	fmt.Printf("%d是4的幂次?%v\n", 5, IsFourPower(5))
	fmt.Printf("%d是4的幂次?%v\n", 6, IsFourPower(6))
	fmt.Printf("%d是4的幂次?%v\n", 7, IsFourPower(7))
	fmt.Printf("%d是4的幂次?%v\n", 8, IsFourPower(8))
	fmt.Printf("%d是4的幂次?%v\n", 16, IsFourPower(16))
}
