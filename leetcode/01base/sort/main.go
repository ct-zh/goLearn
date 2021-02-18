package main

import (
	"fmt"
	Helper "github.com/LannisterAlwaysPaysHisDebts/goLearn/leetcode/base/helper"
	"github.com/LannisterAlwaysPaysHisDebts/goLearn/leetcode/base/sort/insertionSort"
	"github.com/LannisterAlwaysPaysHisDebts/goLearn/leetcode/base/sort/mergeSort"
	"github.com/LannisterAlwaysPaysHisDebts/goLearn/leetcode/base/sort/quickSort"
	"github.com/LannisterAlwaysPaysHisDebts/goLearn/leetcode/base/sort/selectionSort"
	"time"
)

// 不同排序算法的性能比较, 平均时间复杂度
// 冒泡排序(bubble sort)    O(n^2)
// 选择排序(selection sort) O(n^2)
// 插入排序(insertion sort) O(n^2)
// 希尔排序(shell's sort)   O(n^1.5) 时间复杂度下界为 O(n*log2n)
// 快速排序(quick sort)     O(n*logN)
// 归并排序(merge sort)	  O(n*logN)
// 堆排序(heap sort)		  O(n*logN)
// 基数排序(radix sort)	  O(n*log(r)m) r为基数，m为堆数

// 比较内容：
// 1.随机列表排序速度比较
// 2.近乎有序的列表排序速度比较
// 3.元素全部相同的列表排序速度比较
// 做成协程测试，每个排序一个协程
func main() {
	// 在这里填充测试用例
	tests := []struct {
		arr  []int
		text string
	}{
		{
			Helper.GenerateRandArr(100, 0, 999),
			"随机100个数[0-999]",
		},
		{
			Helper.GenerateRandArr(10000, 0, 99999),
			"随机10000个数[0-99999]",
		},
		{
			Helper.GenerateNearlyArray(100, 0),
			"0-99按顺序排列",
		},
		{
			Helper.GenerateNearlyArray(100, 10),
			"0-99随机交换10次",
		},
		{
			Helper.GenerateRandomArray(100, 0, 0),
			"100个0",
		},
	}

	// 在这里填充排序算法
	sortFunctions := map[string]func([]int){
		"选择排序":     selectionSort.SelectionSort,
		"插入排序":     insertionSort.InsertionSort,
		"自顶向下归并排序": mergeSort.MergeSort,
		"自底向上归并排序": mergeSort.MergeSort1,
		"随机快排":     quickSort.QuickSort,
		"双路快排":     quickSort.QuickSort2Ways,
		"三路快排":     quickSort.QuickSort3Ways,
	}

	sortCount := 0
	ch := make(chan string)

	for _, tt := range tests {
		for name, sFunc := range sortFunctions {
			go func(name string, data []int, text string) {
				arr := make([]int, len(data))
				copy(arr, data)
				startTime := time.Now()

				sFunc(arr)

				end := time.Now().Sub(startTime)
				outStr := fmt.Sprintf("%s 消耗时间： %.8fs",
					name, end.Seconds())
				if Helper.CheckSort(arr) {
					outStr += " 例子[" + text + "]检测通过 Pass!"
				} else {
					outStr += " 例子[" + text + "]检测未通过  Wrong!"
				}
				ch <- outStr
			}(name, tt.arr, tt.text)
		}

		sortCount += len(sortFunctions)
	}

	// 异步获取排序结果
	fmt.Printf("总共测试数量：%d条：\n", sortCount)
	for i := 0; i < sortCount; i++ {
		fmt.Printf("测试{%d} %s \n", i+1, <-ch)
	}

	fmt.Println("所有例子测试完毕")
}
