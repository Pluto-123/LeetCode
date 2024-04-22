package main

import (
	"fmt"
	"strconv"
)

// 左 右 中
func evalRPN(tokens []string) int {
	// 逆波兰表达式
	stack := make([]int, 0)
	for i := 0; i < len(tokens); i++ {
		value, err := strconv.Atoi(tokens[i])
		if err != nil {
			// 运算符
			v2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			v1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res := cal(tokens[i], v1, v2)
			stack = append(stack, res)
		} else {
			stack = append(stack, value)
		}

	}
	return stack[len(stack)-1]
}

func cal(method string, v1 int, v2 int) int {
	if method == "+" {
		return v1 + v2
	}
	if method == "/" {
		return v1 / v2
	}
	if method == "*" {
		return v1 * v2
	}
	if method == "-" {
		return v1 - v2
	}
	return 0
}

func main() {
	tokens := []string{"4", "13", "5", "/", "+"}
	fmt.Println(evalRPN(tokens))
}
