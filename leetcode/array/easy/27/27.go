package _7

// https://leetcode-cn.com/problems/remove-element/

// 思路：双指针
// 快指针p1做遍历
// 慢指针p2指向数组中 第一个值为val的位置
// 当p1指向的值不为val时，p1p2的值做交换，p2++
// 细节：1.当数组里只有一个元素时；2.当数组里没有val时
// 时间复杂度O(n)，空间复杂度O(1)
// 执行用时： 0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗： 2.1 MB, 在所有 Go 提交中击败了 36.80% 的用户
func removeElement(nums []int, val int) int {
	// index指向第一个val的位置
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			if nums[index] == val {
				nums[index], nums[i] = nums[i], nums[index]
			}
		}

		// 因为如果nums[index] == val 符合上面的条件，两个数值会发生交换
		// 交换后nums[index] 又 != val 了
		if nums[index] != val {
			index++
		}
	}
	return index
}

// 解法二：照抄的标准答案
// 对比我的解法简洁很多，但是忽略了i=j时仍然做无效交换导致性能浪费的问题
// 各有千秋，不要妄自菲薄
func removeElement2(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
