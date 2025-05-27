# 217. Contains Duplicate

## Problem

Given an integer array `nums`, return `true` if any value appears **at least twice** in the array, and return `false` if every element is distinct.

**Example 1:**

**Input:** nums = \[1,2,3,1\]

**Output:** true

**Explanation:**

The element 1 occurs at the indices 0 and 3.

**Example 2:**

**Input:** nums = \[1,2,3,4\]

**Output:** false

**Explanation:**

All elements are distinct.

**Example 3:**

**Input:** nums = \[1,1,1,3,3,4,3,2,4,2\]

**Output:** true

**Constraints:**

- `1 <= nums.length <= 105`
- `-109 <= nums[i] <= 109`

## Approach 
Simply use a hash map to keep track of which numbers we've already seen. If we've already seen a number, return `true`. Otherwise add to the hashmap and continue iterating. If we reach the end, there are no duplicates.

## Complexity
### Time: `O(n)`
We iterate through the array once

### Space: `O(n)`
We create a hash map that can, at most, contain all the numbers

## Solution

```go
func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	for _, num := range nums {
		if _, exists := seen[num]; exists {
			return true
		}
		seen[num] = true
	}
	return false
}
```
