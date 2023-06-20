package main

import (
	"github.com/jinzhu/copier"
	"math/rand"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

// 功能1 复制切片
func TestCopier_Slice(t *testing.T) {
	Convey("copy slice", t, func() {
		Convey("copy slice int", func() {
			var (
				sliceIntA = []int{1, 2, 3, 4}
				sliceIntB = []int{}
				err       error
			)
			err = copier.Copy(&sliceIntB, &sliceIntA)
			So(err, ShouldBeNil)
			So(sliceIntB, ShouldResemble, sliceIntA)
		})
		Convey("copy slice string", func() {
			var (
				sliceStringA = []string{"1", "2", "3", "4"}
				sliceStringB = []string{}
				err          error
			)
			err = copier.Copy(&sliceStringB, &sliceStringA)
			So(err, ShouldBeNil)
			So(sliceStringB, ShouldResemble, sliceStringA)
		})
		Convey("copy slice interface", func() {
			var (
				sliceInterfaceA = []interface{}{1, "2", 3, "4"}
				sliceInterfaceB = []interface{}{}
				err             error
			)
			err = copier.Copy(&sliceInterfaceB, &sliceInterfaceA)
			So(err, ShouldBeNil)
			So(sliceInterfaceB, ShouldResemble, sliceInterfaceA)
		})
	})
}

// 比较 copier和 for range 的性能区别
// goos: darwin
//goarch: amd64
//cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
//BenchmarkCopier_slice
//BenchmarkCopier_slice-8       	      51	  22054019 ns/op
//BenchmarkCopySliceByRange
//BenchmarkCopySliceByRange-8   	    1231	    968267 ns/op
func BenchmarkCopier_slice(b *testing.B) {
	const testSize = 100000
	genSlice := make([]int, testSize)
	rand.Seed(time.Now().UnixMilli())
	for n := 0; n < testSize; n++ {
		genSlice[n] = rand.Intn(testSize)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var (
			sliceIntB = []int{}
			err       error
		)
		err = copier.Copy(&sliceIntB, &genSlice)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkCopySliceByRange(b *testing.B) {
	const testSize = 100000
	genSlice := make([]int, testSize)
	rand.Seed(time.Now().UnixMilli())
	for n := 0; n < testSize; n++ {
		genSlice[n] = rand.Intn(testSize)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var sliceIntB = []int{}
		for _, i2 := range genSlice {
			sliceIntB = append(sliceIntB, i2)
		}
	}
}
