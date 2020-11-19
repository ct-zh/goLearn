package _80

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii/
// 数组升序 每个元素最多出现两次 要求空间复杂度O(1)

// 思路： 双指针
// p1做遍历，同时计算当前重复元素的个数count; p2放在count>2的位置；
// nums[p2] = nums[p1], p2++, p1++
// 最后返回p2+1
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	count := 1 // 计数器
	index := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count++
		} else { // nums[i] != nums[i-1]
			count = 1
		}

		if count <= 2 {
			if nums[index] != nums[i] {
				nums[index] = nums[i]
			}
			index++
		}
	}

	return index
}
