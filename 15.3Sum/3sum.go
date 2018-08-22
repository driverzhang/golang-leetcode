package ThreeSum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	var ret [][]int
	sort.Ints(nums) // 排序很重要 [-4, -1, -1, 0, 1, 2]

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i-1] == nums[i] { // 排除除了第一个元素外的重复元素
			continue
		}
		j, k := i+1, len(nums)-1 // 1, 5

		for j < k {
			// 首次不会进去
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}

			if nums[i]+nums[j]+nums[k] == 0 {
				ret = append(ret, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			} else if nums[i]+nums[j]+nums[k] < 0 {
				j++
			} else {
				k--
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
