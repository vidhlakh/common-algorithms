package plusone

import "fmt"

// PlusOne add 1 to the arrray . Conside array as digit
func PlusOne(digits []int) []int {
	N := len(digits)
	for i := N - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0

		}
	}

	digits = append([]int{1}, digits...)
	fmt.Println("Digit", digits)
	return digits
}
