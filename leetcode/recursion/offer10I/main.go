package offer10I

// https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/solution/

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return (fib(n-1) + fib(n-2)) % (1e9 + 7)
}

func fib2(n int) int {
	a, b := 0, 1
	for i := 0; i < n-1; i++ {
		a, b = b, (a+b)%1000000007
	}
	return b
}
