package _67

// https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/

// 思路0：嵌套循环暴力破解，时间复杂度O(n^2)

// 思路1： 二分查找
// 时间复杂度O(nlogn)
func twoSum(numbers []int, target int) []int {
	return []int{}
}

// 思路2：对撞指针
// 时间复杂度O(n)
func twoSum2(numbers []int, target int) []int {
	i1 := 0
	i2 := len(numbers) - 1

	for {
		if i1 >= i2 {
			// 不存在答案
			return []int{}
		}
		if numbers[i1]+numbers[i2] == target {
			break
		} else if numbers[i1]+numbers[i2] > target {
			i2--
		} else {
			i1++
		}
	}

	return []int{i1 + 1, i2 + 1}
}
