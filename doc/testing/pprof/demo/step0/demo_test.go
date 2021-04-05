package main

import (
	"bufio"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// step.1
// go test -v -run=^$ -bench=. -benchmem
// BenchmarkHi-8   	  280515	      3717 ns/op	    3044 B/op	      49 allocs/op

// step.2
// go test -v -run=^$ -bench=^BenchmarkHi$ -benchtime=2s -cpuprofile=cpu.prof
// 生成了cpu.prof step0.test

// step.3 分析
// go tool pprof step0.test cpu.prof
// 使用top -cum命令看到regexp.MatchString函数消耗了1.45s

func BenchmarkHi(b *testing.B) {
	req, err := http.ReadRequest(bufio.NewReader(strings.NewReader("GET /hi HTTP/1.0\r\n\r\n")))
	if err != nil {
		b.Fatal(err)
	}
	rw := httptest.NewRecorder()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		handleHi(rw, req)
	}
}
