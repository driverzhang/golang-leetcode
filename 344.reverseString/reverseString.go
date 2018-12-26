package reverseString

func reverseString(s string) string {
	var str []byte
	for i := len(s) - 1; i >= 0; i-- { // 反序打印输出
		str = append(str, s[i])
	}
	return string(str)
}
