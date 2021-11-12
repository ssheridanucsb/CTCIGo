package linkedlists

func kthToLast(head *IntNode, k int) *IntNode {
	//given the head of the linked list return the kthh to last element of the list
	if head == nil {
		return nil
	}

	if k < 0 {
		return nil
	}
	//find the length of the linked list
	len := 1
	n := head
	for n.next != nil {
		len++
		n = n.next
	}

	if k > len {
		return nil
	}

	k2l := len - k

	n = head
	i := 0
	for n.next != nil {
		if i == k2l {
			return n
		}
		i++
		n = n.next
	}
	return n
}
