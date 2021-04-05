package main

import (
	"bufio"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// step.1 可以看到比step0/demo_test.go 快了7倍
// go test -v -run=^$ -bench=. -benchmem
// BenchmarkHi-8   	 2061201ww	       499 ns/op	    3044 B/op	      49 allocs/op

// step.2 go test -v -run=^$ -bench=^BenchmarkHi$ -benchtime=2s -memprofile=mem.prof
// $go tool pprof -sample_index=alloc_space step1.test mem.prof
//

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
	b.ReportAllocs()
}
