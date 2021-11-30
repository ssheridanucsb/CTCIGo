package trees

func subArrayInsert(array []int, root *BSTNode) {
	if len(array) == 0 {
		return
	}
	middle := len(array) / 2
	root = insertBST(root, array[middle])
	subArrayInsert(array[:middle], root)
	subArrayInsert(array[middle+1:], root)
}

func sortedArrayToBST(array []int) *BSTNode {
	if len(array) == 0 {
		return nil
	}
	middle := len(array) / 2
	root := &BSTNode{value: array[middle], left: nil, right: nil}
	subArrayInsert(array[:middle], root)
	subArrayInsert(array[middle+1:], root)
	return root
}
