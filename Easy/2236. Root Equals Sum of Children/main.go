func checkTree(root *TreeNode) bool {
	if root.Left.Val+root.Right.Val == root.Val {
		return true
	}

	return false
}
