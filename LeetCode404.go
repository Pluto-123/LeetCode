package main

func sumOfLeftLeaves(root *TreeNode) int {
	ans := 0
	sumLeft(root, &ans)
	return ans
}

func sumLeft(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		*ans += root.Left.Val
	}
	return sumLeft(root.Left, ans) + sumLeft(root.Right, ans)

}
