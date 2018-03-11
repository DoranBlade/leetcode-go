package tree

import "testing"

func TestValidateBinarySearchTree(t *testing.T) {
	root := TreeNode{Val:0}
	res := isValidBST(&root)
	if !res {
		t.Error("res is false")
	}
}

