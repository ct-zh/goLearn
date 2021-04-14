package cache

// lru.value的实现； 存储真实的缓存值; 只读
// []byte结构比较通用，可以转数字也可以转字符串，更适合网络传输
type ByteView struct {
	b []byte
}

func (v ByteView) Len() int {
	return len(v.b)
}

// 只读，所以获取切片应该是获取切片的复制
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
