# 287. Find the Duplicate Number

## Problem

Given an array of integers `nums` containing `n + 1` integers where each integer is in the range `[1, n]` inclusive.

There is only **one repeated number** in `nums`, return _this repeated number_.

You must solve the problem **without** modifying the array `nums` and using only constant extra space.

**Example 1:**

```
Input: nums = [1,3,4,2,2]
Output: 2

```

**Example 2:**

```
Input: nums = [3,1,3,4,2]
Output: 3

```

**Example 3:**

```
Input: nums = [3,3,3,3,3]
Output: 3
```

**Constraints:**

- `1 <= n <= 105`
- `nums.length == n + 1`
- `1 <= nums[i] <= n`
- All the integers in `nums` appear only **once** except for **precisely one integer** which appears **two or more** times.

**Follow up:**

- How can we prove that at least one duplicate number must exist in `nums`?
- Can you solve the problem in linear runtime complexity?

## Approach
While it may be a little bit unintuitive, this can actually be treated as a linked list problem. The problem states that the numbers are bounded between `[1, n]` and the length of the list is `n+1`. Therefore, we can treat the indices as the values and the values as the next pointer.

For example, taking the input: `[3, 1, 3, 4, 2]`, we have the indices `[0, 1, 2, 3, 4]`. At the `0th` index, we have the value `3`. Therefore, `0 -> 3`. At index `3`, we have `4`, so, `0 -> 3 -> 4`. At index `4` we have `2`, so, `0 -> 3 -> 4 -> 2`. At index `2` we have `3` again so we have a cycle in the linked list:

```
0 -> 3 -> 4 -> 2
     ^         |
     |         |
     ----------
```

From this, we can see in order to solve the problem, we simply need to find the start of the linked list cycle to find the duplicate number in `O(n)` time with `O(1)` space.

To find the start of a linked list cycle, we can use Floyd's algorithm. Find if a linked list cycle exists at all is trivial, we simply iterate with fast and slow pointers and if they intersect, that means we have found a cycle. However, that could be at any point in the cycle so it only tells us if a cycle exists or not. To find the start of the cycle, after we find the intersection point, we initialize a new pointer at the start of the list and iterate both pointers together until they intersect again, this will be the start of our cycle.

## Complexity
### Time: `O(n)`
We iterate over the list once twice.

### Space: `O(1)`
No extra space is created.

## Solution

```go
func findDuplicate(nums []int) int {
	s, f := 0, 0
	for {
		s = nums[s]
		f = nums[nums[f]]
		if s == f {
			break
		}
	}

	// Intersection point found, now find the entrance to the cycle
	s2 := 0
	for {
		s2 = nums[s2]
		s = nums[s]
		if s == s2 {
			return s
		}
	}
}

```
