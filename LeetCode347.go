package main

import (
	"fmt"
)

func topKFrequent2(nums []int, k int) []int {
	res := make([]int, 0)
	mp := map[int]int{}
	d := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		_, ok := mp[nums[i]]
		if ok {
			mp[nums[i]]++
		} else {
			mp[nums[i]] = 1
		}
	}
	for key, value := range mp {
		d = append(d, []int{key, value})
	}
	quickSort4(d, 0, len(d)-1)
	fmt.Println(d)
	for i := 0; i < k; i++ {
		res = append(res, d[(len(d) - 1)][0])
		d = d[:len(d)-1]
	}
	return res
}

func quickSort4(d [][]int, low int, high int) {
	if low > high {
		return
	}
	partition := partition4(d, low, high)
	quickSort4(d, low, partition-1)
	quickSort4(d, partition+1, high)
}

func partition4(d [][]int, low int, high int) int {
	i := low
	j := high
	for i < j {
		for i < j && d[j][1] >= d[low][1] {
			j--
		}
		for i < j && d[i][1] <= d[low][1] {
			i++
		}
		d[i], d[j] = d[j], d[i]
	}
	d[i], d[low] = d[low], d[i]
	return i

}

// 添加一个大根堆的解法
func topKFrequent(nums []int, k int) []int {
	mp := make(map[int]int)
	temp := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		_, ok := mp[nums[i]]
		if ok {
			mp[nums[i]]++
		} else {
			mp[nums[i]] = 1
		}
	}

	for key, value := range mp {
		temp = append(temp, []int{key, value})
	}
	buildMaxHeap347(temp)
	res := make([]int, 0)
	for i := 0; i < k; i++ {
		res = append(res, temp[0][0])
		temp[0] = temp[len(temp)-1]
		temp = temp[:len(temp)-1]
		heapity347(temp, 0)
	}
	return res
}

func buildMaxHeap347(nums [][]int) {
	for i := len(nums) / 2; i >= 0; i-- {
		heapity347(nums, i)
	}
}

func heapity347(nums [][]int, i int) {
	l, r, largest := i*2+1, i*2+2, i
	if l < len(nums) && nums[l][1] > nums[largest][1] {
		largest = l
	}
	if r < len(nums) && nums[r][1] > nums[largest][1] {
		largest = r
	}
	if largest != i {
		nums[largest], nums[i] = nums[i], nums[largest]
		heapity347(nums, largest)
	}
}

func main() {
	s := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	fmt.Println(s)
}
