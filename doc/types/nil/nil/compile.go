package main

import "fmt"

// dlv debug compile.go
// b main.main:6
// c
// disass
func main() {
	var s []int            // nil 切片
	var m map[int]struct{} // nil map
	var c chan int         // nil channel
	var i interface{}      // nil interface
	var f func()           // nil func

	fmt.Printf("nil slice = %v,%T,%p\n", s, s, s)
	fmt.Printf("nil map = %v,%T,%p\n", m, m, m)
	fmt.Printf("nil channel = %v,%T,%p\n", c, c, c)
	fmt.Printf("nil interface = %v,%T,%p\n", i, i, i)
	fmt.Printf("nil func = %T,%p\n", f, f)
}
