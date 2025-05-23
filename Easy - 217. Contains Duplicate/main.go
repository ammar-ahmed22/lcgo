package main

import (
	. "contains-duplicate/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	for _, num := range nums {
		if _, exists := seen[num]; exists {
			return true
		}
		seen[num] = true
	}
	return false
}
// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = bool

var testCases = []*TestCase[ReturnType]{
	NewTestCase(true).WithArgs([]int{1, 2, 3, 1}),
	NewTestCase(false).WithArgs([]int{1, 2, 3, 4}),
	NewTestCase(true).WithArgs([]int{1,1,1,3,3,4,3,2,4,2}),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return containsDuplicate(args[0].([]int))
	})
}
