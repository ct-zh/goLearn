package main

import (
	"fmt"
	Hepler "github.com/LannisterAlwaysPaysHisDebts/goLearn/leetcode/base/helper"
	"github.com/LannisterAlwaysPaysHisDebts/goLearn/leetcode/base/sort/insertionSort"
)

func main() {
	compare := func(item interface{}, target interface{}) bool {
		return item.(int) < target.(int)
	}

	arr := Hepler.GenerateRandomArray(3, 1, 99)

	i := insertionSort.InsertionSort{
		Arr:     arr,
		N:       len(arr),
		Compare: compare,
	}
	fmt.Printf("原数组：%+v\n", arr)
	i.Do3()
	fmt.Printf("转换后：%+v\n", i.Arr)

	//i.P()
}
