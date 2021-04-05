package main

import "testing"

func BenchmarkGetProduct(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			GetOneProduct()
		}
	})
}
