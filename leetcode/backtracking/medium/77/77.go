package _77

import "fmt"

//

func combine(n int, k int) (result [][]int) {
	if n <= 0 || k <= 0 || k > n {
		return
	}

	generateCombinations(n, k, 1, &[]int{}, &result)
	return
}

// 求解C(n, k), 当前已经找到的组合存储在c中；需要从start开始搜索新元素
func generateCombinations(
	n int,           // [1...n]
	k int,           // k个数字
	start int,       // 从n中开始的项
	c *[]int,        // 当前组合的内容
	result *[][]int, // 结果
) {
	if len(*c) == k {
		*result = append(*result, *c)
		return
	}

	// @tag: 回溯法的剪枝
	// 还有 k - len(c)个空位
	// 所以[i...n]中至少有k-len(c)个元素
	// i最多为 n - (k-le n(c)) + 1
	for i := start; i <= n-(k-len(*c))+1; i++ {
		*c = append(*c, i)
		generateCombinations(n, k, i+1, c, result)

		// 回溯
		*c = (*c)[0 : len(*c)-1]
	}

}

func combine2(n int, k int) (result [][]int) {
	if n <= 0 || k <= 0 || k > n {
		return
	}

	generateCombinations2(n, k, 1, []int{}, &result)
	return
}

func generateCombinations2(
	n, k, start int,
	c []int,
	result *[][]int,
) {
	fmt.Printf("n: %d k:%d start: %d c: %+v result: %+v \n", n, k, start, c, result)
	if len(c) == k {
		*result = append(*result, c)
		return
	}

	for i := start; i <= n; i++ {
		c = append(c, i)
		generateCombinations2(n, k, start+1, c, result)
		c = c[0 : len(c)-1]
	}
}
