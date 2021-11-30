package trees

import (
	"testing"
)

func TestSortedArrayToBSTOdd(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	tree := sortedArrayToBST(a)

	c := make(chan int)
	go traverse(tree, c, inOrderTraversal)
	i := 0
	for o := range c {
		if o != a[i] {
			t.Errorf("Expected %d but recieved %d", a[i], o)
		}
		i++
	}
}

func TestSortedArrayToBSTEven(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	tree := sortedArrayToBST(a)

	c := make(chan int)
	go traverse(tree, c, inOrderTraversal)
	i := 0
	for o := range c {
		if o != a[i] {
			t.Errorf("Expected %d but recieved %d", a[i], o)
		}
		i++
	}
}
