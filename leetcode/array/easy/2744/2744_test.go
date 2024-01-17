package l2744

import (
	"testing"
)

func Test_maximumNumberOfStringPairs(t *testing.T) {
	cases := []struct {
		name     string
		words    []string
		expected int
	}{
		{"test1", []string{"abc", "def", "cba", "fed"}, 2},
		{"test2", []string{"cd", "ac", "dc", "ca", "zz"}, 2},
		{"test3", []string{"aaa", "aaa", "aaa", "aaa"}, 6},
		{"test4", []string{"ab", "ba", "cc"}, 1},
		{"test5", []string{"aa", "ab"}, 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := maximumNumberOfStringPairs(c.words)
			if got != c.expected {
				t.Errorf("got: %v, want: %v", got, c.expected)
			}
		})
	}
}
