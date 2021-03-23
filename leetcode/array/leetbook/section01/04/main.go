package _4

import "fmt"

// 见80题

// 相比上一个题目， 增加一个count计数
func removeDuplicates(nums []int) int {
	l := len(nums)
	f := 1
	count := 1

	// 从1开始
	for i := 1; i < l; i++ {
		if nums[i-1] == nums[i] {
			count++
		} else {
			count = 1
		}

		if count <= 2 {
			if nums[f] != nums[i] {
				nums[f] = nums[i]
			}
			f++
		}

		fmt.Println(nums, "  f: ", f, "  i: ", i, " count: ", count)
	}

	return f
}
