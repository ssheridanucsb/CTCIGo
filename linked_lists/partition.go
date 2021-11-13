package linkedlists

func partition(head *IntNode, partition int) *IntNode {
	//returns head of new partitioned list
	if head == nil || head.next == nil {
		return nil
	}

	var leftPartition *IntNode = nil
	var rightPartition *IntNode = nil

	var leftTail *IntNode = nil
	var rightTail *IntNode = nil

	n := head
	for n != nil {
		if n.data < partition {
			if leftPartition == nil {
				leftPartition = NewIntNode(n.data)
				leftTail = leftPartition
			} else {
				leftTail.next = NewIntNode(n.data)
				leftTail = leftTail.next
			}
		} else {
			if rightPartition == nil {
				rightPartition = NewIntNode(n.data)
				rightTail = rightPartition
			} else {
				rightTail.next = NewIntNode(n.data)
				rightTail = rightTail.next
			}
		}
		n = n.next
	}
	if leftPartition == nil {
		return rightPartition
	}

	leftTail.next = rightPartition
	return leftPartition
}
