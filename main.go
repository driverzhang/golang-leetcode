package main

import (
	"fmt"

	longestSubstring "gitHub/golang-leetcode/3.longestSubstring"
)

func main() {
	testsubstring := "abccb"
	int := longestSubstring.LengthOfLongestSubstring(testsubstring)
	fmt.Printf("%+v", int)
}
