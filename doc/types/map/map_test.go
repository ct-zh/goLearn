package _map

import (
	"testing"
)

// map多键索引
// 其实就是传个struct当key，不知道有什么用

type user struct {
	Name string
	Age  int
}

func Test(t *testing.T) {
	u1 := user{Name: "张三", Age: 18}
	u2 := user{Name: "李四", Age: 18}
	m := make(map[user]int)
	m[u1] = 18
	m[u2] = 20
	t.Log(m[u1])
	t.Log(m[u2])
}
