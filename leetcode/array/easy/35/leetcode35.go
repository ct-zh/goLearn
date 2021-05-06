package leetcode35

// https://leetcode-cn.com/problems/search-insert-position/

// 二分搜索

func searchInsert(nums []int, target int) int {
	return search(nums, target, 0, len(nums)-1)
}

func search(nums []int, target int, start, end int) int {
	if end <= start {
		if target > nums[end] {
			return end + 1
		} else {
			return end
		}
	}
	g := (end-start)/2 + start
	if target == nums[g] {
		return g
	} else if target < nums[g] {
		return search(nums, target, start, g)
	} else {
		return search(nums, target, g+1, end)
	}
}

func searchInsert1(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (r-l)/2 + l
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
