package linkedlists

import "testing"

func TestPartition(t *testing.T) {
	head := NewIntNode(7)
	head.appendToTail(8)
	head.appendToTail(9)
	head.appendToTail(4)
	head.appendToTail(5)
	head.appendToTail(6)

	p := partition(head, 7)

	n := p
	for i := 4; i < 9; i++ {
		if n.data != i {
			t.Errorf("Expected %d but got %d", i, p.data)
		}
		n = n.next
	}

}
