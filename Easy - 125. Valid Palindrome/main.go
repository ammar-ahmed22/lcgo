package main

import (
	. "valid-palindrome/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
import "strings"
func isAlphaNumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}
func isPalindrome(s string) bool { 
	s = strings.ToLower(s)
	l := 0
	r := len(s) - 1
	for l < r {
		left := s[l]
		right := s[r]
		if !isAlphaNumeric(left) {
			l++
			continue
		}

		if !isAlphaNumeric(right) {
			r--
			continue
		}
		
		if left != right {
			return false
		}
		l++
		r--
	}
	return true
}
// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = bool

var testCases = []*TestCase[ReturnType]{
	NewTestCase(true).WithArgs("A man, a plan, a canal: Panama"),
	NewTestCase(false).WithArgs("race a car"),
	NewTestCase(true).WithArgs(" "),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return isPalindrome(args[0].(string))
	})
}
