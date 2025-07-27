package main

import (
	. "github.com/ammar-ahmed22/lcgo/testutils"
)

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// <-- DO NOT REMOVE: PROBLEM START -->
func copyRandomList(head *Node) *Node {
	// Create the interweaved list, oldNode -> newNode
	dummy := &Node{}
	curr := head
	tail := dummy
	for curr != nil {
		curr.Next = &Node{
			Val:  curr.Val,
			Next: curr.Next,
		}
		tail.Next = curr
		tail = tail.Next.Next
		curr = curr.Next.Next
	}

	// Create the random connections
	curr = dummy.Next
	for curr != nil {
		if curr.Random != nil {
			curr.Next.Random = curr.Random.Next
		} else {
			curr.Next.Random = nil
		}

		curr = curr.Next.Next
	}

	// Remove the old nodes
	curr = dummy
	for curr != nil && curr.Next != nil {
		if curr.Next != nil {
			tmp := curr.Next // Save the old node
			curr.Next = curr.Next.Next
			tmp.Next = tmp.Next.Next // Re-construct the old list
		} else {
			curr.Next = nil
		}
		curr = curr.Next
	}

	return dummy.Next
}

// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = *Node

var testCases = []*TestCase[ReturnType]{
	// Add test cases here
	// Using twoSum as an example:
	// NewTestCase([]int{0, 1}).WithArgs([]int{2, 7, 11, 15}, 9),
	// You can also give a name to the test case:
	// NewTestCase([]int{0, 1}).WithArgs([]int{2, 7, 11, 15}, 9).WithName("Example Test Case"),
	// You can also provide a custom comparison function:
	// NewTestCase([]int{0, 1}).WithArgs([]int{2, 7, 11, 15}, 9).WithCompareFn(SliceEqualUnordered)
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return copyRandomList(args[0].(*Node))
	})
}
