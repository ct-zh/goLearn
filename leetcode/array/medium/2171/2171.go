package l2171

import "sort"

// 请返回你需要拿出魔法豆的 最少数目。

func minimumRemoval(beans []int) int64 {
	// 自己写遍历
	// 先从小到大排序s 将所有值往s[0]看齐 得到结果a1
	// 再循环将 i 置 0, 其他值往s[1]看齐，得到结果a2
	// 依次操作，然后返回a 的最小值

	sort.Ints(beans)
	var result int64 = 0
	for i := 0; i < len(beans); i++ {
		result += int64(beans[i])
	}
	for i := 0; i < len(beans); i++ {
		var itemResult int64 = 0 // 本次操作的豆子数 初始值
		for m := 0; m < len(beans); m++ {
			if m < i { // 小于i的全拿掉
				itemResult += int64(beans[m])
			} else if m == i { // 跳过当前值
			} else {
				itemResult += int64(beans[m]) - int64(beans[i])
			}
		}
		result = minFn(result, itemResult)
	}

	return result
}

func minimumRemovalV2(beans []int) int64 {
	// 找到所有袋子中魔法豆的总数
	total := int64(0)
	for _, bean := range beans {
		total += int64(bean)
	}

	// 每个袋子中魔法豆的最小数目
	min := total / int64(len(beans))

	// 遍历所有袋子
	for _, bean := range beans {
		// 计算需要拿出魔法豆的数目
		result := total - int64(len(beans))*min + int64(bean) - min

		// 更新最小数目
		min = minFn(min, result)
	}

	return min
}

func minFn(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
