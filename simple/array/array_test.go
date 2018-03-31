package array

import (
	"testing"
)

func Test_Intersect(t *testing.T) {
	sorter, longer := []int{1, 2}, []int {0, 2, 2, 3}
	res := Intersect(longer, sorter)
	if len(res) != 1 {
		t.Error("res length not expect")
	}
}

func TestMoveZero(t *testing.T) {
	nums := []int{1,0}
	moveZeroes(nums)
	if nums[0] != 1 || nums[1] != 0 {
		t.Error("move zero test error")
	}
}
