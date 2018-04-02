package lnked_list

//Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.
//
//Example:
//
//Input: 1->2->4, 1->3->4
//Output: 1->1->2->3->4->4

var head, temp *ListNode
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	for l1 != nil || l2 != nil  {
		if l2 == nil || (l1 != nil && l1.Val <= l2.Val) {
			setNode(l1)
			l1 = l1.Next
			continue
		}
		setNode(l2)
		l2 = l2.Next
	}
	return head
}

func setNode(node *ListNode) {
	if head == nil {
		head = node
		temp = node
	} else {
		temp.Next = node
		temp = node
	}
}
