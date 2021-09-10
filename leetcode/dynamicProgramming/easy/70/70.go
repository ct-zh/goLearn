package _70

// 解题思路：
// 我们用 f(x) 表示爬到第 x 级台阶的方案数，考虑最后一步可能跨了一级台阶，也可能跨了两级台阶，
// 所以我们可以列出如下式子：
// f(x) = f(x-1) + f(x-2)
// 例如，爬到3级台阶的方案数，相当于爬1级台阶的方案数[1] + 爬两级台阶的方案数[1,1]、[2] = 3
// 再例： x = 5, 有
// [1, 1, 1, 1, 1]
// [1, 1, 1, 2], [1, 1, 2, 1], [1, 2, 1, 1], [2, 1, 1, 1]
// [1, 2, 2] [2, 1, 2], [2, 2, 1]
// 所以
// f(5) = 8
// f(5) = f(4) + f(3)
// f(4): [1,1,1,1] [1,1,2], [1,2,1], [2,1,1],[2,2]
// f(5) = 5 + 3 = 8 check!

// 法一，递归
func climbStairsByRecursion(n int) int {
	return _climb(n)
}

func _climb(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return _climb(n-1) + _climb(n-2)
}

// 优化 法二，记忆化搜索;
// 上面的递归明显重复计算了，所以要带记忆去递归
func climbStairs(n int) int {
	// 初始化，没计算的都初始化为 -1
	memo := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		memo[i] = -1
	}
	// 先填充边界条件
	memo[0], memo[1], memo[2] = 0, 1, 2

	// 计算
	for i := 3; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}

	return memo[n]
}

// 法三，动态规划
// 上面记忆话搜索空间复杂度为O(n)，可以进一步优化为O(1)
func climbStairs2(n int) int {
	p, q, r := 0, 1, 2
	for i := 3; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

// 可以跑一下 上面3种方法的benchmark
// 其中，法1不带记忆话的递归，时间复杂度和空间复杂度最高
// 动态规划性能最优
