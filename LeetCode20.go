package main

func isValid(s string) bool {
	stack := make([]byte, 0)
	mp := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else {
			if len(stack) != 0 && mp[stack[len(stack)-1]] == s[i] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	} // for
	if len(stack) == 0 {
		return true
	}
	return false
}
