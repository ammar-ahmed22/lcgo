# 104. Maximum Depth of Binary Tree

## Problem

Given the `root` of a binary tree, return _its maximum depth_.

A binary tree's **maximum depth** is the number of nodes along the longest path from the root node down to the farthest leaf node.

**Example 1:**

![](https://assets.leetcode.com/uploads/2020/11/26/tmp-tree.jpg)

```
Input: root = [3,9,20,null,null,15,7]
Output: 3

```

**Example 2:**

```
Input: root = [1,null,2]
Output: 2

```

**Constraints:**

- The number of nodes in the tree is in the range `[0, 104]`.
- `-100 <= Node.val <= 100`

## Approach
To solve this problem we can use a recursive approach using depth first search. 

We can write a separate recursive function that takes in a node and a depth parameter. If the node is nil, it returns the depth (base case). Otherwise, it returns the maximum of the recursive call on the right and left with the depth being incremented.

In the main function, we check the base case if the root is null (return 0), otherwise, we return the max of our dfs function on the right and left nodes using a starting depth of 1.

## Complexity
### Time: `O(n)`
Where `n` is the number of nodes

### Space: `O(n)`
Recursive calls

## Solution

```go
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

```
