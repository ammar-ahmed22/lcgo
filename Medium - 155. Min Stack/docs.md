# 155. Min Stack

## Problem

Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the `MinStack` class:

- `MinStack()` initializes the stack object.
- `void push(int val)` pushes the element `val` onto the stack.
- `void pop()` removes the element on the top of the stack.
- `int top()` gets the top element of the stack.
- `int getMin()` retrieves the minimum element in the stack.

You must implement a solution with `O(1)` time complexity for each function.

**Example 1:**

```
Input
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

Output
[null,null,null,null,-3,null,0,-2]

Explanation
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin(); // return -3
minStack.pop();
minStack.top();    // return 0
minStack.getMin(); // return -2

```

**Constraints:**

- `-231 <= val <= 231 - 1`
- Methods `pop`, `top` and `getMin` operations will always be called on **non-empty** stacks.
- At most `3 * 104` calls will be made to `push`, `pop`, `top`, and `getMin`.

## Approach
As expected, we'll use a stack to solve this problem, however, the complexity arises with implementing the minimum value logic in constant time.

To solve this issue, we can use two stacks within the `MinStack`. One will be our main `stack` that will track all of the values and the other will be our `minStack` that will track the minimum values.

To start, we initialize both of them. 

For the `push`, we always push to the main `stack`. We only push to `minStack` if the `minStack` is empty (first minimum value) or if the value is less than or equal to the top of `minStack`. This way the `minStack` will always have the smallest value at the top.

For the `pop`, we always pop off the main `stack`. We only pop off the `minStack` when the top of the main `stack` is equal to the top of the `minStack`.

After this, the remaining methods are trivial.

For my specific solution, I implemented a generic, `Stack[R any]` struct to make using array-based stacks a little bit easier syntactically. I could have also used a linked-list based stack for better efficiency but this is sufficient.

## Complexity
### Time: `O(1)`
None of the methods do any iteration at all so it's all constant time.

### Space: `O(n)`
We create two stacks that could potentially be the size of the inputs.

## Solution

```go
type Stack[R any] struct {
	stack []R
}

func (this *Stack[R]) Append(val R) {
	this.stack = append(this.stack, val)
}

func (this *Stack[R]) Len() int {
	return len(this.stack)
}

func (this *Stack[R]) Pop() *R {
	if this.Len() == 0 {
		return nil
	}
	val := this.stack[this.Len() - 1]
	this.stack = this.stack[:this.Len() - 1]
	return &val
}

func (this *Stack[R]) Peek() *R {
	if this.Len() == 0 {
		return nil
	}
	return &this.stack[this.Len() - 1]
}

type MinStack struct {
	stack    Stack[int] 
	minStack Stack[int] 
}

func Constructor() MinStack {
	return MinStack{
		stack: Stack[int]{},
		minStack: Stack[int]{},
	}
}

func (this *MinStack) Push(val int) {
	if this.minStack.Len() == 0 || val <= *this.minStack.Peek() {
		this.minStack.Append(val)
	}
	this.stack.Append(val)
}

func (this *MinStack) Pop() {
	if this.Top() == *this.minStack.Peek() {
		this.minStack.Pop()
	}
	this.stack.Pop()
}

func (this *MinStack) Top() int {
	return *this.stack.Peek() 
}

func (this *MinStack) GetMin() int {
	return *this.minStack.Peek() 
}

```
