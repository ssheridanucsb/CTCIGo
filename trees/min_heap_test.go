package trees

import "testing"

func TestChildren(t *testing.T) {

	l := leftChild(1)
	r := rightChild(1)
	if l != 3 {
		t.Error("expectd 3")
	}
	if r != 4 {
		t.Error("expected 4")
	}
	if parent(l) != 1 {
		t.Error("incorrect parent of left child")
	}

	if parent(r) != 1 {
		t.Error("incorrect parent of right child")
	}
}

func TestMinHeap(t *testing.T) {
	heap := NewMinHeap()
	heap.insert(5)
	heap.insert(4)
	heap.insert(3)
	heap.insert(2)
	heap.insert(1)

	heap.delete(2)

	e := heap.extractMin()
	if e != 1 {
		t.Errorf("expected 1 but recieved %d", e)
	}

	e = heap.extractMin()
	if e != 3 {
		t.Errorf("expected 1 but recieved %d", e)
	}

}
