package insertionSort

import (
	"fmt"
	Helper "github.com/LannisterAlwaysPaysHisDebts/goLearn/leetcode/base/helper"
	"testing"
	"time"
)

func TestInsertionSort_Do(t *testing.T) {
	tests := []struct {
		arr map[int]int
	}{
		{Helper.GenerateRandomArray(100, 0, 999)},
		{Helper.GenerateRandomArray(10000, 0, 99999)},
	}
	for key, tt := range tests {
		startTime := time.Now()

		count := len(tt.arr)
		s := InsertionSort{
			Arr:   tt.arr,
			Count: count,
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

func TestInsertionSort_DoBetter(t *testing.T) {
	tests := []struct {
		arr map[int]int
	}{
		{Helper.GenerateRandomArray(100, 0, 999)},
		{Helper.GenerateRandomArray(10000, 0, 99999)},
	}
	for key, tt := range tests {
		startTime := time.Now()

		count := len(tt.arr)
		s := InsertionSort{
			Arr:   tt.arr,
			Count: count,
		}
		s.DoBetter()

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
