package main

import (
	. "github.com/ammar-ahmed22/lcgo/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
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

// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = int

var testCases = []*TestCase[ReturnType]{
	// Add test cases here
	// Using twoSum as an example:
	// NewTestCase([]int{0, 1}).WithArgs([]int{2, 7, 11, 15}, 9),
	// You can also give a name to the test case:
	// NewTestCase([]int{0, 1}).WithArgs([]int{2, 7, 11, 15}, 9).WithName("Example Test Case"),
	// You can also provide a custom comparison function:
	// NewTestCase([]int{0, 1}).WithArgs([]int{2, 7, 11, 15}, 9).WithCompareFn(SliceEqualUnordered)
	NewTestCase(1).WithArgs([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}),
	NewTestCase(3).WithArgs([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return numIslands(args[0].([][]byte))
	})
}
