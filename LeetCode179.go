package main

import (
	"fmt"
	"strconv"
	"strings"
)

// fail
func largestNumber(nums []int) string {
	// 做一个特殊的排序
	quickSort2(nums, 0, len(nums)-1)

	sb := strings.Builder{}
	for i := 0; i < len(nums); i++ {
		sb.WriteString(strconv.Itoa(nums[i]))
	}
	res := sb.String()

	return reverseString(res)
}

func quickSort2(nums []int, low int, high int) {
	if low >= high {
		return
	}

	partition := partition2(nums, low, high)
	quickSort2(nums, low, partition-1)
	quickSort2(nums, partition+1, high)
}

func partition2(nums []int, low int, high int) int {
	i := low
	j := high
	for i < j {
		for i < j && isBigger(nums[j], nums[low]) {
			j--
		}
		for i < j && isBigger(nums[low], nums[i]) {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[low], nums[i] = nums[i], nums[low]
	return i
}

// 包含 a>=b
func isBigger(a, b int) bool {

	if a == b || b == 0 {
		return true
	}
	if a == 0 {
		return false
	}
	// 比较高位
	for {
		a_high, a_times := getHigh(a)
		b_high, b_times := getHigh(b)

		//if a_times == 0 {
		//	return true
		//}
		//
		//if b_times == 0 {
		//	return false
		//}

		if a_high == b_high {
			a = a - (a_high * a_times)
			b = b - (b_high * b_times)
			continue
		} else {
			return a_high > b_high
		}

	}

	//取出
	return true
}

// 高位值 倍数
func getHigh(num int) (int, int) {
	// 取出最高位
	if num == 0 {
		return 0, 0
	}

	basediv := 1000000000
	h1 := 0
	for {
		h1 = num / basediv
		if h1 == 0 {
			basediv = basediv / 10
		} else {
			break
		}
	}
	return h1, basediv
}

func reverseString(input string) string {
	runes := []rune(input)                            // 将字符串转换为 Unicode 码点数组
	length := len(runes)                              // 获取字符串的长度
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 { // 使用双指针法反转字符串
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes) // 将 Unicode 码点数组转换为字符串并返回
}

func main() {
	//fmt.Println(getHigh(0))
	fmt.Println(isBigger(34, 30))
	//fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))
}
