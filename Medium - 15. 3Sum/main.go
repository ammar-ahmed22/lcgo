package main

import (
	. "github.com/ammar-ahmed22/lcgo/testutils"
	"sort"
)

// <-- DO NOT REMOVE: PROBLEM START -->
import "slices"
func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	var ans [][]int
	for i, curr := range nums {
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		target := curr * -1
		l := i + 1
		r := len(nums) - 1
		for l < r {
			pot := nums[l] + nums[r]
			if pot == target {
				ans = append(ans, []int{curr, nums[l], nums[r]})
				for l < r && nums[l] == nums[l + 1] {
					l++
				}
				for l < r && nums[r] == nums[r - 1] {
					r--
				}
				l++
				r--
			} else if pot < target {
				l++
			} else {
				r--
			}
		}
	}
	return ans
}
// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = [][]int

func customCompareFn(a [][]int, b [][]int) bool {
	for i := range a {
		slices.Sort(a[i])
	}
	for i := range b {
		slices.Sort(b[i])
	}
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
		arrA := a[i]
		arrB := b[i]
		if len(arrA) != len(arrB) {
			return false
		}
		for j := range arrA {
			if arrA[j] != arrB[j] {
				return false
			}
		}
	}

	return true
}

var testCases = []*TestCase[ReturnType]{
	NewTestCase([][]int{{-1, -1, 2}, {-1, 0, 1}}).WithArgs([]int{-1, 0, 1, 2, -1, 4}).WithCompareFn(customCompareFn),
	NewTestCase([][]int{}).WithArgs([]int{0, 1, 1}).WithCompareFn(customCompareFn),
	NewTestCase([][]int{{0, 0, 0}}).WithArgs([]int{0, 0, 0}).WithCompareFn(customCompareFn),
	NewTestCase([][]int{{0, 0, 0}}).WithArgs([]int{0, 0, 0, 0}).WithCompareFn(customCompareFn),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return threeSum(args[0].([]int))
	})
}
