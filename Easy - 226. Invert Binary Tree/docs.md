# 226. Invert Binary Tree

## Problem

Given the `root` of a binary tree, invert the tree, and return _its root_.

**Example 1:**

![](https://assets.leetcode.com/uploads/2021/03/14/invert1-tree.jpg)

```
Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]

```

**Example 2:**

![](https://assets.leetcode.com/uploads/2021/03/14/invert2-tree.jpg)

```
Input: root = [2,1,3]
Output: [2,3,1]

```

**Example 3:**

```
Input: root = []
Output: []

```

**Constraints:**

- The number of nodes in the tree is in the range `[0, 100]`.
- `-100 <= Node.val <= 100`

## Approach
To solve this problem, we can use a recursive approach.

Our base case will be when the node has neither left or right children or the node itself is null, return here.

Otherwise, we set the left side to the right side and the right side to left side and then call the function on both left and right

## Complexity
### Time: `O(n)`
Iterate through the whole tree

### Space: `O(n)`
For the recursive calls.

## Solution

```go
func invertTree(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}

	temp := root.Right
	root.Right = invertTree(root.Left)
	root.Left = invertTree(temp)
	return root
}

```
