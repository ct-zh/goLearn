package main

import "fmt"

// Dave出的测试题，输出什么?
// https://golang.org/ref/spec#Floating-point_literals

func main() {
	__ := 0x7_7p0
	fmt.Println(__)
	fmt.Printf("%+v %T", __, __)
}
