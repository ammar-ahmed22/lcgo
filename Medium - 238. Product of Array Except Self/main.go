package main

import (
	. "product-of-array-except-self/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
func productExceptSelf(nums []int) []int { 
	n := len(nums)
	prefix := make([]int, n)
	prefix[0] = 1
	for i := 1; i < n; i++ {
		prefix[i] = nums[i - 1] * prefix[i - 1]
	}

	suffix := make([]int, n)
	suffix[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		suffix[i] = nums[i + 1] * suffix[i + 1]
	}

	answer := make([]int, n)
	for i := range n {
		answer[i] = prefix[i] * suffix[i]
	}

	return answer
}
// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = []int

var testCases = []*TestCase[ReturnType]{
	NewTestCase([]int{24, 12, 8, 6}).WithArgs([]int{1, 2, 3, 4}),
	NewTestCase([]int{0, 0, 9, 0, 0}).WithArgs([]int{-1, 1, 0, -3, 3}),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return productExceptSelf(args[0].([]int))
	})
}
