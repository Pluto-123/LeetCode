package main

func rotateString(s string, goal string) bool {
	for i := 0; i < len(s); i++ {
		s = s[len(s)-1:] + s[:len(s)-1]
		if s == goal {
			return true
		}
	}
	return false
}
