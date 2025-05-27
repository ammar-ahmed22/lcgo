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

## Approach 
### Brute Force Approach
To "brute force" this, we can create a frequency hash map for all the numbers. Then, we can sort the hash map by the values and return the `k` most frequent elements.
Since we sort the hashmap, this would be an `O(n log n)` algorithm.

### Efficient Approach
In order to do better than the brute force approach, we want to do away with the sorting. In order to do this, we can use the fact that there are, at most, `n` distinct elements. 
So, we can start by, again, creating the hash map to count the frequencies of all the numbers. Next, we'll create a constant size array of size `n` containing empty arrays as it's values.
From this, we can iterate over the hashmap and use the frequency as the index to the array. For example, let's say we have the values: `[1, 1, 1, 2, 2, 3]`. The frequency hash map (zero-based, i.e. zero = one occurrence, one = two occurrences, etc.) would look like:
```
{
    1: 2,
    2: 1,
    3: 0,
}
```
Thus, the frequency array would look like: `[3, 2, 1]`.

Once we have the array created, it's trivial to create the result. We just iterate backwards over the values and continue adding values to the result until it reaches length `k`.


## Complexity
### Time: `O(n)`
We iterate once to create the frequency hash map and once more to create the frequency array. Therefore, `O(n) + O(n) = O(n)`

### Space: `O(n)`
We create the hash map and the frequency array both of which have a maximum size of `n`.

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
