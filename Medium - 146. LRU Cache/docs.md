# 146. LRU Cache

## Problem

Design a data structure that follows the constraints of a **[Least Recently Used (LRU) cache](https://en.wikipedia.org/wiki/Cache_replacement_policies#LRU)**.

Implement the `LRUCache` class:

- `LRUCache(int capacity)` Initialize the LRU cache with **positive** size `capacity`.
- `int get(int key)` Return the value of the `key` if the key exists, otherwise return `-1`.
- `void put(int key, int value)` Update the value of the `key` if the `key` exists. Otherwise, add the `key-value` pair to the cache. If the number of keys exceeds the `capacity` from this operation, **evict** the least recently used key.

The functions `get` and `put` must each run in `O(1)` average time complexity.

**Example 1:**

```
Input
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
Output
[null, null, null, 1, null, -1, null, -1, 3, 4]

Explanation
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // cache is {1=1}
lRUCache.put(2, 2); // cache is {1=1, 2=2}
lRUCache.get(1);    // return 1
lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
lRUCache.get(2);    // returns -1 (not found)
lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
lRUCache.get(1);    // return -1 (not found)
lRUCache.get(3);    // return 3
lRUCache.get(4);    // return 4

```

**Constraints:**

- `1 <= capacity <= 3000`
- `0 <= key <= 104`
- `0 <= value <= 105`
- At most `2 * 105` calls will be made to `get` and `put`.

## Approach
To solve this problem optimally, we can use a doubly linked list in which we keep track of the left and right (head and tail) of the list for our LRU and MRU (most recently used), respectively.

Since we want O(1) retrieval, we'll use a hash map to store the keys alongside a pointer to a node in the list where the key and value will be stored.

When we retrieve an element, we want to update our list to reflect that the retrieval is now the MRU so what we do is remove that specific node from the list and then re-insert at the right (tail) for it to become the MRU.

When we want to insert an element, we create a new node for our list, insert it at the right (MRU). Then, we check if our capacity has been exceeded and remove the node at the left (LRU) and delete the corresponding key from the hashmap.

## Complexity
### Time: `O(1)`
Removing and inserting into the list when we have pointers for the head and tail is constant.

### Space: `O(cap)`
We create a hashmap and a doubly linked list that will have a maximum size of `capacity`.

## Solution

```go
type Node struct {
	Key  int
	Val  int
	Next *Node
	Prev *Node
}

type LRUCache struct {
	left     *Node
	right    *Node
	capacity int
	cache    map[int]*Node
}

func Constructor(capacity int) LRUCache {
	// Left=LRU, right=MRU
	left := &Node{}
	right := &Node{}

	// Connect them to start
	left.Next = right
	right.Prev = left
	return LRUCache{
		capacity: capacity,
		left: left,
		right: right,
		cache: make(map[int]*Node),
	}
}

// Remove the node 
func (this *LRUCache) remove(node *Node) {
	prev, next := node.Prev, node.Next
	prev.Next = next
	next.Prev = prev
}

// Insert this node at right (i.e. in between right.prev and right)
func (this *LRUCache) insert(node *Node) {
	prev, next := this.right.Prev, this.right
	prev.Next = node
	next.Prev = node
	node.Next = next
	node.Prev = prev
}

func (this *LRUCache) Get(key int) int {
	if node, exists := this.cache[key]; exists {
		// Remove this node
		this.remove(node)
		// Insert node at right (MRU)
		this.insert(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.cache[key]; exists {
		// Remove this node (already exists)
		this.remove(node)
	}
	this.cache[key] = &Node{Key: key, Val: value}
	// Insert this node (this.cache[key]) at MRU
	node, _ := this.cache[key]
	this.insert(node)

	if len(this.cache) > this.capacity {
		// Remove the LRU and delete from hash map
		lru := this.left.Next
		this.remove(lru)
		delete(this.cache, lru.Key)
	}
}

```
