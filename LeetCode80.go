package main

func removeDuplicates(nums []int) int {
	// 修改数组的前n个元素
	// 有序数组
	ans := 0
	for i := 0; i < len(nums)-2-ans; i++ {
		j := i + 1
		z := j + 1
		if nums[i] == nums[j] && nums[j] == nums[z] {
			preMove(nums, z)
			ans++
			i--
		}
	}
	return len(nums) - ans
}

func preMove(nums []int, z int) {
	for i := z; i < len(nums)-1; i++ {
		nums[i] = nums[i+1]
	}
}

func main() {
	x := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	removeDuplicates(x)
	//fmt.Println(x)
}
