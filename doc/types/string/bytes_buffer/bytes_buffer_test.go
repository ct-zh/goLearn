package bytes_buffer

import (
	"bytes"
	"testing"
)

func TestBytesBuffer(t *testing.T) {
	b := bytes.Buffer{}
	b.WriteString("hello world")
	t.Logf("%s", b.String())
}

// copy函数的用法
func Test_copy(t *testing.T) {
	buf1 := make([]byte, 0, 64)
	buf1 = buf1[:4] // 初始化前4项
	t.Logf("len = %d, cap = %d, data = %+v", len(buf1), cap(buf1), buf1)
	copy(buf1[0:], "abcd")
	t.Logf("len = %d, cap = %d, data = %+v", len(buf1), cap(buf1), buf1)

	off := 1
	copy(buf1, buf1[off:])
	t.Logf("len = %d, cap = %d, data = %+v", len(buf1), cap(buf1), buf1)
}

func Test_newSlice(t *testing.T) {
	// length = 10 用这种方法申请出来的b2 cap = 16
	length := 10
	b2 := append([]byte(nil), make([]byte, length)...)
	t.Logf("len = %d, cap = %d, data = %+v", len(b2), cap(b2), b2)

	// 这种情况下申请出来的b3 cap = 10
	b3 := make([]byte, length)
	t.Logf("len = %d, cap = %d, data = %+v", len(b3), cap(b3), b3)
}

func Test_tryGrowByReslice(t *testing.T) {
	b := Buffer{
		buf: make([]byte, 0, 10),
	}

	t.Logf("len = %d, cap = %d, data = %+v", len(b.buf), cap(b.buf), b.buf)
	m, ok := b.tryGrowByReslice(10)
	t.Logf("m = %v, ok =%v", m, ok)
	t.Logf("len = %d, cap = %d, data = %+v", len(b.buf), cap(b.buf), b.buf)
}
