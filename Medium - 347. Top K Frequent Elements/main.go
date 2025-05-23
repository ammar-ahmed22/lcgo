package main

import (
	. "top-k-frequent-elements/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)
	for _, num := range nums {
		if _, ok := counts[num]; ok {
			counts[num]++
		} else {
			counts[num] = 0
		}
	}

	n := len(nums)
	freq := make([][]int, n)
	for num, f := range counts {
		freq[f] = append(freq[f], num)
	}

	var result []int
	i := n - 1
	for len(result) < k {
		arr := freq[i]
		for _, num := range arr {
			result = append(result, num)
			if len(result) == k {
				return result
			}
		}
		i--
	}

	return result
}

// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = []int

var testCases = []*TestCase[ReturnType]{
	NewTestCase([]int{1, 2}).WithArgs([]int{1, 1, 1, 2, 2, 3}, 2).WithCompareFn(SliceEqualUnordered),
	NewTestCase([]int{1}).WithArgs([]int{1}, 1).WithCompareFn(SliceEqualUnordered),
	NewTestCase([]int{1, 2}).WithArgs([]int{1, 1, 1, 2, 2, 2, 3, 3}, 2).WithCompareFn(SliceEqualUnordered),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return topKFrequent(args[0].([]int), args[1].(int))
	})
}
