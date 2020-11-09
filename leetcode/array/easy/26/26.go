package _6

// 删除数组里的重复项
//  https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/

// 思路一，双指针
// 指针p1做遍历，
// 指针p2 与p1的值做比较，如果不相等，则指针1的值与 指针2的key+1的值做交换(同样需要判断key+1与p1的值是否相同，防止自己交换自己的浪费操作)
// 细节：1. 如果key+1 == p1 则不进行交换，但是仍然还是需要进行p2++
// 2. 最后返回p2应该+1，因为返回的是子数组的长度，而数组是从0开始的
// 时间复杂读O(n)
// 空间复杂度O(1)
//  执行用时： 8 ms, 在所有 Go 提交中击败了88.14%的用户
//	内存消耗：4.6 MB, 在所有 Go 提交中击败了39.58%的用户
// 内存消耗应该在4.4MB为最佳，不知道怎么降
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	index := 0
	for i := 1; i < len(nums); i++ {
		if nums[index] != nums[i] {
			if index+1 != i {
				nums[index+1], nums[i] = nums[i], nums[index+1]
			}
			index++
		}
	}

	return index + 1
}
