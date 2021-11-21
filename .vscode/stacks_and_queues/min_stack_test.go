package stacksandqueues

import "testing"

func TestMinStack(t *testing.T) {
	minStack := NewIntMinStack()

	if !minStack.isEmpty() {
		t.Error("Expected stack to be empty")
	}

	minStack.push(5)
	minStack.push(3)
	minStack.push(4)
	minStack.push(2)
	minStack.push(6)

	d, _ := minStack.min()
	if d != 2 {
		t.Errorf("Expected min of 2 but recieved %d", d)
	}
	d, _ = minStack.peek()
	if d != 6 {
		t.Errorf("Expectd to peek 6 but recieved %d", d)

	}

	minStack.pop()
	minStack.pop()

	d, _ = minStack.min()
	if d != 3 {
		t.Errorf("Expected new min to be 3 but recieved %d", d)
	}
}
