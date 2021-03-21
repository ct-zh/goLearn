package main

import "fmt"

func main() {
	// 想一想下面几个函数的输出
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}

func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		r += 5
	}()
	return t
}

func f3() (t int) {
	t = 5
	defer func() {
		t += 5
	}()
	return t
}

func f4() (r int) {
	r = 1
	defer func(r int) {
		r = r + 5
	}(r)
	return
}
