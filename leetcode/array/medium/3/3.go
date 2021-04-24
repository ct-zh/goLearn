package main

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

// 思路1： 滑动窗口
// 执行用时： 4 ms, 在所有 Go 提交中击败了 88.63% 的用户
// 内存消耗： 2.7 MB 在所有 Go 提交中击败了68.48%的用户
func lengthOfLongestSubstring1(s string) int {
	freq := [256]int{0} // 存储byte值，用于判断当前字符串是否已经存在于滑动窗口了
	l, r := 0, -1       // 滑动窗口的左右边界，为了保证初始时滑动窗口里没有数据，右边界置为-1
	res := 0            // 当前循环内长度最小的子字符串的长度

	byteS := []byte(s)
	for {
		if l >= len(byteS) { // 窗口左边界不能越界
			break
		}

		// 先判断右边界往后是否还有元素
		// 并且往后的元素不在当前子串中存在相同的字符
		if r+1 < len(byteS) && freq[byteS[r+1]] == 0 {
			freq[byteS[r+1]]++
			r++
		} else {
			// 说明右边界已经达到末尾
			// 或者 右边界往后的元素与当前子串存在相同的字符
			// 尝试缩小字串，看是否符合条件
			freq[byteS[l]]--
			l++
		}

		// 与当前滑动窗口做比较，找出较大的一个
		res = maxInt(res, r-l+1)
	}

	return res
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	l, r, max := 0, -1, 0
	sbyte := []byte(s)

	m := make(map[byte]struct{})
	len2 := len(sbyte)

	for {
		if l >= len2 {
			break
		}

		if r+1 < len2 {
			if _, ok := m[sbyte[r+1]]; !ok {
				m[sbyte[r+1]] = struct{}{}
				r++
				if len(m) > max {
					max = len(m)
				}
			} else {
				delete(m, sbyte[l])
				l++
			}
		} else {
			delete(m, sbyte[l])
			l++
		}

	}

	return max
}
