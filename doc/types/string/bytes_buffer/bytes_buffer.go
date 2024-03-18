package bytes_buffer

import (
	"errors"
	"io"
	"unicode/utf8"
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

// WriteByte 同上
func (b *Buffer) WriteByte(c byte) error {
	b.lastRead = opInvalid
	m, ok := b.tryGrowByReslice(1)
	if !ok {
		m = b.grow(1)
	}
	b.buf[m] = c
	return nil
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

// WriteRune 将 Unicode 代码点 r 的 UTF-8 编码附加到缓冲区，
// 返回其长度和错误，该错误始终为 nil，
// 但包含在内以匹配 bufio.Writer 的 WriteRune。
// 缓冲区根据需要增长；
// 如果它变得太大，WriteRune 将会因 ErrTooLarge 而panic
func (b *Buffer) WriteRune(r rune) (n int, err error) {
	// 由于 rune 类型可以表示负数，这里将其转换为 uint32 类型进行比较。
	// 如果 rune 字符小于 utf8.RuneSelf，则表示它是一个 ASCII 字符。
	// 这种情况下，直接将其转换为 byte 类型并调用 WriteByte 方法写入到 Buffer 中。
	// 返回写入的字节数 (1) 和空错误 (nil) 表示成功。

	// ps.在 Go 语言中，rune 类型是一个 Unicode 代码点，
	// 可以表示任何 Unicode 字符。
	// UTF-8 是 Unicode 字符的编码方式之一，它使用 1 到 4 个字节来表示一个字符。
	// ASCII 字符集是 Unicode 字符集的一个子集，
	// 它使用 7 个比特来表示每个字符，因此 ASCII 字符的代码点范围是 0 到 127。
	// 如果一个 rune 字符的代码点小于 utf8.RuneSelf (128)，
	// 则它只能使用 1 个字节来表示，因此它一定是一个 ASCII 字符。
	if uint32(r) < utf8.RuneSelf {
		b.WriteByte(byte(r))
		return 1, nil
	}
	b.lastRead = opInvalid
	m, ok := b.tryGrowByReslice(utf8.UTFMax) // 尝试通过重新切片的方式扩充缓冲区，最多扩充 utf8.UTFMax 个字节
	if !ok {
		m = b.grow(utf8.UTFMax)
	}
	b.buf = utf8.AppendRune(b.buf[:m], r) // 使用 utf8.AppendRune 函数将 rune 字符编码为 UTF-8 字节序列并写入到 Buffer 的缓冲区中。
	return len(b.buf) - m, nil
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

// Reset 将缓冲器重置为空，
// 但它保留了底层存储以供将来写入使用。
func (b *Buffer) Reset() {
	b.buf = b.buf[:0]
	b.off = 0
	b.lastRead = opInvalid
}

// Truncate 截断缓冲区 b.buf 的长度，使其只包含前 n 个字节
func (b *Buffer) Truncate(n int) {
	if n == 0 { // 如果要截断的字节数为 0，则直接重置缓冲区并返回
		b.Reset()
		return
	}
	b.lastRead = opInvalid    // 将 lastRead 标记为无效，表示缓冲区内容已更改
	if n < 0 || n > b.Len() { // 检查要截断的字节数是否有效，如果无效则触发 panic 错误
		panic("bytes.Buffer: truncation out of range")
	}
	b.buf = b.buf[:b.off+n] // 将缓冲区 b.buf 截断到 n 个字节
}

// Grow 增加缓冲区的容量，以保证至少可以将n个字节写入缓冲区
// 如果n为负数，Grow会panic。
// 如果缓冲区不能增长，它将因ErrTooLarge而死机。
func (b *Buffer) Grow(n int) {
	if n < 0 {
		panic("bytes.Buffer.Grow: negative count")
	}
	m := b.grow(n)
	b.buf = b.buf[:m]
}

// MinRead 是传递给 Read 调用Buffer.ReadFrom的最小切片大小
// 只要 Buffer 的 MinRead 字节数至少超出保存 r 内容所需的字节数，ReadFrom 就不会增加底层缓冲区。
const MinRead = 512

// ReadFrom 从 r 读取数据直到 EOF, 并将其附加到buffer
// 根据需要设置缓冲区。 返回值n是读取的字节数
// 读取期间遇到的除 io.EOF 之外的错误也会返回
// 如果缓冲区变得太大，ReadFrom 会因 ErrTooLarge 而panic
func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error) {
	b.lastRead = opInvalid // 将 lastRead 标记为无效，表示缓冲区内容已更改
	for {
		i := b.grow(MinRead)                // 调用 grow 函数扩充缓冲区，确保有足够的空间来存储读取到的数据
		b.buf = b.buf[:i]                   //
		m, e := r.Read(b.buf[i:cap(b.buf)]) // 调用 Reader 接口的 Read 方法读取数据，并将读取到的数据写入缓冲区
		if m < 0 {                          // 读取到的字节数为负数则panic
			panic(errNegativeRead)
		}

		b.buf = b.buf[:i+m] // 更新缓冲区 b.buf 的长度，使其包含读取到的数据
		n += int64(m)       // 累加读取到的字节总数 n
		if e == io.EOF {
			return n, nil // e is EOF, so return nil explicitly
		}
		if e != nil {
			return n, e
		}
	}
}

// WriteTo 从 Buffer 中读取数据并写入到 io.Writer 接口实现的对象中
// n: 用来保存写入的字节数的变量
func (b *Buffer) WriteTo(w io.Writer) (n int64, err error) {
	b.lastRead = opInvalid
	if nBytes := b.Len(); nBytes > 0 { // 检查缓冲区是否为空
		// 调用 w.Write 方法将缓冲区数据 (b.buf[b.off:]) 写入到 io.Writer 实现的对象中。
		m, e := w.Write(b.buf[b.off:]) //

		// 检查写入的字节数 (m) 是否大于缓冲区中剩余的字节数 (nBytes)
		if m > nBytes {
			panic("bytes.Buffer.WriteTo: invalid Write count")
		}
		b.off += m // 更新 b.off 指针，使其指向缓冲区中尚未写入的数据部分
		n = int64(m)
		if e != nil {
			return n, e
		}
		// 由于 io.Writer 接口的 Write 方法定义不保证写入所有数据，
		// 所以这里再次检查写入的字节数 (m) 是否等于
		// 缓冲区中剩余的字节数 (nBytes)。
		// 如果不相等，则返回已写入的字节数 (n) 和 io.ErrShortWrite 错误。
		if m != nBytes {
			return n, io.ErrShortWrite
		}
	}
	// Buffer is now empty; reset.
	b.Reset()
	return n, nil
}

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

// Next 将缓冲区中接下来的 n 个字节的切片取出来，并将指针指到最新的位置上；
// 如果缓冲区中的字节数少于 n，则 Next 返回整个缓冲区。
// 该切片仅在下次调用读取或写入方法之前有效。
func (b *Buffer) Next(n int) []byte {
	b.lastRead = opInvalid
	m := b.Len()
	if n > m {
		n = m
	}
	data := b.buf[b.off : b.off+n]
	b.off += n
	if n > 0 {
		b.lastRead = opRead
	}
	return data
}

// ReadByte 从缓冲区中读取一个字节并返回
func (b *Buffer) ReadByte() (byte, error) {
	if b.empty() {
		b.Reset()
		return 0, io.EOF
	}
	c := b.buf[b.off]
	b.off++
	b.lastRead = opRead
	return c, nil
}

// ReadRune 从缓冲区中读取一个 rune 字符并返回
func (b *Buffer) ReadRune() (r rune, size int, err error) {
	if b.empty() { // 如果缓冲区为空，则重置缓冲区并返回 io.EOF 错误表示读取结束
		b.Reset()
		return 0, 0, io.EOF
	}
	c := b.buf[b.off]
	if c < utf8.RuneSelf { // 如果是 ASCII 字符，则将其转换为 rune 字符并返回
		b.off++
		b.lastRead = opReadRune1
		return rune(c), 1, nil
	}

	// 从缓冲区中解码一个 rune 字符;
	// r=保存解码后的rune字符，n=保存解码的字节数;
	r, n := utf8.DecodeRune(b.buf[b.off:])
	b.off += n
	b.lastRead = readOp(n)
	return r, n, nil
}

// UnreadRune 撤销上一次使用 ReadRune 函数读取的 rune 字符
func (b *Buffer) UnreadRune() error {
	if b.lastRead <= opInvalid { // 表示上一次的读取操作不是成功的 ReadRune 操作
		return errors.New("bytes.Buffer: UnreadRune: previous operation was not a successful ReadRune")
	}

	// 检查 off 指针是否大于等于上一次读取的字节数
	// off 指针指向缓冲区中尚未读取的数据部分
	if b.off >= int(b.lastRead) {
		b.off -= int(b.lastRead) // 如果 off 指针大于等于上一次读取的字节数，则将 off 指针后退相应的字节数，相当于将读取的 rune 字符放回缓冲区。
	}
	b.lastRead = opInvalid
	return nil
}

var errUnreadByte = errors.New("bytes.Buffer: UnreadByte: previous operation was not a successful read")

// UnreadByte 回滚上次读取字节的操作
func (b *Buffer) UnreadByte() error {
	if b.lastRead == opInvalid {
		return errUnreadByte
	}
	b.lastRead = opInvalid
	if b.off > 0 {
		b.off--
	}
	return nil
}

// ReadString 从缓冲区中读取一直到分隔符为止的字符串并返回
// 参数delim，表示要读取字符串的分隔符 (byte)
func (b *Buffer) ReadString(delim byte) (line string, err error) {
	slice, err := b.readSlice(delim)
	return string(slice), err
}

func (b *Buffer) ReadBytes(delim byte) (line []byte, err error) {
	slice, err := b.readSlice(delim)
	line = append(line, slice...)
	return line, err
}
