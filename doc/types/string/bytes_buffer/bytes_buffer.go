package bytes_buffer

import (
	"errors"
	"io"
)

// 解析bytes.buffer
// 该数据结构是一个封送数据的简单字节缓冲区。

// 最小容量分配大小
const smallBufferSize = 64

// Buffer 是具有读写方法的可变大小的字节缓冲区
// 零值的buffer是可以直接使用的空缓冲区
type Buffer struct {
	buf      []byte //
	off      int    // 读取从 buf[off]开始, 写入从 buf[len(buf)]开始
	lastRead readOp // 最后一次读取操作
}

// ReadOp常量描述在缓冲区上执行的最后一个操作
// 以便UnreadRune和UnreadByte可以检查无效使用情况
// OpReadRuneX常量的选择方式是将其转换为int，它们对应于所读取的rune大小。
type readOp int8

const (
	opRead      readOp = -1 // Any other read operation.
	opInvalid   readOp = 0  // Non-read operation.
	opReadRune1 readOp = 1  // Read rune of size 1.
	opReadRune2 readOp = 2  // Read rune of size 2.
	opReadRune3 readOp = 3  // Read rune of size 3.
	opReadRune4 readOp = 4  // Read rune of size 4.
)

const maxInt = int(^uint(0) >> 1)

// ErrTooLarge 如果无法分配内存来在缓冲区中存储数据，则会传递ErrTooLarge
var ErrTooLarge = errors.New("bytes.Buffer: too large")
var errNegativeRead = errors.New("bytes.Buffer: reader returned negative count from Read")

// WriteString 会将S的内容追加到缓冲区，并根据需要增加缓冲区
// 返回值n是s的长度；err始终为nil
// 如果缓冲区变得太大，则WriteString将panic并返回ErrTooLarge。
func (b *Buffer) WriteString(s string) (n int, err error) {
	b.lastRead = opInvalid

	// 判断当前buf的容量够不够写入s,返回的m是当前buf的长度,也就是开始写入的位置
	// 在该函数里，s对应的容量已经初始化;
	// 例如 buf的len=1 ['a'],s=bc,该函数执行完之后,m返回1, buf为['a', '', ''], 已经给bc初始化了位置
	m, ok := b.tryGrowByReslice(len(s))
	if !ok {
		m = b.grow(len(s)) // 容量不够写入，需要扩容
	}
	return copy(b.buf[m:], s), nil // 使用copy写入
}

// 同上
func (b *Buffer) Write(p []byte) (n int, err error) {
	b.lastRead = opInvalid
	m, ok := b.tryGrowByReslice(len(p))
	if !ok {
		m = b.grow(len(p))
	}
	return copy(b.buf[m:], p), nil
}

// Bytes 返回一个长度为b.Len()的切片
// 该切片保存缓冲区的未读部分
// 该片仅在下一次修改缓冲区之前有效
// (即，仅在下一次调用Read、Write、Reset或Truncate等方法之前使用)
// 下一次缓冲区修改之前，切片会对缓冲区内容产生别名
// 因此对切片的即时更改将影响未来读取的结果。
func (b *Buffer) Bytes() []byte { return b.buf[b.off:] }

// String 以字符串形式返回缓冲区未读部分的内容
// 如果缓冲区为空指针，则返回“<nil>”
// 要更高效地构建字符串，请参见strings.Builder类型。
func (b *Buffer) String() string {
	if b == nil {
		// Special case, useful in debugging.
		return "<nil>"
	}
	return string(b.buf[b.off:])
}

func (b *Buffer) Len() int { return len(b.buf) - b.off }
func (b *Buffer) Cap() int { return cap(b.buf) }

func (b *Buffer) Reset() {
	b.buf = b.buf[:0]
	b.off = 0
	b.lastRead = opInvalid
}

// todo
// Read 从缓冲区读取下一个len(P)字节，或直到缓冲区被耗尽。
// 返回值n是读取的字节数。
// 如果缓冲区没有要返回的数据，则err为io.EOF(除非len(P)为零)；否则为nil。
func (b *Buffer) Read(p []byte) (n int, err error) {
	b.lastRead = opInvalid
	if b.empty() {
		// Buffer is empty, reset to recover space.
		b.Reset()
		if len(p) == 0 {
			return 0, nil
		}
		return 0, io.EOF
	}
	n = copy(p, b.buf[b.off:])
	b.off += n
	if n > 0 {
		b.lastRead = opRead
	}
	return n, nil
}
