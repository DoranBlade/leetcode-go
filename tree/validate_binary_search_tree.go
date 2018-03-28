package tree

var temp *TreeNode
var res = true

func isValidBST(root *TreeNode) bool {
	valid(root)
	return res
}

func valid(node *TreeNode) {
	if node == nil {
		return
	}
	valid(node.Left)
	if temp != nil && temp.Val >= node.Val {
		res = false
		return
	}
	temp = node
	valid(node.Right)
}