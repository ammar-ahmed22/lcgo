# 15. 3Sum

## Problem

Given an integer array nums, return all the triplets `[nums[i], nums[j], nums[k]]` such that `i != j`, `i != k`, and `j != k`, and `nums[i] + nums[j] + nums[k] == 0`.

Notice that the solution set must not contain duplicate triplets.

**Example 1:**

```
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.

```

**Example 2:**

```
Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.

```

**Example 3:**

```
Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.

```

**Constraints:**

- `3 <= nums.length <= 3000`
- `-105 <= nums[i] <= 105`

## Approach
To solve this, we can use a similar approach to [Two Sum II - Input Array Sorted](https://ammarahmed.ca/medium/two-sum-ii-input-array-sorted). 

The idea is that for each number, we want to run the algorithm for [Two Sum II](https://ammarahmed.ca/medium/two-sum-ii-input-array-sorted) on the remaining elements using the current number's negation as the target. However, this will require sorting the input array to start. We don't need to worry about the elements before the current number because the problem states to return the **distinct triplets**. 

Let's go through an example with **Example 1**:

We have the input array: `[-1, 0, 1, 2, -1, 4]`. After sorting, we have: `[-1, -1, 0, 1, 2, 4]`.

Now, for the first iteration, we're treating, `arr[0] = -1` as our current number (`target = 1`) and we'll run the two-pointer algorithm on the remaining elements: `[-1, 0, 1, 2, 4]`. In other words, we're looking in that subset for two numbers that add up to `1` because then the triplet will add up to `0`. We can set up our two pointers and increment/decrement based upon whether the current target is less than or greater than the `target`. We'll find that `-1, 2` add up to `1` (`target`) so we can add that triplet to our list. We'll continue this way until we finish the iteration. See [Two Sum II - Input Array Sorted](https://ammarahmed.ca/medium/two-sum-ii-input-array-sorted) for details on the algorithm we're running on the subset.

Another important note is in regards to duplicate triplets. We can mitigate this by skipping over numbers we've already processed. Since the array is sorted, all duplicates will be beside eachother. Therefore, in the main loop, we can check if the next number is the same as the current and skip the current. In the inner loop, when we find a triplet, we can keep incrementing/decrementing until we are at the last occurrence of the particular number before moving on to finding any more triplets with that target.

### Complexity:
#### Time: `O(n^2)`
Sorting the input array is `n log n`. After that we iterate over the array twice, which is `O(n^2)`. Adding them together, we get `O(n log n) + O(n^2) = O(n^2)` (we drop lower order terms).

#### Space: `O(1)`
No extra space is created.

## Solution

```go
import "slices"
func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	var ans [][]int
	for i, curr := range nums {
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		target := curr * -1
		l := i + 1
		r := len(nums) - 1
		for l < r {
			pot := nums[l] + nums[r]
			if pot == target {
				ans = append(ans, []int{curr, nums[l], nums[r]})
				for l < r && nums[l] == nums[l + 1] {
					l++
				}
				for l < r && nums[r] == nums[r - 1] {
					r--
				}
				l++
				r--
			} else if pot < target {
				l++
			} else {
				r--
			}
		}
	}
	return ans
}
```
