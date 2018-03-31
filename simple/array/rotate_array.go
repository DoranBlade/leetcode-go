package array

func Rotate(nums []int, k int)  {
	length := len(nums)
	temp := make([]int, length)
	copy(temp, nums)
	for i := 0; i < length; i++ {
		nums[(i + k) % length] = temp[i]
	}
}

