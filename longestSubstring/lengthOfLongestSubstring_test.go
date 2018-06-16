package longestSubstring

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	str := "abcdaac"

	data := LengthOfLongestSubstring(str)
	t.Log(data)
}
