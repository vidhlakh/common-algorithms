package reverse

import "fmt"

func ReverseString(s []byte) (output string) {
	fmt.Println("S:", s)
	N := len(s)
	if N <= 1 {
		return string(s)
	}
	i, j := 0, N-1
	for ; i <= N/2; i++ {
		if j >= N/2 {
			s[i], s[j] = s[j], s[i]
		}
		j--
	}
	fmt.Println("After reverse S:", s)
	fmt.Println("Reverse string", string(s))
	return string(s)
}
func reverseStringsimple(s []byte) {
	l := len(s)
	i, j := 0, l-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
