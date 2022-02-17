package strtoint

import "fmt"

func MyAtoi(str string) int {
	var result int
	for i, s := range str {
		fmt.Println("ss", str[i])
		if i == 0 && string(s) == " " {
			continue
		}
	}
	fmt.Println("Result")
	return result
}
