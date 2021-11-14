package linkedlists

import "testing"

func TestIsPalindromeOddTrue(t *testing.T) {
	head := NewIntNode(3)
	head.appendToTail(1)
	head.appendToTail(2)
	head.appendToTail(1)
	head.appendToTail(3)

	if !isPalindrome(head) {
		t.Error("Expected true but got false")
	}
}

func TestIsPalindromeEvenTrue(t *testing.T) {
	head := NewIntNode(3)
	head.appendToTail(1)
	head.appendToTail(2)
	head.appendToTail(2)
	head.appendToTail(1)
	head.appendToTail(3)

	if !isPalindrome(head) {
		t.Error("Expected true but got false")
	}
}

func TestIsPalindromeOddFalse(t *testing.T) {
	head := NewIntNode(0)
	head.appendToTail(1)
	head.appendToTail(2)

	if isPalindrome(head) {
		t.Error("Expected false but got true")
	}
}

func TestIsPalindromeEvenFalse(t *testing.T) {
	head := NewIntNode(0)
	head.appendToTail(1)
	head.appendToTail(2)
	head.appendToTail(3)

	if isPalindrome(head) {
		t.Error("Expected false but got true")
	}
}
