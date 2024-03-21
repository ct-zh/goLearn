package bytes_buffer

import (
	"bytes"
	"io"
)

// tryGrowByReslice是growth的内联版本
// 适用于只需要重新占用内部缓冲区的快速情况
// 它返回应该写入字节的索引以及它是否成功。
func (b *Buffer) tryGrowByReslice(n int) (int, bool) {
	if l := len(b.buf); n <= cap(b.buf)-l { // 判断 当前容量 - 已经分配的内容 是否大于等于 需要写入的长度
		b.buf = b.buf[:l+n] // 可以写入，则将这部分容量初始化并返回true
		return l, true
	}
	return 0, false
}

// grow 增加缓冲区以保证n个字节的空间。
// 返回应该写入字节的索引。
// 如果缓冲区不能增长，它将因ErrTooLarge而panic。
func (b *Buffer) grow(n int) int {
	m := b.Len()

	// 什么情况下会符合这个条件?
	// 如果缓冲区为空（m == 0），但内存并未全部释放（b.off != 0），则重置缓冲区以回收空间。
	if m == 0 && b.off != 0 {
		b.Reset()
	}
	if i, ok := b.tryGrowByReslice(n); ok { // 再次尝试是否不需要扩容
		return i
	}
	if b.buf == nil && n <= smallBufferSize { // 初始化切片，默认容量为64，len=n
		b.buf = make([]byte, n, smallBufferSize)
		return 0
	}

	// 如果已经初始化过了，或者需要初始化一个大切片

	c := cap(b.buf)

	// 如果需要扩充的字节数 n 小于或等于缓冲区容量的一半减去已使用的字节数，
	// 则直接将数据向前移动，腾出空间，避免重新分配新的切片。
	if n <= c/2-m {

		// copy(dst, src)函数将数据从src复制到dst，覆盖dst中已有的数据
		// 这里将buf[off:] 这一部分的数据直接复制到整个buf里面，相当于舍弃掉off之前的已读取数据
		// 比如数据为 [1,2,3,4], 已经Read读取过了1，所以off = 1, copy后结果为 [2,3,4,4]
		// 当前函数最下面会重置buf数据为b.buf[:m+n],此时结果为 [2,3,4], off = 0
		copy(b.buf, b.buf[b.off:])

	} else if c > maxInt-c-n { // 超出了int最大值，报错
		panic(ErrTooLarge)
	} else { // 其他情况下，调用 growSlice 函数来扩充缓冲区
		b.buf = growSlice(b.buf[b.off:], b.off+n)
	}

	// 恢复 b.off 指针和 b.buf 的长度
	b.off = 0
	b.buf = b.buf[:m+n]
	return m
}

// 该函数的作用是扩充切片 b 的容量，使其可以容纳额外的 n 个元素，并保留原始内容。
func growSlice(b []byte, n int) []byte {
	defer func() {
		if recover() != nil {
			panic(ErrTooLarge)
		}
	}()

	// 计算扩充后的容量 c，确保可以容纳 n 个新元素。
	c := len(b) + n

	// 如果扩充后的容量 c 小于当前容量的两倍，则将容量扩充为两倍。
	if c < 2*cap(b) {
		c = 2 * cap(b)
	}

	// 创建一个新的切片 b2，容量为 c。
	// 使用 append 函数将一个空的 byte 切片和一个容量为 c 的 byte 切片拼接起来，形成 b2。
	// ps.这里append出来的切片大小未必等于c，（见Test_newSlice方法）
	//  也就是说这里不使用make来新建切片，而是用这种方式，是为了同时使用append的扩容策略
	b2 := append([]byte(nil), make([]byte, c)...)

	// 将原始切片 b 中的所有元素复制到新切片 b2 中。
	copy(b2, b)

	// 返回一个新的切片，该切片包含原始切片 b 的所有元素，以及额外的 n 个元素的空间。
	return b2[:len(b)]
}

func (b *Buffer) empty() bool { return len(b.buf) <= b.off }

// 从缓冲区中查找分隔符并返回该分隔符之前的子切片
func (b *Buffer) readSlice(delim byte) (line []byte, err error) {
	// IndexByte 返回 b 中 c 的第一个实例的索引，如果 b 中不存在 c，则返回 -1。
	// 使用 IndexByte 函数查找分隔符 (delim) 在缓冲区中 (从 b.off 位置开始) 的索引位置。
	i := bytes.IndexByte(b.buf[b.off:], delim)
	end := b.off + i + 1 // 计算子切片的结束位置,包含分隔符本身(+1)
	if i < 0 {           // 如果小于 0，表示没有找到分隔符，即读取到了缓冲区末尾
		end = len(b.buf)
		err = io.EOF
	}
	line = b.buf[b.off:end] //
	b.off = end             // 更新 b.off 指针，使其指向分隔符之后的下一个字节
	b.lastRead = opRead
	return line, err
}
