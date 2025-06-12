# 739. Daily Temperatures

## Problem

Given an array of integers `temperatures` represents the daily temperatures, return _an array_ `answer` _such that_ `answer[i]` _is the number of days you have to wait after the_ `ith` _day to get a warmer temperature_. If there is no future day for which this is possible, keep `answer[i] == 0` instead.

**Example 1:**

```
Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]

```

**Example 2:**

```
Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]

```

**Example 3:**

```
Input: temperatures = [30,60,90]
Output: [1,1,0]

```

**Constraints:**

- `1 <= temperatures.length <= 105`
- `30 <= temperatures[i] <= 100`

## Approach
To solve this, we can use a stack.

The general premise is that we want to keep track of temperatures until we find a temperature that is larger.

We start off by initializing our answers array to an array of zeroes. This is because we won't create the answer's in order, we'll insert them as we find the larger temperatures.

We also intialize our stack to take arrays as it's value's so we can store both the temperature's and it's index.

Now, we can move on to the iteration. For each number, we check if the current value is greater than the top of the stack. If it is, that means we have found an answer for the number at the top of the stack. We pop the value off the stack, calculate the distance using the index that's stored and insert into the corresponding spot in the answer array. However, since it's possible that the current value will give us an answer for multiple values, we want to keep checking if the current value is greater than the top of the stack and continuing to do the same. If there are no more answers, we can add the current value to the stack.

Let's go through an example with a subset of **Example 1**: `[73, 74, 75, 71, 69, 72, 76]`

To start, intialize answers and stack:
- stack: `[]`, ans: `[0, 0, 0, 0, 0, 0, 0]`
- `i = 0`, `curr = 73`, `stack = []`
    + stack is empty 
    + add `curr` to stack: `[[73, 0]]` 
- `i = 1`, `curr = 74`, `stack = [[73, 0]]`
    + `curr > stack[top]`, pop from stack: `[73, 0]`, `stack = []`, calculate distance (`d = 1 - 0 = 1`), insert at `ans[0]`, `ans = [1, 0, 0, 0, 0, 0, 0]` 
    + stack is empty
    + add `curr` to stack: `[[74, 1]]`
- `i = 2`, `curr = 75`, `stack = [[74, 1]]`
    + `curr > stack[top]`, pop from stack: `[74, 1]`, `stack = []`, calculate distance (`d = 2 - 1 = 1`), insert at `ans[1]`, `ans = [1, 1, 0, 0, 0, 0, 0]`
    + stack is empty
    + add `curr` to stack: `[[75, 2]]` 
- `i = 3`, `curr = 71`, `stack = [[75, 2]]`
    + `curr` is not greater than `stack[top]`
    + add `curr` to stack: `[[75, 2], [71, 3]]`
- `i = 4`, `curr = 69`, `stack = [[75, 2], [71, 3]]`
    + `curr` is not greater than `stack[top]`
    + add `curr` to stack: `[[75, 2], [71, 3], [69, 4]]`
- `i = 5`, `curr = 72`, `stack = [[75, 2], [71, 3], [69, 4]]`
    + `curr > stack[top]`, pop from stack: `[69, 4]`, `stack = [[75, 2], [71, 3]]`, calculate distance (`d = 5 - 4 = 1`), insert at `ans[4]`, `ans = [1, 1, 0, 0, 1, 0, 0]`
    + `curr > stack[top]`, pop from stack: `[71, 3]`, `stack = [[75, 2]]`, calculate distance (`d = 5 - 3 = 2`), insert at `ans[3]`, `ans = [1, 1, 0, 2, 1, 0, 0]`
    + `curr` is not greater than `stack[top]`, add `curr` to stack: `[[75, 2], [72, 5]]`
- `i = 6`, `curr = 76`, `stack = [[75, 2], [72, 5]]`
    + `curr > stack[top]`, pop from stack: `[72, 5]`, `stack = [[75, 2]]`, calculate distance (`d = 6 - 5 = 1`), insert at `ans[5]`, `ans = [1, 1, 0, 2, 1, 1, 0]`
    + `curr > stack[top]`, pop from stack: `[75, 2]`, `stack = []`, calculate distance (`d = 6 - 2 = 4`), insert at `ans[2]`, `ans = [1, 1, 4, 2, 1, 1, 0]`
    + stack is empty, add `curr` to stack: `[[76, 6]]`
- iteration complete, answerr is created

## Complexity
### Time: `O(n)`
We iterate over the values only once. The "inner loop" will never be long enough to give `O(n^2)`.

### Space: `O(n)`
The only extra space we create is the stack.

## Solution

```go
func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	var stack [][]int

	for i, temp := range temperatures {
		for len(stack) > 0 && temp > stack[len(stack) - 1][0] {
			var top []int
			top, stack = stack[len(stack) - 1], stack[:len(stack) - 1]
			d := i - top[1]
			ans[top[1]] = d
		}
		stack = append(stack, []int{ temp, i })
	}

	return ans 
}

```
