package main

import (
	. "github.com/ammar-ahmed22/lcgo/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
func lengthOfLongestSubstring(s string) int {
	count := make(map[byte]int)
	for _, c := range s {
		count[byte(c)] = 0
	}

	var (
		l, r, res int
	)
	for r < len(s) {
		rCount := count[s[r]]
		if rCount == 0 {
			count[s[r]]++
			r++
		} else {
			count[s[l]]--
			l++
		}
		res = max(res, r-l)
	}
	return res
}

// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = int

var testCases = []*TestCase[ReturnType]{
	NewTestCase(3).WithArgs("abcabcbbb"), // { "a": 1, "b": 1, "c": 1 }
	NewTestCase(1).WithArgs("bbbbb"),
	NewTestCase(3).WithArgs("pwwkew"),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return lengthOfLongestSubstring(args[0].(string))
	})
}
