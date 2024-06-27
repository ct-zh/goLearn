package _map

import (
	"math/rand"
	"testing"
	"time"
)

// map多键索引
// 其实就是传个struct当key，不知道有什么用

type user struct {
	Name string
	Age  int
}

func Test(t *testing.T) {
	u1 := user{Name: "张三", Age: 18}
	u2 := user{Name: "李四", Age: 18}
	m := make(map[user]int)
	m[u1] = 18
	m[u2] = 20
	t.Log(m[u1])
	t.Log(m[u2])
}

// fasthttp 常见优化: 尽可能使用slice代替map
// 运行以下benchmark
// goos: darwin
// goarch: amd64
// pkg: github.com/ct-zh/goLearn/doc/types/map
// cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
// BenchmarkSlice
// BenchmarkSlice-8   	49704374	        24.26 ns/op
// BenchmarkMap
// BenchmarkMap-8     	 7388773	       164.5 ns/op
const size = 1000000

// generateSlice generates a slice with random integer values
func generateSlice(size int) []int {
	rand.Seed(time.Now().UnixNano())
	s := make([]int, size)
	for i := range s {
		s[i] = rand.Intn(size)
	}
	return s
}

// generateMap generates a map with random string keys and integer values
func generateMap(size int) map[int]int {
	rand.Seed(time.Now().UnixNano())
	m := make(map[int]int, size)
	for i := 0; i < size; i++ {
		m[i] = rand.Intn(size)
	}
	return m
}

// Benchmark for slice
func BenchmarkSlice(b *testing.B) {
	s := generateSlice(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s[rand.Intn(size)]
	}
}

// Benchmark for map
func BenchmarkMap(b *testing.B) {
	m := generateMap(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m[rand.Intn(size)]
	}
}
