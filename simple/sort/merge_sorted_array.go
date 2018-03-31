package sort

func merge(nums1 []int, m int, nums2 []int, n int)  {
	lastIndex := m + n - 1
	m--
	n--
	for i := lastIndex; i >= 0; i-- {
		if m < 0 {
			nums1[i] = nums2[n]
			n--
			continue
		}
		if n < 0 {
			nums1[i] = nums1[m]
			m--
			continue
		}
		if nums1[m] >= nums2[n] {
			nums1[i] = nums1[m]
			m--
		} else {
			nums1[i] = nums2[n]
			n--
		}
	}
}
