package _861

import (
	"fmt"
	"math"
)

// 贪心算法
// 先保证第一列全为1，再保证每一列1比0多
func matrixScore(A [][]int) int {
	l := len(A)
	if l == 0 {
		return 0
	}
	w := len(A[0])

	// 先把第一列变为1
	for key, value := range A {
		if value[0] != 1 {
			for _key, _value := range A[key] {
				if _value == 1 {
					A[key][_key] = 0
				} else {
					A[key][_key] = 1
				}
			}
		}
	}

	// 0比较多的列进行翻转
	for i := 1; i < w; i++ {
		// flag1 flag0 分别记录1，0的个数
		f1, f0 := 0, 0
		for j := 0; j < l; j++ {
			if A[j][i] == 1 {
				f1++
			} else {
				f0++
			}
		}

		if f0 > f1 {
			for j := 0; j < l; j++ {
				if A[j][i] == 1 {
					A[j][i] = 0
				} else {
					A[j][i] = 1
				}
			}
		}
	}

	//for key := range A {
	//	fmt.Println(A[key])
	//}

	// 计算总和
	sum := 0
	for key := range A {
		itemSum := 0
		for i := w; i > 0; i-- {
			if A[key][w-i] == 1 {
				fmt.Println(w - i)
				itemSum += int(math.Pow(2.0, float64(i-1)))
			}
		}
		//fmt.Println(itemSum)
		sum += itemSum
	}

	return sum
}
