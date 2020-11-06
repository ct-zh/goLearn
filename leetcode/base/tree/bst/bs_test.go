package bst

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 比较非递归和递归写法的二分查找的效率
// 非递归算法在性能上有微弱优势
func TestBinarySearch(t *testing.T) {
	var arr1 []int
	for i := 0; i < 100000; i++ {
		arr1 = append(arr1, i)
	}
	randTarget := rand.New(
		rand.NewSource(
			time.Now().Unix())).
		Int() % 100000

	tests := []struct {
		arr    []int
		target int
	}{
		{arr1, randTarget},
	}
	for key, tt := range tests {
		startTime := time.Now()

		result1 := BinarySearch(tt.arr, len(tt.arr), tt.target)
		fmt.Printf("非递归搜索结果：%d\n", result1)

		end1 := time.Now()
		fmt.Printf("[%d] 消耗时间： %.8fs\n", key, end1.Sub(startTime).Seconds())

		result2 := BinarySearchRecursive(tt.arr, len(tt.arr), tt.target)
		fmt.Printf("递归搜索结果：%d\n", result2)

		end2 := time.Now()
		fmt.Printf("[%d] 消耗时间： %.8fs\n", key, end2.Sub(end1).Seconds())
	}
}
