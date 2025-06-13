package main

import (
	. "search-in-rotated-sorted-array/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
// O(log n)
func rotPoint(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		if (m == 0 || nums[m-1] > nums[m]) && (m == len(nums)-1 || nums[m+1] > nums[m]) {
			return m
		} else if nums[m] < nums[r] {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}

// O(log n)
func binary(nums []int, target, l, r int) int {
	for l <= r {
		m := (l + r) / 2
		if target == nums[m] {
			return m
		} else if target < nums[m] {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}
func search(nums []int, target int) int {
	// Find rotation point, k
	k := rotPoint(nums)

	l, r := 0, len(nums)-1
	if target >= nums[k] && target <= nums[r] {
		// Search from l = k -> r
		l = k
		return binary(nums, target, l, r)
	}

	if target >= nums[l] && (k == 0 || target <= nums[k - 1]) {
		// Search from l -> r = k
		r = k
		return binary(nums, target, l, r)
	}

	return -1
}

// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = int

var testCases = []*TestCase[ReturnType]{
	NewTestCase(4).WithArgs([]int{4, 5, 6, 7, 0, 1, 2}, 0),
	NewTestCase(-1).WithArgs([]int{4, 5, 6, 7, 0, 1, 2}, 3),
	NewTestCase(-1).WithArgs([]int{1}, 0),
	NewTestCase(1).WithArgs([]int{3, 5, 1}, 5),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return search(args[0].([]int), args[1].(int))
	})
}
