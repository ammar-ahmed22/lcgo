package main

import (
	. "github.com/ammar-ahmed22/lcgo/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
import "strings"
func backtrack(stack []string, n, open, closed int, result *[]string) {
	if open == closed && closed == n {
		*result = append(*result, strings.Join(stack, ""))
		return
	}

	if open < n {
		stack = append(stack, "(")
		backtrack(stack, n, open+1, closed, result)
		stack = stack[:len(stack)-1]
	}

	if closed < open {
		stack = append(stack, ")")
		backtrack(stack, n, open, closed+1, result)
		stack = stack[:len(stack)-1]
	}
}
func generateParenthesis(n int) []string {
	var (
		stack []string
		result []string
	)
	backtrack(stack, n, 0, 0, &result)
	return result 
}

// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = []string

var testCases = []*TestCase[ReturnType]{
	NewTestCase([]string{"((()))", "(()())", "(())()", "()(())", "()()()"}).WithArgs(3).WithCompareFn(SliceEqualUnordered),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return generateParenthesis(args[0].(int))
	})
}
