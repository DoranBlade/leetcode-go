package lnked_list

//Reverse a singly linked list.
//
//Hint:
//A linked list can be reversed either iteratively or recursively. Could you implement both?

func reverseList(head *ListNode) *ListNode {
	var reverse *ListNode
	for head != nil {
		next := head.Next
		head.Next = reverse
		reverse = head
		head = next
	}
	return reverse
}
