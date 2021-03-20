package convert

import (
	"encoding/binary"
	"fmt"
	"strconv"
)

// 转换成2进制的字符串
func IntToBin(n int) (result string) {
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return
}

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
