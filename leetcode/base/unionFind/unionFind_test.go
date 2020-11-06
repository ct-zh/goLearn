package unionFind

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestNewUnionFind(t *testing.T) {
	tests := []struct {
		n int
	}{
		{n: 100},
		{n: 10000},
	}

	for key, tt := range tests {
		startTime := time.Now()

		// 并查集测试流程
		uf := NewUnionFind(tt.n)

		// 并行为
		for i := 0; i < tt.n; i++ {
			rand2 := rand.New(rand.NewSource(time.Now().UnixNano()))
			a := rand2.Int() % tt.n
			b := rand2.Int() % tt.n
			err := uf.UnionElements(a, b)
			if err != nil {
				t.Error(err)
			}
		}

		// 查行为
		for i := 0; i < tt.n; i++ {
			rand2 := rand.New(rand.NewSource(time.Now().UnixNano()))
			a := rand2.Int() % tt.n
			b := rand2.Int() % tt.n
			_, err := uf.IsConnected(a, b)
			if err != nil {
				t.Error(err)
			}
		}

		endTime := time.Now().Sub(startTime)
		fmt.Printf("[%d] ops: %d 耗时：%.8fs \n", key, tt.n, endTime.Seconds())
	}
}

func TestNewUf2(t *testing.T) {
	tests := []struct {
		n int
	}{
		{n: 100},
		{n: 10000},
	}

	for key, tt := range tests {
		startTime := time.Now()

		// 并查集测试流程
		uf := NewUf2(tt.n)

		// 并行为
		for i := 0; i < tt.n; i++ {
			rand2 := rand.New(rand.NewSource(time.Now().UnixNano()))
			a := rand2.Int() % tt.n
			b := rand2.Int() % tt.n
			err := uf.UnionElements(a, b)
			if err != nil {
				t.Error(err)
			}
		}

		// 查行为
		for i := 0; i < tt.n; i++ {
			rand2 := rand.New(rand.NewSource(time.Now().UnixNano()))
			a := rand2.Int() % tt.n
			b := rand2.Int() % tt.n
			_, err := uf.IsConnected(a, b)
			if err != nil {
				t.Error(err)
			}
		}

		endTime := time.Now().Sub(startTime)
		fmt.Printf("[%d] ops: %d 耗时：%.8fs \n", key, tt.n, endTime.Seconds())
	}
}

func TestNewUf3(t *testing.T) {
	tests := []struct {
		n int
	}{
		{n: 100},
		{n: 10000},
	}

	for key, tt := range tests {
		startTime := time.Now()

		// 并查集测试流程
		uf := NewUf3(tt.n)

		// 并行为
		for i := 0; i < tt.n; i++ {
			rand2 := rand.New(rand.NewSource(time.Now().UnixNano()))
			a := rand2.Int() % tt.n
			b := rand2.Int() % tt.n
			err := uf.UnionElements(a, b)
			if err != nil {
				t.Error(err)
			}
		}

		// 查行为
		for i := 0; i < tt.n; i++ {
			rand2 := rand.New(rand.NewSource(time.Now().UnixNano()))
			a := rand2.Int() % tt.n
			b := rand2.Int() % tt.n
			_, err := uf.IsConnected(a, b)
			if err != nil {
				t.Error(err)
			}
		}

		endTime := time.Now().Sub(startTime)
		fmt.Printf("[%d] ops: %d 耗时：%.8fs \n", key, tt.n, endTime.Seconds())
	}
}

func TestNewUf4(t *testing.T) {
	tests := []struct {
		n int
	}{
		{n: 100},
		{n: 10000},
	}

	for key, tt := range tests {
		startTime := time.Now()

		// 并查集测试流程
		uf := NewUf4(tt.n)

		// 并行为
		for i := 0; i < tt.n; i++ {
			rand2 := rand.New(rand.NewSource(time.Now().UnixNano()))
			a := rand2.Int() % tt.n
			b := rand2.Int() % tt.n
			err := uf.UnionElements(a, b)
			if err != nil {
				t.Error(err)
			}
		}

		// 查行为
		for i := 0; i < tt.n; i++ {
			rand2 := rand.New(rand.NewSource(time.Now().UnixNano()))
			a := rand2.Int() % tt.n
			b := rand2.Int() % tt.n
			_, err := uf.IsConnected(a, b)
			if err != nil {
				t.Error(err)
			}
		}

		endTime := time.Now().Sub(startTime)
		fmt.Printf("[%d] ops: %d 耗时：%.8fs \n", key, tt.n, endTime.Seconds())
	}
}

func TestNewUf5(t *testing.T) {
	tests := []struct {
		n int
	}{
		{n: 100},
		{n: 10000},
	}

	for key, tt := range tests {
		startTime := time.Now()

		// 并查集测试流程
		uf := NewUf5(tt.n)

		// 并行为
		for i := 0; i < tt.n; i++ {
			rand2 := rand.New(rand.NewSource(time.Now().UnixNano()))
			a := rand2.Int() % tt.n
			b := rand2.Int() % tt.n
			err := uf.UnionElements(a, b)
			if err != nil {
				t.Error(err)
			}
		}

		// 查行为
		for i := 0; i < tt.n; i++ {
			rand2 := rand.New(rand.NewSource(time.Now().UnixNano()))
			a := rand2.Int() % tt.n
			b := rand2.Int() % tt.n
			_, err := uf.IsConnected(a, b)
			if err != nil {
				t.Error(err)
			}
		}

		endTime := time.Now().Sub(startTime)
		fmt.Printf("[%d] ops: %d 耗时：%.8fs \n", key, tt.n, endTime.Seconds())
	}
}
