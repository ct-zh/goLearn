package _125

// https://leetcode-cn.com/problems/valid-palindrome/

// 对撞指针
// 要点： 1. 忽略字母的大小写；2. 只考虑数字和字母；
// 时间复杂度: O(n): 执行4ms，击败75%
// 空间复杂度：O(n)： 消耗3.5mb,击败25%
func isPalindrome(s string) bool {
	if s == "" {
		return true
	}

	s = simplifyStr(s)
	btS := []byte(s)
	i1 := 0
	i2 := len(btS) - 1

	for {
		if i1 >= i2 {
			break
		}

		if btS[i1] != btS[i2] {
			return false
		}
		i1++
		i2--
	}
	return true
}

// 去掉非数字与字母的部分，并将大写字母转换为小写字母
// a-z: 97-122
// A-Z: 65-90
// 0-9: 48-57
func simplifyStr(s string) string {
	btS := []byte(s)

	var nS []byte
	for i := 0; i < len(btS); i++ {
		if (btS[i] >= 48 && btS[i] <= 57) || (btS[i] >= 97 && btS[i] <= 122) {
			nS = append(nS, btS[i])
		} else if btS[i] >= 65 && btS[i] <= 90 {
			nS = append(nS, btS[i]+32)
		}
	}

	return string(nS)
}

// 对撞指针优化版
// 时间复杂度O(n): 0ms,击败100%
// 空间复杂度O(1): 2.8mb，击败52% 最佳2.6mb
func isPalindrome2(s string) bool {
	btS := []byte(s)
	i1 := 0
	i2 := len(btS) - 1
	for {
		if i1 >= i2 {
			break
		}

		if !isalnum(btS[i1]) {
			i1++
			continue
		}
		if !isalnum(btS[i2]) {
			i2--
			continue
		}

		if btS[i1] >= 'A' && btS[i1] <= 'Z' {
			btS[i1] = btS[i1] + 32
		}
		if btS[i2] >= 'A' && btS[i2] <= 'Z' {
			btS[i2] = btS[i2] + 32
		}

		if btS[i1] != btS[i2] {
			return false
		}
		i1++
		i2--
	}

	return true
}

func isalnum(ch byte) bool {
	if (ch >= 'A' && ch <= 'Z') ||
		(ch >= 'a' && ch <= 'z') ||
		(ch >= '0' && ch <= '9') {
		return true
	}
	return false
}
