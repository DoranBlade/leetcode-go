package other

func missingNumber(nums []int) int {
	var total, arrayTotal int64
	for i, v := range nums {
		arrayTotal += int64(v)
		total += int64(i)
	}
	total += int64(len(nums))
	return int(total - arrayTotal)
}
