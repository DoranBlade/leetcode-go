package string

import "strings"

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	s = strings.ToLower(s)
	char := []byte(s)
	start, end := 0, len(char) - 1
	for start <= end {
		if !isWord(char[start]) {
			start++
			continue
		}
		if !isWord(char[end]) {
			end--
			continue
		}
		if char[start] != char[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func isWord(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9')
}