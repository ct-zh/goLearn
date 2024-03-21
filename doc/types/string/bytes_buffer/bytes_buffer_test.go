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

// buffer的使用方法
func TestBytesBuffer1(t *testing.T) {
	// 零值buffer可以直接使用
	b := bytes.Buffer{}
	// 写入字符串
	n, err := b.WriteString("hello world")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("写入字符串的长度： %d", n)

	// 写入字节
	err = b.WriteByte('!')
	if err != nil {
		t.Fatal(err)
	}

	// 获取长度、容量
	t.Logf("buffer len = %d cap = %d", b.Len(), b.Cap())

	// 获取字符串、字节
	t.Logf("get string = %s, get bytes = %s", b.String(), b.Bytes())

	// 读取单个字节
	if bb, err := b.ReadByte(); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("read byte = %c, current buffer = %s", bb, b.String())
		b.UnreadByte()
	}

	// 读取字节切片，会导致buffer的指针变动
	readData := make([]byte, 12)
	if n, err := b.Read(readData); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("read data = %s, current buffer = %s", readData[:n], b.String())
	}

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
