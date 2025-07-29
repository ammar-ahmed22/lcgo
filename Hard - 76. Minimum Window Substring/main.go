package main

import (
	. "github.com/ammar-ahmed22/lcgo/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
	ans := ""
	freq := make(map[byte]int)
	for _, char := range t {
		freq[byte(char)]++
	}

	l, r := 0, len(t)-1

	for i := range len(t) {
		if _, ok := freq[byte(s[i])]; ok {
			freq[byte(s[i])]--
		}
	}

	minWindowLen := len(t)

	for r < len(s) {
		isValid := true
		for _, val := range freq {
			if val > 0 {
				isValid = false
				break
			}
		}

		if isValid {
			windowLen := (r - l + 1)
			if ans == "" || windowLen < len(ans) {
				ans = s[l : r+1]
			}

			if _, ok := freq[byte(s[l])]; ok {
				freq[byte(s[l])]++
			}
			l++

			if windowLen == minWindowLen {
				r++
				if r < len(s) {
					if _, ok := freq[byte(s[r])]; ok {
						freq[byte(s[r])]--
					}
				}
			}
		} else {
			r++
			if r < len(s) {
				if _, ok := freq[byte(s[r])]; ok {
					freq[byte(s[r])]--
				}
			}
		}
	}

	return ans
}

// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = string

var testCases = []*TestCase[ReturnType]{
	// Add test cases here
	// Using twoSum as an example:
	// NewTestCase([]int{0, 1}).WithArgs([]int{2, 7, 11, 15}, 9),
	// You can also give a name to the test case:
	// NewTestCase([]int{0, 1}).WithArgs([]int{2, 7, 11, 15}, 9).WithName("Example Test Case"),
	// You can also provide a custom comparison function:
	// NewTestCase([]int{0, 1}).WithArgs([]int{2, 7, 11, 15}, 9).WithCompareFn(SliceEqualUnordered)
	NewTestCase("BANC").WithArgs("ADOBECODEBANC", "ABC"),
	NewTestCase("a").WithArgs("a", "a"),
	NewTestCase("").WithArgs("a", "aa"),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return minWindow(args[0].(string), args[1].(string))
	})
}
