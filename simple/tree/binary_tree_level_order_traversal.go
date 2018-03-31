package tree

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var nodeContainer [][]*TreeNode
	nodeContainer = append(nodeContainer, []*TreeNode{root})
	for i := 0; i < len(nodeContainer); i++ {
		var temp []*TreeNode
		for j := 0; j < len(nodeContainer[i]); j++ {
			childNodes := getChildrenNode(nodeContainer[i][j])
			if childNodes != nil {
				temp = append(temp, childNodes...)
			}
		}
		if len(temp) != 0 {
			nodeContainer = append(nodeContainer, temp)
		}
	}
	return getValues(nodeContainer)
}

func getChildrenNode(parent *TreeNode) []*TreeNode {
	if parent == nil || (parent.Left == nil && parent.Right == nil) {
		return nil
	}
	result := make([]*TreeNode, 2)
	if parent.Left != nil {
		result = append(result, parent.Left)
	}
	if parent.Right != nil {
		result = append(result, parent.Right)
	}
	return result
}

func getValues(treeNodeList [][]*TreeNode) [][]int {
	var result [][]int
	for _, v := range treeNodeList{
		result = append(result, getValue(v))
	}
	return result
}


func getValue(treeNodes []*TreeNode) []int {
	var result []int
	for _, v := range treeNodes {
		if v != nil {
			result = append(result, v.Val)
		}
	}
	return result
}
