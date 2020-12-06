package _455

import "sort"

// 时间复杂度: O(n) + O(logn)排序 = O(nlogn)
func findContentChildren(g []int, s []int) (res int) {
	sort.Ints(g)
	sort.Ints(s)

	// 指向最大的饼干
	si := len(s) - 1
	// 指向最贪心的小朋友
	gi := len(g) - 1

	for gi >= 0 && si >= 0 {
		// 最大的饼干满足最贪心的小朋友
		if s[si] >= g[gi] {
			res++
			gi--
			si--
		} else { // 不满足
			gi-- // 找当前次贪心的小朋友
		}
	}

	return res
}
