package main

import (
	"fmt"
	"golang/leetCode/longestSubstring"
)

func main() {
	testsubstring := "abccb"
	int := longestSubstring.LengthOfLongestSubstring(testsubstring)
	fmt.Printf("%+v", int)
}
