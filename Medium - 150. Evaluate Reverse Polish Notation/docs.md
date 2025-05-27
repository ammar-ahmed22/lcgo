# 150. Evaluate Reverse Polish Notation

## Problem

You are given an array of strings `tokens` that represents an arithmetic expression in a [Reverse Polish Notation](http://en.wikipedia.org/wiki/Reverse_Polish_notation).

Evaluate the expression. Return _an integer that represents the value of the expression_.

**Note** that:

- The valid operators are `'+'`, `'-'`, `'*'`, and `'/'`.
- Each operand may be an integer or another expression.
- The division between two integers always **truncates toward zero**.
- There will not be any division by zero.
- The input represents a valid arithmetic expression in a reverse polish notation.
- The answer and all the intermediate calculations can be represented in a **32-bit** integer.

**Example 1:**

```
Input: tokens = ["2","1","+","3","*"]
Output: 9
Explanation: ((2 + 1) * 3) = 9

```

**Example 2:**

```
Input: tokens = ["4","13","5","/","+"]
Output: 6
Explanation: (4 + (13 / 5)) = 6

```

**Example 3:**

```
Input: tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
Output: 22
Explanation: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22

```

**Constraints:**

- `1 <= tokens.length <= 104`
- `tokens[i]` is either an operator: `"+"`, `"-"`, `"*"`, or `"/"`, or an integer in the range `[-200, 200]`.

## Approach
This is a very interesting problem. We can use a stack to solve this problem.

The general idea is that whenever we see an operator, we want to use it to evaluate the last two elements and add the result back to our evaluation line.

Let's take the first example: `[2, 1, +, 3, *]`. We'll go iteration by iteration, displaying our stack.
- `i = 0`, `value = 2`, `stack = [2]`
- `i = 1`, `value = 1`, `stack = [2, 1]`
- `i = 2`, `value = +`, `stack = [3]`
    + Since we hit an operator, we pop the last two values from the stack `(2, 1)`, evaluate them (`2 + 1 = 3`) and add back to the stack (`[3]`)
- `i = 3`, `value = 3`, `stack = [3, 3]`
- `i = 4`, `value = *`, `stack = [9]`
    + We hit another operator so we do the same thing, pop the last two values `(3, 3)`, evaluate them (`3 * 3 = 9`), add back to the stack (`[9]`)

Once we complete the iterations, we can simply return the last value in the stack.

For my specific solution, I wrote a helper function to do the evaluation. An important note is on the order of evaluation specifically for subtraction and division. We want to do `b - a` and `b / a`, where `a` is the first value popped and `b` is the second value popped.


## Complexity
### Time: `O(n)`
We iterate through the tokens only once.

### Space: `O(1)`
No extra space is created.

## Solution

```go
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

```
