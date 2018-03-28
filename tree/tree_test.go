package tree

import "testing"

func TestValidateBinarySearchTree(t *testing.T) {
	root := TreeNode{Val:0}
	res := isValidBST(&root)
	if !res {
		t.Error("res is false")
	}
}

func TestBinaryTreeLevelOrder(t *testing.T) {
	root := TreeNode{Val:10}
	root.Left = &TreeNode{Val:2}
	root.Right = &TreeNode{Val:22}
	res := levelOrder(&root)
	if res == nil {
		t.Error("res is nill")
	}
}

