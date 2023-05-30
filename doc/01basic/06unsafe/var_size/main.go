package main

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/pkg/profile"
)

// 如何 获取 变量占用的内存空间大小

func main() {
	mapSizePProf()
}

func mapSize() {
	var m = make(map[int64]struct{})
	for i := 0; i < 1000000; i++ { // 用pprof测试，100w的map[int64]struct 占用大概 16MB
		m[int64(i)] = struct{}{}
	}

	// unsafe.sizeof
	sizeOf1 := unsafe.Sizeof(m)
	fmt.Printf("size of 1 = %d \n", sizeOf1)

	typeOf1 := reflect.TypeOf(m).Size()
	fmt.Printf("type of 1 = %d \n", typeOf1)
}

// 直接运行 go run main.go
// 然后 go tool pprof -http=:1234 mem.pprof
func mapSizePProf() {
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	mapSize()
}
