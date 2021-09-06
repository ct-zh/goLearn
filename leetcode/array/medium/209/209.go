package _209

import Helper "github.com/ct-zh/goLearn/leetcode/basic/helper"

// https://leetcode-cn.com/problems/minimum-size-subarray-sum/

// 思路：滑动窗口
func minSubArrayLen(s int, nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	min, i1, i2 := l+1, 0, -1

	now := 0
	for i1 < l {
		if i2+1 < l && now < s {
			i2++
			now += nums[i2]
		} else {
			// 说明： 右边界已经达到末尾
			// 或者now 大于等于2
			now -= nums[i1]
			i1++
		}

		if now >= s {
			min = Helper.MinInt(i2-i1+1, min)
		}
	}

	if min > l {
		min = 0
	}

	return min
}
