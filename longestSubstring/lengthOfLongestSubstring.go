package longestSubstring

// 支持中文输入类型的字符串
func LengthOfLongestSubstring(s string) int {
	lastOccured := make(map[rune]int)
	start := 0
	maxLength := 0
	//ch: [97 98 99 100 97 97 99] => abcdaac
	//i:  0  1  2  3  4   5  6
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccured[ch]; ok && lastI >= start { // 97:0
			start = lastI + 1 // 1

		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}

		lastOccured[ch] = i
	}
	return maxLength
}
