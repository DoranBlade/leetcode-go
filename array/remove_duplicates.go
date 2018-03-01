package array

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	newLength := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			newLength++
			nums[newLength] = nums[i]
		}
	}
	return newLength
}
