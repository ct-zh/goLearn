package _80

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii/
// 数组升序 每个元素最多出现两次 要求空间复杂度O(1)

// 三指针
func removeDuplicates(nums []int) int {
	l := len(nums)
	if l <= 2 {
		return l
	}

	before, p := 0, 1
	for i := 2; i < l; i++ {
		if nums[i] != nums[p] {
			nums[p+1] = nums[i]
			before++
			p++
		} else if nums[before] != nums[i] {
			nums[p+1] = nums[i]
			before++
			p++
		}
	}

	return p + 1
}

// 思路：双指针
// p1做遍历，同时计算当前重复元素的个数count; p2放在count>2的位置；
// nums[p2] = nums[p1], p2++, p1++
// 最后返回p2+1
func removeDuplicates1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	count := 1
	index := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] { // 维护count的定义：i指向的元素的当前计数个数，如果i发生了改变，则需要重新计数；
			count++
		} else { // nums[i] != nums[i-1]
			count = 1
		}
		if count <= 2 { // 维护p的定义：
			if nums[index] != nums[i] { // 位置在重复元素的第三个
				nums[index] = nums[i]
			}
			index++ // 跟着i遍历,或者始终指向第三个 [1 1 1 1 2] index=2 => [1 1 2 1 2] index=3
		}
	}

	return index
}

// 思路2，双指针, 这种解法思路清晰，适合任何重复n项的题
func removeDuplicates2(nums []int) int {
	p := 0
	for i := 0; i < len(nums); i++ {
		if p < 2 || (nums[i] != nums[p-2]) {
			nums[p] = nums[i]
			p++
		}
	}
	return p
}
