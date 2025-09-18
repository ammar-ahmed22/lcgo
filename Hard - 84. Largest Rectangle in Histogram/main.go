package main

import (
	. "github.com/ammar-ahmed22/lcgo/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
type stack []int

func (s stack) Push(val int) stack {
	return append(s, val)
}

func (s stack) Pop() (stack, int) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) IsEmpty() bool {
	return len(s) == 0
}

func (s stack) Top() int {
	return s[len(s)-1]
}

func largestRectangleArea(heights []int) int {
	maxArea := 0
	st := make(stack, 0)
	for i := range len(heights) + 1 {
		curr := 0
		if i < len(heights) {
			curr = heights[i]
		}

		for !st.IsEmpty() && curr < heights[st.Top()] {
			var hIdx int
			st, hIdx = st.Pop()
			h := heights[hIdx]
			leftIdx := -1
			if !st.IsEmpty() {
				leftIdx = st.Top()
			}
			w := i - leftIdx - 1
			maxArea = max(maxArea, h*w)
		}

		st = st.Push(i)
	}

	return maxArea
}

// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = int

var testCases = []*TestCase[ReturnType]{
	// Add test cases here
	// Using twoSum as an example:
	// NewTestCase([]int{0, 1}).WithArgs([]int{2, 7, 11, 15}, 9),
	// You can also give a name to the test case:
	// NewTestCase([]int{0, 1}).WithArgs([]int{2, 7, 11, 15}, 9).WithName("Example Test Case"),
	// You can also provide a custom comparison function:
	// NewTestCase([]int{0, 1}).WithArgs([]int{2, 7, 11, 15}, 9).WithCompareFn(SliceEqualUnordered)
	NewTestCase(10).WithArgs([]int{2, 1, 5, 6, 2, 3}),
	NewTestCase(4).WithArgs([]int{2, 4}),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return largestRectangleArea(args[0].([]int))
	})
}
