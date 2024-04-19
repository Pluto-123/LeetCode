package main

func twoSum(numbers []int, target int) []int {
	low := 0
	high := len(numbers) - 1
	for low <= high {
		if numbers[low]+numbers[high] == target {
			return []int{low + 1, high + 1}
		}
		if numbers[low]+numbers[high] < target {
			low++
		}

		if numbers[low]+numbers[high] > target {
			high--
		}
	}
	return nil
}
