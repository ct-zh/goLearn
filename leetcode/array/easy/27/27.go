package _7

func removeElementDemo1(nums []int, val int) int {
	// index指向第一个val的位置
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			if nums[index] == val {
				nums[index], nums[i] = nums[i], nums[index]
			}
		}

		// 因为如果nums[index] == val 符合上面的条件，两个数值会发生交换
		// 交换后nums[index] 又 != val 了
		if nums[index] != val {
			index++
		}
	}
	return index
}

// 解法二：照抄的标准答案
// 对比我的解法简洁很多，但是忽略了i=j时仍然做无效交换导致性能浪费的问题
// 各有千秋，不要妄自菲薄
func removeElementDemo2(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

func removeElement(nums []int, val int) int {
	l := len(nums)

	index := 0
	for i := 0; i < l; i++ {
		if nums[i] != val && nums[index] == val {
			nums[i], nums[index] = nums[index], nums[i]
		}
		if nums[index] != val {
			index++
		}
	}

	return index
}
