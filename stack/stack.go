package stacks

import "fmt"

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
func (s *Stack) Push(a string) {
	*s = append(*s, a)
	fmt.Println("S push:", s)
}

func (s *Stack) Pop() {
	if !s.IsEmpty() {
		*s = s[:len(*s)-1]
	}

}

func Stackify(input string) {
	var s *Stack
	//s = strings.Split(input, "")
	for i := 0; i < len(input); i++ {
		if string(input[i]) != "#" {
			s.Push(string(input[i]))
		} else {
			s.Pop()
		}
	}
	return
}
