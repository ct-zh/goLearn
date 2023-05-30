package go_cn

import "fmt"

// go cn 2021时 Dave出的测试题，输出什么?
// https://golang.org/ref/spec#Floating-point_literals

func literals() {
	__ := 0x7_7p0
	fmt.Println(__)
	fmt.Printf("%+v %T", __, __)
}
