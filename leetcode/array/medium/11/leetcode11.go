package leetcode11

// https://leetcode-cn.com/problems/container-with-most-water/

// 对撞指针

func maxArea(height []int) int {
	i1, i2 := 0, len(height)-1
	max := 0
	for i1 < i2 {
		diff := i2 - i1
		var l int
		if height[i1] > height[i2] {
			l = height[i2]
			i2--
		} else {
			l = height[i1]
			i1++
		}

		res := l * diff
		if max < res {
			max = res
		}
	}
	return max
}
