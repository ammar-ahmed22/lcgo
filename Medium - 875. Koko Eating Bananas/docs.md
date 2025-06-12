# 875. Koko Eating Bananas

## Problem

Koko loves to eat bananas. There are `n` piles of bananas, the `ith` pile has `piles[i]` bananas. The guards have gone and will come back in `h` hours.

Koko can decide her bananas-per-hour eating speed of `k`. Each hour, she chooses some pile of bananas and eats `k` bananas from that pile. If the pile has less than `k` bananas, she eats all of them instead and will not eat any more bananas during this hour.

Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.

Return _the minimum integer_ `k` _such that she can eat all the bananas within_ `h` _hours_.

**Example 1:**

```
Input: piles = [3,6,7,11], h = 8
Output: 4

```

**Example 2:**

```
Input: piles = [30,11,23,4,20], h = 5
Output: 30

```

**Example 3:**

```
Input: piles = [30,11,23,4,20], h = 6
Output: 23

```

**Constraints:**

- `1 <= piles.length <= 104`
- `piles.length <= h <= 109`
- `1 <= piles[i] <= 109`

## Approach
We can use a binary search approach to solve this. 

The general premise is that the max speed will always be the max pile. This is because it will ensure that all piles are eaten because if the pile has less than `k`, she will eat all of them. Therefore, the upper limit for the speed is the max of all the piles.

From this, we can conduct a binary search to find the minimum speed. We intiailize our left and right pointers to 1 and max of the piles, respectively. On each iteration, we calculate `k` as the midpoint. Then, we find the time it will take Koko to eat the bananas using this proposed `k` value. We do this by iterating over the piles and adding the `ceil` division of the pile and `k`. Once we have the time it will take, if it is less than or equal to `h`, we have a potential solution. So, we update our resultant variable and try to look for a smaller value by decrementing the right pointer (set to `k - 1`). Otherwise, we want a larger value, so we increment the right pointer (set to `k + 1`). 

## Complexity
### Time: `O(n * log m)`
We start by finding the max of the piles which is `O(n)`. Then, we do a binary search with the upper bound as the max value in the piles, `m` (`O(log m)`). Inside the binary search, we calculate the total time which is `O(n)`. So, we have `O(n) + O(n * log m) = O(2n * log m) = O(n * log m)`.

### Space: `O(1)`
No extra space is created.

## Solution

```go
func minEatingSpeed(piles []int, h int) int {
	var (l, r int)
	for _, p := range piles {
		r = max(r, p)
	}

	var res int
	for l <= r {
		k := (l + r) / 2
		var tot float64
		for _, p := range piles {
			tot += math.Ceil(float64(p) / float64(k))
		}
		if tot <= float64(h) {
			res = k
			r = k - 1
		} else {
			l = k + 1
		}
	}
	return res
}
```
