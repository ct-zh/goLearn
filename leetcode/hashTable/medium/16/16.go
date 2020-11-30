package _16

import (
	"math"
	"sort"
)

// 思路同15题
//
func threeSumClosest(nums []int, target int) int {
	l := len(nums)
	res := math.MaxInt64
	sort.Ints(nums)
	for i := 0; i < l-2; i++ {
		// 跳过重复的数字
		if i >= 1 && nums[i] == nums[i-1] {
			continue
		}

		lt, gt := i+1, l-1
		itemRes := nums[i] + nums[lt] + nums[gt]
		if math.Abs(float64(target-itemRes)) < math.Abs(float64(target-res)) {
			res = itemRes
		}

		for lt < gt {
			// 跳过重复的数字判断
			if lt > i+1 && lt < l-1 && nums[lt] == nums[lt-1] {
				lt++
				continue
			}
			// 跳过重复的数字判断
			if gt < l-1 && gt-1 > lt && nums[gt] == nums[gt+1] {
				gt--
				continue
			}

			itemRes := nums[i] + nums[lt] + nums[gt]
			//fmt.Printf("i: %d lt:%d gt:%d count:%d countdiff%d mix:%d \n", nums[i], nums[lt], nums[gt], itemRes, target-itemRes, target-res)
			if math.Abs(float64(target-itemRes)) < math.Abs(float64(target-res)) {
				res = itemRes
			}
			if itemRes > target { // 太大了
				gt--
			} else {
				lt++
			}
		}
	}
	return res
}
