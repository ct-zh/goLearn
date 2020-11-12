package main

// https://leetcode-cn.com/problems/sort-colors/
// 要考虑用户传进来的数据非法的情况

// 思路一:计数排序
// 只适用于数据量小的时候 ？
// 时间复杂度O(n)
// 空间复杂度，O(k),k为元素的取值范围，这里是3，所以是常数O(1)
func sortColors(nums []int) {
	set := struct {
		r int
		w int
		b int
	}{
		r: 0, w: 0, b: 0,
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			set.r++
		} else if nums[i] == 1 {
			set.w++
		} else if nums[i] == 2 {
			set.b++
		} else {
			panic("数据不符合规范") // 要考虑到用户传的数据非法的情况
		}
	}

	for i := 0; i < len(nums); i++ {
		if set.r > 0 {
			nums[i] = 0
			set.r--
		} else if set.w > 0 {
			nums[i] = 1
			set.w--
		} else {
			nums[i] = 2
		}
	}

}

// 思路二：三路快排
func sortColors2(nums []int) {
	zero := -1       // nums[0...zero] == 0	初始化的时候，没有一个元素是等于0
	two := len(nums) // nums[two...n-1] == 2 初始的时候，没有一个元素是等于2

	for i := 0; i < two; {
		if nums[i] == 1 {
			i++
		} else if nums[i] == 2 {
			two -= 1
			nums[i], nums[two] = nums[two], nums[i]
		} else if nums[i] == 0 {
			zero += 1
			nums[i], nums[zero] = nums[zero], nums[i]
			i++
		} else {
			panic("传入数据非法")
		}
	}

}

func main() {

}
