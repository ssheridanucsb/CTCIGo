package stacksandqueues

import "errors"

type StackIntNode struct {
	data int
	next *StackIntNode
}

func NewStackIntNode(data int) *StackIntNode {
	node := StackIntNode{data: data, next: nil}
	return &node
}

type IntStack struct {
	top  *StackIntNode
	size int
}

func NewIntStack() *IntStack {
	return &IntStack{top: nil, size: 0}
}

func (stack *IntStack) isEmpty() bool {
	return stack.top == nil
}

func (stack *IntStack) pop() (int, error) {
	if stack.isEmpty() {
		return 0, errors.New("emtpy stack")
	}

	item := stack.top.data
	stack.top = stack.top.next
	stack.size--
	return item, nil
}

func (stack *IntStack) peek() (int, error) {
	if stack.isEmpty() {
		return 0, errors.New("emtpy stack")
	}

	return stack.top.data, nil
}

func (stack *IntStack) push(data int) {
	node := NewStackIntNode(data)
	node.next = stack.top
	stack.top = node
	stack.size++
}
