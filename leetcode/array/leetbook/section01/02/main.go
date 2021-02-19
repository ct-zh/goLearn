package _2

// 见27题

// 我的思路：双指针  跟上题的题解一样，遍历nums，不等于val就往前塞
func removeElement(nums []int, val int) int {
	l := len(nums)
	if l <= 0 || l > 100 {
		return 0
	}
	if val < 0 || val > 100 {
		return 0
	}

	// flag从0开始，可能指向第一个等于val元素
	f := 0
	for i := 0; i < l; i++ {
		if nums[i] < 0 || nums[i] > 100 {
			return 0
		}

		// 如果i不为val，则尝试往
		if nums[i] != val {
			if i != f {
				nums[i], nums[f] = nums[f], nums[i]
			}
		}

		// f跟着i一起++ 直到遇到第一个等于val的元素;
		// 如果发生交换，f对应的元素不再等于val，也需要++
		if nums[f] != val {
			f++
		}
	}

	return f
}
