package main

import "fmt"

/**
 * Definition for a Node.
 */
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	// Initialize the first node of the current level
	curr := root

	// While there are nodes at the current level
	for curr != nil {
		// Create a dummy node to form the linked list of the next level
		dummy := &Node{}
		tail := dummy

		// Iterate over all nodes in the current level
		for curr != nil {
			if curr.Left != nil {
				tail.Next = curr.Left
				tail = tail.Next
			}
			if curr.Right != nil {
				tail.Next = curr.Right
				tail = tail.Next
			}
			// Move to the next node in the current level
			curr = curr.Next
		}

		// Move to the next level
		curr = dummy.Next
	}

	return root
}

// Helper function to print the tree level by level
func printTree(root *Node) {
	if root == nil {
		return
	}

	curr := root
	for curr != nil {
		fmt.Print(curr.Val, " -> ")
		if curr.Next != nil {
			fmt.Print(curr.Next.Val)
		} else {
			fmt.Print("nil")
		}
		fmt.Println()
		if curr.Left != nil {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}
}

func main() {
	// Create a test tree: [1,2,3,4,5,null,7]
	root := &Node{Val: 1}
	root.Left = &Node{Val: 2}
	root.Right = &Node{Val: 3}
	root.Left.Left = &Node{Val: 4}
	root.Left.Right = &Node{Val: 5}
	root.Right.Right = &Node{Val: 7}

	// Connect the nodes
	connect(root)

	// Print the tree
	printTree(root)
}
