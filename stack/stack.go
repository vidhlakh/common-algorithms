package main

import (
	"fmt"
	"strings"
)

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
func (s *Stack) Push(a string) {
	fmt.Println("In push")
	*s = append(*s, a)
	fmt.Println("S push:", *s)
}

func (s *Stack) Pop() {
	if !s.IsEmpty() {
		fmt.Println("Pop", s)
		index := len(*s) - 1
		fmt.Println("length", index)
		*s = (*s)[:index]
	}
}

func Stackify(input string) string {
	var s = make(Stack, 0)
	fmt.Println("Input: ", input)
	//s = strings.Split(input, "")
	for _, str := range input {
		if string(str) != "#" {
			s.Push(string(str))
		} else {
			s.Pop()
		}
	}
	fmt.Println("Output: ", s)
	output := strings.Join(s, "")
	return output
}

func main() {
	input := "#"
	output := Stackify(input)
	fmt.Println("Output:", output)
}
