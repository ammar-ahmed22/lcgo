package main

import (
	. "github.com/ammar-ahmed22/lcgo/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
func findDuplicate(nums []int) int {
	s, f := 0, 0
	for {
		s = nums[s]
		f = nums[nums[f]]
		if s == f {
			break
		}
	}

	// Intersection point found, now find the entrance to the cycle
	s2 := 0
	for {
		s2 = nums[s2]
		s = nums[s]
		if s == s2 {
			return s
		}
	}
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
	NewTestCase(2).WithArgs([]int{1, 3, 4, 2, 2}),
	NewTestCase(3).WithArgs([]int{3, 1, 3, 4, 2}),
	NewTestCase(3).WithArgs([]int{3, 3, 3, 3, 3}),
	NewTestCase(9).WithArgs([]int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1}),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return findDuplicate(args[0].([]int))
	})
}
