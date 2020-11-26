package _350

// 思路1 hash 表
// 时间复杂度O(m+n)
// 空间复杂度O(min(m, n))
func intersect(nums1 []int, nums2 []int) []int {
	r := map[int]int{} // 用r来保存nums1里元素的出现次数
	for i := 0; i < len(nums1); i++ {
		r[nums1[i]]++
	}

	var result []int
	for i := 0; i < len(nums2); i++ {
		if r[nums2[i]] > 0 { // 如果该原素在交集里面
			result = append(result, nums2[i])
			r[nums2[i]]--
		}
	}

	return result
}
