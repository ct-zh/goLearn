package simple

import (
	"bytes"
	"encoding/json"
	"sync"
	"testing"
)

// 例子2：高并发请求，需要创建大量相同的对象:

// go test -bench=. -benchmem marshal_test.go

type user struct {
	Name   string
	Age    int
	Remark [1024]byte
}

func BenchmarkUnmarshalByPool(b *testing.B) {
	buf, _ := json.Marshal(user{
		Name: "张三",
		Age:  18,
	})
	pool := sync.Pool{New: func() interface{} {
		return &user{}
	}}

	b.StartTimer() // 重置计数器
	for i := 0; i < b.N; i++ {
		stu := pool.Get().(*user)
		_ = json.Unmarshal(buf, stu)
		pool.Put(stu)
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	buf, _ := json.Marshal(user{
		Name: "张三",
		Age:  18,
	})
	b.StartTimer() // 重置计数器
	for i := 0; i < b.N; i++ {
		stu := &user{}
		_ = json.Unmarshal(buf, stu)
	}
}

func TestUnmarshal(t *testing.T) {
	// 假如json类型的buf是客户端的请求数据
	buf, _ := json.Marshal(user{
		Name: "张三",
		Age:  18,
	})
	//t.Logf("%s", buf)

	obj := unmarshal(buf)
	t.Logf("%+v", obj)
}

func unmarshal(buf []byte) *user {
	stu := &user{}
	_ = json.Unmarshal(buf, stu)
	return stu
}

// 例子3：使用pool来复用buffer对象

var bufferPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

var data = make([]byte, 10000)

func BenchmarkBufferWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := bufferPool.Get().(*bytes.Buffer)
		buf.Write(data)
		buf.Reset()
		bufferPool.Put(buf)
	}
}

func BenchmarkBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var buf bytes.Buffer
		buf.Write(data)
	}
}
