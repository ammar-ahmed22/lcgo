# 121. Best Time to Buy and Sell Stock

## Problem

You are given an array `prices` where `prices[i]` is the price of a given stock on the `ith` day.

You want to maximize your profit by choosing a **single day** to buy one stock and choosing a **different day in the future** to sell that stock.

Return _the maximum profit you can achieve from this transaction_. If you cannot achieve any profit, return `0`.

**Example 1:**

```
Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

```

**Example 2:**

```
Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.

```

**Constraints:**

- `1 <= prices.length <= 105`
- `0 <= prices[i] <= 104`

## Approach
We can use a sliding window approach to solve this. We start by initializing our pointers at the first and second element. If the profit is negative, we track the max profit and increment the right pointer. If the profit is negative, we update the left pointer to the right pointers value because we want this newly found small value and continue to increment our right pointer on every iteration.

## Complexity
### Time: `O(n)`
We iterate over the array once.

### Space: `O(1)`
No extra space is created.

## Solution

```go
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	l := 0
	r := 1
	maxP := 0
	for r < len(prices) {
		p := prices[r] - prices[l]
		if p > 0 {
			maxP = max(maxP, p)
		} else {
			l = r
		}
		r++
	}
	return maxP
}
```
