package convert

import (
	"strconv"
	"strings"
)

// string

// string 转 int32
func StringToRune(s string) (r []rune) {
	for _, i := range s {
		r = append(r, i)
	}
	return r
}

// string 转 string array
func StringSplit(s string) []string {
	return strings.SplitAfter(s, "")
}

// string转换byte
func StringToByte(s string) []byte {
	return []byte(s)
}

func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func StringToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func StringToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
