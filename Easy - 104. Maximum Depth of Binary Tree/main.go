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
func dfs(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}

	return max(dfs(root.Right, depth + 1), dfs(root.Left, depth + 1))
}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(dfs(root.Right, 1), dfs(root.Left, 1))
}

// <-- DO NOT REMOVE: PROBLEM END -->

func main() {
	case1 := TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	Test("Case 1", 3, maxDepth(&case1))

	case2 := TreeNode{Val: 1, Right: &TreeNode{ Val: 2 }}
	Test("Case 2", 2, maxDepth(&case2))
}
