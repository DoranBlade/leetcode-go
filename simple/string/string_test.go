package string

import (
	"testing"
	"strconv"
)

func TestReverse(t *testing.T) {
	res1, res2 := reverseString("abcde"), reverseString("abcd")
	if res1 != "edcba" {
		t.Error("TestReverse expect:edcba, actual:" + res1)
	}
	if res2 != "dcba" {
		t.Error("TestReverse expect: dcba, actual: " + res2)
	}
}

func TestFirstUniqueChar(t *testing.T) {
	res := firstUniqChar("leetcode")
	res1 := firstUniqChar("loveleetcode")
	if res != 0 {
		t.Error("TestFirstUniqueChar expect: 0, actual: " + string(res))
	}
	if res1 != 0 {
		t.Error("TestFirstUniqueChar expect: 0, actual: " + string(res))
	}
}

func TestValidAnagram(t *testing.T) {
	s, r := "anagram", "nagaram"
	res := isAnagram(s, r)
	if !res {
		t.Error("TestValidAnagram expect: true, actual: false")
	}
}

func TestValidPalindrome(t *testing.T) {
	res, res1 := isPalindrome("race a car"), isPalindrome("0p")
	if res || !res1  {
		t.Error("TestValidPalindrome expect: true, actual: false")
	}
}

func TestStringToInteger(t *testing.T) {
	res := 	myAtoi("9223372036854775809")
	if res != 1 {
		t.Error("TestStringToInteger expect: 1, actual: " + strconv.Itoa(res))
	}
}

func TestImplementStr(t *testing.T) {
	res1, res2 := strStr("hello", "ll"), strStr("aaaaa", "bba")
	if res1 != 2 || res2 != -1 {
		t.Error("TestImplementStr error")
	}
}

func TestCountAndSay(t *testing.T) {
	res1, res2 := countAndSay(1), countAndSay(5)
	if res1 != "1" || res2 != "111221" {
		t.Error("TestCountAndSay Error")
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	res1, res2 := longestCommonPrefix([]string{"abcd", "abc", "ab"}), longestCommonPrefix([]string{"abc", "", "abc"})
	if res1 != "ab" || res2 != "" {
		t.Error("TestLongestCommonPrefix Error")
	}
}







