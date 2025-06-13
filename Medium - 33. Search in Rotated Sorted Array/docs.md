# 33. Search in Rotated Sorted Array

## Problem

There is an integer array `nums` sorted in ascending order (with **distinct** values).

Prior to being passed to your function, `nums` is **possibly rotated** at an unknown pivot index `k` ( `1 <= k < nums.length`) such that the resulting array is `[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]` ( **0-indexed**). For example, `[0,1,2,4,5,6,7]` might be rotated at pivot index `3` and become `[4,5,6,7,0,1,2]`.

Given the array `nums` **after** the possible rotation and an integer `target`, return _the index of_ `target` _if it is in_ `nums` _, or_ `-1` _if it is not in_ `nums`.

You must write an algorithm with `O(log n)` runtime complexity.

**Example 1:**

```
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4

```

**Example 2:**

```
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1

```

**Example 3:**

```
Input: nums = [1], target = 0
Output: -1

```

**Constraints:**

- `1 <= nums.length <= 5000`
- `-104 <= nums[i] <= 104`
- All values of `nums` are **unique**.
- `nums` is an ascending array that is possibly rotated.
- `-104 <= target <= 104`

## Approach
To solve this problem efficiently in O(log n) time, we can employ the algorithm used in [153. Find Minimum In Rotated Sorted Array](https://ammarahmed.ca/medium/find-minimum-in-rotated-sorted-array). This is because, in order to do a full binary search on a rotated sorted array, all we need to do is find the rotation point as we did in [153. Find Minimum In Rotated Sorted Array](https://ammarahmed.ca/medium/find-minimum-in-rotated-sorted-array) and then figure out which side of the rotated array the target is in and conduct a simple binary search only on that part of the array.

Let's take modified Example 1: `nums = [4, 5, 6, 7, 0, 1, 2], target = 6`. First we find the rotation point, we can see it occurs at `k = 4, nums[k] = 0`, from this we can deduce that the target should be in the left side of the array because it is less than the value before the rotation point and greater than the left-most value. Therefore, we conduct a simple binary search from `l = 0`, to `r = k` and we'll find our value there. 

The same premise can be applied to the other side as well.

To solve this, I wrote a few helper functions:
- `rotPoint(nums []int) int`: Finds the index of the rotation point of a rotated sorted array as we did in [153. Find Minimum In Rotated Sorted Array](https://ammarahmed.ca/medium/find-minimum-in-rotated-sorted-array)
- `binary(nums []int, target, l, r, int) int`: Conducts binary search with specificed starting parameters.

Using these two functions and the logic described above, we can solve the problem.

## Complexity
### Time: `O(log n)`
We find the rotation point to start which is `O(log n)` and then we conduct binary search on the portion of the rotated array that the value is in which is also `O(log n)`. Thus, `O(log n) + O(log n) = O(2 log n) = O(log n)`.

### Space: `O(1)`
No extra space is created.

## Solution

```go
// O(log n)
func rotPoint(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		if (m == 0 || nums[m-1] > nums[m]) && (m == len(nums)-1 || nums[m+1] > nums[m]) {
			return m
		} else if nums[m] < nums[r] {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}

// O(log n)
func binary(nums []int, target, l, r int) int {
	for l <= r {
		m := (l + r) / 2
		if target == nums[m] {
			return m
		} else if target < nums[m] {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}
func search(nums []int, target int) int {
	// Find rotation point, k
	k := rotPoint(nums)

	l, r := 0, len(nums)-1
	if target >= nums[k] && target <= nums[r] {
		// Search from l = k -> r
		l = k
		return binary(nums, target, l, r)
	}

	if target >= nums[l] && (k == 0 || target <= nums[k - 1]) {
		// Search from l -> r = k
		r = k
		return binary(nums, target, l, r)
	}

	return -1
}

```
