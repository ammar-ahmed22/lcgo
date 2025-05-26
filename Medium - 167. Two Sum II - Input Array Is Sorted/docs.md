# 167. Two Sum II - Input Array Is Sorted

## Problem

Given a **1-indexed** array of integers `numbers` that is already **_sorted in non-decreasing order_**, find two numbers such that they add up to a specific `target` number. Let these two numbers be `numbers[index1]` and `numbers[index2]` where `1 <= index1 < index2 <= numbers.length`.

Return _the indices of the two numbers,_ `index1` _and_ `index2` _, **added by one** as an integer array_ `[index1, index2]` _of length 2._

The tests are generated such that there is **exactly one solution**. You **may not** use the same element twice.

Your solution must use only constant extra space.

**Example 1:**

```
Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].

```

**Example 2:**

```
Input: numbers = [2,3,4], target = 6
Output: [1,3]
Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].

```

**Example 3:**

```
Input: numbers = [-1,0], target = -1
Output: [1,2]
Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].

```

**Constraints:**

- `2 <= numbers.length <= 3 * 104`
- `-1000 <= numbers[i] <= 1000`
- `numbers` is sorted in **non-decreasing order**.
- `-1000 <= target <= 1000`
- The tests are generated such that there is **exactly one solution**.

## Approach
The problem states that we must use constant space, therefore, we cannot use a hashmap as we did for the original two sum. 

In order to solve this, we can use the fact that it is sorted and there is exactly one solution. 

We can use a two-pointer approach to solve this. We initialize two pointers, one on each side of the array and iterate as long as they don't cross eachother. On each iteration, we can check if the two numbers add up to the target, if they do, we return the solution. Otherwise, we check if the result is less than the target; that means we need a larger number, so we increment our left pointer because the array is sorted and we want a larger value. If the result is greater than the target, we want a smaller number, so we decrement the right pointer because, again, the array is sorted and we want a smaller value.

Through this, we are guranteed to find the solution in the iteration because the tests are generated so that there is exactly one solution.

The problem also states that the array is **1-indexed** so we want to make sure to increment our index values when returning the solution.
### Complexity
#### Time: `O(n)`
We iterate over the input array only once.

#### Space: `O(1)`
We do not create any additional space.

## Solution

```go
func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1
	for l < r {
		curr := numbers[l] + numbers[r]
		if curr == target {
			return []int{l + 1, r + 1}
		} else if curr < target {
			l++
		} else {
			r--
		}
	}
	return []int{}
}
```
