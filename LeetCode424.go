package main

// 超时
func characterReplacement(s string, k int) int {
	ans := 0
	//diff := 0
	left := 0
	right := min(len(s)-1, k-1)
	for right < len(s) {
		c := checker1(s, left, right, k)
		if c == true {
			ans = max(ans, right-left+1)
			right++
		}
		if c == false {
			left++
		}
	}
	return ans
}

func checker1(s string, start, end, k int) bool {
	//mp := map[byte]int{}
	if k >= end-start+1 {
		return true
	}
	mp := map[byte]bool{} // 用来标识使用过的主字符
	times := k
	for i := start; i <= end; i++ {
		if mp[s[i]] {
			continue
		}
		mp[s[i]] = true
		for j := start; j <= end; j++ {
			if s[j] != s[i] {
				times--
			}
		} //for
		if times >= 0 {
			return true
		}
		times = k
	} // for
	return false
}
