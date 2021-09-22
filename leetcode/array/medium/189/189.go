package _89

// 法一 暴力解法 超出时间限制 O(n^k)
func rotate(nums []int, k int) {
	l := len(nums)
	for i := 0; i < k; i++ {
		tmp := nums[l-1]
		for m := l - 1; m > 0; m-- {
			nums[m] = nums[m-1]
		}
		nums[0] = tmp
	}
}

// 法二 额外开一个空间进行移动
// 时间：O(n), 空间：O(n)
func rotate2(nums []int, k int) {
	l := len(nums)
	nums2 := make([]int, l)
	for i, num := range nums {
		nums2[(i+k)%l] = num
	}
	copy(nums, nums2)
}

// 翻转法：
// 1234567 翻转 => 7654321
// k=3 => 765  4321
// 翻转回来 => 5671234
// 时间O(n),空间O(1)
func rotate3(nums []int, k int) {
	l := len(nums)
	k %= l
	for i := 0; i < l/2; i++ {
		nums[i], nums[l-1-i] = nums[l-1-i], nums[i]
	}
	for i := 0; i < k/2; i++ {
		nums[i], nums[k-1-i] = nums[k-1-i], nums[i]
	}
	for i := 0; i < (l-k)/2; i++ {
		nums[i+k], nums[l-1-i] = nums[l-1-i], nums[i+k]
	}
}

// 替换
// 例1. 1234567 (i+k)%l
// => 1到4，4到7，7到3 (6+3)%7,3到6，6到2(5+3)%7，2到5，5到1;
// => 5671234
func rotate4(nums []int, k int) {
	l := len(nums)
	k %= l
	if l%2 == 0 && k%2 == 0 {
		for i := 0; i < l/2; i++ {
			m := (i + k) % l
			nums[i], nums[m] = nums[m], nums[i]
		}
	} else {
		tmp := 0
		tmpVal := nums[tmp]
		for i := 0; i < l; i++ {
			m := (tmp + k) % l
			nums[m], tmp, tmpVal = tmpVal, m, nums[m]
		}
	}
}
