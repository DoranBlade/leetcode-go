package array

func plusOne(digits []int) []int {
	lastIndex := -1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			lastIndex = i
			digits[i] += 1
			break
		} else {
			digits[i] = 0
		}
	}
	if lastIndex == -1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
