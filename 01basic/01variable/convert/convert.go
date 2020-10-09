package convert

import (
	"fmt"
	"strconv"
	"strings"
)

func ToBin(n int) (result string) {
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return
}

// toString 其他类型转成string，见[strconv包]

// int 转换

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func IntToString2(i int) string {
	return fmt.Sprintf("%d", i)
}

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

// string
//  [strings 包] 有string的常用操作

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

func StringToByte(s string) []byte {
	return []byte(s)
}

func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func StringToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}
