package stacksandqueues

import "testing"

func TestQueue(t *testing.T) {
	q := NewIntQueue()

	if !q.isEmpty() {
		t.Error("Expected queue to be empty")
	}

	q.add(1)
	q.add(2)

	p, err := q.remove()
	if err != nil {
		t.Error("expected no error on remove")
	}

	if p != 1 {
		t.Error("expected p to be 1")
	}

	p, err = q.peek()

	if err != nil {
		t.Error("expected no error on peek")
	}

	if p != 2 {
		t.Error("expected to see 2")
	}
}
