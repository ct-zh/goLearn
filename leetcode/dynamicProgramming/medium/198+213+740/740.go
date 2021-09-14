package _98_213_740

// 该题可转换思路，根据t2例子：[2, 2, 3, 3, 3, 4]
// 假如获取了分数3，则2和4都不能获取了，想要拿到更高分数，就只能把剩下的3全拿了
// 此时题目可以转换为：[4(2+2), 9(3+3+3), 4] 的打家劫舍问题，不能偷相邻的房屋

func deleteAndEarn(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	maxInt := 0
	for _, num := range nums {
		maxInt = max(num, maxInt)
	}

	rob := make([]int, maxInt+1)
	for _, num := range nums {
		rob[num] += num
	}

	p, q := rob[0], max(rob[1], rob[0])
	for i := 2; i <= maxInt; i++ {
		p, q = q, max(rob[i]+p, q)
	}

	return q
}
