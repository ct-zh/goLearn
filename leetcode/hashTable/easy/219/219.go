package _219

// 滑动窗口  这里的滑动窗口不需要维护左右边界，而是维护长度小于等于k
// 循环数组，对于当前元素i，查找滑动窗口内是否存在相同的值
// 时间复杂度 O(n)
// 空间复杂度 O(k)
func containsNearbyDuplicate(nums []int, k int) bool {

	// 记录值的集合set （滑动窗口）
	record := map[int]struct{}{}
	for i := 0; i < len(nums); i++ {

		// 如果在区间内存在相同的值，直接return
		if _, ok := record[nums[i]]; ok {
			return true
		}

		// 将i的值加入set中，同时保证set中值的个数小于等于k
		record[nums[i]] = struct{}{}
		if len(record) == k+1 {
			delete(record, nums[i-k]) // 删除掉滑动窗口的第一个值
		}
	}
	return false
}
