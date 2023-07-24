package xxhash

import (
	"crypto/sha512"
	"hash/crc64"
	"hash/fnv"
	"testing"

	"github.com/cespare/xxhash/v2"
)

// sha512 生成的hash长度最长，安全性更高，但是也需要更高的计算资源
func TestCrypto_Sha512(t *testing.T) {
	h := sha512.New()
	h.Write([]byte("input data"))
	hash := h.Sum(nil)
	t.Logf("sha512: %x", hash)
}

// 简单快速的哈希算法，适用于非加密场景，常用于数据校验和快速哈希比较
// 不适合用于安全性要求较高的场景，因为其易于遭受哈希碰撞
func TestHash_Fnv(t *testing.T) {
	h := fnv.New64()
	h.Write([]byte("input data"))
	hash := h.Sum64()
	t.Logf("fnv: %x", hash)
}

// crc64循环冗余校验算法，常用于数据校验，特别是在存储和传输数据时进行错误检测
// 提供较高的数据校验能力，因此更适用于数据完整性验证等场景
// 算法计算速度较快，但相比FNV稍微复杂一些 也不适合用于安全性要求较高的场景
func TestHash_Crc64(t *testing.T) {
	table := crc64.MakeTable(crc64.ECMA)
	hashVal := crc64.Checksum([]byte("input data"), table)
	t.Logf("crc64: %x", hashVal)
}

func TestCespare_XxHash(t *testing.T) {
	d := xxhash.New()
	d.WriteString("input data")
	d.Sum64()
	t.Logf("xxhash: %+v", d.Sum64())
}

// go test -bench=. -benchmem

func BenchmarkFnv(b *testing.B) {
	data := []byte("Hello, world!") // 要计算哈希值的数据
	hash := fnv.New64()

	b.ResetTimer()

	b.ReportAllocs() // 报告内存分配情况

	for i := 0; i < b.N; i++ {
		hash.Write(data)
		_ = hash.Sum64()
		hash.Reset()
	}

	b.SetBytes(int64(len(data))) // 设置每次迭代的数据大小
}

func BenchmarkCrc64(b *testing.B) {
	data := []byte("Hello, world!") // 要计算哈希值的数据
	table := crc64.MakeTable(crc64.ECMA)

	b.ResetTimer()

	b.ReportAllocs() // 报告内存分配情况

	for i := 0; i < b.N; i++ {
		_ = crc64.Checksum(data, table)
	}

	b.SetBytes(int64(len(data))) // 设置每次迭代的数据大小
}

func BenchmarkXxHash(b *testing.B) {
	data := []byte("Hello, world!") // 要计算哈希值的数据

	b.ResetTimer()

	b.ReportAllocs() // 报告内存分配情况

	for i := 0; i < b.N; i++ {
		_ = xxhash.Sum64(data)
	}

	b.SetBytes(int64(len(data))) // 设置每次迭代的数据大小
}
