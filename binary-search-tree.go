// You can edit this code!
// Click here and start typing.
package main

import "fmt"

var count int

// Node represents the components of the binary search tree

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert method will add node to the tree
// node should not already be there in the tree
func (n *Node) Insert(k int) {

	if n.Key < k {
		// move to the right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		// move to the left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	} else {
		fmt.Println("The value already exists.")
	}

}

// Search method will take in a value
// and Return if it exists in the tree

func (n *Node) Search(k int) bool {

	count++

	if n == nil {
		return false
	}

	if n.Key < k {
		// move to the right
		return n.Right.Search(k)

	} else if n.Key > k {
		// move to the left
		return n.Left.Search(k)
	}

	return true

}

func main() {
	tree := &Node{Key: 100}
	fmt.Println(tree)
	tree.Insert(200)
	tree.Insert(300)
	tree.Insert(52)
	tree.Insert(19)
	tree.Insert(76)
	tree.Insert(150)
	tree.Insert(310)
	tree.Insert(7)
	tree.Insert(24)
	tree.Insert(88)
	tree.Insert(276)
	fmt.Println(tree)

	fmt.Println(tree.Search(400))
	fmt.Println(count)
}
