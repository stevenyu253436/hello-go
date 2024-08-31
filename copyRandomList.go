package main

import "fmt"

// Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// Map to store the original node to its copy
	nodeMap := make(map[*Node]*Node)

	// Step 1: First pass to create all the nodes with the correct values
	current := head
	for current != nil {
		nodeMap[current] = &Node{Val: current.Val}
		current = current.Next
	}

	// Step 2: Second pass to assign the next and random pointers
	current = head
	for current != nil {
		if current.Next != nil {
			nodeMap[current].Next = nodeMap[current.Next]
		}
		if current.Random != nil {
			nodeMap[current].Random = nodeMap[current.Random]
		}
		current = current.Next
	}

	// Return the new head
	return nodeMap[head]
}

func main() {
	// Example usage:
	node1 := &Node{Val: 7}
	node2 := &Node{Val: 13}
	node3 := &Node{Val: 11}
	node4 := &Node{Val: 10}
	node5 := &Node{Val: 1}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	node2.Random = node1
	node3.Random = node5
	node4.Random = node3
	node5.Random = node1

	// Copy the list
	copiedList := copyRandomList(node1)

	// Print the copied list for verification
	current := copiedList
	for current != nil {
		fmt.Printf("Node Val: %d", current.Val)
		if current.Random != nil {
			fmt.Printf(", Random Val: %d", current.Random.Val)
		} else {
			fmt.Print(", Random is nil")
		}
		fmt.Println()
		current = current.Next
	}
}
