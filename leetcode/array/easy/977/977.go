package _77

// https://leetcode-cn.com/problems/squares-of-a-sorted-array/

func sortedSquares(nums []int) []int {
	i1, i2, i3 := 0, len(nums)-1, len(nums)-1
	answer := make([]int, len(nums))
	for i1 <= i2 {
		if nums[i1]*nums[i1] > nums[i2]*nums[i2] {
			answer[i3] = nums[i1] * nums[i1]
			i1++
		} else {
			answer[i3] = nums[i2] * nums[i2]
			i2--
		}
		i3--
	}
	return answer
}
