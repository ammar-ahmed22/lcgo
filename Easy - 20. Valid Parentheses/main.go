package main

import (
	. "valid-parentheses/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
func isValid(s string) bool {
	p := map[rune]rune{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	var stack []rune
	for _, c := range s {
		closing, isOpening := p[c]
		if isOpening {
			stack = append(stack, closing)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack) - 1]
			if top != c {
				return false
			}
			stack = stack[:len(stack) - 1]
		}
	}

	return len(stack) == 0
}
// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = bool

var testCases = []*TestCase[ReturnType]{
	NewTestCase(true).WithArgs("()"),
	NewTestCase(true).WithArgs("()[]{}"),
	NewTestCase(false).WithArgs("(]"),
	NewTestCase(true).WithArgs("([])"),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return isValid(args[0].(string))
	})
}
