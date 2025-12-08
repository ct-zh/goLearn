package main

func add(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11 int) (b1, b2, b3, b4, b5, b6, b7, b8, b9, b10, b11 int) {
	return a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11
}

func main() {
	// 使用 16 进制方便在汇编/调试器中观察
	// 预期寄存器分配 (AMD64):
	// 0x11 -> RAX
	// 0x22 -> RBX
	// 0x33 -> RCX
	// 0x44 -> RDI
	// 0x55 -> RSI
	// 0x66 -> R8
	// 0x77 -> R9
	// 0x88 -> R10
	// 0x99 -> R11
	// 0xAA -> Stack (栈)
	// 0xBB -> Stack (栈)
	a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11 := add(
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xAA, 0xBB,
	)
	println(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
}