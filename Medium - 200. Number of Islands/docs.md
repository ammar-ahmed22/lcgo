# 200. Number of Islands

## Problem

Given an `m x n` 2D binary grid `grid` which represents a map of `'1'` s (land) and `'0'` s (water), return _the number of islands_.

An **island** is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

**Example 1:**

```
Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1

```

**Example 2:**

```
Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3

```

**Constraints:**

- `m == grid.length`
- `n == grid[i].length`
- `1 <= m, n <= 300`
- `grid[i][j]` is `'0'` or `'1'`.

## Approach
To solve this, we can use a depth-first search approach.

The idea is that we iterate through the grid and when we see a `1` we mark it as a 0, then recursively check all of it's neighbours and do the same until the entire island has been marked as `0`. Then, we can increment our islands counter.

This ensures that we only process every element once because they will be marked as zeros if it's part of a previously processed island.


## Complexity
### Time: `O(m * n)`
We iterate through the entire grid once.

### Space: `O(1)`
No extra space is created.

## Solution

```go
func dfs(grid[][]byte, r, c, rows, cols int) {
	directions := [][]int{
		{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0},
	}

	if (r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] == '0') {
		return
	}

	grid[r][c] = '0'
	for _, dir := range directions {
		dx, dy := dir[0], dir[1]
		dfs(grid, r + dy, c + dx, rows, cols)
	}

}
func numIslands(grid [][]byte) int {
	ROWS, COLS := len(grid), len(grid[0])

	islands := 0

	for r, row := range grid {
		for c, cell := range row {
			if cell == '1' {
				// dfs
				dfs(grid, r, c, ROWS, COLS)
				islands++
			}
		} 
	}

	return islands
}

```
