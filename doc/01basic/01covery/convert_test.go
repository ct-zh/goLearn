package convery

import (
	"fmt"
	"strconv"
	"testing"
)

// int64转换成string的两种方法，基准测试
// 1.strconv.FormatInt
// 2.Sprintf
func BenchmarkInt64ConvertToString1_FormatInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var ii int64 = 10
		_ = strconv.FormatInt(ii, 10)
	}
}
func BenchmarkInt64ConvertToString2_Sprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var ii int64 = 10
		_ = fmt.Sprintf("%d", ii)
	}
}
