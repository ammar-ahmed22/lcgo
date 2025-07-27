package main

import (
	. "github.com/ammar-ahmed22/lcgo/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
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
// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = bool

var testCases = []*TestCase[ReturnType]{
	NewTestCase(true).WithArgs([][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}}, 3),
	NewTestCase(false).WithArgs([][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}}, 13),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return searchMatrix(args[0].([][]int),args[1].(int))
	})
}
