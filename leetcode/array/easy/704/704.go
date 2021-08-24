package _04m

// 给定一个n个元素有序的（升序）整型数组nums 和一个目标值target ，写一个函数搜索nums中的 target，如果目标值存在返回下标，否则返回 -1。

//链接：https://leetcode-cn.com/problems/binary-search

func search(nums []int, target int) int {
	right := len(nums) - 1
	left := 0
	for left <= right {
		i := (right-left)/2 + left
		if target == nums[i] {
			return i
		} else if target > nums[i] {
			left = i + 1
		} else {
			right = i - 1
		}
	}
	return -1
}
