# 143. Reorder List

## Problem

You are given the head of a singly linked-list. The list can be represented as:

```
L0 → L1 → … → Ln - 1 → Ln

```

_Reorder the list to be on the following form:_

```
L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …

```

You may not modify the values in the list's nodes. Only nodes themselves may be changed.

**Example 1:**

![](https://assets.leetcode.com/uploads/2021/03/04/reorder1linked-list.jpg)

```
Input: head = [1,2,3,4]
Output: [1,4,2,3]

```

**Example 2:**

![](https://assets.leetcode.com/uploads/2021/03/09/reorder2-linked-list.jpg)

```
Input: head = [1,2,3,4,5]
Output: [1,5,2,4,3]

```

**Constraints:**

- The number of nodes in the list is in the range `[1, 5 * 104]`.
- `1 <= Node.val <= 1000`

## Approach
In order to solve this optimally, we can realize that we are essentially doing a merge on the first and second halves of the list with the second half being reversed.

If we have, `1 -> 2 -> 3 -> 4 -> 5 -> nil`, the answer is `1 -> 5 -> 2 -> 4 -> 3 -> nil`. Taking the second half and reversing it, we have `5 -> 4 -> nil`, with the first half being `1 -> 2 -> 3 -> nil`. So, it's easy to see how merging these will give us our answer.

In order to actually accomplish this, we have 3 steps:
1. Find the middle
2. Reverse the second half
3. Merge

### Finding the middle
To find the middle easily, we can use fast and slow pointers, when the fast pointer reaches the end, the slow pointer will be at the center.

### Reversing
Once the middle is at found at the slow pointer, we want to reverse the list by using the slow pointers next value as our current. We'll want to separate the lists out as well by breaking the link (set `slow.Next = nil`)

### Merging
To merge the lists, we create two pointers, `first` and `second` initialized to the `head` and `reversed`. We iterate while the `reversed` is not nil (because `reversed` will be the smaller list). On each iteration, we save the next values for each pointer in temporary variables (`tmp1` and `tmp2`), we set `first`'s next value to be `second`. Then, we set the `second`'s next value to be `tmp1`. This is our merge step. After that we increment the pointers but settings `first = tmp1` and `second = tmp2`.  

## Complexity
### Time: `O(n)`
We iterate over the list once, 3 times.

### Space: `O(1)`
No extra space is created.

## Solution

```go
func reorderList(head *ListNode) {
	// Find middle of list with slow and fast pointers
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Reverse the second half of the list
	curr := slow.Next
	slow.Next = nil
	reversed := slow.Next
	for curr != nil {
		tmp := curr.Next
		curr.Next = reversed
		reversed = curr
		curr = tmp
	}

	// Merge the first and second half (reversed) of the list
	first, second := head, reversed
	for second != nil {
		tmp1, tmp2 := first.Next, second.Next
		first.Next = second
		second.Next = tmp1
		first = tmp1
		second = tmp2
	}
}
```
