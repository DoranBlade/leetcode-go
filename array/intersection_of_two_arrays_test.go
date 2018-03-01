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