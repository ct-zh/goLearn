package _198

// 暴力解法： 遍历所有可能性，筛选掉会出发报警的； O((2^n)*n)

// 递归解法
// 考虑偷取[0...n-1]范围内的所有房子
// 偷取0： 问题变成偷取[2...n-1]范围内的房子
// 根据对状态的定义，决定状态的转移：
// f(0) = max{v(0) + f(2), v(1) + f(3), v(2) + f(4), ...,
// 			v(n-3) + f(n-1), v(n-2), v(n-1)} (状态转移方程)
func rob(nums []int) int {

}

type solution struct {
}

// 考虑抢劫nums[index...len(nums)]这个范围内的所有房子
func (s solution) tryRob(nums *[]int, index int) int {

}
