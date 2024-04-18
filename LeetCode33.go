package main

func search(nums []int, target int) int {
	res := binarySearch(nums, target)
	return res
}

func binarySearch(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[low] <= nums[mid] {
			// 左边有序
			if target >= nums[low] && target <= nums[mid] {
				// target在左边的有序区间
				high = mid - 1
				continue
			} else {
				// target在右边的无序区间
				low = mid + 1
				continue
			}
		}
		if nums[mid] <= nums[high] {
			// 右边有序
			if target >= nums[mid] && target <= nums[high] {
				// target在右边的有序区间
				low = mid + 1
				continue
			} else {
				high = mid - 1
				continue
			}
		}

	}

	return -1
}

//
//func main() {
//	s := search([]int{4, 5, 6, 7, 0, 1, 2}, 0)
//	fmt.Println(s)
//}
