package findDuplicates

import "fmt"

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// FindDuplicates find duplicates in an array
//[1,2,4,5,6]
// 
func FindDuplicates(nums []int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {

		index := Abs(nums[i]) - 1
		fmt.Println("Index:", index)
		if nums[index] < 0 {
			result = append(result, index+1)
		} else {
			nums[index] = nums[index] * -1
		}
		fmt.Println("Nums:", nums)
	}
	return result
}
