# 74. Search a 2D Matrix

## Problem

You are given an `m x n` integer matrix `matrix` with the following two properties:

- Each row is sorted in non-decreasing order.
- The first integer of each row is greater than the last integer of the previous row.

Given an integer `target`, return `true` _if_ `target` _is in_ `matrix` _or_ `false` _otherwise_.

You must write a solution in `O(log(m * n))` time complexity.

**Example 1:**

![](https://assets.leetcode.com/uploads/2020/10/05/mat.jpg)

```
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
Output: true

```

**Example 2:**

![](https://assets.leetcode.com/uploads/2020/10/05/mat2.jpg)

```
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
Output: false

```

**Constraints:**

- `m == matrix.length`
- `n == matrix[i].length`
- `1 <= m, n <= 100`
- `-104 <= matrix[i][j], target <= 104`

## Approach
We can use a nested binary search for this. Since the rows are sorted by the first element. We run an initial binary search on the rows to find which row would contain our value. Then we run a binary search on that row to see if it has the value.

To do the binary search on the rows, we can check if the value is contained inside the first and last values in the row. If it is, that means we've found our row. Otherwise, we continue with binary search using the first element as our check.

One important note is to remember to return `false` if the inner binary search does not find the answer because it should be in that row and it's not that means it's not in the matrix. Not doing this will lead to an infinite loop.

### Complexity
#### Time: `O(log(m * n))`
We do two nested binary searches. One on the matrix and the other on the row.

#### Space: `O(1)`
No extra space was created.

## Solution

```go
func searchMatrix(matrix [][]int, target int) bool {
	top := 0
	bot := len(matrix) - 1
	for top <= bot {
		c := (top + bot) / 2
		center := matrix[c]
		if target >= center[0] && target <= center[len(center) - 1] {
			// binary search on center
			l := 0
			r := len(center) - 1
			for l <= r {
				m := (l + r) / 2
				if target == center[m] {
					return true
				} else if target < center[m] {
					r = m - 1
				} else {
					l = m + 1
				}
			}
			return false
		} else if target < center[0] {
			bot = c - 1
		} else {
			top = c + 1
		}
	}
	return false
}
```
