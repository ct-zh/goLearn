package leetcode53

// 如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。

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
