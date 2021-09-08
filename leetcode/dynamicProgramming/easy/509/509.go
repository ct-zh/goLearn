package _09

// stupid
func fib(n int) int {
	var f []int
	f = append(f, 0)
	f = append(f, 1)
	for i := 2; i <= n; i++ {
		f = append(f, f[i-1]+f[i-2])
	}
	return f[n]
}

func fib2(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	p, q, r := 0, 0, 1
	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}
