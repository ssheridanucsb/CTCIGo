package linkedlists

import "math"

func sumList(head1 *IntNode, head2 *IntNode) int {
	//takes linked lists representing numbers is reverse order
	n1 := head1
	n2 := head2

	sum := 0
	place := 1
	for {
		if n1 == nil && n2 == nil {
			break
		}

		if n1 != nil {
			sum += n1.data * place
			n1 = n1.next
		}

		if n2 != nil {
			sum += n2.data * place
			n2 = n2.next
		}
		place *= 10
	}
	return sum
}

func sumListForward(head1 *IntNode, head2 *IntNode) int {
	//takes linked lists representing numbers in forward order and sums them
	len1 := 0
	len2 := 0

	n1 := head1
	n2 := head2

	for n1 != nil {
		len1++
		n1 = n1.next
	}

	for n2 != nil {
		len2++
		n2 = n2.next
	}

	placeOne := int(math.Pow(10, float64(len1-1)))
	placeTwo := int(math.Pow(10, float64(len2-1)))

	sum := 0

	n1 = head1
	n2 = head2

	for n1 != nil {
		sum += n1.data * placeOne
		placeOne = placeOne / 10
		n1 = n1.next
	}

	for n2 != nil {
		sum += n2.data * placeTwo
		placeTwo = placeTwo / 10
		n2 = n2.next
	}
	return sum
}
