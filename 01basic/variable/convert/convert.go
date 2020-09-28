package convert

import "strconv"

func ToBin(n int) (result string) {
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return
}
