// This is heap data structure realised with golang generics.
// Attractiveness of this package is it simplicity to use,
// because need to declare only one function -- Comparator.

package heap

type Heap[T comparable] struct {
	heap       []T
	Comparator func(T, T) bool
}

// siftUp find right place for element moving it from bottom to top layers
// by swapping parent with their child.
func (h *Heap[T]) siftUp(i int) {
	for i > 0 && !h.Comparator(h.heap[(i-1)/2], h.heap[i]) {
		h.heap[(i-1)/2], h.heap[i] = h.heap[i], h.heap[(i-1)/2]
		i = (i - 1) / 2
	}
}

// siftDown find right place for element moving it from top to bottom layers
// by swapping parent with their child.
func (h *Heap[T]) siftDown(i int) {
	for i < len(h.heap) {
		maxIndex := i

		leftChild := 2*i + 1
		if leftChild < len(h.heap) && h.Comparator(h.heap[leftChild], h.heap[maxIndex]) {
			maxIndex = leftChild
		}

		rightChild := 2*i + 2
		if rightChild < len(h.heap) && h.Comparator(h.heap[rightChild], h.heap[maxIndex]) {
			maxIndex = rightChild
		}

		if i != maxIndex {
			h.heap[i], h.heap[maxIndex] = h.heap[maxIndex], h.heap[i]
			i = maxIndex
		} else {
			break
		}
	}

}

// Insert allow to insert one ore more elemnts into heap.
func (h *Heap[T]) Insert(p ...T) {
	if len(p) > 1 {
		h.heap = append(h.heap, p...)
		for i := (len(h.heap) - 1) / 2; i >= 0; i-- {
			h.siftDown(i)
		}
	} else {
		h.heap = append(h.heap, p[0])
		h.siftUp(len(h.heap) - 1)
	}
}

// Top return the top element of the heap.
func (h *Heap[T]) Top() T {
	return h.heap[0]
}

// ExtractTop retrun top element and delete them from heap.
func (h *Heap[T]) ExtractTop() T {
	result := h.heap[0]
	h.heap[0] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	h.siftDown(0)
	return result
}

// Remove erase element which has index i from heap.
func (h *Heap[T]) Remove(i int) {
	h.heap[i] = h.Top()
	h.siftUp(i)
	h.ExtractTop()
}

// ChangePriority set for i-th elemnt p value.
func (h *Heap[T]) ChangePriority(i int, p T) {
	old := h.heap[i]
	h.heap[i] = p
	if !h.Comparator(old, p) {
		h.siftUp(i)
	} else {
		h.siftDown(i)
	}
}

// GetHeap return heap as it is.
func (h *Heap[T]) GetHeap() []T {
	return h.heap
}

// Size return size of the heap
func (h *Heap[T]) Size() int {
	return len(h.heap)
}
