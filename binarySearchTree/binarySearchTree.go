package binarySearchTree

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

var Count int

func (n *Node) Insert(k int) {
	if k < n.Key {
		//move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		}
		n.Left.Insert(k)
	} else if k > n.Key {
		// move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		}
		n.Right.Insert(k)
	}
}

func (n *Node) Search(k int) bool {
	Count++
	if n == nil {
		return false
	}

	if k < n.Key {
		// move left
		n.Left.Search(k)

	} else if k > n.Key {
		// move right
		n.Right.Search(k)
	}
	return true
}
