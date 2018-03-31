package array

func containsDuplicate(nums []int) bool {
	container := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		_, ok := container[nums[i]]
		if ok {
			return true
		}
		container[nums[i]] = true
	}
	return false
}
