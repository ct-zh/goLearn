package main

import "fmt"

// https://leetcode-cn.com/problems/move-zeroes/

// 方法一，额外开辟一个数组。
// 注意这个方法是不符合题目条件的，仅仅用来引导思路
// 时间复杂度O(n)
// 空间复杂度O(n)
func moveZeroes(nums []int) {
	l := len(nums)
	if l <= 0 {
		return
	}

	// 将非零的数组都添加进新的数组
	var tmp []int
	for i := 0; i < l; i++ {
		if nums[i] != 0 {
			tmp = append(tmp, nums[i])
		}
	}

	for i := 0; i < len(tmp); i++ {
		nums[i] = tmp[i]
	}
	for i := len(tmp); i < l; i++ {
		nums[i] = 0
	}
}

// 方法二，增加单个索引
// 时间复杂度O(n)
// 空间复杂度O(1)
func moveZeroes2(nums []int) {
	if len(nums) <= 0 {
		return
	}
	index := 0 // 从[0...index)的元素均为非0元素

	// 遍历到第i个元素后，保证[0...i]中所有非0元素
	// 都按照顺序排列在[0...k)中
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[index] = nums[i]
			index += 1
		}
	}

	// 将nums剩余的位置放置为0
	for i := index; i < len(nums); i++ {
		nums[i] = 0
	}

}

//
func moveZeroes3(nums []int) {
	if len(nums) <= 0 {
		return
	}
	index := 0 // 从[0...index)的元素均为非0元素

	// 遍历到第i个元素后，保证[0...i]中所有非0元素
	// 都按照顺序排列在[0...k)中
	// 同时[k...i]为0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if index != i { // 防止自己交换自己
				nums[index], nums[i] = nums[i], nums[index]
			}
			index += 1
		}
	}
}

func main() {
	tests := []struct {
		arr    []int
		result []int
	}{
		{
			arr:    []int{0, 1, 0, 3, 12},
			result: []int{1, 3, 12, 0, 0},
		},
		{
			arr:    []int{0, 0, 0, 0, 0, 0},
			result: []int{0, 0, 0, 0, 0, 0},
		},
		{
			arr:    []int{1, 1, 1, 1},
			result: []int{1, 1, 1, 1},
		},
		{
			arr:    []int{-959151711, 623836953, 209446690, -1950418142, 0 - 1626162038},
			result: []int{-959151711, 623836953, 209446690, -1950418142, -1626162038, 0},
		},
	}

	for key, tt := range tests {
		moveZeroes3(tt.arr)
		for i := 0; i < len(tt.arr); i++ {
			if tt.result[i] != tt.arr[i] {
				panic(fmt.Sprintf("error key: %d result:%+v answer: %+v", key, tt.arr, tt.result))
			}
		}
	}
	fmt.Println("OK!")
}
