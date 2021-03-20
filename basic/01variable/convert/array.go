package convert

import "strings"

// string 数组转 string  explode
func ArrayToString(s []string) string {
	return strings.Join(s, "")
}
