package string

//Write a function that takes a string as input and returns the string reversed.
//
//Example:
//Given s = "hello", return "olleh".
func reverseString(s string) string {
	charArray := []byte(s)
	var left, right int
	if len(charArray)%2 == 0 {
		right = len(charArray) / 2
		left = right - 1
	} else {
		left = len(charArray)/2
		right = left
	}
	for ; right < len(charArray); left, right = left-1, right+1 {
		if left == right {
			continue
		}
		charArray[left], charArray[right] = charArray[right], charArray[left]
	}
	return string(charArray)
}
