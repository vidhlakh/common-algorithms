package reverse

import (
	"fmt"
	"math"
)

func ReverseInt(x int) int {
	INT_MAX := math.MaxInt32
	INT_MIN := math.MinInt32
	fmt.Println("Int max", INT_MAX, INT_MIN)
	if x == 0 {
		return x
	}
	var temp int
	for x != 0 {
		r := x % 10
		x = x / 10
		temp = temp*10 + r
	}
	fmt.Println("Temp:", temp)

	if temp > INT_MAX || temp < INT_MIN {

		return 0
	}

	return temp
}
func ReverseIntSimple(x int) int {
	var nums, newnums int
	MaxInt32 := 1<<31 - 1
	MinInt32 := -1 << 31
	fmt.Println("Int max", MaxInt32, MinInt32)
	for x != 0 {
		a := x % 10
		newnums = nums*10 + a

		nums = newnums
		x = x / 10

		if nums > MaxInt32 || nums < MinInt32 {
			return 0
		}
	}

	return nums
}
