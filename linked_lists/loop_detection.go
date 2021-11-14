package linkedlists

func loopDetection(head *IntNode) *IntNode {
	seen := make(map[*IntNode]int)
	n := head
	for n != nil {
		_, ok := seen[n]
		if ok {
			return n
		}
		seen[n]++
		n = n.next
	}
	return nil
}
