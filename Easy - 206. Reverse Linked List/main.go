package main

import (
	"fmt"
	. "reverse-linked-list/testutils"
	"strings"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func (l *ListNode) toString() string {
	var builder strings.Builder
	curr := l
	for curr != nil {
		builder.WriteString(fmt.Sprintf("%d -> ", curr.Val))
		curr = curr.Next
	}
	builder.WriteString("nil")
	return builder.String()
}

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
