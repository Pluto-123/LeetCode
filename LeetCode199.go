package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TODO
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	temp := make([]*TreeNode, 0)
	temp = append(temp, root)
	res := make([]int, 0)
	length := 2
	for len(temp) != 0 {
		//fmt.Println(temp[0])
		length--

		//res = append(res, temp[0].Val)
		if temp[0].Left != nil {
			temp = append(temp, temp[0].Left)
		}
		if temp[0].Right != nil {
			temp = append(temp, temp[0].Right)
		}
		temp = temp[1:]

	}

	return res
}

//func toe()  {
//
//}
