package stacksandqueues

import "testing"

func TestQFromStacks(t *testing.T) {
	q := NewQFromStacks()
	if !q.isEmpty() {
		t.Error("Expected empty stack")
	}

	q.add(1)
	q.add(2)

	d, _ := q.remove()
	if d != 1 {
		t.Errorf("Expected 2 but recieved %d", d)
	}

	q.add(3)

	d, _ = q.remove()
	if d != 2 {
		t.Errorf("Expected 3 but recieved %d", d)
	}

	d, _ = q.peek()
	if d != 3 {
		t.Errorf("Expected 2 but recieved %d", d)
	}

	d, _ = q.remove()
	if d != 3 {
		t.Errorf("Expected 2 but recieved %d", d)
	}

}
