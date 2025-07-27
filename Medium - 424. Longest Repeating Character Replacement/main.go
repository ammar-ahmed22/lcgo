package main

import (
	. "github.com/ammar-ahmed22/lcgo/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
func characterReplacement(s string, k int) int {
	var freq [26]int
	left, maxF, maxLen := 0, 0, 0

	for right := range len(s) {
		freq[s[right]-'A']++
		maxF = max(maxF, freq[s[right]-'A'])

		for (right-left+1)-maxF > k {
			freq[s[left]-'A']--
			left++
		}

		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
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
	NewTestCase(4).WithArgs("ABAB", 2),
	NewTestCase(4).WithArgs("AABABBA", 1),
	NewTestCase(4).WithArgs("AAAA", 0),
	NewTestCase(2).WithArgs("ABCDE", 1),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return characterReplacement(args[0].(string), args[1].(int))
	})
}
