package findDuplicates

import "fmt"

func FirstUniqChar(s string) int {
	N := len(s)
	var m = make(map[string]int)
	for i := 0; i < N; i++ {

		m[string(s[i])] += 1

	}
	fmt.Println("map:", m)
	for i := 0; i < N; i++ {
		if v, ok := m[string(s[i])]; ok && v == 1 {

			fmt.Println("char present", string(s[i]), i)
			return i
		}
	}

	return -1
}

func FirstUniqCharSimple(s string) int {
	indices := make([]int, 26)

	for i := 0; i < len(s); i++ {
		c := s[i] - 'a'
		v := indices[c]
		if v == 0 {
			indices[c] = i + 1
		} else if v > 0 {
			// mark duplicate
			indices[c] = -1
		}
	}

	var min = len(s)
	for i := 0; i < len(indices); i++ {
		k := indices[i]
		if k > 0 && k-1 <= min {
			min = k - 1
		}
	}

	if min >= len(s) {
		return -1
	}

	return min
}
