package main

// 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	// 假设是 完全二叉树
	invert226(root)
	return root
}

func invert226(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	invert226(root.Left)
	invert226(root.Right)
}
