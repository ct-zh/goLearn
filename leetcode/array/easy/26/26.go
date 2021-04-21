package _6

func removeDuplicatesDemo(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	index := 0
	for i := 1; i < len(nums); i++ {
		if nums[index] != nums[i] {
			if index+1 != i {
				nums[index+1], nums[i] = nums[i], nums[index+1]
			}
			index++
		}
	}

	return index + 1
}

func removeDuplicates(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return l
	}

	index := 0
	for i := 1; i < l; i++ {
		if index != i {
			if nums[i] != nums[index] {
				if index+1 != i {
					nums[i], nums[index+1] = nums[index+1], nums[i]
				}
				index++
			}
		}
	}

	return index + 1
}
