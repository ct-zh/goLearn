package selectionSort

import (
	"testing"
)

type Student struct {
	name  string
	score int
}

func (s *Student) less(other Student) bool {
	return s.score < other.score
}

// 测试内容
// 1. 随机生成测试用例
// 2. 排序成功的判断(check 函数)
// 3. 平均运行时间：(时间复杂度测试)
func TestSelectionSort(t *testing.T) {

}
