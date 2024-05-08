package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countAndSay(n int) string {

	dp := make([][]int, 31)
	dp[1] = []int{1}
	for i := 2; i <= 30; i++ {
		dp[i] = cal38(dp[i-1])
	}
	fmt.Println(dp)
	sb := strings.Builder{}
	for i := 0; i < len(dp[n]); i++ {
		sb.WriteString(strconv.Itoa(dp[n][i]))
	}
	return sb.String()
}

func cal38(pre []int) []int {
	ans := make([]int, 0)

	for i := 0; i < len(pre); i++ {
		for j := i; ; {
			if j == len(pre) {
				// 退出双重循环
				ans = append(ans, j-i)
				ans = append(ans, pre[i])
				i = j
				break
			}

			if pre[i] == pre[j] {
				j++
				continue
			}
			if pre[i] != pre[j] {
				ans = append(ans, j-i)
				ans = append(ans, pre[i])
				i = j - 1
				break
			}
		} // forj
	} // fori
	return ans
}

func main() {
	countAndSay(20)
}
