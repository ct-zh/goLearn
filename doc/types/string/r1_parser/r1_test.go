package r1

import (
	"math/rand"
	"testing"
	"time"
)

type ttt struct {
	val int64
}

// goos: darwin
// goarch: amd64
// pkg: github.com/ct-zh/goLearn/doc/types/string/r1_parser
// cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
// BenchmarkMapValue
// BenchmarkMapValue-8   	 5195098	       221.6 ns/op
// BenchmarkMapPtr
// BenchmarkMapPtr-8     	 4190518	       260.3 ns/op

func BenchmarkMapValue(b *testing.B) {
	m := make(map[int64]ttt)
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		val := rand.Int63n(1000000)

		v, ok := m[val]
		if ok {
			v.val += val
		} else {
			v.val = val
		}
		m[val] = v
	}
}

func BenchmarkMapPtr(b *testing.B) {
	m := make(map[int64]*ttt)
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		val := rand.Int63n(1000000)

		v := m[val]
		if v != nil {
			v.val += val
		} else {
			v = &ttt{val: val}
		}
		m[val] = v
	}
}
