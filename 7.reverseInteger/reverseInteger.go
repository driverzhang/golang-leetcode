package reverseInteger

// myself
func Reverse(x int) int {
	Max := 1<<31 - 1
	Min := -Max + 1
	y := 0

	for x != 0 {
		// 一个负数余10 仍未负数 => 不需要再判断正负区分符号！
		y = y*10 + x%10
		x = x / 10
	}

	if y < Min || y > Max {
		return 0
	} else {
		return y
	}
}

// authority
func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}

	result := 0
	for x > 0 {
		bit := x % 10
		result = result*10 + bit
		x /= 10
	}

	if result > 1<<31-1 || result <= -1<<31 {
		return 0
	}

	return result * sign
}
