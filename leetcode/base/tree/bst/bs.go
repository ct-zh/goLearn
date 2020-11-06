package bst

// 二分查找法,在有序数组arr中,查找target
// 如果找到target,返回相应的索引index
// 如果没有找到target,返回-1

// arr: 有序数组; n: 数组长度; target:  要查找的值
func BinarySearch(arr []int, n int, target int) int {

	//  在arr[l...r]之中查找target
	l := 0
	r := n - 1
	for {
		if l > r { // 越界
			break
		}
		mid := (r-l)/2 + l // 不这样求(l+r)/2中值的原因: 防止极端情况下  r + l 越界

		if arr[mid] == target { // 刚好是中值的情况
			return mid
		}

		if arr[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

// 用递归的方式写二分查找法
func BinarySearchRecursive(arr []int, n int, target int) int {
	return _binarySearchRecursive(arr, 0, n-1, target)
}

func _binarySearchRecursive(arr []int, l int, r int, target int) int {
	if l > r { // 越界
		return -1
	}
	mid := (r-l)/2 + l // 不这样求(l+r)/2中值的原因: 防止极端情况下  r + l 越界
	if target == arr[mid] {
		return mid
	} else if arr[mid] > target {
		return _binarySearchRecursive(arr, l, mid-1, target)
	} else {
		return _binarySearchRecursive(arr, mid+1, r, target)
	}
}
