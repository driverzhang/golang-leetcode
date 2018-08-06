package main

import (
	"fmt"
	"gitHub/golang-leetcode/longestSubstring"
)

func main() {
	testsubstring := "abccb"
	int := longestSubstring.LengthOfLongestSubstring(testsubstring)
	fmt.Printf("%+v", int)
}
