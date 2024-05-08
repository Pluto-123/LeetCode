package main

// 恢复二叉搜索树
func recoverTree(root *TreeNode) {
	s := make([]int, 0)
	midOrder(root, &s)
	link := make([]int, 0)
	//sort.Ints()
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			link = append(link, s[i])
		}
	}
	replaceTree(root, link[0], link[1])
}

func replaceTree(root *TreeNode, a int, b int) {
	if root == nil {
		return
	}
	if root.Val == a {
		root.Val = b
	}
	if root.Val == b {
		root.Val = a
	}
	replaceTree(root.Left, a, b)
	replaceTree(root.Right, a, b)
}

func midOrder(root *TreeNode, s *[]int) {
	if root == nil {
		return
	}
	midOrder(root.Left, s)
	*s = append(*s, root.Val)
	midOrder(root.Right, s)

}
