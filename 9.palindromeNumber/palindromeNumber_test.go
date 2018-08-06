package palindromeNumber

import "testing"

func TestPalindromeNumber(t *testing.T) {
	x := 12321
	exist := IsPalindrome(x)
	t.Log(exist)
}
