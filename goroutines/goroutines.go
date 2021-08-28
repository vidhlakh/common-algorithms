package main

import (
	"fmt"
	"time"
)

func addname(name chan string) {
	name <- "vidhya"
}
func hello(name chan string) {
	data := <-name
	fmt.Println("Hello world", data)
}
func hello2(name chan string) {

	fmt.Println("Hello world goroutine")
}
func main() {
	name := make(chan string)
	go hello(name)
	go addname(name)
	go hello2(name)
	fmt.Println("main function")
	time.Sleep(1 * time.Second)
}
