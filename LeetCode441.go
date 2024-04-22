package main

func arrangeCoins(n int) int {
	ans := 0
	i := 1
	for n >= 0 {
		n = n - i
		i++
		ans++
		if n < 0 {
			ans--
			break
		}
	}
	return ans
}
