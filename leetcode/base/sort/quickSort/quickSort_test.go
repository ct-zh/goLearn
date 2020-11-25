package quickSort

import (
	"fmt"
	Helper "github.com/LannisterAlwaysPaysHisDebts/goLearn/leetcode/base/helper"
	"testing"
	"time"
)

func TestQuickSort_Do(t *testing.T) {
	tests := []struct {
		arr map[int]int
	}{
		{Helper.GenerateRandomArray(10, 0, 999)},
		{Helper.GenerateRandomArray(100, 0, 999)},
		{Helper.GenerateRandomArray(10000, 0, 99999)},
	}
	for key, tt := range tests {
		startTime := time.Now()

		count := len(tt.arr)
		s := QuickSort{
			Arr: tt.arr,
			N:   count,
		}
		s.Do()

		// check
		err := false
		for i := 1; i < count; i++ {
			if s.Arr[i-1] > s.Arr[i] {
				err = true
			}
		}
		if err {
			t.Errorf("[%d] 排序算法有问题：\n 原始数据：   %+v \n  排序后的数据： %+v\n", key, tt.arr, s.Arr)
		}

		end := time.Now().Sub(startTime)
		fmt.Printf("[%d] 消耗时间： %.8fs\n", key, end.Seconds())
	}
}

func TestQuickSort_Do2(t *testing.T) {
	tests := []struct {
		arr map[int]int
	}{
		{Helper.GenerateRandomArray(100, 0, 999)},
		{Helper.GenerateRandomArray(10000, 0, 99999)},
	}
	for key, tt := range tests {
		startTime := time.Now()

		count := len(tt.arr)
		s := QuickSort{
			Arr: tt.arr,
			N:   count,
		}
		s.Do2()

		// check
		err := false
		for i := 1; i < count; i++ {
			if s.Arr[i-1] > s.Arr[i] {
				err = true
			}
		}
		if err {
			t.Errorf("[%d] 排序算法有问题：\n 原始数据：   %+v \n  排序后的数据： %+v\n", key, tt.arr, s.Arr)
		}

		end := time.Now().Sub(startTime)
		fmt.Printf("[%d] 消耗时间： %.8fs\n", key, end.Seconds())
	}
}

func TestQuickSort_Do3(t *testing.T) {
	tests := []struct {
		arr map[int]int
	}{
		{Helper.GenerateRandomArray(100, 0, 999)},
		{Helper.GenerateRandomArray(10000, 0, 99999)},
	}
	for key, tt := range tests {
		startTime := time.Now()

		count := len(tt.arr)
		s := QuickSort{
			Arr: tt.arr,
			N:   count,
		}
		s.Do3()

		// check
		err := false
		for i := 1; i < count; i++ {
			if s.Arr[i-1] > s.Arr[i] {
				err = true
			}
		}
		if err {
			t.Errorf("[%d] 排序算法有问题：\n 原始数据：   %+v \n  排序后的数据： %+v\n", key, tt.arr, s.Arr)
		}

		end := time.Now().Sub(startTime)
		fmt.Printf("[%d] 消耗时间： %.8fs\n", key, end.Seconds())
	}
}
