package heap

import "fmt"

type Comparator func(x int, y int) bool

type Heap struct {
	vals []int
	comp Comparator
}

func (h *Heap) insert(value int) {
	h.vals = append(h.vals, value)
	h.heapify()
}

func (h *Heap) removeTop() {
	h.vals = h.vals[1:len(h.vals)]
	h.heapify()
}

func (h *Heap) heapify() {
	vals := h.vals
	length := len(vals)
	for i := length - 1; i >= 0; i-- {
		if i == 0 {
			return
		}
		parentIndex := (i - 1) / 2
		if !h.comp(vals[parentIndex], vals[i]) {
			swap(&vals[parentIndex], &vals[i])
		}
	}
}

func swap(first *int, second *int) {
	temp := *second
	*second = *first
	*first = temp
}

func (h *Heap) top() int {
	return h.vals[0]
}

func lessComparator(x int, y int) bool {
	return x <= y
}

func greaterComparator(x int, y int) bool {
	return x >= y
}

func TestHeap() {
	minHeap := Heap{}
	minHeap.comp = lessComparator
	minHeap.insert(100)
	minHeap.insert(2)
	minHeap.insert(200)
	minHeap.insert(1)
	minHeap.insert(4)
	minHeap.insert(5)
	printTop(&minHeap)
	minHeap.removeTop()
	printTop(&minHeap)

	maxHeap := Heap{}
	maxHeap.comp = greaterComparator
	maxHeap.insert(100)
	maxHeap.insert(2)
	maxHeap.insert(200)
	maxHeap.insert(1)
	maxHeap.insert(4)
	maxHeap.insert(5)
	printTop(&maxHeap)
	maxHeap.removeTop()
	printTop(&maxHeap)
}

func printTop(minHeap *Heap) {
	fmt.Printf("Top value %d \n", minHeap.top())
}
