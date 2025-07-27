package main

import (
	. "github.com/ammar-ahmed22/lcgo/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
func twoSum(nums []int, target int) []int { 
	numMap := make(map[int]int)
	
	for i, num := range nums {
		if idx, ok := numMap[target - num]; ok {
			return []int{idx, i}
		}
		numMap[num] = i
	}
	return []int{}
}
// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = []int

var testCases = []*TestCase[ReturnType]{
	NewTestCase([]int{0, 1}).WithArgs([]int{2, 7, 11, 15}, 9).WithCompareFn(SliceEqualUnordered),
	NewTestCase([]int{1, 2}).WithArgs([]int{3, 2, 4}, 6).WithCompareFn(SliceEqualUnordered),
	NewTestCase([]int{0, 1}).WithArgs([]int{3, 3}, 6).WithCompareFn(SliceEqualUnordered),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return twoSum(args[0].([]int),args[1].(int))
	})
}
