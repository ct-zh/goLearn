package object

import "testing"

// question 1
func TestMa(t *testing.T) {
	m := Ma{}
	m.getName() // 会输出什么？

	m2 := &Ma{}
	m2.getName()
}

// question 2
func TestQuote(t *testing.T) {
	myTest()
}
