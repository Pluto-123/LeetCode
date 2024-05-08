package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 寻找最低公共祖先
	return findLCA(root, p, q)
}

func findLCA(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}

	left := findLCA(root.Left, p, q)
	right := findLCA(root.Right, p, q)

	// 如果左右子树分别包含p和q，则当前节点为最低公共祖先
	if left != nil && right != nil {
		return root
	}

	// 如果只有左子树包含p或q，则返回左子树中找到的节点
	if left != nil {
		return left
	}

	// 如果只有右子树包含p或q，则返回右子树中找到的节点
	return right
}
