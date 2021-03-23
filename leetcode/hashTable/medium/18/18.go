package _18

import "sort"

// https://leetcode-cn.com/problems/4sum/

// 思路1：基于三数之和(15题) 外面多套一层循环
// 通过三次剪枝，速度从18ms提升到4ms
// 执行用时： 4 ms, 在所有 Go 提交中击败了 97.42% 的用户
// 内存消耗： 2.8 MB, 在所有 Go 提交中击败了 90.77% 的用户
// 时间复杂度 O(n^3 + nlogn) 排序 + 三重循环
// 空间复杂度 O(n) 应该是用于排序了
func fourSum(nums []int, target int) (res [][]int) {
	l := len(nums)
	sort.Ints(nums)

	// 剪枝3. i 到 l-3就结束了
	for i := 0; i < l-3; i++ {
		// 剪枝1. 跳过重复的数字，防止重复判断
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 剪枝2。 提前跳过不可能的组合
		// 最大的三个值加上i也小于target
		if nums[i]+nums[l-1]+nums[l-2]+nums[l-3] < target {
			continue
		}
		// i 加上最小的三个值也大于target，说明剩下的所有组合都大于target
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}

		// 多套一层j， 四数之和就变成了求 i1+i2 = target-i-j
		// 剪枝3. j 到 l-2就结束了
		for j := i + 1; j < l-2; j++ {
			// 剪枝1. 依然也要跳过重复的数
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			itemSum := target - nums[i] - nums[j]

			// 剪枝2. 提前跳过不可能的组合, 即i+j 加最大的两个数也小于target 或者i+j加剩下两个最小的数也大于target
			if nums[l-1]+nums[l-2] < itemSum || nums[j+1]+nums[j+2] > itemSum {
				continue
			}

			// 双指针从两端开始检索, 直至相遇
			i1 := j + 1
			i2 := l - 1
			for i1 < i2 {
				// 剪枝1. 依然也要跳过重复的数
				if i1 > j+1 && nums[i1] == nums[i1-1] {
					i1++
					continue
				}
				if i2 < l-1 && nums[i2] == nums[i2+1] {
					i2--
					continue
				}

				if nums[i1]+nums[i2] == itemSum { // 刚好相等
					res = append(res, []int{nums[i], nums[j], nums[i1], nums[i2]})
					i1++
					i2--
				} else if nums[i1]+nums[i2] > itemSum {
					i2--
				} else {
					i1++
				}
			}
		}
	}

	return
}

// todo: 思路2 DFS
