package main

import (
	. "github.com/ammar-ahmed22/lcgo/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
func reverseList(head *ListNode) *ListNode { 
	curr := head
	var prev *ListNode
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	return prev
}
// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = *ListNode

var original = &ListNode{ Val: 1, Next: &ListNode{ Val: 2, Next: &ListNode{ Val: 3, Next: &ListNode{ Val: 4, Next: &ListNode{ Val: 5 }}}}}
var reversed = &ListNode{ Val: 5, Next: &ListNode{ Val: 4, Next: &ListNode{ Val: 3, Next: &ListNode{ Val: 2, Next: &ListNode{ Val: 1 }}}}}

var testCases = []*TestCase[ReturnType]{
	NewTestCase(reversed).WithArgs(original),
	NewTestCase(&ListNode{ Val: 2, Next: &ListNode{ Val: 1 }}).WithArgs(&ListNode{ Val: 1, Next: &ListNode{ Val: 2 }}),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return reverseList(args[0].(*ListNode))
	})
}
