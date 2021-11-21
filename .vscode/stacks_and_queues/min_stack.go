package stacksandqueues

import "errors"

type MinStackIntNode struct {
	data        int
	next        *MinStackIntNode
	subStackMin *MinStackIntNode
}

func NewMinStackIntNode(data int) *MinStackIntNode {
	return &MinStackIntNode{data: data, next: nil, subStackMin: nil}
}

type IntMinStack struct {
	top *MinStackIntNode
}

func NewIntMinStack() *IntMinStack {
	return &IntMinStack{top: nil}
}

func (stack *IntMinStack) isEmpty() bool {
	return stack.top == nil
}

func (stack *IntMinStack) pop() (int, error) {
	if stack.isEmpty() {
		return 0, errors.New("emtpy stack")
	}

	item := stack.top.data
	stack.top = stack.top.next
	return item, nil
}

func (stack *IntMinStack) peek() (int, error) {
	if stack.isEmpty() {
		return 0, errors.New("emtpy stack")
	}

	return stack.top.data, nil
}

func (stack *IntMinStack) push(data int) {
	node := NewMinStackIntNode(data)
	node.next = stack.top

	if stack.top == nil || node.data < stack.top.subStackMin.data {
		node.subStackMin = node
	} else {
		node.subStackMin = stack.top.subStackMin
	}

	stack.top = node

}

func (stack *IntMinStack) min() (int, error) {
	if stack.isEmpty() {
		return 0, errors.New("emtpy stack")
	}

	return stack.top.subStackMin.data, nil
}
