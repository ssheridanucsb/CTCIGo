package linkedlists

import "testing"

func TestLoopDetectionTrue(t *testing.T) {
	head := NewIntNode(0)
	head.next = NewIntNode(1)
	head.next.next = NewIntNode(2)
	head.next.next.next = head.next

	l := loopDetection(head)

	if l == nil {
		t.Error("Expected loop node but recieved nil")
	}

	if l != nil && l.data != 1 {
		t.Errorf("Expected loop node to be 1 but got %d", l.data)
	}
}

func TestLoopDetectionFalse(t *testing.T) {
	head := NewIntNode(0)
	head.appendToTail(1)
	head.appendToTail(2)

	l := loopDetection(head)

	if l != nil {
		t.Errorf("Expected nil return but recieved %d", l.data)
	}
}
