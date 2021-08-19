package main

import (
	"fmt"

	"github.com/vidhlakh/common-algorithms/anagram"
	"github.com/vidhlakh/common-algorithms/stacks"
)

func main() {
	// result := palindrome.IsNumberPalindrome(2002)
	// fmt.Println("Result", result)
	// if result {
	// 	fmt.Println("Number is a palindrome ")
	// } else {
	// 	fmt.Println("Number is not a palindrome ")
	// }
	// result := palindrome.IsStringPalindrome("RaceCAR")
	// fmt.Println("Result", result)
	// if result {
	// 	fmt.Println("String is a palindrome ")
	// } else {
	// 	fmt.Println("String is not a palindrome ")
	// }
	result := anagram.IsAnagram("actia", "acati")
	fmt.Println("Result", result)
	if result {
		fmt.Println("Strings are anagrams ")
	} else {
		fmt.Println("String are not anagrams ")
	}
	input := "abc#"
	stacks.Stackify(input)
}
