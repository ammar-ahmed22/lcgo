# 704. Binary Search

## Problem

Given an array of integers `nums` which is sorted in ascending order, and an integer `target`, write a function to search `target` in `nums`. If `target` exists, then return its index. Otherwise, return `-1`.

You must write an algorithm with `O(log n)` runtime complexity.

**Example 1:**

```
Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4

```

**Example 2:**

```
Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1

```

**Constraints:**

- `1 <= nums.length <= 104`
- `-104 < nums[i], target < 104`
- All the integers in `nums` are **unique**.
- `nums` is sorted in ascending order.

## Approach
This is a very simple algorithm that everyone should know.

The idea is that, since the values are sorted, we can throw away half off the values on each iteration. We start by initializing two pointers, one on each side. On each iteration, we check the midpoint value of the two pointers. If the mid point equals the target, we return the value. If it doesn't that means it's either on the left side or the right side depending on if the middle value is larger or greater than the target. If the middle value is larger than the target, that means the value is on the left side (smaller side) of the array. So, we can set our right pointer to one less than the middle. If the middle value is smaller than the target, that means the value is on the right side (larger side) of the array. So, we can set our left pointer to one more than the middle. If we complete the iterations, that means the value is not in the array.

## Complexity
### Time: `O(log n)`
Since we're halving the array on each iteration, this is a log based 2 algorithm.

### Space: `O(1)`
No extra space is created.

## Solution

```go
func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		} else if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return -1
}
```
