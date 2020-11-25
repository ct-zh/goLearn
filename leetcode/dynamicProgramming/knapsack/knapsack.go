package knapsack

import (
	Helper "github.com/LannisterAlwaysPaysHisDebts/goLearn/leetcode/base/helper"
)

// 0-1 背包问题
// 有一个背包,容量为C.现有n种不同的物品编号为0...n-1,
// 其中每一件物品的重量为w(i),价值为v(i).
// 问可以向这个背包中盛放哪些物品,使得在不超过背包容量的基础上,
// 物品的总价值最大.

// 暴力解法: 每一件物品都可以放进背包,也可以不放进背包.  O((2^n)*n)

// F(n, C)考虑将n个物品放入容量为C的背包,使得价值最大
// 情况一,不需要考虑第i件物品,则问题转换为F(i-1, c)
// 情况二, 需要考虑第i件物品,则问题转换为v(i) + F(i-1, c-w(i))
// F(i, c) = max( F(i-1, c),  v(i) + F(i-1, c-w(i))  )

// w为每件物品的总重量；v为价值；背包容量为C
// 需要返回物品的最大总价值
// 时间复杂度 O(n*C)  空间复杂度O(n*C)
func knapsack(w []int, v []int, C int) int {
	wLen := len(w)
	k := knapsack1{
		memo: make([][]int, wLen),
	}
	for i := 0; i < wLen; i++ {
		k.memo[i] = make([]int, C+1)
		for j := 0; j <= C; j++ {
			k.memo[i][j] = -1
		}
	}
	//defer fmt.Println(k.memo)
	return k.bestValue(w, v, wLen-1, C)
}

type knapsack1 struct {
	memo [][]int // 因为有两个约束条件:物品、剩余容积; 所以memo是一个二维数组,第一个key代表物品,第二个key代表剩余容量. 即memo[i][j] 代表第i个物品放入容量为j的背包时可以得到的最大的价值
}

// 用[0...index]的物品，填充容积为c的背包的最大价值
func (k *knapsack1) bestValue(w []int, v []int, index int, c int) int {
	// 如果当前已经没有了物品/没有了容量
	if index < 0 || c <= 0 {
		return 0
	}

	if k.memo[index][c] != -1 {
		return k.memo[index][c]
	}

	// 情况1, 无视该物品
	res := k.bestValue(w, v, index-1, c)
	if c >= w[index] { // 如果该物品可以装进背包
		// 情况2，包含该物品
		res2 := k.bestValue(w, v, index-1, c-w[index]) + v[index]
		res = Helper.MaxInt(res, res2)
	}

	k.memo[index][c] = res
	return res
}

// 方法二： 动态规划
// 时间复杂度 O(n*C)  空间复杂度O(n*C)
func knapsack2(w []int, v []int, C int) int {
	wLen := len(w)
	if wLen == 0 {
		return 0
	}

	memo := make([][]int, wLen)
	for i := 0; i < wLen; i++ {
		memo[i] = make([]int, C+1)
		for j := 0; j <= C; j++ {
			memo[i][j] = -1
		}
	}

	// 先填充第一个物品放入容量为j的背包中的可能性
	for j := 0; j <= C; j++ {
		if j >= w[0] {
			memo[0][j] = v[0]
		} else {
			memo[0][j] = 0
		}
	}

	for i := 1; i < wLen; i++ {
		for j := 0; j <= C; j++ {
			memo[i][j] = memo[i-1][j]
			if j >= w[i] {
				memo[i][j] = Helper.MaxInt(memo[i][j], v[i]+memo[i-1][j-w[i]])
			}
		}
	}
	return memo[wLen-1][C]
}

// 优化1 空间复杂度到 O(C)
func knapsack3(w []int, v []int, C int) int {
	wLen := len(w)
	if wLen == 0 {
		return 0
	}

	memo := make([][]int, 2)
	for i := 0; i < 2; i++ {
		memo[i] = make([]int, C+1)
		for j := 0; j <= C; j++ {
			memo[i][j] = -1
		}
	}

	for j := 0; j <= C; j++ {
		if j >= w[0] {
			memo[0][j] = v[0]
		} else {
			memo[0][j] = 0
		}
	}

	for i := 1; i < wLen; i++ {
		for j := 0; j <= C; j++ {
			memo[i%2][j] = memo[(i-1)%2][j]
			if j >= w[i] {
				memo[i%2][j] = Helper.MaxInt(memo[i%2][j], v[i]+memo[(i-1)%2][j-w[i]])
			}
		}
	}
	return memo[(wLen-1)%2][C]
}

// 优化2 只使用一维数组 空间时间都减小了
func knapsack4(w []int, v []int, C int) int {
	wLen := len(w)
	if wLen == 0 {
		return 0
	}

	memo := make([]int, C+1)
	for j := 0; j <= C; j++ {
		memo[j] = -1
	}

	for j := 0; j <= C; j++ {
		if j >= w[0] {
			memo[j] = v[0]
		} else {
			memo[j] = 0
		}
	}

	for i := 1; i < wLen; i++ {
		for j := C; j >= w[i]; j-- {
			memo[j] = Helper.MaxInt(memo[j], v[i]+memo[j-w[i]])
		}
	}
	return memo[C]
}
