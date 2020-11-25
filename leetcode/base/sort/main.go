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
		arr  map[int]int
		text string
	}{
		//{
		//	Helper.GenerateRandomArray(100, 0, 999),
		//	"0-999 随机100个数",
		//},
		{
			Helper.GenerateRandomArray(10000, 0, 99999),
			"0-99999 随机10000个数",
		},
		//{
		//	Helper.GenerateNearlyArray(100, 0),
		//	"0-99按顺序排列",
		//},
		//{
		//	Helper.GenerateNearlyArray(100, 10),
		//	"0-99随机交换10次",
		//},
		//{
		//	Helper.GenerateRandomArray(100, 0, 0),
		//	"100个0",
		//},
	}

	// 在这里填充排序算法
	sortFunc := []func(map[int]int, string, chan string){
		selectionS,
		insertionS,
		mergeS1,
		mergeS2,
		quickS1,
		quickS2,
		quickS3,
	}

	sortCount := 0
	ch := make(chan string)

	for _, tt := range tests {
		for _, sFunc := range sortFunc {
			// cp map data 因为map是引用类型的
			go sFunc(Helper.CopyData(tt.arr), tt.text, ch)
		}

		sortCount += len(sortFunc)
	}

	// 异步获取排序结果
	fmt.Printf("总共测试数量：%d条：\n", sortCount)
	for i := 0; i < sortCount; i++ {
		fmt.Printf("测试{%d} %s \n", i+1, <-ch)
	}

	fmt.Println("所有例子测试完毕")
}

// 测试排序算法是否通过,默认排序是从小到大
func check(arr map[int]int, count int) bool {
	for i := 1; i < count; i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}

// 选择排序
func selectionS(arr map[int]int, text string, ch chan string) {
	startTime := time.Now()

	s := selectionSort.SelectionSort{
		Arr:   arr,
		Count: len(arr),
	}
	s.Do()

	end := time.Now().Sub(startTime)
	outStr := fmt.Sprintf("选择排序 消耗时间： %.8fs", end.Seconds())

	if check(s.Arr, s.Count) {
		outStr += " 例子[" + text + "]检测通过 Pass!"
	} else {
		outStr += " 例子[" + text + "]检测未通过  Wrong!"
	}

	ch <- outStr
}

// 插入排序
func insertionS(arr map[int]int, text string, ch chan string) {

	startTime := time.Now()

	s := insertionSort.InsertionSort{
		Arr:   arr,
		Count: len(arr),
	}
	s.DoBetter()

	end := time.Now().Sub(startTime)
	outStr := fmt.Sprintf("插入排序 消耗时间： %.8fs", end.Seconds())

	if check(s.Arr, s.Count) {
		outStr += " 例子[" + text + "]检测通过 Pass!"
	} else {
		outStr += " 例子[" + text + "]检测未通过  Wrong!"
	}

	ch <- outStr
}

// 自顶向下归并排序
func mergeS1(arr map[int]int, text string, ch chan string) {
	startTime := time.Now()

	s := mergeSort.MergeSort{
		Arr: arr,
		N:   len(arr),
	}
	s.Do()

	end := time.Now().Sub(startTime)
	outStr := fmt.Sprintf("自顶向下归并排序 消耗时间： %.8fs", end.Seconds())

	if check(s.Arr, s.N) {
		outStr += " 例子[" + text + "]检测通过 Pass!"
	} else {
		outStr += " 例子[" + text + "]检测未通过  Wrong!"
	}

	ch <- outStr
}

// 自底向上归并排序
func mergeS2(arr map[int]int, text string, ch chan string) {
	startTime := time.Now()

	s := mergeSort.MergeSort{
		Arr: arr,
		N:   len(arr),
	}
	s.Do()

	end := time.Now().Sub(startTime)
	outStr := fmt.Sprintf("自底向上归并排序 消耗时间： %.8fs", end.Seconds())

	if check(s.Arr, s.N) {
		outStr += " 例子[" + text + "]检测通过 Pass!"
	} else {
		outStr += " 例子[" + text + "]检测未通过  Wrong!"
	}

	ch <- outStr
}

// 随机快排
func quickS1(arr map[int]int, text string, ch chan string) {
	startTime := time.Now()

	s := quickSort.QuickSort{
		Arr: arr,
		N:   len(arr),
	}
	s.Do()

	end := time.Now().Sub(startTime)
	outStr := fmt.Sprintf("随机快排 消耗时间： %.8fs", end.Seconds())

	if check(s.Arr, s.N) {
		outStr += " 例子[" + text + "]检测通过 Pass!"
	} else {
		outStr += " 例子[" + text + "]检测未通过  Wrong!"
	}

	ch <- outStr
}

// 双路快排
func quickS2(arr map[int]int, text string, ch chan string) {
	startTime := time.Now()

	s := quickSort.QuickSort{
		Arr: arr,
		N:   len(arr),
	}
	s.Do2()

	end := time.Now().Sub(startTime)
	outStr := fmt.Sprintf("双路快排 消耗时间： %.8fs", end.Seconds())

	if check(s.Arr, s.N) {
		outStr += " 例子[" + text + "]检测通过 Pass!"
	} else {
		outStr += " 例子[" + text + "]检测未通过  Wrong!"
	}

	ch <- outStr
}

// 三路快排
func quickS3(arr map[int]int, text string, ch chan string) {
	startTime := time.Now()

	s := quickSort.QuickSort{
		Arr: arr,
		N:   len(arr),
	}
	s.Do3()

	end := time.Now().Sub(startTime)
	outStr := fmt.Sprintf("三路快排 消耗时间： %.8fs", end.Seconds())

	if check(s.Arr, s.N) {
		outStr += " 例子[" + text + "]检测通过 Pass!"
	} else {
		outStr += " 例子[" + text + "]检测未通过  Wrong!"
	}

	ch <- outStr
}
