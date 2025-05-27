# 20. Valid Parentheses

## Problem

Given a string `s` containing just the characters `'('`, `')'`, `'{'`, `'}'`, `'['` and `']'`, determine if the input string is valid.

An input string is valid if:

1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.
3. Every close bracket has a corresponding open bracket of the same type.

**Example 1:**

**Input:** s = "()"

**Output:** true

**Example 2:**

**Input:** s = "()\[\]{}"

**Output:** true

**Example 3:**

**Input:** s = "(\]"

**Output:** false

**Example 4:**

**Input:** s = "(\[\])"

**Output:** true

**Constraints:**

- `1 <= s.length <= 104`
- `s` consists of parentheses only `'()[]{}'`.

## Approach
We can use a stack to solve this problem.

We iterate through the string, whenever we see an opening bracket, we push the corresponding closing bracket to the stack. When we see a closing bracket, we check if the top of the stack is the same bracket, if it's not, we can return `false` early. Otherwise, we pop from the stack and continue iterating. If we reach the end and the stack is empty, the parentheses are valid. If the stack is not empty, that means something was not closed and we can return `false`.

## Complexity
### Time: `O(n)`
We iterate over the input once.

### Space: `O(n)`
We create a stack that can possibly contain all the elements of the input.

## Solution

```go
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
```
