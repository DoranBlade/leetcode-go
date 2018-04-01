package string

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var container = [26]int{}
	sChar, tChar := []byte(s), []byte(t)
	for _, c := range sChar {
		container[c-'a']++
	}
	for _, c := range tChar {
		container[c-'a']--
	}
	for _, v := range container {
		if v != 0 {
			return false
		}
	}
	return true
}
