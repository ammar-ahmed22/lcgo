# 128. Longest Consecutive Sequence

## Problem

Given an unsorted array of integers `nums`, return _the length of the longest consecutive elements sequence._

You must write an algorithm that runs in `O(n)` time.

**Example 1:**

```
Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.

```

**Example 2:**

```
Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9

```

**Example 3:**

```
Input: nums = [1,0,1,2]
Output: 3

```

**Constraints:**

- `0 <= nums.length <= 105`
- `-109 <= nums[i] <= 109`

## Approach
We can start by creating a hashmap to store which numbers occur in the array. After this, we iterate over the numbers. For each number, we check if the value one less than it occurs in the array via the hashmap. If it doesn't, we know we are the start of a sequence. Once we know we are at the start of a sequence, we iterate as long as there is a value exactly one more than that in the array, keep track of the count. We update the max value with this count. 

### Edge Cases
The first edge case is when there are no values in the array, we want to simply return `0`. 

The next edge case occurs with the time complexity. If there are duplicates of the start of a sequence, without proper handling, we'll run unnecessary iterations. For example, if a sequence starts with 0 and has a length of 100, and we have two zeroes in the array, we'll run that 100 length iteration twice. To mitigate this, we want to mark in the hash map if the start of a sequence has already been processed and check for that mark before iterating over the sequence and override it. 

In our case, we do this by marking with a boolean value in the hashmap since we are using a boolean map.

## Complexity
### Time: `O(n)`
It might look like it's not `O(n)` because we have a nested loop, however, there is an important consideration. We only go down that nested loop if it's the start of a sequence, otherwise, we keep marching forward. Therefore, in the worst case, when all values are consecutive, we only go down the nested loop once, maintaining `O(n)`

### Space: `O(n)`
We create a hashmap of size `n`.

## Solution

```go
func longestConsecutive(nums []int) int {
	// Edge case check
	if len(nums) == 0 {
		return 0
	}

	// Creating hash map
	freq := make(map[int]bool)
	for _, num := range nums {
		if _, exists := freq[num]; !exists {
			freq[num] = true
		}
	}

	// Intialize to 1 because we already handled the case if its zero
	ans := 1
	for _, num := range nums {
		// If there is no previous value, start of a sequence
		if _, prevExists := freq[num - 1]; !prevExists {
			if !freq[num] { // Check if the current value is marked as false (we've already processed this sequence start)
				continue
			}
			freq[num] = false // Mark the sequence start as processed

			// Find the length of the sequence
			count := 0
			_, hasNext := freq[num + count]
			for hasNext {
				count++
				_, hasNext = freq[num + count]
			}
			// Update answer
			ans = max(ans, count)
		}
	}

	return ans
}

```
