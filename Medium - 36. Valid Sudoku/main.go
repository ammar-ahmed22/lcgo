package main

import (
	. "valid-sudoku/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
func hasDuplicate(values []byte) bool {
	seen := make(map[byte]bool)
	for _, val := range values {
		if val == '.' {
			continue
		}
		if _, exists := seen[val]; exists {
			return true
		}
		seen[val] = true
	}
	return false
}
func isValidSudoku(board [][]byte) bool {
	n := len(board)
	// Checking rows
	for _, row := range board {
		if hasDuplicate(row) {
			return false
		}
	}

	// Checking columns
	for i := range n {
		col := make([]byte, n)
		for j := range n {
			col[j] = board[j][i]
		}
		if hasDuplicate(col) {
			return false
		}
	}

	// Checking sub-boxes
	for brow := range 3 {
		for bcol := range 3 {
			var values []byte
			for row := range 3 {
				for col := range 3 {
					values = append(values, board[3 * brow + row][3 * bcol + col])
				}
			}
			if hasDuplicate(values) {
				return false
			}
		}
	}

	return true
}

// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = bool

var test1 = [][]byte{
	{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
}

var test2 = [][]byte{
	{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
}

var testCases = []*TestCase[ReturnType]{
	NewTestCase(true).WithArgs(test1),
	NewTestCase(false).WithArgs(test2),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return isValidSudoku(args[0].([][]byte))
	})
}
