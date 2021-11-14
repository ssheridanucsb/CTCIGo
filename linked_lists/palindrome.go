package linkedlists

func isPalindrome(head *IntNode) bool {
	//takes linked list and checkes if its a palindrome

	var endNode *IntNode = nil
	startNode := head
	n := head

	for {

		for n.next != endNode {
			n = n.next
		}

		if startNode.data != n.data {
			return false
		}
		endNode = n
		startNode = startNode.next
		n = startNode

		if startNode.next == endNode {
			//for even length lists
			return startNode.next.data == endNode.data
		}
		if startNode == endNode {
			//for odd length lists
			break
		}

	}
	return true
}
