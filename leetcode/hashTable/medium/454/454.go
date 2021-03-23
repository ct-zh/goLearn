package _454

// 四个数组做查找，每个数组500个元素,如果暴力破解，需要500^4次
// 将两个数组的合加入查找表中，就只需要500^2次查找，而空间复杂度会变成500^2
// 时间复杂度 O(n^2)
// 空间复杂度 O(n^2)
// 执行用时： 64 ms, 在所有 Go 提交中击败了 72.79% 的用户
// 内存消耗： 22.7 MB, 在所有 Go 提交中击败了 31.69% 的用户
func fourSumCount(A []int, B []int, C []int, D []int) int {
	record := map[int]int{}
	for i := 0; i < len(C); i++ {
		for j := 0; j < len(D); j++ {
			record[C[i]+D[j]]++
		}
	}

	res := 0
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			if times, ok := record[0-A[i]-B[j]]; ok {
				res += times
			}
		}
	}
	return res
}
