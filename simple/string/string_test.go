package string

import "testing"

func TestReverse(t *testing.T) {
	res1, res2 := reverseString("abcde"), reverseString("abcd")
	if res1 != "edcba" {
		t.Error("TestReverse expect:edcba, actual:" + res1)
	}
	if res2 != "dcba" {
		t.Error("TestReverse expect: dcba, actual: " + res2)
	}
}

