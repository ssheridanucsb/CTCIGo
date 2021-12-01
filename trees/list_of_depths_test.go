package trees

import "testing"

func TestListOfDepths(t *testing.T) {
	bst := constructTestTree()

	dm := listOfDepths(bst)

	l1 := dm[0]
	if l1.bstNode.value != 10 {
		t.Error("unexpeccted root value")
	}

	if len(dm) != 3 {
		t.Error("incorrect number of levels")
	}

}
