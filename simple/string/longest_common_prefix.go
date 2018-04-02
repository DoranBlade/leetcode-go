package string

func longestCommonPrefix(strs []string) string {
	res := ""
	minIndex := min(strs)
	if minIndex != -1 {
		count := 0
		for i := range strs[minIndex] {
			if !contain(strs, i) {
				break
			}
			count++
		}
		res += strs[minIndex][:count]
	}
	return res
}

func min(strs [] string) int {
	min := -1
	for i, v := range strs {
		if i == 0 {
			min = i
		}
		if len(v) < len(strs[min]) {
			min = i
		}
	}
	return min
}

func contain(strs []string, currentIndex int) bool {
	var temp uint8
	for i, v := range strs {
		if i == 0 {
			temp = v[currentIndex]
			continue
		}
		if v[currentIndex] != temp {
			return false
		}
	}
	return true
}
