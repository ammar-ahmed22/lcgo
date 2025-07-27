package main

import (
	. "github.com/ammar-ahmed22/lcgo/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		} else if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return -1
}
// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = int

var testCases = []*TestCase[ReturnType]{
	NewTestCase(4).WithArgs([]int{-1,0,3,5,9,12}, 9),
	NewTestCase(-1).WithArgs([]int{-1,0,3,5,9,12}, 12),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return search(args[0].([]int),args[1].(int))
	})
}
