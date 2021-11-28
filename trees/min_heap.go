package trees

type MinHeap struct {
	heap []int
}

func NewMinHeap() *MinHeap {
	h := make([]int, 0)
	return &MinHeap{heap: h}
}

func parent(i int) int {
	return int((i - 1) / 2)
}

func leftChild(i int) int {
	return 2*i + 1
}
func rightChild(i int) int {
	return 2*i + 2
}

func (h *MinHeap) upHeap() {
	l := len(h.heap) - 1
	for {
		p := parent(l)
		//if we've reached the top or the heap property is satisfied break
		if p < 0 || h.heap[p] <= h.heap[l] {
			break
		} else {
			//otherwise swap with the parent
			h.heap[p], h.heap[l] = h.heap[l], h.heap[p]
			l = p
		}
	}
}

func (h *MinHeap) downHeap(index int, length int) {
	left := leftChild(index)
	right := rightChild(index)
	smallest := index

	if left <= length && h.heap[left] < h.heap[index] {
		smallest = left
	} else if right <= left && h.heap[right] < h.heap[left] {
		smallest = right
	}

	if smallest != index {
		h.heap[index], h.heap[smallest] = h.heap[smallest], h.heap[index]
		h.downHeap(smallest, length)
	}
}

func (h *MinHeap) insert(value int) {
	h.heap = append(h.heap, value)
	h.upHeap()
}

func (h *MinHeap) extractMin() int {
	min := h.heap[0]
	last := len(h.heap) - 1
	h.heap[0], h.heap[last] = h.heap[last], h.heap[0]
	h.downHeap(0, last+1)
	return min
}

func (h *MinHeap) delete(value int) {
	for idx, v := range h.heap {
		if v == value {
			//swap element with last element
			last := len(h.heap) - 1
			tmp := h.heap[last]
			h.heap[idx], h.heap[last] = h.heap[last], h.heap[idx]
			h.heap = h.heap[:last]

			if tmp < v {
				h.upHeap()
			} else {
				h.downHeap(idx, last)
			}
			break
		}
	}
}
