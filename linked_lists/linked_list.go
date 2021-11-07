package linkedlists

type IntNode struct {
	next *IntNode
	data int
}

func NewIntNode(data int) *IntNode {

	node := IntNode{next: nil, data: data}
	return &node
}

func (n *IntNode) appendToTail(data int) {
	endNode := NewIntNode(data)
	cur := n
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = endNode
}

func deleteNode(head *IntNode, data int) *IntNode {
	//no head
	if head == nil {
		return nil
	}
	n := head
	if n.data == data {
		//moved head
		return head.next
	}

	for n.next != nil {
		if n.next.data == data {
			n.next = n.next.next
			//head didn't change
			return head
		}
		n = n.next
	}
	return head
}
