package strings_builder

import (
	"bytes"
	"strings"
	"testing"
)

func TestStringBuilder(t *testing.T) {
	b := strings.Builder{}
	b.WriteString("hello world")
	t.Logf("%s", b.String())
}

func BenchmarkStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sb = strings.Builder{}
		sb.WriteString("hello world!")
	}
}

func BenchmarkBytesBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var bb = bytes.Buffer{}
		bb.WriteString("hello world!")
	}
}
