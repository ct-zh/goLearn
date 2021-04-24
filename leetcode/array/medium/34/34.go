package _34

import (
	Helper "github.com/LannisterAlwaysPaysHisDebts/goLearn/leetcode/basic/helper"
)

// 思路1，二分算法
// 执行用时： 8 ms, 在所有 Go 提交中击败了 90.40% 的用户
// 内存消耗： 4.1 MB, 在所有 Go 提交中击败了 10.19% 的用户
func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	l := len(nums)
	if l <= 0 {
		return res
	}

	min, max := find(nums, 0, l-1,
		l, -1, target)
	if min != l {
		res[0] = min
	}
	if max != -1 {
		res[1] = max
	}

	return res
}

func find(arr []int, l int, r int, min int, max int, target int) (int, int) {
	if l > r {
		return min, max
	}
	mid := (r-l)/2 + l
	if target == arr[mid] {
		min = Helper.MinInt(min, mid)
		max = Helper.MaxInt(max, mid)
	}

	min1, max1 := find(arr, l, mid-1, min, max, target)
	min2, max2 := find(arr, mid+1, r, min, max, target)
	return Helper.MinInt2(min, min1, min2), Helper.MaxInt2(max, max1, max2)
}
