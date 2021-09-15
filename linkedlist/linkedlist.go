package linkedlist

import "fmt"

// Node in the linked list holds data and next pointer
type Node struct {
	Data int
	Next *Node
}

// LinkList holds head which is pointing to first node
type LinkList struct {
	Head   *Node
	Length int
}

func (l *LinkList) Prepend(n *Node) {
	temp := l.Head
	l.Head = n
	l.Head.Next = temp
	l.Length++
}

// Delete the value
func (l *LinkList) Delete(value int) {
	if l.Length == 0 {
		return
	}
	if l.Head.Data == value {
		l.Head = l.Head.Next
		l.Length--
		return
	}

	temp := l.Head
	var deletedElement int
	for temp.Next.Data != value {
		if temp.Next.Next == nil {
			return
		}
		temp = temp.Next
	}
	deletedElement = temp.Next.Data
	temp.Next = temp.Next.Next
	l.Length--

	fmt.Printf("Del element is %d\n", deletedElement)
}

func (l LinkList) Get() {
	temp := l.Head
	for l.Length != 0 {

		fmt.Printf("%d ", temp.Data)
		temp = temp.Next
		l.Length--
	}
	fmt.Printf("\n")
}
