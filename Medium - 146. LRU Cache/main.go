package main

import (
	. "github.com/ammar-ahmed22/lcgo/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
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

// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = LRUCache

var testCases = []*TestCase[ReturnType]{
	// Add test cases here
	// Using twoSum as an example:
	// NewTestCase([]int{0, 1}).WithArgs([]int{2, 7, 11, 15}, 9),
	// You can also give a name to the test case:
	// NewTestCase([]int{0, 1}).WithArgs([]int{2, 7, 11, 15}, 9).WithName("Example Test Case"),
	// You can also provide a custom comparison function:
	// NewTestCase([]int{0, 1}).WithArgs([]int{2, 7, 11, 15}, 9).WithCompareFn(SliceEqualUnordered)
}

func main() {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	Test("Get(1) = 1", 1, lru.Get(1))
	lru.Put(3, 3)
	Test("Get(2) = -1", -1, lru.Get(2))
	lru.Put(4, 4)
	Test("Get(1) = -1", -1, lru.Get(1))
	Test("Get(3) = 3", 3, lru.Get(3))
	Test("Get(4) = 4", 4, lru.Get(4))
}
