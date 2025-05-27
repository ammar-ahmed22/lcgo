package main

import (
	. "best-time-to-buy-and-sell-stock/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	l := 0
	r := 1
	maxP := 0
	for r < len(prices) {
		p := prices[r] - prices[l]
		if p > 0 {
			maxP = max(maxP, p)
		} else {
			l = r
		}
		r++
	}
	return maxP
}
// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = int

var testCases = []*TestCase[ReturnType]{
	NewTestCase(5).WithArgs([]int{7, 1, 5, 3, 6, 4}),
	NewTestCase(0).WithArgs([]int{7, 6, 4, 3, 1}),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return maxProfit(args[0].([]int))
	})
}
