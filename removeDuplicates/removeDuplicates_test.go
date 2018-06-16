package removeDuplicates

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 2, 3, 3}

	data := RemoveDuplicates(nums)
	t.Log(data)
}
