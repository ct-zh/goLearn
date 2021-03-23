package _46

import "fmt"

// 排列问题
// perms(nums[0...n-1])={取出一个数字}+Perms(nums[{0...n-1}-这个数字])
func permute(nums []int) (res [][]int) {
	if len(nums) == 0 {
		return
	}

	used := make([]bool, len(nums))
	generatePermutation(nums, 0, []int{}, &res, used)
	return
}

// p中保存了一个有index个元素的排列
// 向这个排列的末尾添加第index+1个元素，获得一个有index+1个元素的排列
func generatePermutation(
	nums []int,
	index int,
	p []int,
	result *[][]int,
	used []bool) {
	fmt.Printf("result: %+v   p: %+v  index: %d  used: %+v \n", result, p, index, used)

	if index == len(nums) {
		fmt.Printf("\nbefore: %+v\n", result)
		*result = append(*result, p)
		fmt.Printf(" after: %+v p:%+v \n\n", result, p)
		return
	}

	for i := 0; i < len(nums); i++ {
		// nums[i]是否在p中 => 不能重复
		if !used[i] {
			p = append(p, nums[i])
			used[i] = true
			generatePermutation(nums, index+1, p, result, used)

			// 执行完之后需要删除当前数据i
			p = p[0 : len(p)-1] // 删除最后一项
			fmt.Printf("before result: %+v used: %+v i: %d \n", result, used, i)
			used[i] = false     // 第i项数据设置为false
			fmt.Printf("after  result: %+v used: %+v i: %d \n", result, used, i)
		}
	}
	return
}
