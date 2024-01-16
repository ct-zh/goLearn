package l13

import "testing"

func Test_romanToInt(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected int
	}{
		{"test1", "III", 3},
		{"test2", "IV", 4},
		{"test3", "IX", 9},
		{"test4", "LVIII", 58},
		{"test5", "MCMXCIV", 1994},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := romanToInt(c.input)
			if result != c.expected {
				t.Fatalf("expected: %d, but got: %d, input: %s", c.expected, result, c.input)
			}
		})
	}
}
