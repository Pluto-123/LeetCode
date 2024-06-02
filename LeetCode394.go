package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 字符串解码
func decodeString(s string) string {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == ']' {
			bytes := getBytes(&stack)
			times := getTimes(&stack)
			for t := 0; t < times; t++ {
				for b := 0; b < len(bytes); b++ {
					stack = append(stack, bytes[b])
				}
			}
			continue
		} else {
			stack = append(stack, s[i])
		}
	}
	sb := strings.Builder{}
	sb.Write(stack)
	return sb.String()
}

func getTimes(stack *[]byte) int {
	//t, _ := strconv.Atoi(string((*stack)[len(*stack)-1]))
	//*stack = (*stack)[:len(*stack)-1]
	t := 0
	beishu := 1
	for {
		if len(*stack) != 0 && (*stack)[len(*stack)-1] >= '0' && (*stack)[len(*stack)-1] <= '9' {
			num, _ := strconv.Atoi(string((*stack)[len(*stack)-1]))
			t += num * beishu
			beishu *= 10
			*stack = (*stack)[:len(*stack)-1]
		} else {
			break
		}
	}
	return t
}

func getBytes(stack *[]byte) []byte {
	res := make([]byte, 0)
	for i := len(*stack) - 1; i >= 0; i-- {
		if (*stack)[i] != '[' {
			res = append(res, (*stack)[i])
			*stack = (*stack)[:len(*stack)-1]
			continue
		} else {
			*stack = (*stack)[:len(*stack)-1]
			break
		}
	}
	ans := make([]byte, 0)
	for i := len(res) - 1; i >= 0; i-- {
		ans = append(ans, res[i])
	}
	return ans
}

func main() {
	s := decodeString("100[leetcode]")
	fmt.Println(s)
}
