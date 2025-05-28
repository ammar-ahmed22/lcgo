# 206. Reverse Linked List

## Problem

Given the `head` of a singly linked list, reverse the list, and return _the reversed list_.

**Example 1:**

![](https://assets.leetcode.com/uploads/2021/02/19/rev1ex1.jpg)

```
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]

```

**Example 2:**

![](https://assets.leetcode.com/uploads/2021/02/19/rev1ex2.jpg)

```
Input: head = [1,2]
Output: [2,1]

```

**Example 3:**

```
Input: head = []
Output: []

```

**Constraints:**

- The number of nodes in the list is the range `[0, 5000]`.
- `-5000 <= Node.val <= 5000`

**Follow up:** A linked list can be reversed either iteratively or recursively. Could you implement both?

## Approach
To solve this iteratively, we can iterater over the list by intializing a pointer to the head and iterating while that pointer is not null. On each iteration, we start off by saving the current value's next node to a temp variable. Then, we assing the current value's next reference to our `previous` node that was also initialized at the start. After that we set the previous value to the current node and assign the current node to that temp variable we saved to continue iterating. 

## Complexity
### Time: `O(n)`
We iterate through the list once.

### Space: `O(1)`
No extra space created.

## Solution

```go
func reverseList(head *ListNode) *ListNode { 
	curr := head
	var prev *ListNode
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	return prev
}
```
