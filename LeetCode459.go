package main

func repeatedSubstringPattern(s string) bool {
	for i := 1; i < len(s)/2+1; i++ {
		s1 := s[i:] + s[0:i]
		if s1 == s {
			return true
		}
	}
	return false
}

func main() {
	repeatedSubstringPattern("abab")
}
