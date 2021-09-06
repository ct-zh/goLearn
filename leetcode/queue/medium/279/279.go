package _279

import (
	"fmt"
	Helper "github.com/ct-zh/goLearn/leetcode/basic/helper"
	"github.com/ct-zh/goLearn/leetcode/basic/queue"
	"math"
)

// 没有解怎么办？=> 不可能没有解, 因为完全平方数里最小为1， 所有正整数n都能由1构成

// 完全背包问题的思路：物品i的重量为i^2(即把物品看成完全平方数)，所有物品的价值为1
// 问题转换成将容量为n的背包装满，价值最小的装法。

// 思路1： 记忆化搜索；
// 例如求 7、8、9, 知F(0)=0,
// F(7) = min(F(7-1)+1, F(7-4)+1)
// F(8) = min(F(8-1)+1, F(8-4)+1)
// F(9) = min(F(9-1)+1, F(9-4)+1, F(9-9)+1)
// 执行用时： 80 ms, 在所有 Go 提交中击败了 25.28% 的用户
// 内存消耗： 7.1 MB, 在所有 Go 提交中击败了 13.32% 的用户
// 时间复杂度
// 空间负载度 O(n) 额外开辟了一个n的空间
func numSquares1(n int) int {
	if n <= 0 {
		panic("error")
	}

	s := solution{
		memo: make([]int, n+1),
	}

	// 初始化 0 = 0
	s.memo[0] = 0
	for i := 1; i <= n; i++ {
		s.memo[i] = -1
	}

	return s.square(n)
}

type solution struct {
	memo []int
}

// 求组成数n的最少完全平方数
func (s *solution) square(n int) int {
	if s.memo[n] >= 0 {
		return s.memo[n]
	}

	sq := math.Sqrt(float64(n))
	res := -1
	for i := 1; i <= int(sq); i++ {
		// 将n分成n-i*i 与 i*i, i*i是一个完全平方数，需要加1，
		// 问题转换为求n-i*i最少组成的完全平方数
		item := s.square(n-i*i) + 1
		if res == -1 {
			res = item
		} else {
			res = Helper.MinInt(item, res)
		}
		//fmt.Printf("n=%d i=%d deep=%d res=%d \n", n, i, deep+1, res)
	}

	s.memo[n] = res
	return res
}

// 思路1 的动态规划写法
// 执行用时： 40 ms , 在所有 Go 提交中击败了 55.58% 的用户
// 内存消耗： 6.1 MB , 在所有 Go 提交中击败了 47.22% 的用户
// 时间复杂度 O(n*sqrt(n))
// 空间复杂度 O(n)
func numSquares2(n int) int {
	if n <= 0 {
		panic("error")
	}

	memo := make([]int, n+1)
	memo[0] = 0

	for i := 1; i <= n; i++ {
		sq := math.Sqrt(float64(i))
		memo[i] = -1
		for j := 1; j <= int(sq); j++ {
			if memo[i] == -1 {
				memo[i] = memo[i-j*j] + 1
			} else {
				memo[i] = Helper.MinInt(memo[i-j*j]+1, memo[i])
			}
		}
	}

	return memo[n]
}

// 思路1 的动态规划版本, 去掉math包的写法
// 执行用时： 36 ms , 在所有 Go 提交中击败了 61.28% 的用户
// 内存消耗： 6 MB , 在所有 Go 提交中击败了 62.22% 的用户
// 时间复杂度 O(n*sqrt(n)), 外层循环n，内层循环根号n
// 空间复杂度 O(n)
func numSquares2Better(n int) int {
	if n <= 0 {
		panic("error")
	}

	memo := make([]int, n+1)
	memo[0] = 0

	// 记录小于等于n的所有完全平方数
	var nums []int
	for i := 1; i*i <= n; i++ {
		nums = append(nums, i*i)
	}

	for i := 1; i <= n; i++ {
		memo[i] = -1
		for _, value := range nums {
			if value > i {
				break
			}
			if memo[i] == -1 {
				memo[i] = memo[i-value] + 1
			} else {
				memo[i] = Helper.MinInt(memo[i-value]+1, memo[i])
			}
		}
	}

	fmt.Println(memo)

	return memo[n]
}

// 思路2
// 对问题做建模，转换成图论的问题
// 从n到0，每个数组表示一个节点
// 如果两个数字x到y相差一个完全平方树，则连接一条边
// 我们得到了一个无权图
// 原问题就转换成了 求这个无权图中从n到0的最短路径问题
type pair struct {
	p1 int // 表示数字是多少
	p2 int // 图中经过了几段路径走到了这个数字
}

func numSquares(n int) int {
	if n <= 0 {
		panic("参数非法")
	}

	q := queue.Queue{}
	q.Push(pair{
		p1: n,
		p2: 0, // 从n走到n，一步都不需要走
	})

	// 用于过滤重复的节点
	visited := map[int]bool{}
	visited[n] = true

	for !q.IsEmpty() {
		top := q.Pop()
		num := top.(pair).p1  // 表示数字是多少
		step := top.(pair).p2 // 表示从n到该数字，已经走了几步了

		// num这个数还能承受一个完全平方数
		// 完全平方数： i*i 即[1,4,9,16,...]
		for i := 1; ; i++ {
			a := num - i*i
			if a < 0 { // 计算的结果不能小于0
				break
			}
			if a == 0 { // 说明刚好到0点，直接返回
				return step + 1
			}
			if _, ok := visited[a]; !ok {
				q.Push(pair{
					p1: a,        // 移动到下一个数
					p2: step + 1, // 步数+1
				})
				visited[a] = true
			}
		}
	}

	return 0
}
