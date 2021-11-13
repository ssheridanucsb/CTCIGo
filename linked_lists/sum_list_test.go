package linkedlists

import "testing"

func TestSumList(t *testing.T) {
	//730 + 629 = 1359
	head1 := NewIntNode(0)
	head1.appendToTail(3)
	head1.appendToTail(7)

	head2 := NewIntNode(9)
	head2.appendToTail(2)
	head2.appendToTail(6)

	s := sumList(head1, head2)
	if s != 1359 {
		t.Errorf("expected 1359 but recieved %d", s)
	}
}

func TestSumListForward(t *testing.T) {
	//730 + 629 = 1359

	head1 := NewIntNode(7)
	head1.appendToTail(3)
	head1.appendToTail(0)

	head2 := NewIntNode(6)
	head2.appendToTail(2)
	head2.appendToTail(9)

	s := sumListForward(head1, head2)
	if s != 1359 {
		t.Errorf("expected 1359 but recieved %d", s)
	}
}
