package main

import "fmt"

func main() {
	type pos [2]int
	a := pos{5, 7}
	b := pos{5, 7}

	fmt.Println(a == b)
	// What is output here?
	// A. true
	// B. false
	// C. 编译错误

}
