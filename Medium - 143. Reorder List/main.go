package main

import (
	. "reorder-list/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
func reorderList(head *ListNode) {
	// Find middle of list with slow and fast pointers
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Reverse the second half of the list
	curr := slow.Next
	slow.Next = nil
	reversed := slow.Next
	for curr != nil {
		tmp := curr.Next
		curr.Next = reversed
		reversed = curr
		curr = tmp
	}

	// Merge the first and second half (reversed) of the list
	first, second := head, reversed
	for second != nil {
		tmp1, tmp2 := first.Next, second.Next
		first.Next = second
		second.Next = tmp1
		first = tmp1
		second = tmp2
	}
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
	NewTestCase(ListFromSlice([]int{1, 4, 2, 3}).String()).WithArgs(ListFromSlice([]int{1, 2, 3, 4})),
	NewTestCase(ListFromSlice([]int{1, 5, 2, 4, 3}).String()).WithArgs(ListFromSlice([]int{1, 2, 3, 4, 5})),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		reorderList(args[0].(*ListNode))
		return args[0].(*ListNode).String()
	})
}
