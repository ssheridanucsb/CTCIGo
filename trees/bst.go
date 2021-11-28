package trees

type BSTNode struct {
	value int
	left  *BSTNode
	right *BSTNode
}

func insertBST(root *BSTNode, value int) *BSTNode {
	//insert node into bst
	if root == nil {
		return &BSTNode{value: value, left: nil, right: nil}
	}
	if value > root.value {
		root.right = insertBST(root.right, value)
	} else {
		root.left = insertBST(root.left, value)
	}

	return root
}

func deleteBSTNode(root *BSTNode, value int) *BSTNode {
	//delete node fomr bst

	//base case
	if root == nil {
		return nil
	}

	if root.value > value {
		//go left
		root.left = deleteBSTNode(root.left, value)
	} else if root.value < value {
		//go right
		root.right = deleteBSTNode(root.right, value)
	} else {
		// delete node
		if root.right == nil {
			return root.left
		} else if root.left == nil {
			return root.right
		}
		rightSmallest := root.right
		for rightSmallest.left != nil {
			rightSmallest = rightSmallest.left
		}
		rightSmallest.left = root.left
		return root.right
	}
	return root
}

func searchBST(root *BSTNode, value int) *BSTNode {
	if root == nil || root.value == value {
		return root
	}

	if root.value < value {
		return searchBST(root.right, value)
	}
	return searchBST(root.left, value)
}

func traverse(root *BSTNode, ouput chan int, method func(node *BSTNode, ouput chan int)) {
	defer close(ouput)
	method(root, ouput)
}

func inOrderTraversal(node *BSTNode, output chan int) {
	if node != nil {
		inOrderTraversal(node.left, output)
		output <- node.value
		inOrderTraversal(node.right, output)
	}
}

func preOrderTraversal(node *BSTNode, output chan int) {
	if node != nil {
		output <- node.value
		preOrderTraversal(node.left, output)
		preOrderTraversal(node.right, output)
	}
}

func postOrderTraversal(node *BSTNode, output chan int) {
	if node != nil {
		postOrderTraversal(node.left, output)
		postOrderTraversal(node.right, output)
		output <- node.value
	}
}
