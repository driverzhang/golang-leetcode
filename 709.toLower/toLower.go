package toLower

func ToLowerCase(str string) string {
	vv := []rune(str)
	for i, v := range vv {
		if v <= 90 && v >= 65 {
			v += 32
			vv[i] = v
		}
	}

	return string(vv)
}

// 官方答案
func toLowerCase(str string) string {
	rst := []rune{}
	for _, v := range str {
		if v <= 90 && v >= 65 {
			rst = append(rst, v+32)
		} else {
			rst = append(rst, v)
		}
	}
	return string(rst)
}
