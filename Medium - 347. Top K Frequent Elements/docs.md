# 347. Top K Frequent Elements

## Problem

Given an integer array `nums` and an integer `k`, return _the_ `k` _most frequent elements_. You may return the answer in **any order**.

**Example 1:**

```
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]

```

**Example 2:**

```
Input: nums = [1], k = 1
Output: [1]

```

**Constraints:**

- `1 <= nums.length <= 105`
- `-104 <= nums[i] <= 104`
- `k` is in the range `[1, the number of unique elements in the array]`.
- It is **guaranteed** that the answer is **unique**.

**Follow up:** Your algorithm's time complexity must be better than `O(n log n)`, where n is the array's size.

## Solution Notes
### Brute Force Approach
- We can create a frequency hash map for all the numbers
- Then, we can sort the hash map by the values and return the `k` most frequent elements
- This would be an `O(n log n)` algorithm because of the sorting
- We can do better, ergo the follow-up

### Efficient Approach
- We can use the idea that there can be at most `n` distinct elements
    + i.e. If we have all unique numbers, all the frequencies are the same and we can return any `k` numbers as the response
- How can we use this?
    + Create a hash map to count the frequencies of all the numbers
    + Then, create an array of size `n` with empty arrays as it's values
        - We'll use each index of the array as the frequency value and add the value to that array 
        - i.e. Let's take Example 1:
        - Our hashmap will look like: `{ 1: 3, 2: 2, 3: 1}`
        - Our array will be populated as such: `[[], [3], [2], [1], [], []]`
        - As we can see, at index 1, we have the value `3` because `3` occurs 1 time
        - At index 2, we have the value `2`, because `2` occurs twice
        - At index 3, we have the value `1`, because `1` occurs once
        - The reason we use an 2D array is to account for numbers that have the same frequency
    + From this, we can iterate backwards and create our resultant array, until it is length `k`
- This removes the requirement of sorting and we only iterative over the input array twice -> `O(n) + O(n) = O(n)`


### Complexity
#### Time: `O(n)`
- We iterate over the input array once to create the frequency hash map
- Iterate again to create the frequency array
- `O(n) + O(n) = O(n)`

#### Space: `O(n)`
- Create a frequency hash map -> `O(n)`
- Create a frequency array -> `O(n)`
- `O(n) + O(n) = O(n)`

## Solution

```go
func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)
	for _, num := range nums {
		if _, ok := counts[num]; ok {
			counts[num]++
		} else {
			counts[num] = 0
		}
	}

	n := len(nums)
	freq := make([][]int, n)
	for num, f := range counts {
		freq[f] = append(freq[f], num)
	}

	var result []int
	i := n - 1
	for len(result) < k {
		arr := freq[i]
		for _, num := range arr {
			result = append(result, num)
			if len(result) == k {
				return result
			}
		}
		i--
	}

	return result
}

```
