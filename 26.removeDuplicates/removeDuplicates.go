package removeDuplicates

import (
	"fmt"

	"go.zhuzi.me/go/log"
)

// myself
func RemoveDuplicates165(nums []int) int { // [1, 2, 3, 3]
	i := 0
	for j := 1; j < len(nums); j++ {
		log.Error(nums[i], nums[j])
		if nums[i] != nums[j] {
			i++
			fmt.Println(nums[i], nums[j])
			nums[i] = nums[j]
		}
	}
	log.Error(nums)
	return i + 1
}

func RemoveDuplicates84(nums []int) int {
	length, newLength := len(nums), 1
	for i := 1; i < length; i++ {
		if nums[i-1] != nums[i] {
			nums[newLength] = nums[i]
			newLength++
		}
	}
	return newLength
}
