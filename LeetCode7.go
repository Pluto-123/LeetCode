package main

import "math"

func reverse(x int) int {
	// 去除尾部的0
	if x == 0 {
		return 0
	}
	for x%10 == 0 {
		x = x / 10
	}
	absx := myAbs(x)
	res := make([]int, 0)
	for {
		if absx < 10 {
			res = append(res, absx)
			break
		}
		res = append(res, absx%10)
		absx = absx / 10

	}
	ans := 0
	times := 1
	for i := len(res) - 1; i >= 0; i-- {
		ans += res[i] * times
		times *= 10

		if float64(ans) > math.Pow(2.0, 31.0) {
			return 0
		}
	}
	if x < 0 {
		return ans * (-1)
	}
	return ans
}

func myAbs(x int) int {
	if x < 0 {
		return x * (-1)
	}
	return x
}

func main() {
	reverse(1)
}
