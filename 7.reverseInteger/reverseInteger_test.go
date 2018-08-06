package reverseInteger

import "testing"

func TestReverse(t *testing.T) {
	number := -102929
	num := Reverse(number)
	t.Logf("%+v", num)
}
