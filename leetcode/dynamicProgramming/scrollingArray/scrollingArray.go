package main

// 滚动数组
import "fmt"

// 正常打印斐波那契数列
func normalFib(n int) {
	fib := make([]int, n)
	fib[0] = 1
	fib[1] = 1
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	fmt.Println(fib[len(fib)-3:])
}

// 使用滚动数组打印斐波那契数列
func scrollArrayFib(n int) {
	fib := make([]int, 3)
	fib[1] = 1
	fib[2] = 1
	for i := 2; i < n; i++ {
		fib[0] = fib[1]
		fib[1] = fib[2]
		fib[2] = fib[0] + fib[1]
	}
	fmt.Println(fib)
}

// 使用滚动数组打印斐波那契数列
func scrollArrayFib2(n int) {
	fib := make([]int, 3)
	fib[0] = 1
	fib[1] = 1
	for i := 2; i < n; i++ {
		fib[i%3] = fib[(i-1)%3] + fib[(i-2)%3]
	}
	fmt.Println(fib[n%3], fib[(n-2)%3], fib[(n-1)%3])
}

func main() {
	normalFib(10)
	scrollArrayFib(10)
	scrollArrayFib2(10)
}
