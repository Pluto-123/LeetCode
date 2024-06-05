package main

func longestPalindrome2(s string) int {
	mp := map[byte]int{}
	// 初始化map
	for i := 0; i < len(s); i++ {
		_, ok := mp[s[i]]
		if ok {
			mp[s[i]]++
		} else {
			mp[s[i]] = 1
		}
	}
	res := 0
	flag := false // 用来表示回文中心有没有被使用
	for _, value := range mp {
		if value%2 == 0 {
			res += value
		} else {
			// 奇数
			if flag == true {
				res += value - 1
			}
			if flag == false {
				res += value
				flag = true
			}
		}
	}
	return res
}
