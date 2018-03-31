package array

//Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.
//
//For example, given nums = [0, 1, 0, 3, 12], after calling your function, nums should be [1, 3, 12, 0, 0].
//
//Note:
//You must do this in-place without making a copy of the array.
//Minimize the total number of operations.
func moveZeroes(nums []int)  {
	lastIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 && i != lastIndex {
			nums[lastIndex] = nums[i]
		}
		if nums[i] != 0 {
			lastIndex++
		}
	}
	for i := lastIndex; i < len(nums); i++{
		nums[i] = 0
	}
}