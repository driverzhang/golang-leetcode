package toLower

import "testing"

func TestToLowerCaset(t *testing.T) {
	str := "HWFFolle World"
	data := ToLowerCase(str)
	t.Log(data)
}
