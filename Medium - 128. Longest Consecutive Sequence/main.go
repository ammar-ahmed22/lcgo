package main

import (
	. "github.com/ammar-ahmed22/lcgo/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
func longestConsecutive(nums []int) int {
	// Edge case check
	if len(nums) == 0 {
		return 0
	}

	// Creating hash map
	freq := make(map[int]bool)
	for _, num := range nums {
		if _, exists := freq[num]; !exists {
			freq[num] = true
		}
	}

	// Intialize to 1 because we already handled the case if its zero
	ans := 1
	for _, num := range nums {
		// If there is no previous value, start of a sequence
		if _, prevExists := freq[num - 1]; !prevExists {
			if !freq[num] { // Check if the current value is marked as false (we've already processed this sequence start)
				continue
			}
			freq[num] = false // Mark the sequence start as processed

			// Find the length of the sequence
			count := 0
			_, hasNext := freq[num + count]
			for hasNext {
				count++
				_, hasNext = freq[num + count]
			}
			// Update answer
			ans = max(ans, count)
		}
	}

	return ans
}

// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = int

var testCases = []*TestCase[ReturnType]{
	NewTestCase(4).WithArgs([]int{100, 4, 200, 1, 3, 2}),
	NewTestCase(9).WithArgs([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}),
	NewTestCase(3).WithArgs([]int{1, 0, 1, 2}),
	NewTestCase(1).WithArgs([]int{0}),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return longestConsecutive(args[0].([]int))
	})
}
