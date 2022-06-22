package main

import (
	"container/heap"
	"fmt"
	"math"
)

type MaxHeap []int

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Swap(i, j int) {
	slice := *h
	slice[i], slice[j] = slice[j], slice[i]
}

func (h *MaxHeap) Less(i, j int) bool {
	slice := *h
	return slice[i] > slice[j]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	result := old[len(old)-1]
	*h = old[:len(old)-1]
	return result
}

func minimumDeviation(nums []int) int {
	minimum := math.MaxInt32
	h := &MaxHeap{}
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 1 {
			nums[i] = nums[i] * 2
		}
		minimum = min(minimum, nums[i])
		heap.Push(h, nums[i])
	}
	fmt.Println(h)
	// dev := nums[len(nums)-1] - minimum
	// for nums[len(nums)-1]%2==0{
	//     curr := heap.Pop(&h).(int)
	//     fmt.Println(curr)
	//     if curr - minimum < dev{
	//         dev=curr-minimum
	//     }
	//     heap.Push(&h,curr/2)
	// }
	return 0
}
