package main

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	ans := make([]int, 0)

	h := &IHeap{}
	heap.Init(h)

	return ans
}

type IHeap [][2]int

func (h IHeap) Len() int {
	return len(h)
}
func (h IHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}
func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
