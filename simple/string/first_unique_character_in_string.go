package string

func firstUniqChar(s string) int {
	container := [26]int{}
	for _, c := range s {
		container[c - 'a']++
	}
	for i, c := range s {
		if container[c - 'a'] == 1 {
			return i
		}
	}
	return -1
}
