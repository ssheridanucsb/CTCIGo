package linkedlists

func deleteMiddleNode(middleNode *IntNode) {
	//shift all following nodes back one and remove last node
	if middleNode == nil || middleNode.next == nil {
		return
	}

	currentNode := middleNode
	nextNode := middleNode.next

	for nextNode != nil {
		currentNode.data = nextNode.data
		currentNode = nextNode
		nextNode = nextNode.next
	}

	currentNode.next = nil
}
