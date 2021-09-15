package _1maxSubArr

import (
	Helper "github.com/ct-zh/goLearn/leetcode/basic/helper"
)

// 仍然是 最大子序和问题
// 对于一个环形数组求最大子序和，直接将该题目拆分成两种情况：
//1. 如果最大子序在数组内，则该题目转换成最基础的求数组最大子序和的问题；
//2. 如果最大子序和横跨了数组，则说明最大子序包括了数组的首节点`nums[0]`与数组的尾节点`nums[l-1]`，问题是否能转换为求数组最小子序和的问题？

func maxSubarraySumCircular(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return nums[0]
	}
	maxEnding, maxSum, minEnding, minSum := nums[0], nums[0], nums[0], nums[0]
	for i := 1; i < l; i++ {
		maxEnding = max(nums[i], maxEnding+nums[i])
		maxSum = max(maxSum, maxEnding)
		minEnding = Helper.MinInt(nums[i], nums[i]+minEnding)
		minSum = Helper.MinInt(minEnding, minSum)
	}

	max2 := Helper.SumInt(nums) - minSum
	if max2 == 0 {
		max2 = maxSum
	}

	return max(maxSum, max2)
}
