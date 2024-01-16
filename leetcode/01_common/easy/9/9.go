package l9

import (
	"strconv"
)

// 能不将整数转为字符串来解决这个问题吗？
func isPalindrome(x int) bool {
	if x < 0 { // 负数直接返回false
		return false
	}
	xStr := strconv.Itoa(x)
	left, right := 0, len(xStr)-1 // 将整数拆成两个数，逐位对比
	for left < right {
		if xStr[left] != xStr[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 优化：1、不对字符串进行转换；2、只用一个变量，直接与本体对比
func isPalindromeV2(x int) bool {
	// 负数和以0结尾的非零数不是回文数
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	reversed := 0
	for x > reversed {
		reversed = reversed*10 + x%10
		x /= 10
	}
	// 当数字长度为奇数时，去掉中间的数字
	return x == reversed || x == reversed/10
}
