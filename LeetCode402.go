package main

func removeKdigits(num string, k int) string {
	stack := make([]byte, 0)
	for i := 0; i < len(num); i++ {
		if len(stack) == 0 {
			stack = append(stack, num[i])
			continue
		}

	}
	return "nil"
}
