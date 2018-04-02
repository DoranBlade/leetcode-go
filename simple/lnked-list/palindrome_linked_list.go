package lnked_list

//Given a singly linked list, determine if it is a palindrome.
//
//Follow up:
//Could you do it in O(n) time and O(1) space?

var temp1 *ListNode
func isPalindrome(head *ListNode) bool {
	temp1 = head
	return palind(head)
}
func palind(node *ListNode) bool {
	if node == nil {
		return true
	}
	result := palind(node.Next) && temp1.Val == node.Val
	temp1 = temp1.Next
	return result
}