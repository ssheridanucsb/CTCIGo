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
				leftPartition = n
				leftTail = leftPartition
			} else {
				leftTail.next = n
				leftTail = leftTail.next
			}
			n = n.next
			leftTail.next = nil
		} else {
			if rightPartition == nil {
				rightPartition = n
				rightTail = rightPartition
			} else {
				rightTail.next = n
				rightTail = rightTail.next
			}
			n = n.next
			rightTail.next = nil
		}
	}
	if leftPartition == nil {
		return rightPartition
	}

	leftTail.next = rightPartition
	return leftPartition
}
