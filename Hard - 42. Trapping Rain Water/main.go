package main

import (
	. "github.com/ammar-ahmed22/lcgo/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
func trap(height []int) int {
	l, r, ans := 0, len(height) - 1, 0
	maxL, maxR := height[l], height[r]
	for l < r {
		if maxL <= maxR {
			l++
			maxL = max(maxL, height[l])
			ans += maxL - height[l]
		} else {
			r--
			maxR = max(maxR, height[r])
			ans += maxR - height[r]
		}
	}
	return ans
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
	NewTestCase(6).WithArgs([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}),
	NewTestCase(9).WithArgs([]int{4, 2, 0, 3, 2, 5}),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return trap(args[0].([]int))
	})
}
