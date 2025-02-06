package main

// https://leetcode.cn/problems/count-substrings-that-can-be-rearranged-to-contain-a-string-i

// 给你两个字符串 word1 和 word2 。
// 如果一个字符串 x 重新排列后，word2 是重排字符串的 前缀 ，那么我们称字符串 x 是 合法的 。
// 请你返回 word1 中 合法 子字符串 的数目。
func validSubstringCount(word1 string, word2 string) int64 {

	count := int64(0)
	n1 := len(word1)

	// 统计 word2 中每个字符的频率
	freq2 := make(map[rune]int)
	for _, ch := range word2 {
		freq2[ch]++
	}

	// 遍历 word1 的所有子字符串
	for i := 0; i < n1; i++ {
		freq1 := make(map[rune]int)
		for j := i; j < n1; j++ {
			freq1[rune(word1[j])]++
			// 检查当前子字符串的频率是否满足 word2 的前缀条件
			if isValidPrefix(freq1, freq2) {
				count++
			}
		}
	}

	return count
}
