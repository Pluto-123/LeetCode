package main

func twoSum(nums []int, target int) []int {
	mp := map[int]int{}
	ans := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		_, ok := mp[target-nums[i]]
		if ok {
			ans = append(ans, mp[target-nums[i]])
			ans = append(ans, i)
			break
		} else {
			mp[nums[i]] = i
		}
	}
	return ans
}

func main() {
	twoSum([]int{3, 2, 4}, 6)
}
