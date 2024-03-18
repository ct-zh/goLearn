package strings_builder

import (
	"unicode/utf8"
	"unsafe"
)

// Builder 用Write方法有效地构建字符串
// 它最大限度地减少了内存复制
// 零值即可使用
// 请勿复制非零Builder
type Builder struct {
	addr *Builder
	buf  []byte
}

// 用于欺骗 Go 语言的逃逸分析器
// 实际上是一个恒等函数,即输入和输出的指针是同一个指针
// 但是，由于 noescape 函数的存在，逃逸分析器会认为输出指针和输入指针不是同一个，从而避免将指针分配到堆内存上
// (go:nosplit命令)它指示编译器尽量避免将该函数拆分到多个基本块 (basic block) 中，以提高性能。
// (go:nocheckptr命令)它指示编译器在调用该函数时不进行指针检查，这可能会带来安全隐患，需要谨慎使用。
//
//go:nosplit
//go:nocheckptr
func noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)              // 将输入指针 (p) 转换为 uintptr 类型，该类型表示无符号的内存地址
	return unsafe.Pointer(x ^ 0) // 将 x (内存地址) 进行按位异或 (^) 操作，实际上并没有改变内存地址本身 然后将结果转换为 unsafe.Pointer 类型并返回
}

func (b *Builder) copyCheck() {
	if b.addr == nil {
		// 它是一个解决 Go 语言逃逸分析器缺陷的临时方案 (hack)
		// 这段注释表示这是一个临时性的解决方案，等到 issue 7921 (可能与改进逃逸分析器有关) 被修复后，应该将代码改回原来的样子。
		// 如果 b.addr 为 nil，则将 b 本身通过 noescape 函数处理后转换为 *Builder 类型并赋值给 b.addr。
		// 这是一种欺骗逃逸分析的手段，让 b 不会被分配到堆内存上。
		b.addr = (*Builder)(noescape(unsafe.Pointer(b)))
	} else if b.addr != b { // 这表示 Builder 结构体被复制了，但是 b.addr 字段没有被正确复制
		panic("strings: illegal use of non-zero Builder copied by value")
	}
}

// String returns the accumulated string.
func (b *Builder) String() string {
	// 使用 unsafe.SliceData 函数获取 b.buf 字段的底层数据指针和长度
	// 然后使用 unsafe.String 函数将底层数据指针和长度转换为 string 类型并返回
	return unsafe.String(unsafe.SliceData(b.buf), len(b.buf))
}

func (b *Builder) Write(p []byte) (int, error) {
	b.copyCheck()
	b.buf = append(b.buf, p...)
	return len(p), nil
}

func (b *Builder) WriteString(s string) (int, error) {
	b.copyCheck()
	b.buf = append(b.buf, s...)
	return len(s), nil
}

func (b *Builder) WriteByte(c byte) error {
	b.copyCheck()
	b.buf = append(b.buf, c)
	return nil
}

func (b *Builder) WriteRune(r rune) (int, error) {
	b.copyCheck()
	n := len(b.buf)
	b.buf = utf8.AppendRune(b.buf, r)
	return len(b.buf) - n, nil
}

// Growth 将缓冲区复制到一个新的、更大的缓冲区
// 以便在 len(b.buf) 之外至少有 n 个字节的容量。
func (b *Builder) grow(n int) {
	buf := make([]byte, len(b.buf), 2*cap(b.buf)+n)
	copy(buf, b.buf)
	b.buf = buf
}

// Grow 会增加 b 的容量，以保证在 Grow(n) 之后，至少可以将 n 个字节写入 b，而无需再次分配
// 如果 n 为负数，Grow 会panic。
func (b *Builder) Grow(n int) {
	b.copyCheck()
	if n < 0 {
		panic("strings.Builder.Grow: negative count")
	}
	if cap(b.buf)-len(b.buf) < n {
		b.grow(n)
	}
}

func (b *Builder) Len() int { return len(b.buf) }

func (b *Builder) Cap() int { return cap(b.buf) }

func (b *Builder) Reset() {
	b.addr = nil
	b.buf = nil
}
