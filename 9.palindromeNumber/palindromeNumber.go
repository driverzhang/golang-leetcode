package palindromeNumber

// 解法一：官方推挤，将数字去半截在逆转后半部分 进行对比
func IsPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		// 核心逻辑代码 %10可以拿到最后一位数字 *10 + %10 就能拿到倒转的后面的数字
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	// 当数字长度为奇数时，我们可以通过 revertedNumber/10 去除处于中位的数字。
	// 例如，当输入为 12321 时，在 while 循环的末尾我们可以得到 x = 12，revertedNumber = 123，
	// 由于处于中位的数字不影响回文（它总是与自己相等），所以我们可以简单地将其去除。
	return x == revertedNumber || x == revertedNumber/10

}

// 解法二：取全进来的数字整体进行逆转 进行对比
func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}
	tmp := x
	var result int
	for tmp > 0 {
		result = result*10 + (tmp % 10)
		tmp /= 10
	}
	return result == x
}
