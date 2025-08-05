# 42. Trapping Rain Water

## Problem

Given `n` non-negative integers representing an elevation map where the width of each bar is `1`, compute how much water it can trap after raining.

**Example 1:**

![](https://assets.leetcode.com/uploads/2018/10/22/rainwatertrap.png)

```
Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.

```

**Example 2:**

```
Input: height = [4,2,0,3,2,5]
Output: 9

```

**Constraints:**

- `n == height.length`
- `1 <= n <= 2 * 104`
- `0 <= height[i] <= 105`

## Approach
To start, we need to understand how to calculate the amount of water that can be trapped at any given point. The amount of water that can be trapped on any given block is bounded by the smaller wall. As in, the smaller max value on the left or right will bound the amount of water that can be trapped at that point. 

For example, if we have `[2, 1, 0, 2, 3]` and we want to calculate the amount of water trapped at `height[2] = 0`, the max value on the left is `2` and the max value on the right is `3`. Therefore, the bounding value is `2` and we subtract the height at `height[2] = 0` from that, so, `2 - 0 = 2`. Therefore, the amount of water that can be trapped at `i = 2` is `2`.

### Linear Space Approach
From the above, we can easily devise a linear space and time complexity solution. We simply calculat the max left and max right values at each index in the height array in their own separate arrays. Then, we iterate over the heights and check the min value between the two max left and max right arrays and calculate the amount of water trapped. If it's greater than zero, we add it to our total.

However, we can do better than linear space.

### Constant Space Approach
We can do away with tracking all of the max left and right values by using a two-pointer approach. We can have pointers on the left and right side of the heights alongside tracking the max left and max right values. We only move the pointer that has the smaller max value and whenever we move the pointer, we update our answer and max value. This works because when `maxL <= maxR`, we know that any water at the left pointer is bounded by `maxL` (since there's already a taller wall `maxR` on the right). We move the pointer with the smaller max value, update the max, and add the trapped water.

## Complexity
### Time: `O(n)`
We iterate over the input once.

### Space: `O(1)`
No extra space is created.

## Solution

```go
func trap(height []int) int {
	l, r, ans := 0, len(height) - 1, 0
	maxL, maxR := height[l], height[r]
	for l < r {
		if maxL <= maxR {
			l++
			maxL = max(maxL, height[l])
			ans += maxL - height[l]
		} else {
			r--
			maxR = max(maxR, height[r])
			ans += maxR - height[r]
		}
	}
	return ans
}

```
