package _345

// 时间复杂度O(n): 执行0ms，击败100%
// 新开辟了一个变量btS 空间复杂度O(n)：内存消耗4mb，击败68%
func reverseVowels(s string) string {
	btS := []byte(s)
	i1 := 0
	i2 := len(s) - 1
	for {
		if i1 >= i2 {
			break
		}

		if !isVowel(btS[i1]) {
			i1++
			continue
		}

		if !isVowel(btS[i2]) {
			i2--
			continue
		}

		btS[i1], btS[i2] = btS[i2], btS[i1]
		i1++
		i2--
	}
	return string(btS)
}

func isVowel(s byte) bool {
	if s == 'a' || s == 'A' || s == 'e' || s == 'E' ||
		s == 'i' || s == 'I' || s == 'o' || s == 'O' ||
		s == 'u' || s == 'U' {
		return true
	}
	return false
}
