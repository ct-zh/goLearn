package _344

// https://leetcode-cn.com/problems/reverse-string/

// 要求空间复杂度为O(1)
func reverseString(s []byte) {
	i1 := 0
	i2 := len(s) - 1
	for {
		if i1 >= i2 {
			break
		}
		s[i1], s[i2] = s[i2], s[i1]
		i1++
		i2--
	}
}
