package linkedlists

import "testing"

func TestIntersectionTrue(t *testing.T) {
	h1 := NewIntNode(1)
	h1.appendToTail(2)

	h2 := h1
	h2.appendToTail(3)

	i := intersection(h1, h2)

	if i == nil {
		t.Error("Expected Intersection but recieved nil")
	}
}

func TestIntersectionFalse(t *testing.T) {
	h1 := NewIntNode(1)
	h1.appendToTail(2)

	h2 := NewIntNode(4)
	h2.appendToTail(3)

	i := intersection(h1, h2)

	if i != nil {
		t.Error("Expected nil but recieved intersection")
	}
}
