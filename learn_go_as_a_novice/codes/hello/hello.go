package main

// go tool compile -S ./hello.go|grep "hello.go:6"
func main() {
	var a = "hello world"
	var b = []byte(a) // runtime.stringtoslicebyte(SB)
	println(b)
}
