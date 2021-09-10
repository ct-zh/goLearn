package _46

// https://leetcode-cn.com/problems/min-cost-climbing-stairs/

// 同70题；
// 到x层最小的体力消耗，为 (x-1层cost) + (x-1到x层cost) 与 (x-2层cost) + (x-2层到x层cost) 的最小值
// 状态转移方程为: f(x) = min( f(x-1) + cost(x-1), f(x-2) + cost(x-2) )

// 举例， 如 test1 [10, 15, 20]
// 则到达楼顶 f(3) = min( f(2) + cost(2), f(1) + cost(1) )
// => min(10 + 20, 0 + 15) => 15
// 走法：[1, 2, 1] 消费30 , 或者 [2, 2] 消费 15

// f(2) = min(f(1) + cost(1), f(0) + cost(0))
// =>  min(0 + 15, 0 + 10) => 10
// 走法：[2, 1] 消费15 或者 [1, 2] 消费10

// f(1) = min(f(0) + cost(0), 0)
// => min(10, 0) => 0
// 走法： [1, 1] 消费10 或者 [2] 消费0

// 法1 记忆话搜索
func minCostClimbingStairs(cost []int) int {
	top := len(cost) // 楼顶

	mem := make([]int, top+1)
	for i := 2; i <= top; i++ {
		mem[i] = min(mem[i-1]+cost[i-1], mem[i-2]+cost[i-2])
	}

	return mem[top]
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

// 法2 动态规划
func minCostClimbingStairsDynamic(cost []int) int {
	p, q := 0, 0
	for i := 2; i < len(cost)+1; i++ {
		p, q = q, min(p+cost[i-2], q+cost[i-1])
	}
	return q
}
