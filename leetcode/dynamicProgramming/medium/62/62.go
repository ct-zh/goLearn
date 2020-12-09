package _62

// 动态规划 抄答案
// 因为只能向下或者向右走， 对于某个格子f(i, j), 他的来源只可能是上一格或者左一格
// 因此有动态转移方程：f(i, j) = f(i-1, j) + f(i, j-1)
// 剪枝： 如果i=0或者j=0 则动态转移方程变为: f(0,j) = f(0, j-1) 和 f(i,0)=f(i-1,0)
// 时间复杂度： O(mn)
// 空间复杂度: O(mn) 可以用滚动数组进行优化
func uniquePaths1(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		// 到边界的点只有一条路，直接往下走或者直接往右走
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

// dfs： 超时
func uniquePaths(m int, n int) int {
	// 从0, 0开始 到 m, n
	s := solution{
		// 方向： 只能往下或者往右
		d: [4][2]int{
			{0, 1},
			{1, 0},
		},
		m: m,
		n: n,
	}

	return s.dfs(0, 0, 0)
}

type solution struct {
	d    [4][2]int
	m, n int
}

func (s solution) dfs(x, y int, sum int) int {
	if x == s.m-1 && y == s.n-1 {
		return sum + 1
	}
	for i := 0; i < 2; i++ {
		newX := s.d[i][0] + x
		newY := s.d[i][1] + y
		if newX >= 0 && newX < s.m && newY >= 0 && newY < s.n {
			sum = s.dfs(newX, newY, sum)
		}
	}
	return sum
}
