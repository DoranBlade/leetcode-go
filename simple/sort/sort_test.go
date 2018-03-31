package sort

import "testing"

func Test_Merge(t *testing.T) {
	nums1, nums2 := []int{4,5,6,0,0,0}, []int{1,2,3}
	merge(nums1, 3, nums2, 3)
	if !equals(nums1, []int{1, 2, 3, 4, 5, 6}) {
		t.Error("merge sort slice test error")
	}
}

func equals(nums1 []int, nums2 []int) bool {
	if len(nums1) != len(nums2) {
		return false
	}
	for i := 0; i < len(nums1); i++ {
		if nums1[i] != nums2[i] {
			return false
		}
	}
	return true
}

