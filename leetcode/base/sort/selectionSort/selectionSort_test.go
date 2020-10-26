package selectionSort

import (
	"fmt"
	Helper "github.com/LannisterAlwaysPaysHisDebts/goLearn/leetcode/base/helper"
	"testing"
	"time"
)

// 测试内容
// 1. 随机生成测试用例
// 2. 排序成功的判断(check 函数)
// 3. 平均运行时间：(时间复杂度测试)
func TestSelectionSort(t *testing.T) {
	tests := []struct {
		arr map[int]int
	}{
		{Helper.GenerateRandomArray(100, 0, 999)},
		{Helper.GenerateRandomArray(10000, 0, 99999)},
	}
	for key, tt := range tests {
		startTime := time.Now()

		count := len(tt.arr)
		s := SelectionSort{
			Arr:   tt.arr,
			Count: count,
		}
		s.Do()

		// check
		for i := 1; i < count; i++ {
			if s.Arr[i-1] > s.Arr[i] {
				t.Errorf("[%d] 排序算法有问题：\n 原始数据：   %+v \n  排序后的数据： %+v\n", key, tt.arr, s.Arr)
			}
		}

		// 运行时间
		duration := time.Now().Sub(startTime)
		fmt.Printf("[%d] 消耗时间： %.8fs\n", key, duration.Seconds())
	}
}
