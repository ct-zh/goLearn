package _70

// 记忆化搜索
func climbStairs(n int) int {
	s := solution{memo: make([]int, n+1)}
	for i := 0; i < n+1; i++ {
		s.memo[i] = -1
	}
	return s.memo[n]
}

type solution struct {
	memo []int
}

func (s solution) calcWays(n int) int {
	if n == 0 || n == 1 { // // 一阶台阶都不走或者只走一节
		return 1
	}

	if s.memo[n] == -1 {
		s.memo[n] = s.calcWays(n-1) + s.calcWays(n-2)
	}

	return s.memo[n]
}

// 动态规划
func climbStairs2(n int) int {
	memo := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		memo[i] = -1
	}
	memo[0], memo[1] = 1, 1
	for i := 2; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[n]
}
