package main

import (
	"fmt"
	"math"
)

func getMinimumDifference(root *TreeNode) int {
	ans := math.MaxInt32
	var f func(root *TreeNode, res *[]int)
	f = func(root *TreeNode, res *[]int) {
		if root == nil {
			return
		}
		f(root.Left, res)
		*res = append(*res, root.Val)
		f(root.Right, res)
	}
	res := make([]int, 0)
	f(root, &res)
	fmt.Println(res)
	//ans = res[1] - res[0]
	for i := 0; i < len(res)-1; i++ {
		ans = min(ans, res[i+1]-res[i])
	}
	return ans
}
