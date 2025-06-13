package main

import (
	. "find-minimum-in-rotated-sorted-array/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
func findMin(nums []int) int {
	l, r := 0, len(nums) - 1
	for l <= r {
		m := (l + r) / 2
		if (m == 0 || nums[m - 1] > nums[m]) && (m == len(nums) - 1 || nums[m + 1] > nums[m]) {
			return nums[m]
		} else if nums[m] < nums[r] {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return nums[l]
}
// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = int

var testCases = []*TestCase[ReturnType]{
	NewTestCase(1).WithArgs([]int{3, 4, 5, 1, 2}),
	NewTestCase(0).WithArgs([]int{4, 5, 6, 7, 0, 1, 2}),
	NewTestCase(11).WithArgs([]int{11, 13, 15, 17}),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return findMin(args[0].([]int))
	})
}
