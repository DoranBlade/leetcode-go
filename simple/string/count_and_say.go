package string

import (
	"bytes"
)

func countAndSay(n int) string {
	if n == 0 {
		return ""
	}
	result := "1"
	for i := 1; i < n; i++ {
		result = conv(result)
	}
	return result
}

func conv(str string) string {
	var count byte = 1
	var result bytes.Buffer
	for i := 0; i < len(str)-1; i++ {
		if str[i] != str[i+1] {
			result.Write(split(str[i], count))
			count = 1
		} else if i == len(str)-2 {
			count++
			result.Write(split(str[i], count))
		} else {
			count++
		}
	}
	if result.String() == "" {
		result.Write(split(str[0], count))
	}
	if len(str) >= 2 && str[len(str)-1] != str[len(str)-2] {
		result.Write(split(str[len(str) - 1], count))
	}
	return result.String()
}

func split(v byte, count byte) []byte {
	return []byte{count + '0', v}
}
