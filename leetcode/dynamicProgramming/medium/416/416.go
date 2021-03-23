package _416

// 背包问题的变种
// 在n个物品中选定一定物品，填满sum/2背包
// F(n, C) 考虑将n个物品填满容量为C的背包
// 状态转移方程：
// F(i, c) = F(i-1, c) || F(i-1, c-w(i))
// 时间复杂度 O(n*sum/2) = O(n*sum )

func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum%2 != 0 { // 不能被平分
		return false
	}

	s := solution{
		memo: make([][]int, len(nums)),
	}
	for i := 0; i < len(nums); i++ {
		s.memo[i] = make([]int, sum/2+1) //背包总大小
		for j := 0; j < sum/2+1; j++ {
			s.memo[i][j] = -1
		}
	}

	// 尝试对数字进行分割， 数字的个数len(nums)-1, 背包大小为sum/2
	return s.tryP(nums, len(nums)-1, sum/2)
}

type solution struct {
	// memo[i][c]表示使用索引为[0...i]的这些元素，是否可以完全填充一个容量为c的背包
	// 用-1表示未计算； 0 表示不可以填充； 1 表示可以填充
	memo [][]int
}

// 使用nums[0...index] 是否可以完全填充一个容量为sum的背包
func (s *solution) tryP(nums []int, index int, sum int) bool {
	if sum == 0 { // 已经填充满了
		return true
	}
	if sum < 0 || index < 0 { // 背包装不下||没有物品可选，背包还没满
		return false
	}
	if s.memo[index][sum] != -1 {
		return s.memo[index][sum] == 1
	}

	if s.tryP(nums, index-1, sum) ||
		s.tryP(nums, index-1, sum-nums[index]) {
		s.memo[index][sum] = 1
		return true
	} else {
		s.memo[index][sum] = 0
		return false
	}
}

// 动态规划
func canPartition2(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}

	n := len(nums)
	C := sum / 2

	memo := make([]bool, C+1)

	for i := 0; i <= C; i++ {
		memo[i] = nums[0] == i
	}
	for i := 1; i < n; i++ {
		for j := C; j >= nums[i]; j-- {
			memo[j] = memo[j] || memo[j-nums[i]]
		}
	}

	return memo[C]
}
