package ThreeSum

import (
	"sort"
)

// func fourSum(nums []int, target int) [][]int {
// 	var resp [][]int
// 	sort.Ints(nums)
// 	first := 0
// 	secend := first + 1
// 	threeth := secend + 1
// 	last := len(nums) - 1

// 	firstNum, secendNum, threethNum, lastNum := nums[first], nums[secend], nums[threeth], nums[last]

// 	for ; first <= last-1; first++ {

// 	}

// 	return resp

// }

func threeSum(nums []int) [][]int {
	var ret [][]int
	sort.Ints(nums) // 排序很重要 [-4, -1, -1, 0, 1, 2]

	for first := 0; first < len(nums)-2; first++ {
		firstNum := nums[first]

		if first > 0 && nums[first-1] == firstNum { // 排除除了第一个元素外的重复元素
			continue
		}
		second, last := first+1, len(nums)-1 // 1, 5

		for second < last {
			secondNum, lastNums := nums[second], nums[last]
			sum := firstNum + secondNum + lastNums

			// 首次不会进去
			if second > first+1 && secondNum == nums[second-1] {
				second++
				continue
			}

			if sum == 0 {
				ret = append(ret, []int{firstNum, secondNum, lastNums})
				second++
				last--
			} else if sum < 0 {
				second++
			} else {
				last--
			}
		}
	}
	return ret
}

// 优解
func threeSum02(nums []int) [][]int {
	sort.Ints(nums)

	r := make([][]int, 0)
	l := len(nums)
	for i := 0; i < l-2; {
		a := nums[i]
		cnt := -a
		j := i + 1
		k := l - 1
		for {
			if j >= k {
				break
			}
			n := nums[j] + nums[k]
			if n == cnt {
				r = append(r, []int{nums[i], nums[j], nums[k]})
				j0 := j
				k0 := k
				for j < l-1 {
					j++
					if nums[j0] != nums[j] {
						break
					}
				}
				for k > j {
					k--
					if nums[k0] != nums[k] {
						break
					}
				}
			} else if n > cnt {
				k--
			} else if n < cnt {
				j++
			}
		}
		for i < l-2 {
			i++
			if nums[i] != a {
				break
			}
		}
	}
	return r
}
