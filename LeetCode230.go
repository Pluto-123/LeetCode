package main

func kthSmallest(root *TreeNode, k int) int {
	res := make([]int, 0)

	midTravel(root, &res)
	return res[k-1]
}

func midTravel(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	midTravel(root.Left, res)
	*res = append(*res, root.Val)
	midTravel(root.Right, res)

}
