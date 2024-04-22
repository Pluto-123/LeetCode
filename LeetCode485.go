package main

func findMaxConsecutiveOnes(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			length := calc(nums, i)
			ans = max(ans, length)
			i = i + length - 1
		}
	}
	return ans
}

func calc(nums []int, index int) int {
	length := 0
	for i := index; i < len(nums); i++ {
		if nums[i] == 1 {
			length++
		} else {
			break
		}
	}
	return length
}
