package stacksandqueues

import "errors"

type QueueIntNode struct {
	data int
	next *QueueIntNode
}

func NewQueueIntNode(data int) *QueueIntNode {
	return &QueueIntNode{data: data, next: nil}
}

type IntQueue struct {
	first *QueueIntNode
	last  *QueueIntNode
}

func NewIntQueue() *IntQueue {
	return &IntQueue{first: nil, last: nil}
}

func (q *IntQueue) add(item int) {
	t := NewQueueIntNode(item)
	if q.last != nil {
		q.last.next = t
	}
	q.last = t
	if q.first == nil {
		q.first = q.last
	}

}

func (q *IntQueue) remove() (int, error) {
	if q.first == nil {
		return 0, errors.New("empty queue")
	}

	d := q.first.data
	q.first = q.first.next
	if q.first == nil {
		q.last = nil
	}
	return d, nil
}

func (q *IntQueue) peek() (int, error) {
	if q.first == nil {
		return 0, errors.New("emtpy queue")
	}
	return q.first.data, nil

}

func (q *IntQueue) isEmpty() bool {
	return q.first == nil
}
