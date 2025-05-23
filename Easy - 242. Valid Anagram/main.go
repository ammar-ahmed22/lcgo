package main

import (
	. "valid-anagram/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
func isAnagram(s string, t string) bool { 
	if len(s) != len(t) {
		return false
	}

	var freqS [26]int
	var freqT [26]int
	n := len(s)
	for i := range n {
		freqS[s[i] - 'a']++
		freqT[t[i] - 'a']++
	}

	for i := range 26 {
		if freqS[i] != freqT[i] {
			return false
		}
	}

	return true
}
// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = bool

var testCases = []*TestCase[ReturnType]{
	NewTestCase(true).WithArgs("anagram", "nagaram"),
	NewTestCase(false).WithArgs("rat", "car"),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return isAnagram(args[0].(string),args[1].(string))
	})
}
