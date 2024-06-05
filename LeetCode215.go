package main

func findKthLargest(nums []int, k int) int {
	buildMaxHeap215(nums)
	for i := 0; i < k-1; i++ {
		nums[0] = nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		heapity215(nums, 0)
	}
	return nums[0]
}

func buildMaxHeap215(a []int) {
	for i := len(a) / 2; i >= 0; i-- {
		heapity215(a, i)
	}
}

func heapity215(a []int, i int) {
	l, r, largest := i*2+1, i*2+2, i
	if l < len(a) && a[l] > a[largest] {
		largest = l
	}
	if r < len(a) && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		heapity215(a, largest)
	}
}
