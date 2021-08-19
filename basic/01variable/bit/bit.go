package main

import "fmt"

func main() {
	// 按位或运算
	// 参加运算的两个数只要两个数中的一个为1，结果就为1。
	fmt.Printf("12 | 10 的十进制结果是%d, 二进制结果是%b\n", 12|10, 12|10)

	// 按位异或运算
	fmt.Printf("12 ^ 10 的十进制结果是%d, 二进制结果是%b\n", 12^10, 12^10)

	// 按位与运算
	fmt.Printf("12 & 10 的十进制结果是%d, 二进制结果是%b\n", 12&10, 12&10)

	// 位移 左移运算
	fmt.Println(12 << 1)
	fmt.Println(12 << 2)
	fmt.Println(12 << 3)
	fmt.Println(12 << 4)

	// 右移运算
	fmt.Println(12 >> 1)
	fmt.Println(12 >> 2)
	fmt.Println(12 >> 3)
}
