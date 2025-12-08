package main

func use() {
	// 随意计算，破坏寄存器
	_ = 1 + 2 + 3
}

func main() {
	x := 10
	use()
	_ = x
}
