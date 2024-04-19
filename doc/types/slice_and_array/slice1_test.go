package slice

import "testing"

func Test_foo(t *testing.T) {
	foo1()
	foo2()
	foo3()
}

func BenchmarkFoo1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foo1()
	}
}

func BenchmarkFoo2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foo2()
	}
}

func BenchmarkFoo3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foo3()
	}
}
