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

## Solution Notes
- We can use a hash map to solve this
- Create a hashmap with the key as the number and boolean for the value
- Iterate over the numbers
    + If the number exists in the hash map -> return `true` (we've already seen this number, duplicate)
    + If it doesn't add it to the hash map
    + If the loop complets, all numbers are unique -> return `false`

### Complexity
#### Time: `O(n)`
- We iterate through the array once

#### Space: `O(n)`
- We create a hash map that can, at most, contain all the numbers

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
