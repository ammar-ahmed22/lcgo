package main

import (
	. "github.com/ammar-ahmed22/lcgo/testutils"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// <-- DO NOT REMOVE: PROBLEM START -->
func invertTree(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}

	temp := root.Right
	root.Right = invertTree(root.Left)
	root.Left = invertTree(temp)
	return root
}

// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = *TreeNode

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
	// RunTestCases(testCases, func(args ...any) ReturnType {
	// 	return /** * Definition for a binary tree node. * type TreeNode struct { * Val int * Left *TreeNode * Right *TreeNode * } */ func invertTree(args[0].(*TreeNode))
	// })
}
