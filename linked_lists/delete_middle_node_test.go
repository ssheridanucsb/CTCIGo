package linkedlists

import "testing"

func TestDeleteMiddleNode(t *testing.T) {
	head := NewIntNode(0)
	head.appendToTail(1)
	head.appendToTail(2) // middle node
	head.appendToTail(3)

	middleNode := head.next.next
	deleteMiddleNode(middleNode)
	if head.data != 0 {
		t.Errorf("Expected head to be 0 but got %d", head.data)
	}
	if head.next.data != 1 {
		t.Errorf("Expected second node to be 1 but got %d", head.next.data)
	}
	if head.next.next.data != 3 {
		t.Errorf("Expected Tail to be 3 but got %d", head.next.next.data)
	}
}

func TestDeleteMiddleNodeLongInput(t *testing.T) {
	head := NewIntNode(0)
	head.appendToTail(1)
	head.appendToTail(2) // middle node
	head.appendToTail(3)
	head.appendToTail(8)
	head.appendToTail(4)
	head.appendToTail(5)
	head.appendToTail(6)

	middleNode := head.next.next.next.next

	deleteMiddleNode(middleNode)
	n := head
	for i := 0; i < 7; i++ {
		if n.data != i {
			t.Errorf("Incorrect node, received %d but expeted %d", n.data, i)
		}
		n = n.next
	}
}
