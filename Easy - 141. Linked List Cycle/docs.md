# 141. Linked List Cycle

## Problem

Given `head`, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the `next` pointer. Internally, `pos` is used to denote the index of the node that tail's `next` pointer is connected to. **Note that `pos` is not passed as a parameter**.

Return `true` _if there is a cycle in the linked list_. Otherwise, return `false`.

**Example 1:**

![](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist.png)

```
Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).

```

**Example 2:**

![](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test2.png)

```
Input: head = [1,2], pos = 0
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.

```

**Example 3:**

![](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test3.png)

```
Input: head = [1], pos = -1
Output: false
Explanation: There is no cycle in the linked list.

```

**Constraints:**

- The number of the nodes in the list is in the range `[0, 104]`.
- `-105 <= Node.val <= 105`
- `pos` is `-1` or a **valid index** in the linked-list.

**Follow up:** Can you solve it using `O(1)` (i.e. constant) memory?

## Approach
### Brute Force Approach
Solving this with the "brute force" appraoch is quite trivial. We simply use a hash set to store each visited node while we're iterating. If we see a node we've already seen, that means there is a cycle. Otherwise the iteration will complete and there is no cycle. In terms of time complexity, the is `O(n)` since we iterate once but it uses a hash set which is also `O(n)`.

### Optimal Approach
The optimal approach to solving this is by using a slow and fast pointer approach. If there is no cycle, the iteration will complete. Otherwise, the fast pointer will keep going over the cycle over and over until the slow pointer meets it somewhere. If they meet, we return true, otherwise the iteration will complete and return false.

## Complexity
### Time: `O(n)`
We iterate over the input array once, the amount of times the fast pointer will go around the cycle is bounded by the separation between the slow and fast pointers when it starts going around.

### Space: `O(1)`
No extra space.

## Solution

```go
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
```
