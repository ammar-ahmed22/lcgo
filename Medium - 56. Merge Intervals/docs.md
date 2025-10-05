# 56. Merge Intervals

## Problem

Given an array of `intervals` where `intervals[i] = [starti, endi]`, merge all overlapping intervals, and return _an array of the non-overlapping intervals that cover all the intervals in the input_.

**Example 1:**

```
Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].

```

**Example 2:**

```
Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.

```

**Example 3:**

```
Input: intervals = [[4,7],[1,4]]
Output: [[1,7]]
Explanation: Intervals [1,4] and [4,7] are considered overlapping.

```

**Constraints:**

- `1 <= intervals.length <= 104`
- `intervals[i].length == 2`
- `0 <= starti <= endi <= 104`

## Approach
To solve this problem, we can start by sorting the intervals by the start. After this, we iterate through using a sliding window approach processing each interval with the next one.

To check if these two intervals will overlap, we check if the end of the first interval is greater than or equal to the start of the next interval. If they are, we merge these by taking the minimum of both starts and the maximum of both ends.

Practically, to avoid having to delete stuff from an array which will be an O(n) operation everytime we want to delete an interval from the array, we can use top of our resultant list as the the current and the next will be what we are iterating starting from `next = 1`.

For example, if we take **Example 1**, `[[1, 3], [2, 6], [8, 10], [15, 18]]`

We'll start by adding the first interval to our result, `result = [[1, 3]]`

Then we start iterating with `i = 1`, with `i = 1`, our `next` is `next = [2, 6]` and our `curr` is the top of the result, `curr = result.top() = [1, 3]`.

Since should be merged, we pop the top off of our result and add the merged interval, `result = [[1, 6]]` and we move on to the next. If there is no merge required, we simply push the next on to the result.


## Complexity
### Time: `O(n log n)`
Sorting the array is `n log n` and then we iterate over the intervals once, `O(n)`

### Space: `O(1)`
Space for the result is not accounted for so we don't create any new space.

## Solution

```go
import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		curr := result[len(result)-1]
		next := intervals[i]

		if curr[1] >= next[0] {
			// merge
			result = result[:len(result)-1] // pop off the top
			merged := []int{min(curr[0], next[0]), max(curr[1], next[1])}
			result = append(result, merged)
		} else {
			result = append(result, next)
		}
	}

	return result
}

```
