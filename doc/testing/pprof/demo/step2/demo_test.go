package main

import (
	"bufio"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// go test -v -run=^$ -bench=^BenchmarkHi$ -benchtime=2s -memprofile=mem.prof

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
