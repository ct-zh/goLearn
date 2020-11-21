package convert

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
)

// 转换成2进制的字符串
func ToBin(n int) (result string) {
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return
}

// toString 其他类型转成string，见[strconv包]

// int 转换

// int转换字符串
func IntToString(i int) string {
	return strconv.Itoa(i)
}

// int转换字符串
func IntToString2(i int) string {
	return fmt.Sprintf("%d", i)
}

// int64 转换字符串
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

// string 数组转 string  explode
func ArrayToString(s []string) string {
	return strings.Join(s, "")
}

func ByteToString(b []byte) string {
	return string(b)
}

// go里大小端的问题： binary.BigEndian/ binary.LettleEndian
func ByteToUInt64(n []byte) uint64 {
	//buffer := bytes.NewBuffer(n)
	//var x int32
	//binary.Read(buffer, binary.BigEndian, &x)
	//return int(x)

	return binary.BigEndian.Uint64(n)
}
