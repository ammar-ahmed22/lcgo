package main

import (
	. "reorder-list/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
func reorderList(head *ListNode) { }
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
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		reorderList(args[0].(*ListNode))
		return args[0].(*ListNode).toString()
	})
}
