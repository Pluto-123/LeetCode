package main

func numOfSubarrays(arr []int, k int, threshold int) int {
	left := 0
	right := k - 1
	ans := 0
	for right < len(arr) {
		if avg(arr, left, right) >= threshold {
			ans++
		}
		left++
		right++
	}
	return ans
}

func avg(arr []int, left int, right int) int {
	sum := 0
	for i := left; i <= right; i++ {
		sum += arr[i]
	}
	return sum / (right - left + 1)

}
