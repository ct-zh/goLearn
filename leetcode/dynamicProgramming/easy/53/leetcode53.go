package leetcode53

// 动态规划解题思路：
// 假设，求f(x)是第x个数结尾的最大子序和；
// 那么问题可以转变为比较f(x)大还是f(x-1)大；并且f(x) = num[x] + f(x-1);
// 状态转移方程为 f(x) = max(f(x), f(x-1)) = max(f(x-1) + num[x], f(x-1))
// 边界条件为 f(0) = num[0]

func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}
