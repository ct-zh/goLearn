package _98_213

// f(x)代表第x间房可能的最高偷窃金额，那么
// f(x) = max( f(x-2) + nums[x], f(x-1))
// 边界条件： x = 1, x = 2
// 两个题目不管是遍历1遍还是两遍，都是O(n)的时间复杂度，O(1)的空间复杂度

func rob(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	if l == 2 {
		return max(nums[0], nums[1])
	}

	p, q, r := 0, nums[0], nums[1]
	for i := 2; i < l; i++ {
		p = q
		q = r
		r = max(p+nums[i], q)
	}

	return r
}

// 对于213题，因为首尾相连，计算区间会存在两种情况
// (l为切片长度)：[0, l - 2], [1, l - 1]

func rob2(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	if l == 2 {
		return max(nums[0], nums[1])
	}

	p1, q1, r1 := 0, nums[0], nums[1]
	for i := 2; i <= l-2; i++ {
		p1 = q1
		q1 = r1
		r1 = max(p1+nums[i], q1)
	}

	p2, q2, r2 := 0, nums[1], nums[2]
	for i := 3; i <= l-1; i++ {
		p2 = q2
		q2 = r2
		r2 = max(p2+nums[i], q2)
	}

	return max(r1, r2)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
