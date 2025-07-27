package main

import (
	. "github.com/ammar-ahmed22/lcgo/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	var stack [][]int

	for i, temp := range temperatures {
		for len(stack) > 0 && temp > stack[len(stack) - 1][0] {
			var top []int
			top, stack = stack[len(stack) - 1], stack[:len(stack) - 1]
			d := i - top[1]
			ans[top[1]] = d
		}
		stack = append(stack, []int{ temp, i })
	}

	return ans 
}

// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = []int

var testCases = []*TestCase[ReturnType]{
	NewTestCase([]int{1, 1, 4, 2, 1, 1, 0, 0}).WithArgs([]int{73, 74, 75, 71, 69, 72, 76, 73}),
	NewTestCase([]int{1, 1, 1, 0}).WithArgs([]int{30, 40, 50, 60}),
	NewTestCase([]int{1, 1, 0}).WithArgs([]int{30, 60, 90}),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return dailyTemperatures(args[0].([]int))
	})
}
