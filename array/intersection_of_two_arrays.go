package array

func Intersect(nums1 []int, nums2 []int) []int {
	slow, fast := 0, 0
	var longer, sorter *[]int
	if len(nums1) > len(nums2) {
		longer = &nums1
		sorter = &nums2
	} else {
		longer = &nums2
		sorter = &nums1
	}
	matched := make([]int, len(*sorter))
	for ; fast < len(*longer); fast++ {
		for j := 0; j < len(*sorter); j++ {
			if matched[j] == 0 && (*sorter)[j] == (*longer)[fast] {
				matched[j] = 1
				slow++
				break
			}
		}
	}
	res := make([]int, slow)
	for i,j := 0,0; i < len(matched); i++ {
		if matched[i] == 1 {
			res[j] = (*sorter)[i]
			j++
		}
	}
	return  res
}