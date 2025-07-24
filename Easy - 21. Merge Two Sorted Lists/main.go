package main

import (
	. "merge-two-sorted-lists/testutils"
	"reflect"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func listFromSlice(nums []int) *ListNode {
	if len(nums) == 0 {
		return &ListNode{}
	}
	list := ListNode{
		Val: nums[0],
	}
	curr := &list
	for i := 1; i < len(nums); i++ {
		curr.Next = &ListNode{
			Val: nums[i],
		}
		curr = curr.Next
	}
	return &list
}

// <-- DO NOT REMOVE: PROBLEM START -->
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	curr1, curr2 := list1, list2
	for curr1 != nil && curr2 != nil {
		var v int
		if curr1.Val <= curr2.Val {
			v = curr1.Val
			curr1 = curr1.Next
		} else {
			v = curr2.Val
			curr2 = curr2.Next
		}
		tail.Next = &ListNode{ Val: v }
		tail = tail.Next
	}
	if curr2 != nil {
		tail.Next = curr2
	}
	if curr1 != nil {
		tail.Next = curr1
	}
	return dummy.Next
}

// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = *ListNode

func compare(a *ListNode, b *ListNode) bool {
	return reflect.DeepEqual(*a, *b)
}

var testCases = []*TestCase[ReturnType]{
	NewTestCase(listFromSlice([]int{1, 1, 2, 3, 4, 4})).WithArgs(listFromSlice([]int{1, 2, 4}), listFromSlice([]int{1, 3, 4})).WithCompareFn(compare),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return mergeTwoLists(args[0].(*ListNode), args[1].(*ListNode))
	})
}
