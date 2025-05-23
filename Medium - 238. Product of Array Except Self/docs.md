# 238. Product of Array Except Self

## Problem

Given an integer array `nums`, return _an array_ `answer` _such that_ `answer[i]` _is equal to the product of all the elements of_ `nums` _except_ `nums[i]`.

The product of any prefix or suffix of `nums` is **guaranteed** to fit in a **32-bit** integer.

You must write an algorithm that runs in `O(n)` time and without using the division operation.

**Example 1:**

```
Input: nums = [1,2,3,4]
Output: [24,12,8,6]

```

**Example 2:**

```
Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]

```

**Constraints:**

- `2 <= nums.length <= 105`
- `-30 <= nums[i] <= 30`
- The input is generated such that `answer[i]` is **guaranteed** to fit in a **32-bit** integer.

**Follow up:** Can you solve the problem in `O(1)` extra space complexity? (The output array **does not** count as extra space for space complexity analysis.)

## Solution Notes
- We can use prefix and suffix product arrays to solve this
- The prefix product array would be the product of all the numbers before the current number
    + i.e. `prefix[i] = product of all the numbers before nums[i]`
- The suffix produc array would be the product of all the number after the current number
    + i.e. `suffix[i] = product of all the numbers after nums[i]`
- Therefore, the `answer[i] = prefix[i] * suffix[i]` because we want the product of all the number except itself
- To create the prefix and suffix arrays:
    + For the prefix array, we start with `1` as the first value because there is nothing before the first number
        * Iterate from `1 to n`, multiply, `nums[i - 1] * prefix[i - 1]` and set it to `prefix[i]`
    + For the suffix array, we start with `1` as the last value because there is nothing after the last number
        * Iterate from `n-2 to 0`, multiple, `nums[i + 1] * suffix[i + 1]` and set it to `suffix[i]`
- Finally, we can create the answer by multiplying them together

### Complexity
#### Time: `O(n)`
- Iterating over the input array twice, one after the other

#### Space: `O(n)`
- We create two arrays of size `n`, prefix and suffix


## Solution

```go
func productExceptSelf(nums []int) []int { 
	n := len(nums)
	prefix := make([]int, n)
	prefix[0] = 1
	for i := 1; i < n; i++ {
		prefix[i] = nums[i - 1] * prefix[i - 1]
	}

	suffix := make([]int, n)
	suffix[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		suffix[i] = nums[i + 1] * suffix[i + 1]
	}

	answer := make([]int, n)
	for i := range n {
		answer[i] = prefix[i] * suffix[i]
	}

	return answer
}
```

## Solution

```go
func productExceptSelf(nums []int) []int { 
	n := len(nums)
	prefix := make([]int, n)
	prefix[0] = 1
	for i := 1; i < n; i++ {
		prefix[i] = nums[i - 1] * prefix[i - 1]
	}

	suffix := make([]int, n)
	suffix[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		suffix[i] = nums[i + 1] * suffix[i + 1]
	}

	answer := make([]int, n)
	for i := range n {
		answer[i] = prefix[i] * suffix[i]
	}

	return answer
}
```
