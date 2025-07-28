# 2. Add Two Numbers

## Problem

You are given two **non-empty** linked lists representing two non-negative integers. The digits are stored in **reverse order**, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

**Example 1:**

![](https://assets.leetcode.com/uploads/2020/10/02/addtwonumber1.jpg)

```
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

```

**Example 2:**

```
Input: l1 = [0], l2 = [0]
Output: [0]

```

**Example 3:**

```
Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]

```

**Constraints:**

- The number of nodes in each linked list is in the range `[1, 100]`.
- `0 <= Node.val <= 9`
- It is guaranteed that the list represents a number that does not have leading zeros.

## Approach
To solve this problem, we are essentially just doing addition as you would do on paper. Whenever the resulting number is greater than or equal to 10, we put the new nodes value as the difference from 10 and carry-over the 1. We continue this on an on until we have constructed the whole list being careful with the edge cases.

To start, we create our dummy node to house our resultant list and set our `tail` pointer to it. We also create pointers for each input list, `curr1` and `curr2`. We iterate while both `curr1` and `curr2` are not nil because we want to keep the pointers in sync.

On each iteration, we calculate the sum of the values at each pointer, `curr1` and `curr2`. If the sum is less than 10, we simply create a new node with that value and assign it to the `tail.Next`. If the sum is greater than 10, this is where we have to handle "carrying". We create a new node with the value of `sum - 10`, then we carry over `1`. It will always be `1` because we are adding single digit numbers always in which the largest possible sum is 18. To do the "carrying", we add `1` to next node of whichever pointer's next value is not `nil`. If both of the next values are `nil`, we have reached the end of both of the lists at the same time, so we create another node for next value of the new node we are creating and assign it's value to `1`.

Finally, once the iteration is complete, we have to handle the edge cases of the lists not being the same length. We iterate while each of `curr1` and `curr2` pointers are not nil separately. In each iteration, we check the value of the node, if it's less than 10 (single digit), we simply create a new node at the tail. Otherwise, we do the same carrying over logic as above.The reason the values in this iteration can be more than 10 is because of potential carrying over we've done in the previous loop.

After all this, we return `dummy.Next` as our result.

## Complexity
### Time: `O(n)`
We iterate over each input list once.

### Space: `O(1)`
No extra space is created except for the output list.

## Solution

```go
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	curr1, curr2 := l1, l2
	for curr1 != nil && curr2 != nil {
		sum := curr1.Val + curr2.Val
		node := &ListNode{}
		if sum - 10 >= 0 {
			node.Val = sum - 10
			if curr1.Next != nil {
				curr1.Next.Val++
			} else if curr2.Next != nil {
				curr2.Next.Val++
			} else {
				node.Next = &ListNode{
					Val: 1,
				}
			}
		} else {
			node.Val = sum
		}
		tail.Next = node
		curr1 = curr1.Next
		curr2 = curr2.Next
		tail = tail.Next
	}

	for curr1 != nil {
		node := &ListNode{}
		if curr1.Val < 10 {
			node.Val = curr1.Val
		} else {
			node.Val = curr1.Val - 10
			if curr1.Next != nil {
				curr1.Next.Val++
			} else {
				node.Next = &ListNode{Val: 1}
			}
		}
		tail.Next = node
		tail = tail.Next
		curr1 = curr1.Next
	}

	for curr2 != nil {
		node := &ListNode{}
		if curr2.Val < 10 {
			node.Val = curr2.Val
		} else {
			node.Val = curr2.Val - 10
			if curr2.Next != nil {
				curr2.Next.Val++
			} else {
				node.Next = &ListNode{Val: 1}
			}
		}
		tail.Next = node
		tail = tail.Next
		curr2 = curr2.Next
	}

	return dummy.Next
}
```
