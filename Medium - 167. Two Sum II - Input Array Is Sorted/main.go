package main

import (
	. "github.com/ammar-ahmed22/lcgo/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1
	for l < r {
		curr := numbers[l] + numbers[r]
		if curr == target {
			return []int{l + 1, r + 1}
		} else if curr < target {
			l++
		} else {
			r--
		}
	}
	return []int{}
}
// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = []int

var testCases = []*TestCase[ReturnType]{
	NewTestCase([]int{1,2}).WithArgs([]int{2, 7, 11, 15}, 9).WithCompareFn(SliceEqualUnordered),
	NewTestCase([]int{1, 3}).WithArgs([]int{2, 3, 4}, 6).WithCompareFn(SliceEqualUnordered),
	NewTestCase([]int{1, 2}).WithArgs([]int{-1, 0}, -1).WithCompareFn(SliceEqualUnordered),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return twoSum(args[0].([]int),args[1].(int))
	})
}
