package _137

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	m1, m2, m3, r := 0, 0, 1, 1
	for i := 3; i <= n; i++ {
		m1 = m2
		m2 = m3
		m3 = r
		r = m1 + m2 + m3
	}
	return r
}
