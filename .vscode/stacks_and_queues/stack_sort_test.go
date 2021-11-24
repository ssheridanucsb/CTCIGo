package stacksandqueues

import "testing"

func TestStackSort(t *testing.T) {
	s := NewIntStack()

	s.push(1)
	s.push(2)
	s.push(3)
	s.push(4)
	s.push(5)

	sorted := stackSort(s)

	for i := 1; i < 6; i++ {
		p, _ := sorted.pop()
		if p != i {
			t.Errorf("Expected %d, but recieved %d", i, p)
		}
	}
}

func TestStackSortComplex(t *testing.T) {
	s := NewIntStack()

	s.push(5)
	s.push(2)
	s.push(3)
	s.push(4)
	s.push(1)

	sorted := stackSort(s)

	for i := 1; i < 6; i++ {
		p, _ := sorted.pop()
		if p != i {
			t.Errorf("Expected %d, but recieved %d", i, p)
		}
	}
}
