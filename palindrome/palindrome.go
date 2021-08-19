package palindrome

import (
	"fmt"
	"strings"
)

func IsNumberPalindrome(num int) bool {
	input := num
	if num < 10 {
		return true
	}
	var rev int
	for num > 0 {
		rem := num % 10
		rev = rev*10 + rem
		num = num / 10
		fmt.Println("rev:", rev)
	}
	fmt.Println("Reversed value", rev)

	return rev == input

}

func IsStringPalindrome(str string) bool {
	actual := strings.ToLower(str)
	var output string
	if len(str) <= 1 {
		return true
	}

	for i := len(actual) - 1; i >= 0; i-- {

		fmt.Println("Str:", string(actual[i]))
		output += string(actual[i])

	}
	fmt.Println("Out", output)
	return output == actual
}
