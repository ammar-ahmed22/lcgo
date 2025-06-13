# 153. Find Minimum in Rotated Sorted Array

## Problem

Suppose an array of length `n` sorted in ascending order is **rotated** between `1` and `n` times. For example, the array `nums = [0,1,2,4,5,6,7]` might become:

- `[4,5,6,7,0,1,2]` if it was rotated `4` times.
- `[0,1,2,4,5,6,7]` if it was rotated `7` times.

Notice that **rotating** an array `[a[0], a[1], a[2], ..., a[n-1]]` 1 time results in the array `[a[n-1], a[0], a[1], a[2], ..., a[n-2]]`.

Given the sorted rotated array `nums` of **unique** elements, return _the minimum element of this array_.

You must write an algorithm that runs in `O(log n) time`.

**Example 1:**

```
Input: nums = [3,4,5,1,2]
Output: 1
Explanation: The original array was [1,2,3,4,5] rotated 3 times.

```

**Example 2:**

```
Input: nums = [4,5,6,7,0,1,2]
Output: 0
Explanation: The original array was [0,1,2,4,5,6,7] and it was rotated 4 times.

```

**Example 3:**

```
Input: nums = [11,13,15,17]
Output: 11
Explanation: The original array was [11,13,15,17] and it was rotated 4 times.

```

**Constraints:**

- `n == nums.length`
- `1 <= n <= 5000`
- `-5000 <= nums[i] <= 5000`
- All the integers of `nums` are **unique**.
- `nums` is sorted and rotated between `1` and `n` times.

## Approach
To solve this problem, we can employ the premise that we are essentially looking for the rotation point of the array. The rotation point would be the minimum of the array. 

To do this, we can employ binary search to fit within the time complexity constraint of the problem. 

With binary search, the first thing we want to do is figure out our return case (i.e. the case where we have found the rotation point). The rotation point in a rotated sorted array would mean that the number before it and after it would be larger than it because if they are both not larger, then the rotation point is somewhere else.

For example, taking the array: `[4, 5, 6, 7, 0, 1, 2]`, if we look at the number `6`, the number before it is smaller than it and the number after is larger (so it's in order), thus it's not the rotation point. Looking at `0`, we can see that the number before (`6`) and the number after (`1`) are both larger than it, so that is the rotation point.

Therefore, in our binary search we do this check (while also checking that there is a number before and after, if there isn't skip that check).

If the midpoint of our binary search is not the rotation point, then we want to move our search to the part of the array that will contain it. For this, we check if the midpoint is less than the right pointer, if it is, that means the right side of our array is trending upwards so our rotation point is in the left side of the array, otherwise it's in the right side.

If after the iterations, we have not found the rotation point through the midpoint, then that means it's at the left pointer (because we went past it with our divisions) so we return the value at the left pointer.

## Complexity
### Time: `O(log n)`
We do a binary search

### Space: `O(1)`
No extra space created.

## Solution

```go
func findMin(nums []int) int {
	l, r := 0, len(nums) - 1
	for l <= r {
		m := (l + r) / 2
		if (m == 0 || nums[m - 1] > nums[m]) && (m == len(nums) - 1 || nums[m + 1] > nums[m]) {
			return nums[m]
		} else if nums[m] < nums[r] {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return nums[l]
}
```
