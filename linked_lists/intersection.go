package linkedlists

func intersection(headOne *IntNode, headTwo *IntNode) *IntNode {
	//takes two linked lists and returns the intersecting node
	if headOne == nil || headTwo == nil {
		return nil
	}

	nodes := make(map[*IntNode]int)

	n := headOne
	for n != nil {
		nodes[n]++
		n = n.next
	}

	n = headTwo
	for n != nil {
		nodes[n]++
		n = n.next
	}

	for key, val := range nodes {
		if val > 1 {
			return key
		}
	}
	return nil
}
