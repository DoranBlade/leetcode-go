package tree

func maxDepth(root *TreeNode) int {
	return childDepth(root)
}

func childDepth(parent *TreeNode) int {
	if parent == nil {
		return 0
	}
	left := childDepth(parent.Left) + 1
	right := childDepth(parent.Right) + 1
	if left > right {
		return left
	}
	return right
}