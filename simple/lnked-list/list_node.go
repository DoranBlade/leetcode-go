package lnked_list

type ListNode struct {
	Val int
	Next *ListNode
}

func createNode(nums []int) *ListNode {
	var header, temp *ListNode
	for i, v := range nums {
		current := &ListNode{v, nil}
		if i == 0 {
			header = current
			temp = current
			continue
		}
		temp.Next = current
		temp = temp.Next
	}
	return header
}

func valueNode(header *ListNode) []int {
	var container []int
	for header != nil  {
		container = append(container, header.Val)
		header = header.Next
	}
	return container
}