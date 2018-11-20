package longestSubstring

import "fmt"

// 支持中文输入类型的字符串
func LengthOfLongestSubstring(s string) int {
	lastOccured := make(map[rune]int)
	start := 0
	maxLength := 0
	//ch: [97 98 99 100 97 97 99] => abcdaac
	//i:  0  1  2  3  4   5  6
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccured[ch]; ok && lastI >= start { // 97:0 98:1 99:2 100:3 97:4 97:5 99:6
			fmt.Println(lastI, start)
			start = lastI + 1 // 1

		}
		if i-start+1 > maxLength { // 1>0 2>1 3>2 4>3  4>4? 2>4?
			maxLength = i - start + 1
		}

		lastOccured[ch] = i // map[97:0, 98:1, 99:2, 100:3, 97:4, 97:6]
	}
	// fmt.Printf("%+v", lastOccured) // map[98:1 99:7 100:3 101:5 97:6]
	return maxLength
}
