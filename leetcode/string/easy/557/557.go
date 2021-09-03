package _57

func reverseWords(s string) string {
	bs := []byte(s)
	i1, i2, l := 0, 0, len(s)
	for i1 < l && i2 <= l {
		if i2 < l && bs[i2] != ' ' {
			i2++
			continue
		}
		m, n := i1, i2-1
		for m < n {
			bs[m], bs[n] = bs[n], bs[m]
			m++
			n--
		}
		i2++
		i1 = i2
	}
	return string(bs)
}
