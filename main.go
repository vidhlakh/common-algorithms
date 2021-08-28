package main

import (
	"fmt"
	"sort"

	"github.com/vidhlakh/common-algorithms/anagram"
	"github.com/vidhlakh/common-algorithms/intersection"
	"github.com/vidhlakh/common-algorithms/stringmanipulate"
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
	stringmanipulate.Stackify(input)

	n1 := []int{1, 2, 1, 2}
	n2 := []int{1, 0, 2}
	sort.Ints(n1)
	sort.Ints(n2)
	intersectionArray := intersection.IntersectionOfTwoSortedArrays(n1, n2)
	fmt.Println("Intersection of sorted array:", intersectionArray)
}
