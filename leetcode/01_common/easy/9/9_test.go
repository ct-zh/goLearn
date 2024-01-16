package l9

import "testing"

var testCases = []struct {
	name     string
	x        int
	expected bool
}{
	{"case1", 121, true},
	{"case2", -121, false},
	{"case3", 10, false},
	{"case4", 12321, true},
}

func Test_isPalindrome(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isPalindrome(tc.x)
			if result != tc.expected {
				t.Errorf("isPalindrome(%d) = %t, want %t", tc.x, result, tc.expected)
			}
		})
	}
}

func Test_isPalindromeV2(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isPalindromeV2(tc.x)
			if result != tc.expected {
				t.Errorf("isPalindrome(%d) = %t, want %t", tc.x, result, tc.expected)
			}
		})
	}
}
