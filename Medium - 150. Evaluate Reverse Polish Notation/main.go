package main

import (
	. "evaluate-reverse-polish-notation/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
import "strconv"

func eval(a, b int, op string) int {
	if op == "+" {
		return a + b
	} else if op == "*" {
		return a * b
	} else if op == "-" {
		return b - a
	} else {
		return b / a
	}
}

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, s := range tokens {
		if s == "+" || s == "*" || s == "-" || s == "/" {
			// evaluate
			var a, b int
			a, stack = stack[len(stack)-1], stack[:len(stack)-1]
			b, stack = stack[len(stack)-1], stack[:len(stack)-1]
			stack = append(stack, eval(a, b, s))
		} else {
			num, _ := strconv.Atoi(s)
			stack = append(stack, num)
		}
	}

	return stack[len(stack)-1]
}

// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = int

var testCases = []*TestCase[ReturnType]{
	NewTestCase(9).WithArgs([]string{"2", "1", "+", "3", "*"}),
	NewTestCase(6).WithArgs([]string{"4", "13", "5", "/", "+"}),
	NewTestCase(22).WithArgs([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return evalRPN(args[0].([]string))
	})
}
