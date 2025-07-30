package main

import (
	"container/list"

	. "github.com/ammar-ahmed22/lcgo/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
func maxSlidingWindow(nums []int, k int) []int {
	ans := []int{}
	dq := list.New()
	l, r := 0, 0

	for r < len(nums) {
		for dq.Len() > 0 && nums[dq.Front().Value.(int)] < nums[r] {
			dq.Remove(dq.Front())
		}
		dq.PushFront(r)

		if l > dq.Back().Value.(int) {
			dq.Remove(dq.Back())
		}

		if (r + 1) >= k {
			ans = append(ans, nums[dq.Back().Value.(int)])
			l++
		}
		r++
	}

	return ans
}

// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = []int

var testCases = []*TestCase[ReturnType]{
	// Add test cases here
	// Using twoSum as an example:
	// NewTestCase([]int{0, 1}).WithArgs([]int{2, 7, 11, 15}, 9),
	// You can also give a name to the test case:
	// NewTestCase([]int{0, 1}).WithArgs([]int{2, 7, 11, 15}, 9).WithName("Example Test Case"),
	// You can also provide a custom comparison function:
	// NewTestCase([]int{0, 1}).WithArgs([]int{2, 7, 11, 15}, 9).WithCompareFn(SliceEqualUnordered)
	NewTestCase([]int{3, 3, 5, 5, 6, 7}).WithArgs([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return maxSlidingWindow(args[0].([]int), args[1].(int))
	})
}
