package main

func add(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11 int) (b1, b2, b3, b4, b5, b6, b7, b8, b9, b10, b11 int) {
	return a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11
}

func main() {
	a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11 := add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	println(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
}
