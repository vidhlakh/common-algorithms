package main

import (
	"fmt"

	"github.com/vidhlakh/common-algorithms/anagram"
	"github.com/vidhlakh/common-algorithms/rotatematrix"
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
	// input := "abc#"
	// stack.Stackify(input)

	// n1 := []int{1, 2, 1, 2}
	// n2 := []int{1, 0, 2}
	// intersectionArray2 := intersection.IntersectionOfTwoArrays(n1, n2)
	// fmt.Println("Intersection of two array:", intersectionArray2)
	// // sort.Ints(n1)
	// // sort.Ints(n2)
	// intersectionArray := intersection.IntersectionOfTwoSortedArrays(n1, n2)
	// fmt.Println("Intersection of sorted array:", intersectionArray)
	// a := []int{}
	// b := []int{}
	// if reflect.DeepEqual(a, b) {
	// 	fmt.Println("Equal")
	// }
	// //pluone
	// plusoneresult := plusone.PlusOne([]int{1, 2, 3})
	// fmt.Println("Plus one :", plusoneresult)
	// // move zeroes
	// movezeroesResult := movezeros.MoveZeroesInplace([]int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0})
	// fmt.Println("Move zeroes :", movezeroesResult)

	// twosumsResult := twosums.TwoSumBrute([]int{3, 2, 4}, 6)
	// fmt.Println("Twosums", twosumsResult)

	// // valid sudoku
	// sudokuInput := [][]string{
	// 	{"5", "3", ".", ".", "7", ".", ".", ".", "."},
	// 	{"6", ".", ".", "1", "9", "5", ".", ".", "."},
	// 	{".", "9", "8", ".", ".", ".", ".", "6", "."},
	// 	{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
	// 	{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
	// 	{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
	// 	{".", "6", ".", ".", ".", ".", "2", "8", "."},
	// 	{".", ".", ".", "4", "1", "9", ".", ".", "5"},
	// 	{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	// }

	// for i, col := range sudokuInput {
	// 	Map := map[string]int{}
	// 	for j := range col {
	// 		e := sudokuInput[j][i]
	// 		fmt.Println("j,i:", j, i)
	// 		if e != "." {

	// 			if _, ok := Map[e]; ok {
	// 				fmt.Println("Duplicate present")
	// 			} else {
	// 				Map[e] = 1
	// 			}
	// 		}
	// 	}
	// }

	// // LinkedList
	// ls := linkedlist.LinkList{}
	// node1 := &linkedlist.Node{Data: 48}
	// node2 := &linkedlist.Node{Data: 37}
	// node3 := &linkedlist.Node{Data: 25}
	// node4 := &linkedlist.Node{Data: 15}
	// node5 := &linkedlist.Node{Data: 10}
	// node6 := &linkedlist.Node{Data: 5}
	// ls.Prepend(node1)
	// ls.Prepend(node2)
	// ls.Prepend(node3)
	// ls.Prepend(node4)
	// ls.Prepend(node5)
	// ls.Prepend(node6)
	// fmt.Println("List:", ls)
	// ls.Get()
	// ls.Delete(100)
	// ls.Get()

	// // Binary Search Tree
	// tree := &binarySearchTree.Node{Key: 100}
	// tree.Insert(52)
	// tree.Insert(7)
	// tree.Insert(24)
	// tree.Insert(19)
	// tree.Insert(76)
	// tree.Insert(88)
	// tree.Insert(203)
	// tree.Insert(150)
	// tree.Insert(276)
	// tree.Insert(310)

	// fmt.Println("Tree:", tree)
	// fmt.Println("Count:", binarySearchTree.Count)
	// binarySearchTree.Count = 0
	// tree.Search(76)
	// fmt.Println("Count:", binarySearchTree.Count)
	// binarySearchTree.Count = 0
	// tree.Search(203)
	// fmt.Println("Count:", binarySearchTree.Count)
	// binarySearchTree.Count = 0
	// tree.Search(52)
	// fmt.Println("Count:", binarySearchTree.Count)
	// binarySearchTree.Count = 0
	// tree.Search(276)
	// fmt.Println("Count:", binarySearchTree.Count)

	// //solid polymorphism
	// // declaring a rectangle instance
	// rectangle := solid.Rectangle{

	// 	Length: 10.5,
	// 	Width:  12.25,
	// }

	// // declaring a square instance
	// square := solid.Square{

	// 	Side: 15.0,
	// }

	// // printing the calculated result
	// fmt.Printf("Area of rectangle: %.3f unit sq.\n", rectangle.Area())
	// fmt.Printf("Area of square: %.3f unit sq.\n", square.Area())

	// // reverse a string

	// reverseinput := "hello"
	// reverse.ReverseString([]byte(reverseinput))

	// revint := 1534236469
	// _ = reverse.ReverseIntSimple(revint)

	// //find unique char
	// findDuplicates.FirstUniqChar("loveleetcode")

	//Rotate matrix
	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotatematrix.RotateMatrix(input)
}
