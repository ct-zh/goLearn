package leetcode55_45

// 方法一 阅读理解，把0当坑，不能跳过坑则返回false
func canJump(nums []int) bool {
	l := len(nums)
	if l <= 1 {
		return true
	}
	maxPlace := 0
	for i, num := range nums {
		if num == 0 && i >= maxPlace { // 只有当num为0时才会跳不出去
			return false
		}
		if i+num > maxPlace {
			maxPlace = i + num
		}
		if maxPlace >= l-1 { // 可以跳到最后一层，直接return
			return true
		}
	}
	return true
}
