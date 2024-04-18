package main

import "fmt"

// 加油站
func canCompleteCircuit(gas []int, cost []int) int {
	// 写一个过滤函数
	//base := 0
	//for i := 0; i < len(cost)-6; i++ {
	//	if gas[i] == cost[i] {
	//		gas = gas[1:]
	//		cost = cost[1:]
	//		base++
	//	} else {
	//		break
	//	}
	//}

	length := len(cost)
	remainGas := 0
	// 遍历起始点
	for start := 0; start < len(cost); start++ {
		//remainGas += gas[start]
		if cost[start] > gas[start] {
			continue
		} else {
			remainGas = gas[start] - cost[start]
		}

		nowIndex := (start + 1) % length

		for time := 0; time < length; time++ {
			// 如果到达原点
			if nowIndex == start {
				return start
			}
			remainGas += gas[nowIndex]
			remainGas -= cost[nowIndex]
			if remainGas < 0 {
				start = nowIndex
				start--
				break
			}
			nowIndex = (nowIndex + 1) % length
		}
	}
	return -1
}

func main() {
	gas := []int{2, 3, 4}
	cost := []int{3, 4, 3}
	s := canCompleteCircuit(gas, cost)
	fmt.Println(s)
}
