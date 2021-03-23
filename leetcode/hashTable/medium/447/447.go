package _447

// 固定点i 问题就转换为基于点i遍历整个矩阵，获取到各个点与i距离的集合
// 再根据这个集合计算可能的组合方式
//
// 问题1 获取各个点的距离，计算方法是 根号((x1-x0)^2+(y1-y0)^2)
// 这个根号会导致出现浮点数。解决的办法是不开这个根号，那么集合的定义就变成了各个点到i距离的平方的集合
//
// 时间复杂度 O(n^2)
// 空间复杂度 O(n)
func numberOfBoomerangs(points [][]int) (res int) {
	// 以点i作为中心
	for i := 0; i < len(points); i++ {

		// key：点i与其他点的距离；
		// value： 这个距离出现的次数；
		record := map[int]int{}
		for j := 0; j < len(points); j++ {
			if j != i {
				record[dis(points[i], points[j])]++
			}
		}

		// 遍历集合record，找到出现次数在两次以上的距离
		// 即是存在至少两点到点i的距离相同
		// 计算这些点的排列组合：
		// 如果有3个点，就是 3*2=6种可能
		// 如果有4个点，就是 4*3=12种可能
		// 以此类推
		for _, v := range record {
			if v >= 2 {
				res += v * (v - 1)
			}
		}
	}
	return res
}

// 距离的平方
func dis(pa []int, pb []int) int {
	return (pa[0]-pb[0])*(pa[0]-pb[0]) + (pa[1]-pb[1])*(pa[1]-pb[1])
}
