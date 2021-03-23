package leetbookSection1

// 见283题

// leetbook题解：
// 只要把数组中所有的非零元素，按顺序给数组的前段元素位赋值，剩下的全部直接赋值 0。

// 我的题解: 双指针， 第一个指针指向第一个非0的元素，第二个指针遍历
func moveZeroes(nums []int) {
	l := len(nums)
	p := 0
	for i := 0; i < l; i++ {
		if nums[i] != 0 {
			for {
				if p >= i {
					break
				} else if nums[p] != 0 {
					p++
					continue
				} else {
					nums[p], nums[i] = nums[i], nums[p]
					break
				}
			}
		}
	}
}
