package _220

import "math"

// 思路 滑动窗口, 对比219 改成二叉树实现
// 这里应该要一棵平衡的二叉搜索树
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	record := createTree(k)

	for i := 0; i < len(nums); i++ {
		if record.findOver(nums[i])-nums[i] <= t {
			return true
		}
		if nums[i]-record.findLess(nums[i]) <= t {
			return true
		}

		record.add(nums[i])
	}

	return false
}

type tree struct {
	data     []int
	len      int
	capacity int
}

func createTree(capacity int) *tree {
	return &tree{
		data:     make([]int, math.Pow(float64(2), float64(capacity-1))),
		len:      0,
		capacity: capacity,
	}
}

// 获取超过v的最小值
func (t *tree) findOver(v int) int {

}

// 获取未达到v的最大值
func (t *tree) findLess(v int) int {

}

func (t *tree) add(v int) {
	if t.len == t.capacity+1 {
		t.data = t.data[1:]
	}

	// 从1开始
	if t.len == 0 {
		t.data[1] = v
	}


	t.len++
}
