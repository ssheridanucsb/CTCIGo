package linkedlists

import "testing"

func TestNewIntNode(t *testing.T) {

	n := NewIntNode(40)
	if n.data != 40 || n.next != nil {
		t.Error("Node was not initialized correctly")
	}
}

func TestAppendNode(t *testing.T) {
	head := NewIntNode(0)
	head.appendToTail(1)
	head.appendToTail(2)

	n := head
	for i := 0; i < 3; i++ {
		if n.data != i {
			t.Errorf("Expected data to equal %d", i)
		}
		n = n.next
	}
}

func TestDeleteNodeFromTail(t *testing.T) {
	head := NewIntNode(0)
	head.appendToTail(1)

	head = deleteNode(head, 1)

	if head.next != nil {
		t.Error("Expected head.next to be nil")
	}
}

func TestDeleteNodeFromHead(t *testing.T) {
	head := NewIntNode(0)
	head.appendToTail(1)

	head = deleteNode(head, 0)
	if head.data != 1 || head.next != nil {
		t.Error("Expected head.data to be 1 and head.next to be nil")
	}
}

func TestDeleteNodeFromMiddle(t *testing.T) {
	head := NewIntNode(0)
	head.appendToTail(1)
	head.appendToTail(2)

	head = deleteNode(head, 1)

	if head.data != 0 || head.next.data != 2 {
		t.Error("Expected 1 to be delted")
	}
}
