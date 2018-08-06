package twoSum

import "testing"

func TestTwoSun12(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6

	data := TwoSum12(nums, target)

	t.Logf("%+v", data)

}

func TestTwoSun4(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6

	data := TwoSum4(nums, target)

	t.Logf("%+v", data)

}
