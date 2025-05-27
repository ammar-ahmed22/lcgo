# 11. Container With Most Water

## Problem

You are given an integer array `height` of length `n`. There are `n` vertical lines drawn such that the two endpoints of the `ith` line are `(i, 0)` and `(i, height[i])`.

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return _the maximum amount of water a container can store_.

**Notice** that you may not slant the container.

**Example 1:**

![](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/07/17/question_11.jpg)

```
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

```

**Example 2:**

```
Input: height = [1,1]
Output: 1

```

**Constraints:**

- `n == height.length`
- `2 <= n <= 105`
- `0 <= height[i] <= 104`

## Approach
We can use two pointers to solve this problem. Obviously, we want to track the maximum area that we see but without checking every possible combination. The optimal way to do this is by having two pointers on each side and always moving the pointers toward the larger bar. This works because we want to maximize width as well as height and consistently moving towards the larger bar while starting at the maximum width will ensure this.

### Complexity
#### Time: `O(n)`
We iterate through the array once

#### Space: `O(1)`
No extra space is created

## Solution

```go
func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	maxArea := 0
	for l < r {
		w := r - l
		h := min(height[l], height[r])
		a := w * h
		maxArea = max(maxArea, a)
		if height[l] == height[r] {
			l++
			r--
		} else if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxArea
}
```
