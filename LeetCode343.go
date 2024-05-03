package main

import "math"

func integerBreak(n int) int {
	// 拆分成i个数字后的最大乘积
	link := make([]float64, n+1)
	for k := 2; k <= n; k++ {
		if n%k == 0 {
			link[k] = math.Pow(float64(n/k), float64(k))
			continue
		}
		if n%k != 0 {
			link[k] = calc2(n, k)
		}
	}
	m := 0.0
	for i := 0; i < len(link); i++ {
		m = maxfloat(m, link[i])
	}
	return int(m)
}

func maxfloat(m float64, f float64) float64 {
	if m > f {
		return m
	}
	return f
}

func calc2(n int, k int) float64 {
	tmp := make([]int, k)
	for i := 0; i < len(tmp); i++ {
		tmp[i] = n / k
	}
	start := 0
	for i := 0; i < n%k; i++ {
		tmp[start]++
		start++
	}
	ans := 1
	for i := 0; i < len(tmp); i++ {
		ans *= tmp[i]
	}
	return float64(ans)
}
