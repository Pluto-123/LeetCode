package main

import "fmt"

func strStr(haystack string, needle string) int {
	next := make([]int, len(needle))
	next[0] = 0
	i := 1
	j := 0
	for i < len(needle) {
		if needle[i] == needle[j] {
			next[i] = j + 1
			i++
			j++
		} else {
			// 不等于
			if j > 0 {
				j = next[j-1]
			} else {
				next[i] = 0
				i++
			}
		}
	}
	fmt.Println(next)

	// 匹配
	i = 0
	j = 0
	for {
		// 成功跳出判断
		if j == len(needle) {
			return i - len(needle)
		}
		if i == len(haystack) {
			break
		}
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			if j == 0 {
				i++
			}
			if j > 0 {
				j = next[j-1]
			}
		}

	}
	//ans := 0
	return -1
}

func main() {
	strStr("a", "abaabaabaaba")
}
