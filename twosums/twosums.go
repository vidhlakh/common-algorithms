package twosums

// TwoSumBrute add 2 elemnt of array to get targer
func TwoSumBrute(nums []int, target int) []int {
	N := len(nums)
	i := 0
	for i < N-1 {
		for j := 1; j < N; j++ {
			if i != j && nums[i]+nums[j] == target {

				return []int{i, j}
			}
		}
		i++
	}
	return []int{}
}

// TwoSumUsingHashMap add 2 elemnt of array to get target
func TwoSumUsingHashMap(nums []int, target int) []int {
	N := len(nums)
	var keyStore = make(map[int]int)
	for i := 0; i < N; i++ {
		keyToFind := target - nums[i]
		if value, ok := keyStore[keyToFind]; !ok {

			return []int{i, value}
		} else {
			keyStore[nums[i]] = i
		}
	}

	return []int{}
}
