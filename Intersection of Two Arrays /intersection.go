package intersection

import (
	"fmt"
)

func IntersectionOfTwoSortedArrays(n1 []int, n2 []int) []int {

	var res []int

	N1 := len(n1)
	N2 := len(n2)
	fmt.Println("a:", n1, "n2:", n2)
	i := 0
	j := 0
	for i < N1 && j < N2 {
		if n1[i] > n2[j] {
			j++

		} else if n2[j] > n1[i] {
			i++
		} else {
			res = append(res, n1[i])
			i++
			j++
		}
	}
	return res
}
