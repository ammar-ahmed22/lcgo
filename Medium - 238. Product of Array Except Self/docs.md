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

## Approach 
To solve this, we can use prefix and suffix product arrays. The prefix product array is an array that contains the product of all the values before the `ith` value. The suffix product array is the same thing but for values after the `ith` value. We can use this to solve the problem because we want the product of eveything except itself so we want to multiple the product before the `ith` value with the product after the `ith` value. The prefix and suffix arrays will handle this for us. 

To create the prefix array, we create an array of size `n` with `1` as the first value (there are no values before the first one so it defaults to 1). Then, we iterate from `1 to n` and set `prefix[i] = nums[i - 1] * prefix[i - 1]`. 

For the suffix array, we create an array of size `n` with `1` as the last value (there are no values after the last one so it defaults to 1). Then, we iterate from `n-2 to 0` and set `suffix[i] = nums[i + 1] * suffix[i + 1]`.

Finally, we create the answer by multiplying the product and suffix arrays (`answer[i] = prefix[i] * suffix[i]`).

## Complexity
### Time: `O(n)`
We iterate over the input array twice, one after the other.

### Space: `O(n)`
We create two arrays of size `n`

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
