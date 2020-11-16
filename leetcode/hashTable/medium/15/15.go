package _15

// https://leetcode-cn.com/problems/3sum/

// 思路： 排序 + 双指针
// 先将数组排序，再进行循环，这样就将问题变为了在子列表中
// 是否存在nums[a] + nums[b] = -nums[i] 的问题，这里可以使用双指针来解决
// 时间复杂度为 排序O(nlogn) 外层循环O(n),内层循环O(n),即O(n^2)
// 空间复杂度为 排序空间一般为O(logn),但是额外写了一个顺序数组，所以为O(n)
// 执行用时： 44 ms, 在所有 Go 提交中击败了 35.85% 的用户
// 内存消耗： 7.3 MB, 在所有 Go 提交中击败了 28.55% 的用户
func threeSum(nums []int) [][]int {
	l := len(nums)

	// 排序
	quickSort(nums, 0, l-1)

	var result [][]int
	for i := 0; i < l; i++ {

		// 跳过重复的值
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		lt := i + 1
		gt := l - 1
		for {
			// 跳过重复的值
			if lt+1 < l && lt > i+1 && nums[lt] == nums[lt-1] {
				lt++
				continue
			}
			if gt-1 > lt && gt < l-1 && nums[gt] == nums[gt+1] {
				gt--
				continue
			}

			if lt >= gt {
				break
			}

			target := -nums[i]
			if nums[lt]+nums[gt] == target {
				result = append(result, []int{nums[i], nums[lt], nums[gt]})
				lt++
				gt--
			} else if nums[lt]+nums[gt] > target { // 太大了
				gt--
			} else { // 太小了
				lt++
			}
		}
	}

	return result
}

// 快排
func quickSort(arr []int, l int, r int) {
	if l >= r {
		return
	}

	lt, gt := partition(arr, l, r)
	quickSort(arr, l, lt)
	quickSort(arr, gt, r)
}

func partition(arr []int, l int, r int) (lt int, gt int) {
	flag := arr[l]

	lt = l
	gt = r + 1
	index := l + 1
	for {
		if index >= gt {
			break
		}

		if flag == arr[index] {
			index++
		} else if flag > arr[index] {
			arr[index], arr[lt+1] = arr[lt+1], arr[index]
			lt++
			index++
		} else {
			arr[index], arr[gt-1] = arr[gt-1], arr[index]
			gt--
		}
	}

	arr[lt], arr[l] = arr[l], arr[lt]

	return lt - 1, gt
}

// 时间优化： i大于0，大于i的数必大于0，后面的数相加必不可能等于0
// 执行用时： 28 ms , 在所有 Go 提交中击败了 99.04% 的用户
func threeSum2(nums []int) [][]int {
	l := len(nums)
	var res [][]int
	quickSort(nums, 0, l-1)
	for i := 0; i < l-2; i++ {
		// 优化1. i大于0，大于i的数必大于0，后面的数相加必不可能等于0
		if nums[i] > 0 {
			break
		}
		// 跳过重复的
		if i >= 1 && nums[i] == nums[i-1] {
			continue
		}
		lt, gt := i+1, l-1
		for lt < gt {
			n1, n2, n3 := nums[i], nums[lt], nums[gt]
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				for lt < gt && nums[lt] == n2 {
					lt++
				}
				for lt < gt && nums[gt] == n3 {
					gt--
				}
			} else if n1+n2+n3 < 0 {
				lt++
			} else {
				gt--
			}
		}
	}
	return res
}
