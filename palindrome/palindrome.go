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

func IsPalindromeSimple(s string) bool {
	var answer bool
	var charList1, charList2 []int

	// get char code list
	for _, charByte := range strings.ToLower(s) {
		charCode := int(charByte)
		if 48 <= charCode && charCode <= 57 || 97 <= charCode && charCode <= 122 {
			charList1 = append(charList1, int(charCode))
			charList2 = append(charList2, int(charCode))
		}
	}

	// reverse
	for i, j := 0, len(charList2)-1; i < j; i, j = i+1, j-1 {
		charList2[i], charList2[j] = charList2[j], charList2[i]
	}

	// check palindrome
	if len(charList1) == len(charList2) {
		answer = true
		for i := 0; i < len(charList1); i++ {
			if charList1[i] != charList2[i] {
				answer = false
				break
			}
		}
	} else {
		answer = false
	}

	return answer
}
