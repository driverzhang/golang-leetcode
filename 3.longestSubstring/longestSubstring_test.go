package longestSubstring

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	str := "abcdaeac"

	data := LengthOfLongestSubstring(str)
	t.Log(data)
}
