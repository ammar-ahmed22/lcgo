package main

import (
	. "github.com/ammar-ahmed22/lcgo/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	curr1, curr2 := l1, l2
	for curr1 != nil && curr2 != nil {
		sum := curr1.Val + curr2.Val
		node := &ListNode{}
		if sum - 10 >= 0 {
			node.Val = sum - 10
			if curr1.Next != nil {
				curr1.Next.Val++
			} else if curr2.Next != nil {
				curr2.Next.Val++
			} else {
				node.Next = &ListNode{
					Val: 1,
				}
			}
		} else {
			node.Val = sum
		}
		tail.Next = node
		curr1 = curr1.Next
		curr2 = curr2.Next
		tail = tail.Next
	}

	for curr1 != nil {
		node := &ListNode{}
		if curr1.Val < 10 {
			node.Val = curr1.Val
		} else {
			node.Val = curr1.Val - 10
			if curr1.Next != nil {
				curr1.Next.Val++
			} else {
				node.Next = &ListNode{Val: 1}
			}
		}
		tail.Next = node
		tail = tail.Next
		curr1 = curr1.Next
	}

	for curr2 != nil {
		node := &ListNode{}
		if curr2.Val < 10 {
			node.Val = curr2.Val
		} else {
			node.Val = curr2.Val - 10
			if curr2.Next != nil {
				curr2.Next.Val++
			} else {
				node.Next = &ListNode{Val: 1}
			}
		}
		tail.Next = node
		tail = tail.Next
		curr2 = curr2.Next
	}

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
	NewTestCase(ListFromNums(7, 0, 8).String()).WithArgs(ListFromNums(2, 4, 3), ListFromNums(5, 6, 4)),
	NewTestCase(ListFromNums(1, 3, 2, 1).String()).WithArgs(ListFromNums(9, 8, 7), ListFromNums(2, 4, 4)),
	NewTestCase(ListFromNums(2, 4, 3, 4).String()).WithArgs(ListFromNums(1, 2, 3, 4), ListFromNums(1, 2)),
	NewTestCase(ListFromNums(8, 9, 9, 9, 0, 0, 0, 1).String()).WithArgs(ListFromNums(9, 9, 9, 9, 9, 9, 9), ListFromNums(9, 9, 9, 9)),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		result := addTwoNumbers(args[0].(*ListNode),args[1].(*ListNode))
		if result == nil {
			return "nil"
		}
		return result.String()
	})
}
