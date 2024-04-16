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
	backTrace(digits, mp, &trace, &res)
	return res
}

func backTrace(digits string, mp map[byte][]byte, trace *[]byte, res *[]string) {
	if len(digits) == len(*trace) {
		*res = append(*res, string(*trace))
		return
	}
	for i := 0; i < len(digits); i++ {
		//if inside(trace, digits[i], mp) {
		//	continue
		//}

	}
}
