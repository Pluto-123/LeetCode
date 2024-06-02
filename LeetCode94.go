package main

func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	inorder94(root, &ans)
	return ans
}

func inorder94(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	inorder94(root.Left, ans)
	*ans = append(*ans, root.Val)
	inorder94(root.Right, ans)
}
