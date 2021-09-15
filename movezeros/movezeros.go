package movezeros

import (
	"fmt"
)

// MoveZeroes using additional space move zeroes to the end
func MoveZeroes(nums []int) []int {
	fmt.Println("Input:", nums)
	N := len(nums)
	res := make([]int, 0)
	count := 0
	for i := 0; i < N; i++ {
		if nums[i] == 0 {
			count += 1
		} else {
			res = append(res, nums[i])
		}
	}
	for j := 0; j < count; j++ {
		res = append(res, 0)
	}
	fmt.Println("Result ", res)
	return res
}

// MoveZeroesInplace inplace  move zeroes to the end
func MoveZeroesInplace(nums []int) []int {
	fmt.Println("Input:", nums)
	if len(nums) == 0 {
		return nums
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
				fmt.Println("i,j", nums[i], nums[j])
			}
			fmt.Println("I:", i, j)
			j++
		}
	}
	return nums
}
