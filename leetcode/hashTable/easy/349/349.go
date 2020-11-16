package _349

// 思路1 使用set结构
// 时间复杂度 O(n+m)
// 空间复杂度 最大为 O(n+m)
func intersection(nums1 []int, nums2 []int) []int {
	tmp := make(map[int]int)
	for _, v := range nums1 {
		tmp[v] = 1
	}

	s := make(map[int]int)
	for _, v := range nums2 {
		if tmp[v] == 1 {
			s[v] = 1
		}
	}

	var res []int
	for v := range s {
		res = append(res, v)
	}

	return res
}

// 优化为空结构体： go语言中空结构体占用内存极小
func intersection2(nums1 []int, nums2 []int) []int {
	tmp := map[int]struct{}{}
	for _, v := range nums1 {
		tmp[v] = struct{}{}
	}

	s := map[int]struct{}{}
	for _, v := range nums2 {
		if _, has := tmp[v]; has {
			s[v] = struct{}{}
		}
	}

	var res []int
	for v := range s {
		res = append(res, v)
	}

	return res
}

// 思路3 排序 + 双指针
// 时间复杂度为 O(mlogm+nlogn) 两个排序
// 空间复杂度为 O(logm + logn) 仍然是用于两个排序
func intersection3(nums1 []int, nums2 []int) (result []int) {
	quickSort(nums1, 0, len(nums1)-1)
	quickSort(nums2, 0, len(nums2)-1)

	for i1, i2 := 0, 0; i1 < len(nums1) && i2 < len(nums2); {
		n1, n2 := nums1[i1], nums2[i2]
		if n1 == n2 {

			// 经过排序后的数组是递增的，如果想去重，只要保证当前值比result里最后一个值还大就行了
			if result == nil || n1 > result[len(result)-1] {
				result = append(result, n1)
			}
			i1++
			i2++
		} else if n1 < n2 {
			i1++
		} else { // n1 > n2
			i2++
		}
	}

	return
}

// 快排
func quickSort(arr []int, l int, r int) {

}
