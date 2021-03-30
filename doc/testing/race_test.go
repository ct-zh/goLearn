package testing

import (
	"testing"
)

//

func TestSet(t *testing.T) {
	fn := func(s Set) {
		s.Add("aaa")
		if !s.Has("aaa") {
			t.Log("key is not exist")
		}
		s.Delete("aaa")
	}

	s := Set{}
	go fn(s)

}

func BenchmarkSet(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		s := Set{}
		for pb.Next() {
			s.Add("aaa")
			if !s.Has("aaa") {
				b.Log("key is not exist")
			}
			s.Delete("aaa")
		}
	})
}

type Set map[string]struct{}

func (s Set) Has(key string) bool {
	_, ok := s[key]
	return ok
}

func (s Set) Add(key string) {
	s[key] = struct{}{}
}

func (s Set) Delete(key string) {
	delete(s, key)
}
