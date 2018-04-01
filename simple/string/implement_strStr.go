package string

func strStr(haystack string, needle string) int {
	needleL := len(needle)
	if needleL == 0 {
		return 0
	}
	l := len(haystack) - len(needle)
	for i := 0; i <= l; i++ {
		if haystack[i] == needle[0] && equal(haystack[i:needleL + i], needle) {
			return i
		}
	}
	return -1
}

func equal(nums1 string, nums2 string) bool {
	for i := 1; i < len(nums1); i++ {
		if nums2[i] != nums1[i] {
			return false
		}
	}
	return true
}
