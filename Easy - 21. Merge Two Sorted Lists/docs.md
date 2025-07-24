# 21. Merge Two Sorted Lists

## Problem

You are given the heads of two sorted linked lists `list1` and `list2`.

Merge the two lists into one **sorted** list. The list should be made by splicing together the nodes of the first two lists.

Return _the head of the merged linked list_.

**Example 1:**

![](https://assets.leetcode.com/uploads/2020/10/03/merge_ex1.jpg)

```
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]

```

**Example 2:**

```
Input: list1 = [], list2 = []
Output: []

```

**Example 3:**

```
Input: list1 = [], list2 = [0]
Output: [0]

```

**Constraints:**

- The number of nodes in both lists is in the range `[0, 50]`.
- `-100 <= Node.val <= 100`
- Both `list1` and `list2` are sorted in **non-decreasing** order.

## Approach
We can start off by creating a dummy list node that will store our resultant list in it's `next` pointer. Then, we take a pointer to it as the tail of our list. We also initialize pointers for list1 and list2. 

We start by iterating while the list1 and list2 pointers are not nil. We check the values for them and whichever is less, we set the tail's next pointer to a new node with that value and move the tail pointer to the new node. We also move the pointer for the list that we took the value from to its next node.

When we complete the iteration, there may be one list that we haven't iterated through fully so we check both pointers and if either of them is not nil, we set the tail's next pointer to the remaining nodes of that list.

Finally, we return the next pointer of the dummy node which contains our created head of the merged list.

## Complexity
### Time: `O(n)`
We iterate through both lists once, where `n` is the total number of nodes in both lists combined.

### Space: `O(1)`
No extra space is used except for the resultant list which is not accounted for in the analysis.

## Solution

```go
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	curr1, curr2 := list1, list2
	for curr1 != nil && curr2 != nil {
		var v int
		if curr1.Val <= curr2.Val {
			v = curr1.Val
			curr1 = curr1.Next
		} else {
			v = curr2.Val
			curr2 = curr2.Next
		}
		tail.Next = &ListNode{ Val: v }
		tail = tail.Next
	}
	if curr2 != nil {
		tail.Next = curr2
	}
	if curr1 != nil {
		tail.Next = curr1
	}
	return dummy.Next
}

```
