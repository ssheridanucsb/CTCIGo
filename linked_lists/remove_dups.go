package linkedlists

func removeDups(head *IntNode) *IntNode {
	//removes duplicates from linked list
	// brute force solution using map buffer
	if head.next == nil {
		return head
	}

	dupCount := make(map[int]int)

	n := head
	for n.next != nil {
		dupCount[n.data]++
		n = n.next
	}

	for key, val := range dupCount {
		if val > 1 {
			for i := 0; i < val-1; i++ {
				head = deleteNode(head, key)
			}
		}
	}
	return head
}

func removeDupsNoBuff(head *IntNode) {
	if head == nil || head.next == nil {
		return
	}

	n := head

	for n.next != nil {
		p := n
		d := n.data
		for p.next != nil {
			if p.next.data == d {
				p.next = p.next.next
				continue
			}
			p = p.next
		}
		n = n.next
	}
}
