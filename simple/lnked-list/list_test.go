package lnked_list

import (
	"testing"
	"fmt"
)

func TestRemoveNthNode(t *testing.T) {
	input := createNode([]int{1,2,3,4,5})
	res := removeNthFromEnd(input, 5)
	fmt.Print(valueNode(res))
}

func TestReverseList(t *testing.T) {
	input := createNode([]int{1,2,3,4,5})
	res := reverseList(input)
	fmt.Print(valueNode(res))
}

func TestMergeTwoSortedList(t *testing.T) {
	input1 := createNode([]int{-30,-27,-26,-25,-23,-23,-22,-20,-18,-16,-16,-12,-9,-7,-1,-1,-1,0,8,8,10,16,18,19,19,19,20,22,22,24,29})
	input2 := createNode([]int{-30,-29,-26,-23,-15,-14,-14,-13,-13,18,22})
	res := mergeTwoLists(input1, input2)
	fmt.Print(valueNode(res))
}

func TestPalindrome(t *testing.T) {
	input := createNode([]int{1})
	res := isPalindrome(input)
	fmt.Print(res)
}


