# 239. Sliding Window Maximum

## Problem

You are given an array of integers `nums`, there is a sliding window of size `k` which is moving from the very left of the array to the very right. You can only see the `k` numbers in the window. Each time the sliding window moves right by one position.

Return _the max sliding window_.

**Example 1:**

```
Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
Output: [3,3,5,5,6,7]
Explanation:
Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7

```

**Example 2:**

```
Input: nums = [1], k = 1
Output: [1]

```

**Constraints:**

- `1 <= nums.length <= 105`
- `-104 <= nums[i] <= 104`
- `1 <= k <= nums.length`

## Approach
### Brute Force
To solve this with brute force, we can simply find the max for each window and update our answer. This would be an `O(k * (n - k))` algorithm.

### Efficient
We can use a deque to solve this problem because we want to be able to efficiently remove from the front and back. The idea is that we only care about tracking a new number if it's larger than the ones we've previously seen but we also want to remove numbers once we're done processing that window. Therefore, we need to be able to add and remove efficiently from the front and back of a data structure, hence a deque.

The way we can do this is using a deque to track indices instead of actual values. To start, we initialize our deque and our left and right pointers. We iterate using the right pointer. On each iteration, we first do a check where if the deque has elements and the top of the deque is less than the value at the right pointer, we pop from the front of the deque until the condition is no longer true. Then, we add the right index. Effectively, we are removing any values before the current value that are less than it because we no longer care about them. After this, we check if the left pointer is greater than the value at the back of the deque, if it is that means that value is no longer in the window so we can remove it. From there, we check if our window has reached the size of `k` and add the value at the back of the deque to our answer and slide our window over. 

## Complexity
### Time: `O(n)`
We iterate over the input only once.

### Space: `O(n)`
We create a deque which can have `n` elements.

## Solution

```go
func maxSlidingWindow(nums []int, k int) []int {
	ans := []int{}
	dq := list.New()
	l, r := 0, 0

	for r < len(nums) {
		for dq.Len() > 0 && nums[dq.Front().Value.(int)] < nums[r] {
			dq.Remove(dq.Front())
		}
		dq.PushFront(r)

		if l > dq.Back().Value.(int) {
			dq.Remove(dq.Back())
		}

		if (r + 1) >= k {
			ans = append(ans, nums[dq.Back().Value.(int)])
			l++
		}
		r++
	}

	return ans
}

```
