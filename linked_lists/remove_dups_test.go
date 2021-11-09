package linkedlists

import "testing"

func TestRemoveDup(t *testing.T) {
	head := NewIntNode(0)
	head.appendToTail(1)
	head.appendToTail(1)
	head.appendToTail(1)
	head.appendToTail(2)

	head = removeDups(head)

	if head.next.data != 1 || head.next.next.data != 2 {
		t.Error("Failed to remove duplicates")
	}
}

func TestRemoveDupNoBuff(t *testing.T) {
	head := NewIntNode(0)
	head.appendToTail(1)
	head.appendToTail(1)
	head.appendToTail(1)
	head.appendToTail(2)

	removeDupsNoBuff(head)

	if head.next.data != 1 || head.next.next.data != 2 {
		t.Error("Failed to remove duplicates")
	}
}
