package main

// fail
func letterCombinations(digits string) []string {
	res := make([]string, 0)
	mp := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	} // 2 ==> [a,b,c]
	trace := make([]byte, 0) // 把trace设置成切片比较好修改
	checked := map[byte]bool{}
	backTrace(digits, mp, &trace, &res, checked)
	return res
}

func backTrace(digits string, mp map[byte][]byte, trace *[]byte, res *[]string, checked map[byte]bool) {

	if len(*trace) == len(digits) {
		tmp := make([]byte, len(*trace))
		copy(tmp, *trace)
		*res = append(*res, string(tmp))
		return
	}
	for i := 0; i < len(digits); i++ {
		// 把东西取出来

	}
}
