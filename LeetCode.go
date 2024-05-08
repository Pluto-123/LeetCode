package main

func flatten(root *TreeNode) {
	nums := make([]int, 0)
	preTravel(root, &nums)
	ans := &TreeNode{}
	ansi := ans
	for i := 0; i < len(nums); i++ {
		newNode := &TreeNode{
			Val:   nums[i],
			Left:  nil,
			Right: nil,
		}
		ansi.Right = newNode
		ansi = ansi.Right
	}
	root = ans.Right
	//if root == nil {
	//	root.Left = nil
	//
	//}
}

func preTravel(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	*nums = append(*nums, root.Val)
	preTravel(root.Left, nums)
	preTravel(root.Right, nums)
}
