package main

import (
	"fmt"
	"reflect"

	"github.com/vidhlakh/common-algorithms/anagram"
	"github.com/vidhlakh/common-algorithms/intersection"
	"github.com/vidhlakh/common-algorithms/plusone"
	"github.com/vidhlakh/common-algorithms/stack"
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
	stack.Stackify(input)

	n1 := []int{1, 2, 1, 2}
	n2 := []int{1, 0, 2}
	intersectionArray2 := intersection.IntersectionOfTwoArrays(n1, n2)
	fmt.Println("Intersection of two array:", intersectionArray2)
	// sort.Ints(n1)
	// sort.Ints(n2)
	intersectionArray := intersection.IntersectionOfTwoSortedArrays(n1, n2)
	fmt.Println("Intersection of sorted array:", intersectionArray)
	a := []int{}
	b := []int{}
	if reflect.DeepEqual(a, b) {
		fmt.Println("Equal")
	}
	//pluone
	plusoneresult := plusone.PlusOne([]int{1, 2, 3})
	fmt.Println("Plus one :", plusoneresult)
}
