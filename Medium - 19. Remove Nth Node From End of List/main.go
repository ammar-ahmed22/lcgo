package main

import (
	. "github.com/ammar-ahmed22/lcgo/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	right := head
	for n > 0 {
		right = right.Next
		n--
	}

	dummy := &ListNode{}
	dummy.Next = head
	left := dummy
	for right != nil {
		left = left.Next
		right = right.Next
	}
	left.Next = left.Next.Next
	return dummy.Next
}
// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = string 

var testCases = []*TestCase[ReturnType]{
	// Add test cases here
	// Using twoSum as an example:
	// NewTestCase([]int{0, 1}).WithArgs([]int{2, 7, 11, 15}, 9),
	// You can also give a name to the test case:
	// NewTestCase([]int{0, 1}).WithArgs([]int{2, 7, 11, 15}, 9).WithName("Example Test Case"),
	// You can also provide a custom comparison function:
	// NewTestCase([]int{0, 1}).WithArgs([]int{2, 7, 11, 15}, 9).WithCompareFn(SliceEqualUnordered)
	NewTestCase(ListFromSlice([]int{1, 2, 3, 5}).String()).WithArgs(ListFromSlice([]int{1, 2, 3, 4, 5}), 2),
	NewTestCase("nil").WithArgs(ListFromSlice([]int{1}), 1),
	NewTestCase(ListFromSlice([]int{1}).String()).WithArgs(ListFromSlice([]int{1, 2}), 1),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		result := removeNthFromEnd(args[0].(*ListNode),args[1].(int))
		if result == nil {
			return "nil"
		}
		return result.String()
	})
}
