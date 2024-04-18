package main

// 每个整数位于 0 到 255 之间组成，且不能含有前导 0
// fail
func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	path := make([]byte, 0)
	visited := make([]bool, len(s)) // 用来标识是否访问过
	backTrace2(s, &path, visited, &res)
	return res
}

func backTrace2(s string, path *[]byte, visited []bool, res *[]string) {
	if finished(*path) {
		*res = append(*res, string(*path))
		return
	}
	for i := 0; i < len(s); i++ {
		if visited[i] {
			continue
		}
		*path = append(*path, s[i])

	}
}

func finished(bytes []byte) bool {
	return true
}
