package _976

// 思路1： 暴力解法，算出所有面积不为0的三角形周长，然后排序
// 检查是否能成为三角形 => 较小两个值之和必须大于第三个值
// 面积公式： s = 1/2 * a * h   a为三角形的底，h为三角形的高

// 思路2： 贪心算法， 直接从最大值开始枚举符合条件的另外两边
// 反证法：因为 a + b > c 如果数组最大值c 还大于c-1 + c-2 的值，那么就说明往前的所有边都不会符合条件，说明该边不在最大周长的三角形中，因此排除
func largestPerimeter(A []int) int {
	if len(A) > 15 {
		quicksort(0, len(A)-1, A)
	} else {
		insertSort(A)
	}
	for i := len(A) - 1; i >= 2; i-- {
		// 是一个合法的三角形
		if A[i-2]+A[i-1] > A[i] {
			return A[i] + A[i-1] + A[i-2]
		}
	}
	return 0
}

func insertSort(arr []int) {
	for index := 1; index < len(arr); index++ {
		saver := arr[index]
		n := index
		for ; n > 0 && saver < arr[n-1]; n-- {
			arr[n] = arr[n-1]
		}
		arr[n] = saver
	}
}

func quicksort(start int, end int, arr []int) {
	if end <= start {
		return
	}
	p := partition(start, end, arr)
	quicksort(start, p-1, arr)
	quicksort(p+1, end, arr)
}

func partition(start, end int, arr []int) int {
	f := arr[start]
	p := start
	for i := start + 1; i <= end; i++ {
		if arr[i] < f {
			arr[p+1], arr[i] = arr[i], arr[p+1]
			p++
		}
	}

	arr[start], arr[p] = arr[p], arr[start]
	return p
}
