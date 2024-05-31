package main

func moveZeroes(nums []int) {
	// 同向双指针
	i := 0
	j := 0
	for i < len(nums) && j < len(nums) {
		// i --> 0
		for nums[i] != 0 {
			i++
			if i == len(nums) {
				return
			}
		}
		// j --> 1
		for nums[j] == 0 {
			j++
			if j == len(nums) {
				return
			}
		}
		if j > i {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			j = i + 1
		}
	}
}

func main() {
	moveZeroes([]int{1, 0})
}
