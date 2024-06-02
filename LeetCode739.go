package main

func dailyTemperatures(temperatures []int) []int {
	stack := make([][]int, 0)
	ans := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		if len(stack) == 0 {
			stack = append(stack, []int{temperatures[i], i})
			continue
		}
		for temperatures[i] > stack[len(stack)-1][0] {
			index := stack[len(stack)-1][1]
			ans[index] = i - index
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, []int{temperatures[i], i})
				continue
			}

		} // for
		if temperatures[i] < stack[len(stack)-1][0] {
			stack = append(stack, []int{temperatures[i], i})
			continue
		}
	}
	return ans
}
