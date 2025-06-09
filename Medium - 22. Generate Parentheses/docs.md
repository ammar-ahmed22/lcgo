# 22. Generate Parentheses

## Problem

Given `n` pairs of parentheses, write a function to _generate all combinations of well-formed parentheses_.

**Example 1:**

```
Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]

```

**Example 2:**

```
Input: n = 1
Output: ["()"]

```

**Constraints:**

- `1 <= n <= 8`

## Approach
We can solve this problem recursively by using the idea of open and closed parentheses. We can see that the number of open and closed parentheses in a well-formed parentheses string must be equal to `n`. As in, `open = closed = n`. However, when generating the string, the number of closed parentheses must not exceed the number of open parentheses. For example, if we have `"(())"`, the number of `open` and `closed` parentheses is equal to 2. Therefore, we cannot add a closed parentheses here because closed must always be less than or equal to open.

Using this principle, we have the following rules for our recursive function:
- **Base case**: If `open == closed == n`, we are finished with the formation and can add the parentheses string to our result
- `open < n`: If `open < n`, then we can add more open parentheses
- `closed < open`: If `closed < open`, we can add more closing parentheses

Using these rules alongside a stack that is passed with each recursive function, we can form our parentheses.

## Complexity
### Time: `O(4^n)`
Since we have 2 possibilites at each step and there are `2n` steps, we have `2 ^ 2n = 4^n`

### Space: `O(n)`
Since it's recursive, we'll be creating `2n` calls on the stack at a time.

## Solution

```go
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

```
