# 84. Largest Rectangle in Histogram

## Problem

Given an array of integers `heights` representing the histogram's bar height where the width of each bar is `1`, return _the area of the largest rectangle in the histogram_.

**Example 1:**

![](https://assets.leetcode.com/uploads/2021/01/04/histogram.jpg)

```
Input: heights = [2,1,5,6,2,3]
Output: 10
Explanation: The above is a histogram where width of each bar is 1.
The largest rectangle is shown in the red area, which has an area = 10 units.

```

**Example 2:**

![](https://assets.leetcode.com/uploads/2021/01/04/histogram-1.jpg)

```
Input: heights = [2,4]
Output: 4

```

**Constraints:**

- `1 <= heights.length <= 105`
- `0 <= heights[i] <= 104`

## Approach

We'll use a stack.

The idea is to maintain a monotonic increasing stack of indices into `heights`. Each index on the stack represents a bar whose rectangle hasn't been fully "closed" yet because we haven't seen a shorter bar to its right. As soon as we do see a shorter bar, we can finalize areas for all taller bars that were waiting on the stack.

Key points:
- Store indices, not heights. Indices let us compute widths precisely.
- Keep the stack increasing by height: `heights[stack[0]] <= heights[stack[1]] <= ...`.
- When the current height is smaller than the height at the top index on the stack, we pop and compute an area for that popped height, because the current index is the first smaller bar to its right.
- The next element now on top of the stack (after popping) is the first smaller bar to the left. That gives us the left boundary. If the stack becomes empty, it means there was no smaller bar to the left, so the left boundary is before the beginning.

Algorithm:
1. Initialize `maxArea = 0` and an empty stack of indices `st`.
2. Iterate `i` from `0` to `len(heights)` inclusive. Treat `i == len(heights)` as a sentinel height `0` to flush the stack at the end.
3. Let `curr = heights[i]` if `i < n`, otherwise `curr = 0`.
4. While `st` is not empty and `curr < heights[st.top]`:
   - Pop `hIdx := st.pop()` and let `h := heights[hIdx]`.
   - The new top of the stack (if any) is the index of the first smaller bar to the left; call it `leftIdx`. If the stack is empty, `leftIdx = -1`.
   - The first smaller bar to the right is at `i` (the current index).
   - Width = `i - leftIdx - 1`.
   - Area = `h * width`. Update `maxArea`.
5. Push `i` onto the stack.
6. After the loop, `maxArea` is the answer.

Quick walkthrough on `[2, 1, 5, 6, 2, 3]`:
- Push 0 (height 2), then see 1 < 2 → pop 2: left boundary is empty → width `= 1 - (-1) - 1 = 1`, area `= 2`.
- Push 1 (height 1).
- Push 2 (5), push 3 (6).
- See height 2 < 6 → pop 6 at idx 3: left smaller is idx 2 (5) still on stack → width `= 4 - 2 - 1 = 1`, area `= 6`.
- Still 2 < 5 → pop 5 at idx 2: left smaller is idx 1 (height 1) → width `= 4 - 1 - 1 = 2`, area `= 10` (current max).
- Push 4 (2), push 5 (3).
- Sentinel 0 at end flushes remaining bars similarly.

## Complexity
### Time: `O(n)`
Each index is pushed and popped at most once.

### Space: `O(n)`
The stack can hold up to `n` indices in the increasing run.


## Solution

```go
type stack []int

func (s stack) Push(val int) stack {
	return append(s, val)
}

func (s stack) Pop() (stack, int) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) IsEmpty() bool {
	return len(s) == 0
}

func (s stack) Top() int {
	return s[len(s)-1]
}

func largestRectangleArea(heights []int) int {
	maxArea := 0
	st := make(stack, 0)
	for i := range len(heights) + 1 {
		curr := 0
		if i < len(heights) {
			curr = heights[i]
		}

		for !st.IsEmpty() && curr < heights[st.Top()] {
			var hIdx int
			st, hIdx = st.Pop()
			h := heights[hIdx]
			leftIdx := -1
			if !st.IsEmpty() {
				leftIdx = st.Top()
			}
			w := i - leftIdx - 1
			maxArea = max(maxArea, h*w)
		}

		st = st.Push(i)
	}

	return maxArea
}

```
