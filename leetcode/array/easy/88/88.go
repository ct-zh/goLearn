package _88

// https://leetcode-cn.com/problems/merge-sorted-array/

// 思路1 暴力破解 O(n+(m+n)log(m+n))
// 用时12ms 击败100%
// 内存消耗2.3MB 击败68% =>  最少内存消耗:2.28MB
// 这里如果调用随机函数来做随机快排，会大幅度加大内存消耗到6.3mb
func merge(nums1 []int, m int, nums2 []int, n int) {
	// 合并两个数组
	for i := 0; i < n; i++ {
		nums1[i+m] = nums2[i]
	}

	// 再对nums1快排
	// 闭区间[0, m+n-1]
	sort(nums1, 0, m+n-1)
}

func sort(arr []int, l int, r int) {
	if l >= r {
		return
	}
	lt, gt := partition(arr, l, r)
	sort(arr, l, lt)
	sort(arr, gt, r)
}

func partition(arr []int, l int, r int) (int, int) {
	// 锚定点flag
	flag := arr[l]

	// 定义边界lt gt，初始化的值保证在循环列表外
	lt := l
	gt := r + 1

	// 定义游标p遍历数组，并维护定义：
	// [l+1...lt] < flag; [lt+1...gt-1] = flag; [gt,r] >flag
	p := l + 1
	for {
		if p >= gt {
			break
		}

		if arr[p] == flag {
			p++
		} else if arr[p] < flag {
			arr[lt+1], arr[p] = arr[p], arr[lt+1]
			p++
			lt++
		} else {
			arr[gt-1], arr[p] = arr[p], arr[gt-1]
			gt--
		}
	}

	arr[lt], arr[l] = arr[l], arr[lt]

	return lt - 1, gt
}

// 思路2 经典双指针
// 注意这里的思路应该是从后往前循环
// 时间复杂度O(m+n)
// 空间复杂度O(1)
func merge2(nums1 []int, m int, nums2 []int, n int) {
	i1 := m - 1
	i2 := n - 1
	last := m + n - 1

	for {
		if i1 < 0 && i2 < 0 {
			break
		}
		if i2 < 0 || (i1 >= 0 && nums1[i1] > nums2[i2]) {
			nums1[last] = nums1[i1]
			i1--
		} else { // i1 < 0 || i2
			nums1[last] = nums2[i2]
			i2--
		}

		last--
	}
}
