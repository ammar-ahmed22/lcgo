# 138. Copy List with Random Pointer

## Problem

A linked list of length `n` is given such that each node contains an additional random pointer, which could point to any node in the list, or `null`.

Construct a [**deep copy**](https://en.wikipedia.org/wiki/Object_copying#Deep_copy) of the list. The deep copy should consist of exactly `n` **brand new** nodes, where each new node has its value set to the value of its corresponding original node. Both the `next` and `random` pointer of the new nodes should point to new nodes in the copied list such that the pointers in the original list and copied list represent the same list state. **None of the pointers in the new list should point to nodes in the original list**.

For example, if there are two nodes `X` and `Y` in the original list, where `X.random --> Y`, then for the corresponding two nodes `x` and `y` in the copied list, `x.random --> y`.

Return _the head of the copied linked list_.

The linked list is represented in the input/output as a list of `n` nodes. Each node is represented as a pair of `[val, random_index]` where:

- `val`: an integer representing `Node.val`
- `random_index`: the index of the node (range from `0` to `n-1`) that the `random` pointer points to, or `null` if it does not point to any node.

Your code will **only** be given the `head` of the original linked list.

**Example 1:**

![](https://assets.leetcode.com/uploads/2019/12/18/e1.png)

```
Input: head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
Output: [[7,null],[13,0],[11,4],[10,2],[1,0]]

```

**Example 2:**

![](https://assets.leetcode.com/uploads/2019/12/18/e2.png)

```
Input: head = [[1,1],[2,1]]
Output: [[1,1],[2,1]]

```

**Example 3:**

**![](https://assets.leetcode.com/uploads/2019/12/18/e3.png)**

```
Input: head = [[3,null],[3,0],[3,null]]
Output: [[3,null],[3,0],[3,null]]

```

**Constraints:**

- `0 <= n <= 1000`
- `-104 <= Node.val <= 104`
- `Node.random` is `null` or is pointing to some node in the linked list.

## Approach
To solve this problem, we can create mapping from old nodes to new nodes and then use the old nodes when creating the random connections in the new nodes.

To start, we'll create a new list that has old and new nodes interweaved. i.e.: `7 -> 7' -> 13 -> 13' -> 11 -> 11' -> 10 -> 10' -> 1 -> 1' -> nil`.

From this, we can iterate over only the old nodes and set the connections on the `curr.Next.Random` (new node) to be `curr.Random.Next` because the next node will always be the corresponding new node.

After this, we remove the old nodes and re-construct the old list.

## Complexity
### Time: `O(n)`
We iterate over the input array once, many times.

### Space: `O(1)`
The only extra space created is for the result.

## Solution

```go
func copyRandomList(head *Node) *Node {
	// Create the interweaved list, oldNode -> newNode
	dummy := &Node{}
	curr := head
	tail := dummy
	for curr != nil {
		curr.Next = &Node{
			Val:  curr.Val,
			Next: curr.Next,
		}
		tail.Next = curr
		tail = tail.Next.Next
		curr = curr.Next.Next
	}

	// Create the random connections
	curr = dummy.Next
	for curr != nil {
		if curr.Random != nil {
			curr.Next.Random = curr.Random.Next
		} else {
			curr.Next.Random = nil
		}

		curr = curr.Next.Next
	}

	// Remove the old nodes
	curr = dummy
	for curr != nil && curr.Next != nil {
		if curr.Next != nil {
			tmp := curr.Next // Save the old node
			curr.Next = curr.Next.Next
			tmp.Next = tmp.Next.Next // Re-construct the old list
		} else {
			curr.Next = nil
		}
		curr = curr.Next
	}

	return dummy.Next
}

```
