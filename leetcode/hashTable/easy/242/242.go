package _242

func isAnagram(s string, t string) bool {
	saver := map[rune]int{}
	for _, value := range []rune(s) {
		saver[value]++
	}

	for _, value := range []rune(t) {
		if saver[value] > 0 {
			saver[value]--
		} else {
			return false
		}
	}

	for _, value := range saver {
		if value > 0 {
			return false
		}
	}
	return true
}
