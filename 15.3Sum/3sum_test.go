package ThreeSum

import "testing"

func TestThreeSum(t *testing.T) {
	// b := []int{-2, 0, 0, 2, 2}
	a := []int{-1, 0, 1, 2, -1, -4}
	data := threeSum(a)
	t.Logf("%+v", data)
}
