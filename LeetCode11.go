package main

func maxArea(height []int) int {
	res := 0
	left := 0
	right := len(height) - 1
	for left < right {
		res = max(res, (right-left)*min(height[left], height[right]))
		if height[left] <= height[right] {
			left++
			continue
		}
		if height[left] > height[right] {
			right--
			continue
		}
	}
	return res
}
