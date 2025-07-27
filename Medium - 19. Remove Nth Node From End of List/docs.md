# 19. Remove Nth Node From End of List

## Problem

Given the `head` of a linked list, remove the `nth` node from the end of the list and return its head.

**Example 1:**

![](https://assets.leetcode.com/uploads/2020/10/03/remove_ex1.jpg)

```
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]

```

**Example 2:**

```
Input: head = [1], n = 1
Output: []

```

**Example 3:**

```
Input: head = [1,2], n = 1
Output: [1]

```

**Constraints:**

- The number of nodes in the list is `sz`.
- `1 <= sz <= 30`
- `0 <= Node.val <= 100`
- `1 <= n <= sz`

**Follow up:** Could you do this in one pass?

## Approach
### Brute Force
To solve this with brute force, we can simply conver the list to an array, remove the node at `len(arr) - n` and reconstruct the result. This would be `O(n)` space.

### Two Pass Approach
We can remove the space constraint using a two pass approach. In the first pass, we calculate the length of the list, `N`. In the second pass, we remove the node at `N - n`. 

To do this, we can use a dummy node before the head of the list alongside tracking a counter while we iterate over the list starting at the dummy node. When we get to `N - n`, we will be at the node before the node to remove at which point we can set it's `Next` value to `Next.Next`. 

### One Pass Approach
To solve this in one pass as the problem asks, we can use two pointers. If we separate the pointers by `n`, when the right pointer gets to the end, the left pointer will be on the node to remove. Therefore, if we want to get it on the node before, we can use the same dummy node approach as above.

First, we move our right pointer to the `n`th node starting from the head. Next, we create a dummy node and assign it's `Next` value to be the head of our input list. Then, we set our left pointer at the dummy node. From this, we iterate until the right pointer reaches the end of the array. After that, our left pointer will be at the node before the one to remove so we can remove it with `left.Next = left.Next.Next`.

## Complexity
### Time: `O(n)`
We iterate over the input exactly once.

### Space: `O(1)`
No extra space is created.

## Solution

```go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	right := head
	for n > 0 {
		right = right.Next
		n--
	}

	dummy := &ListNode{}
	dummy.Next = head
	left := dummy
	for right != nil {
		left = left.Next
		right = right.Next
	}
	left.Next = left.Next.Next
	return dummy.Next
}
```
