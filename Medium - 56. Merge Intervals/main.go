package main

import (
	. "github.com/ammar-ahmed22/lcgo/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		curr := result[len(result)-1]
		next := intervals[i]

		if curr[1] >= next[0] {
			// merge
			result = result[:len(result)-1] // pop off the top
			merged := []int{min(curr[0], next[0]), max(curr[1], next[1])}
			result = append(result, merged)
		} else {
			result = append(result, next)
		}
	}

	return result
}

// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = [][]int

func compare(a, b [][]int) bool {
	sort.Slice(a, func(i, j int) bool {
		return a[i][0] < a[j][0]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i][0] < b[j][0]
	})

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		currA := a[i]
		currB := b[i]
		if len(currA) != 2 || len(currB) != 2 {
			return false
		}
		if currA[0] != currB[0] {
			return false
		}
		if currA[1] != currB[1] {
			return false
		}
	}
	return true
}

var testCases = []*TestCase[ReturnType]{
	// Add test cases here
	// Using twoSum as an example:
	// NewTestCase([]int{0, 1}).WithArgs([]int{2, 7, 11, 15}, 9),
	// You can also give a name to the test case:
	// NewTestCase([]int{0, 1}).WithArgs([]int{2, 7, 11, 15}, 9).WithName("Example Test Case"),
	// You can also provide a custom comparison function:
	// NewTestCase([]int{0, 1}).WithArgs([]int{2, 7, 11, 15}, 9).WithCompareFn(SliceEqualUnordered)
	NewTestCase([][]int{{1, 6}, {8, 10}, {15, 18}}).WithArgs([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}).WithCompareFn(compare),
	NewTestCase([][]int{{1, 5}}).WithArgs([][]int{{1, 4}, {4, 5}}).WithCompareFn(compare),
	NewTestCase([][]int{{1, 7}}).WithArgs([][]int{{4, 7}, {1, 4}}).WithCompareFn(compare),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return merge(args[0].([][]int))
	})
}
