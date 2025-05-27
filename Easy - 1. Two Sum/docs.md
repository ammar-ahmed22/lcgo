# 1. Two Sum

## Problem

Given an array of integers `nums` and an integer `target`, return _indices of the two numbers such that they add up to `target`_.

You may assume that each input would have **_exactly_ one solution**, and you may not use the _same_ element twice.

You can return the answer in any order.

**Example 1:**

```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

```

**Example 2:**

```
Input: nums = [3,2,4], target = 6
Output: [1,2]

```

**Example 3:**

```
Input: nums = [3,3], target = 6
Output: [0,1]

```

**Constraints:**

- `2 <= nums.length <= 104`
- `-109 <= nums[i] <= 109`
- `-109 <= target <= 109`
- **Only one valid answer exists.**

**Follow-up:** Can you come up with an algorithm that is less than `O(n2)` time complexity?

## Approach 
We can use a hashmap to solve this problem. The idea is that we want to check if there is another number in the array that can be added to the current number to reach the target. Therefore, we can iterate over the numbers and check if the complement (`target - number`) exists in the hashmap, if it does, we return the two indices. Otherwise, we add the number to the hashmap
with the index as its value. Since it's guranteed that there is a solution, we don't need to worry about what to return if the iteration completes (it never will).

## Complexity
### Time: `O(n)`
We only iterate over the input array once

### Space: `O(n)`
We create a hash map that could, at most, contain all the numbers


## Solution

```go
func twoSum(nums []int, target int) []int { 
	numMap := make(map[int]int)
	
	for i, num := range nums {
		if idx, ok := numMap[target - num]; ok {
			return []int{idx, i}
		}
		numMap[num] = i
	}
	return []int{}
}
```
