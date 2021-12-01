package trees

type bstListNode struct {
	bstNode *BSTNode
	next    *bstListNode
}

func (n *bstListNode) appendToTail(node *BSTNode) {
	endNode := &bstListNode{bstNode: node, next: nil}
	cur := n
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = endNode
}

func depth(node *BSTNode, d int, dm map[int]*bstListNode) {
	if node == nil {
		return
	}
	n, ok := dm[d]
	if !ok {
		dm[d] = &bstListNode{bstNode: node, next: nil}
	} else {
		n.appendToTail(node)
	}
	d++
	depth(node.left, d, dm)
	depth(node.right, d, dm)
}

func listOfDepths(root *BSTNode) map[int]*bstListNode {
	//find all nodes at each depth and return a linked list for all the nodes at each depth
	if root == nil {
		return nil
	}
	dm := make(map[int]*bstListNode)
	depth(root, 0, dm)
	return dm
}
