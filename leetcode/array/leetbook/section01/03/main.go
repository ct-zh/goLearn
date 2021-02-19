package _3

// 见26题

// 我的思路：双指针，见26题
func removeDuplicates(nums []int) int {
	l := len(nums)
	f := 0
	for i := 0; i < l; i++ {
		// 边界问题， f+1 不能等于l
		if nums[i] != nums[f] {
			if f+1 != i && f+1 < l {
				nums[i], nums[f+1] = nums[f+1], nums[i]
			}
			f++
		}
	}
	return f + 1
}
