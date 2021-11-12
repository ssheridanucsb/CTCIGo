package linkedlists

import "testing"

func Test2ndToLast(t *testing.T) {
	head := NewIntNode(0)
	head.appendToTail(1)
	head.appendToTail(2)
	head.appendToTail(3)

	k2l := kthToLast(head, 2)

	if k2l.data != 2 {
		t.Errorf("Expected the second to last element to be 2 but go %d", k2l.data)
	}
}

func Test3rdToLast(t *testing.T) {
	head := NewIntNode(0)
	head.appendToTail(1)
	head.appendToTail(2)
	head.appendToTail(3)

	k2l := kthToLast(head, 3)

	if k2l.data != 1 {
		t.Errorf("Expected the second to last element to be 1 but go %d", k2l.data)
	}
}

func TestToLast(t *testing.T) {
	head := NewIntNode(0)
	head.appendToTail(1)
	head.appendToTail(2)
	head.appendToTail(3)

	k2l := kthToLast(head, 0)

	if k2l.data != 3 {
		t.Errorf("Expected the second to last element to be 3 but got %d", k2l.data)
	}
}
