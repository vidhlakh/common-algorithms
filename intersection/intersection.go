package intersection

import (
	"fmt"
	"sort"
)

func IntersectionOfTwoSortedArrays(n1 []int, n2 []int) []int {
	sort.Ints(n1)
	sort.Ints(n2)
	res := []int{}

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

// Extra space using maps
func IntersectionOfTwoArrays(n1 []int, n2 []int) []int {

	m := map[int]int{}
	res := []int{}
	for _, n := range n1 {
		m[n]++
	}
	for _, n := range n2 {
		if m[n] > 0 {
			res = append(res, n)
			m[n]--
		}
	}
	return res
}
