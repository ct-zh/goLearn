package leetcode55_45

func jump(nums []int) int {
	l := len(nums)
	maxPlace := 0 // 能跳最远的位置
	end := 0      // 当前不跳跃，可以存在的最远位置
	step := 0     // 跳跃次数
	for i, num := range nums {
		if i > end { // 如果当前节点不在「不跳跃可以存在在最远位置」内，则跳跃，并更新
			end = maxPlace
			step += 1
		}
		if i+num > maxPlace { // 更新最远能跳位置
			maxPlace = i + num
		}
		//fmt.Println("i=", i, "end=",end, "step=",step, "maxPlace=",maxPlace)
		if maxPlace >= l-1 { // 如果最终节点已经在跳跃范围内了，则直接跳跃
			return step + 1
		}
	}
	return step
}
