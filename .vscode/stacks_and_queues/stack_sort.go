package stacksandqueues

func stackSort(stack *IntStack) *IntStack {
	//sort stack with min elements at top using only another stack
	if stack.isEmpty() {
		return stack
	}

	tmpStack := NewIntStack()

	for {
		p, err := stack.pop()
		if err != nil {
			break
		}

		for {
			peek, err := tmpStack.peek()
			if err != nil || peek <= p {
				tmpStack.push(p)
				break
			} else {
				pop, _ := tmpStack.pop()
				stack.push(pop)
			}
		}
	}

	for {
		p, err := tmpStack.pop()
		if err != nil {
			break
		}

		stack.push(p)
	}

	return stack
}
