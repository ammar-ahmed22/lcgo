package main

import (
	. "github.com/ammar-ahmed22/lcgo/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	maxArea := 0
	for l < r {
		w := r - l
		h := min(height[l], height[r])
		a := w * h
		maxArea = max(maxArea, a)
		if height[l] == height[r] {
			l++
			r--
		} else if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxArea
}
// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = int

var testCases = []*TestCase[ReturnType]{
	NewTestCase(49).WithArgs([]int{1,8,6,2,5,4,8,3,7}),
	NewTestCase(1).WithArgs([]int{1, 1}),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return maxArea(args[0].([]int))
	})
}
