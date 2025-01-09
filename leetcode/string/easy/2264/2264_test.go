package main

import "testing"

func Test_largestGoodInteger(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{name: "test1", args: "6777133339", want: "777"},
		{name: "test2", args: "2300019", want: "000"},
		{name: "test3", args: "42352338", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestGoodInteger(tt.args); got != tt.want {
				t.Errorf("largestGoodInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
