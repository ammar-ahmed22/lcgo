package main

import (
	. "github.com/ammar-ahmed22/lcgo/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	// Populate frequency map with s1 values
	var freqMap [26]int
	for _, char := range s1 {
		freqMap[char-'a']++
	}

	l := 0
	r := len(s1)

	// Populate frequency map with chars in s2 from l to r
	for i := range r {
		// Entering "window", subtract from map
		freqMap[s2[i]-'a']--
	}

	for r < len(s2) {
		// Check if freqMap is all zeros
		hasNonZero := false
		for _, f := range freqMap {
			if f != 0 {
				hasNonZero = true
				break
			}
		}
		if !hasNonZero {
			return true
		}

		// Leaving window
		freqMap[s2[l]-'a']++
		// Entering window
		freqMap[s2[r]-'a']--
		l++
		r++
	}

	for _, f := range freqMap {
		if f != 0 {
			return false
		}
	}
	return true
}

// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = bool

var testCases = []*TestCase[ReturnType]{
	// Add test cases here
	// Using twoSum as an example:
	// NewTestCase([]int{0, 1}).WithArgs([]int{2, 7, 11, 15}, 9),
	// You can also give a name to the test case:
	// NewTestCase([]int{0, 1}).WithArgs([]int{2, 7, 11, 15}, 9).WithName("Example Test Case"),
	// You can also provide a custom comparison function:
	// NewTestCase([]int{0, 1}).WithArgs([]int{2, 7, 11, 15}, 9).WithCompareFn(SliceEqualUnordered)
	NewTestCase(true).WithArgs("ab", "eidbaooo"),
	NewTestCase(false).WithArgs("ab", "eidboaoo"),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return checkInclusion(args[0].(string), args[1].(string))
	})
}
