package main

import (
	. "koko-eating-bananas/testutils"
	"math"
)

// <-- DO NOT REMOVE: PROBLEM START -->
func minEatingSpeed(piles []int, h int) int {
	var (l, r int)
	for _, p := range piles {
		r = max(r, p)
	}

	var res int
	for l <= r {
		k := (l + r) / 2
		var tot float64
		for _, p := range piles {
			tot += math.Ceil(float64(p) / float64(k))
		}
		if tot <= float64(h) {
			res = k
			r = k - 1
		} else {
			l = k + 1
		}
	}
	return res
}
// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = int

var testCases = []*TestCase[ReturnType]{
	NewTestCase(4).WithArgs([]int{3, 6, 7, 11}, 8),
	NewTestCase(30).WithArgs([]int{30, 11, 23, 4, 20}, 5),
	NewTestCase(23).WithArgs([]int{30, 11, 23, 4, 20}, 6),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return minEatingSpeed(args[0].([]int), args[1].(int))
	})
}
