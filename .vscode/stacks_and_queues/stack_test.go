package stacksandqueues

import "testing"

func TestStack(t *testing.T) {
	s := NewIntStack()

	if !s.isEmpty() {
		t.Error("Expected stack to be empty")
	}

	s.push(1)
	if s.isEmpty() {
		t.Error("Expected stack to not be empty")
	}

	s.push(2)

	p, err := s.pop()

	if err != nil {
		t.Error("Did not expect error when popping")
	}
	if p != 2 {
		t.Error("Expected to pop 2")
	}

	p, err = s.peek()
	if err != nil {
		t.Error("Did not expect error when peeking")
	}

	if p != 1 {
		t.Error("Expected to peek 1")
	}
}
