// 测试用函数
package Hepler

import (
	"math/rand"
	"time"
)

// 生成指定范围的整数数组
func GenerateRandomArray(n int, rangeL int, rangeR int) map[int]interface{} {
	if rangeL > rangeR {
		return nil
	}

	// rand
	arr := map[int]interface{}{}

	rand2 := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < n; i++ {
		// 指定范围随机数生成的标准写法
		arr[i] = rand2.Int()%(rangeR-rangeL+1) + rangeL
	}

	return arr
}

func GenerateNearlyArray(n int, swapTimes int) map[int]interface{} {
	var arr map[int]interface{}
	for i := 0; i < n; i++ {
		arr[i] = i
	}

}
