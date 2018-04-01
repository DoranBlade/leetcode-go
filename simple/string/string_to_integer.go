package string

import (
	"math"
	"strings"
)

func myAtoi(str string) int {
	str = strings.TrimLeft(str, " ")
	chars := []byte(str)
	seqFlag := false
	var result int64
	for i, v := range chars{
		if !isSeq(v) && !isNum(v) && i == 0 {
			return 0
		}
		if isSeq(v) && i == 0 {
			seqFlag = true
			continue
		}
		if isNum(v) {
			result = result * 10 + int64(v - '0')
			if result > math.MaxInt32 + 1 {
				break
			}
		} else {
			break
		}
	}
	if seqFlag && chars[0] == '-' {
		result = 0 - result
	}
	if result > math.MaxInt32 {
		return math.MaxInt32
	}
	if result < math.MinInt32 {
		return math.MinInt32
	}
	return int(result)
}

func isSeq(b byte) bool {
	return b == '+' || b == '-'
}

func isNum(b byte) bool {
	return b >= '0' && b <= '9'
}