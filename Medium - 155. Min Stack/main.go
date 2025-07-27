package main

import (
	. "github.com/ammar-ahmed22/lcgo/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
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

// <-- DO NOT REMOVE: PROBLEM END -->

func main() {
	// We'll have to run our own tests for this
	ms := Constructor()
	ms.Push(-2)
	ms.Push(0)
	ms.Push(-3)
	minVal := ms.GetMin()
	Test("GetMin(1/2)", -3, minVal)
	ms.Pop()
	top := ms.Top()
	Test("Top(1/1)", 0, top)
	minVal = ms.GetMin()
	Test("GetMin(2/2)", -2, minVal)
}
