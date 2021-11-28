package trees

import (
	"testing"
)

func constructTestTree() *BSTNode {
	root := &BSTNode{value: 10, left: nil, right: nil}
	root = insertBST(root, 15)
	root = insertBST(root, 14)
	root = insertBST(root, 16)
	root = insertBST(root, 5)
	root = insertBST(root, 4)
	root = insertBST(root, 6)
	return root
}

func TestInOrderTraversal(t *testing.T) {
	root := constructTestTree()

	c := make(chan int)

	go traverse(root, c, inOrderTraversal)
	inOrder := []int{4, 5, 6, 10, 14, 15, 16}
	index := 0
	for v := range c {
		expected := inOrder[index]
		if v != expected {
			t.Errorf("Expected %d but recieved %d", expected, v)
		}
		index++
	}
}

func TestPreOrderTraversal(t *testing.T) {
	root := constructTestTree()

	c := make(chan int)

	go traverse(root, c, preOrderTraversal)
	inOrder := []int{10, 5, 4, 6, 15, 14, 16}
	index := 0
	for v := range c {
		expected := inOrder[index]
		if v != expected {
			t.Errorf("Expected %d but recieved %d", expected, v)
		}
		index++
	}
}

func TestPostOrderTraversal(t *testing.T) {
	root := constructTestTree()

	c := make(chan int)

	go traverse(root, c, postOrderTraversal)
	inOrder := []int{4, 6, 5, 14, 16, 15, 10}
	index := 0
	for v := range c {
		expected := inOrder[index]
		if v != expected {
			t.Errorf("Expected %d but recieved %d", expected, v)
		}
		index++
	}
}

func TestSearch(t *testing.T) {
	root := constructTestTree()

	v := searchBST(root, 200)
	if v != nil {
		t.Error("Expect to not find 200")
	}

	v = searchBST(root, 15)
	if v.value != 15 {
		t.Errorf("expected 15 but recieved %d", v.value)
	}
}

func TestDelete(t *testing.T) {
	root := constructTestTree()
	root = deleteBSTNode(root, 15)

	s := searchBST(root, 15)
	if s != nil {
		t.Error("expected to not find 15")
	}
}
