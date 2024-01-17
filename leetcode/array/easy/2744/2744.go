package l2744

// 给你一个下标从 0 开始的数组 words ，数组中包含 互不相同 的字符串。
// 如果字符串 words[i] 与字符串 words[j] 满足以下条件，我们称它们可以匹配：
// 字符串 words[i] 等于 words[j] 的反转字符串。
// 0 <= i < j < words.length
// 请你返回数组 words 中的 最大 匹配数目。
// 注意，每个字符串最多匹配一次。

func maximumNumberOfStringPairs(words []string) int {
	result := 0
	for start := 0; start < len(words); start++ {
		end := len(words) - 1 // 每次从最右边开始匹配
		for start < end {
			if len(words[start]) != len(words[end]) { // 长度不同直接退
				end--
				continue
			}
			add := true
			for i := 0; i < len(words[end]); i++ {
				if words[start][i] != words[end][len(words[end])-1-i] {
					add = false
					break
				}
				//fmt.Printf("[%d]start = %v end=%v \n", i, words[start][i], words[end][len(words[end])-1-i])
			}
			if add {
				//fmt.Printf("add result start=%d end=%d \n", start, end)
				result++
			}
			end--
		}
	}
	return result
}
