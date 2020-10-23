package insertionSort

import "fmt"

// 插入排序
type InsertionSort struct {
	Arr     map[int]interface{}
	N       int
	Compare func(item interface{}, target interface{}) bool // 比较函数
}

// 插入排序 版本1
// 这个版本的排序对比选择排序，因为会提前break掉，所以看起来性能要优于选择排序
// 实际上测试出来的这个插入排序效率要低于选择排序，因为插入排序在每次做比较时都会swap一次，产生了内存操作
// 建议写个test测试一下
func (i *InsertionSort) Do() {
	// 因为第一个元素arr[0]不需要做判断， 所以循环是从1开始的
	for m := 1; m < i.N; m++ {
		// 寻找元素 arr[m] 合适的插入位置

		// 每次与前面一个数做比较，因为最后是与arr[0]做比较，所以这里 n > 0就可以了
		for n := m; n > 0; n-- {
			if i.Compare(i.Arr[n], i.Arr[n-1]) { // 每次与前一个元素比较
				i.swap(n, n-1)
			} else {
				break // 插入排序相对选择排序的优点：可能会提前结束
			}
		}

		// 可以简化为：
		//for n := m; n > 0 && i.Compare(i.Arr[n], i.Arr[n-1]); n-- {
		//	i.swap(n, n-1)
		//}
	}

	fmt.Println(i.Arr)
}

// 插入排序 优化版
//
func (i *InsertionSort) Do2() {
	for m := 1; m < i.N; m++ {
		save := i.Arr[m]
		fmt.Printf("Save = %d\n", save)

		for n := m; n > 0; n-- { // n: 1
			// 如果不符合条件，说明 n 应该在 n-1 的前面
			// 相对上面的版本，这里我们不交换两个数据，而是仅仅将 n - 1 赋给 n

			if n > 0 && i.Compare(save, i.Arr[n-1]) {
				i.Arr[n] = i.Arr[n-1]
				fmt.Printf("比较：Save: %d 小于 Target(Arr[n-1]): %d 交换后：Arr[n]=%d  Arr[n-1]=%d \n",
					save, i.Arr[n-1], i.Arr[n], i.Arr[n-1])
			} else {
				fmt.Printf("比较：Save: %d 大于 Target(Arr[n-1]): %d 交换： Arr[n]=%d \n", save, i.Arr[n-1], i.Arr[n])
				i.Arr[n] = save
			}
		}
	}

	fmt.Printf("结果： %+v\n", i.Arr)
}

// 优化后的插排比选择排要快很多， 尤其在数组中重复数据较多的情况下
func (i *InsertionSort) Do3() {
	for m := 1; m < i.N; m++ {
		elem := i.Arr[m]
		var n int
		for n = m; n > 0 && i.Compare(elem, i.Arr[n-1]); n-- {
			i.Arr[n] = i.Arr[n-1]
		}
		i.Arr[n] = elem
	}
}

// 交换两个的值
func (i *InsertionSort) swap(k1 int, k2 int) {
	save := i.Arr[k1]
	i.Arr[k1] = i.Arr[k2]
	i.Arr[k2] = save

	//fmt.Printf("交换： k1: %d k2: %d \n", k1, k2)
}

func (i *InsertionSort) P() {
	for k, v := range i.Arr {
		fmt.Printf("Key: %d Value: %+v \n", k, v)
	}
	fmt.Println("=========")
}
