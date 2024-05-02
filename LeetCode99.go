package main

func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil && root.Right != nil {
		if root.Left.Val > root.Val {
			// 交换
		} else if root.Right.Val < root.Val {
			// 交换
		} else {
			recoverTree(root.Left)
			recoverTree(root.Right)
		}
	}
	if root.Left == nil {

	}
}
