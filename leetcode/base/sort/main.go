package main

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
}
