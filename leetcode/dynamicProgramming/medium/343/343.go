package _343

import Helper "github.com/LannisterAlwaysPaysHisDebts/goLearn/leetcode/base/helper"

// 暴力解法： 回溯遍历将一个数分割的所有可能性 O(2^n)

func integerBreak(n int) int {
	if n < 2 {
		return 0
	}

	s := solution{memo: make([]int, n+1)}
	for i := 0; i <= n; i++ {
		s.memo[i] = -1
	}

	return s.breakInt(n)
}

type solution struct {
	memo []int
}

// 求将n进行分割（至少分割两部分），可以获得最大乘积
// n=4 i从1开始：1 + 3, 2 + 2, 3 + 1
func (s solution) breakInt(n int) int {
	if n == 1 {
		return 1
	}

	if s.memo[n] != -1 {
		return s.memo[n]
	}

	res := -1

	// 将n分割成 i + (n - i)
	for i := 1; i <= n-1; i++ {
		// 先直接判断 i 与 (n - i) 的乘积
		res = Helper.MaxInt(i*(n-i), res)

		// 再分割n-i，获取n-i分割后乘积的最大值，再乘以i
		res = Helper.MaxInt(i*s.breakInt(n-i), res)
	}

	s.memo[n] = res
	return res
}

// 动态规划: 自底向上解决问题
func integerBreak2(n int) int {
	if n < 2 {
		return 0
	}

	// memo[i] 表示将数字i分割（至少分割成两部分）后得到的最大乘积
	memo := make([]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = -1
	}

	// 先解决小的问题
	memo[1] = 1
	for i := 2; i <= n; i++ {
		// 分割memo[i]
		for j := 1; j <= i-1; j++ {
			// j + (i-j)
			memo[i] = Helper.MaxInt2(memo[i], j*(i-j), j*memo[i-j])
		}
	}

	return memo[n]
}
