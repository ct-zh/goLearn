package selectionSort

import (
	"math/rand"
	"testing"
	"time"
)

// 生成指定范围的整数数组
func generateRandomArray(n int, rangeL int, rangeR int) map[int]interface{} {
	if rangeL > rangeR {
		return nil
	}

	// rand
	arr := map[int]interface{}{}

	rand2 := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < n; i ++ {
		// 指定范围随机数生成的标准写法
		arr[i] = rand2.Int() % (rangeR - rangeL + 1) + rangeL
	}

	return arr
}

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
