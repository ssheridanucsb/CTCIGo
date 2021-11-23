package stacksandqueues

type QFromStacks struct {
	leftStack  *IntStack
	rightStack *IntStack
	left       bool
}

func NewQFromStacks() *QFromStacks {
	ls := NewIntStack()
	rs := NewIntStack()

	return &QFromStacks{leftStack: ls, rightStack: rs, left: true}
}

func (q *QFromStacks) isEmpty() bool {
	if q.left {
		return q.leftStack.isEmpty()
	}
	return q.rightStack.isEmpty()
}

func (q *QFromStacks) add(data int) {
	if !q.left {
		for {
			d, err := q.rightStack.pop()
			if err != nil {
				break
			}
			q.leftStack.push(d)
		}
		q.left = true
	}
	q.leftStack.push(data)
}

func (q *QFromStacks) remove() (int, error) {
	if q.left {
		for {
			d, err := q.leftStack.pop()
			if err != nil {
				break
			}
			q.rightStack.push(d)
		}
		q.left = false
	}

	return q.rightStack.pop()
}

func (q *QFromStacks) peek() (int, error) {
	if q.left {
		for {
			d, err := q.leftStack.pop()
			if err != nil {
				break
			}
			q.rightStack.push(d)
			q.left = false
		}
	}

	return q.rightStack.peek()
}
