package _1

// 查找表。将v前面的元素放入查找表中，对于某个元素判断target - v是否存在
// (如果全部元素放进查找表，会有相同元素的问题)
// 时间复杂度与空间复杂度都是 O(n)
func twoSum(nums []int, target int) []int {
	record := map[int]int{}

	for i := 0; i < len(nums); i++ {

		complement := target - nums[i]

		// 注意
		if _, ok := record[complement]; ok {
			var res []int
			if i > record[complement] {
				res = []int{record[complement], i}
			} else {
				res = []int{i, record[complement]}
			}
			return res
		}

		record[nums[i]] = i
	}

	panic("no result")
}
