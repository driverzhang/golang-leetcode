package twoSum

// myself
// 两遍 哈希表
func TwoSum12(nums []int, target int) []int {
	numsMap := make(map[int]int)
	testMap := make(map[int]int)

	var intRp = make([]int, 2)

	for jj, v := range nums {
		numsMap[v] = v
		testMap[v] = jj
	}
	for i, v := range nums {
		if s, ok := numsMap[target-v]; ok && testMap[s] != i {
			intRp = []int{i, testMap[s]}
			break
		}
	}
	return intRp
}

// 一遍 哈希表
func TwoSum4(nums []int, target int) []int { // [3,2,4] 6
	result := make([]int, 2)
	m := make(map[int]int)

	for i, v := range nums {
		if _, ok := m[target-v]; ok { // m[3] m[4] m[2]
			result[1] = i
			result[0] = m[target-nums[i]]
			break
		}
		m[v] = i // 省去了 for range => m[3]0  m[2]1 m[4]2
	}
	return result
}

// func TwoSum4(nums []int, target int) []int { // [3,2,4] 6
// 	result := make([]int, 2)
// 	m := make(map[int]int)
// 	log.Error(m)
// 	for i := 0; i < len(nums); i++ {
// 		if _, ok := m[target-nums[i]]; ok { // m[3] m[4] m[2]
// 			result[1] = i
// 			result[0] = m[target-nums[i]]
// 			break
// 		}
// 		m[nums[i]] = i // 省去了 for range => m[3]0  m[2]1 m[4]2
// 	}
// 	return result
// }
