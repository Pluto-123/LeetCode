package main

func isSymmetric(root *TreeNode) bool {
	return checker(root.Left, root.Right)
}

func checker(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	return checker(a.Left, b.Right) && checker(a.Right, b.Left)
}
