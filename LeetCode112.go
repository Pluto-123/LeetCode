package main

func hasPathSum(root *TreeNode, targetSum int) bool {

	return check(root, 0, targetSum)
}

func check(root *TreeNode, traceSum int, targetSum int) bool {
	if root == nil {
		return false
	}
	traceSum += root.Val
	if root.Left == nil && root.Right == nil {
		// 叶子节点
		if targetSum == traceSum {
			return true
		} else {
			return false
		}
	} else {
		// 非叶子节点
		return check(root.Left, traceSum, targetSum) == true || check(root.Right, traceSum, targetSum) == true
	}
}
