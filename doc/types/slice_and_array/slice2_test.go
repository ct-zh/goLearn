package slice

import (
	"math/rand"
	"runtime"
	"testing"
	"time"
)

// 随机生成n个整数并写入切片
func generateWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

// 打印出当前的内存占用
func printMem(t *testing.T) {
	t.Helper()
	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)
	t.Logf("%.2f MB", float64(rtm.Alloc)/1024./1024.)
}

func testLastChars(t *testing.T, f func([]int) []int) {
	t.Helper()
	ans := make([][]int, 0)
	for k := 0; k < 100; k++ { // 生成100个大小为1M的切片
		origin := generateWithCap(128 * 1024) // 1M = (64/8 * 128 * 1024)字节
		ans = append(ans, f(origin))
	}
	printMem(t)
	_ = ans
}

func TestLastCharsBySlice(t *testing.T) { testLastChars(t, lastNumsBySlice) }

func TestLastCharsByCopy(t *testing.T) { testLastChars(t, lastNumsByCopy) }
