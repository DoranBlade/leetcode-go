package lnked_list

//Given a linked list, remove the nth node from the end of list and return its head.
//
//For example,
//
//Given linked list: 1->2->3->4->5, and n = 2.
//
//After removing the second node from the end, the linked list becomes 1->2->3->5.
//Note:
//Given n will always be valid.
//Try to do this in one pass.

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var container []*ListNode
	temp := head
	for temp != nil {
		container = append(container, temp)
		temp = temp.Next
	}
	if len(container) == 0 {
		return nil
	}
	if len(container)-1-n < 0 {
		return head.Next
	}
	preNode := container[len(container) - 1 - n]
	preNode.Next = preNode.Next.Next
	return head
}