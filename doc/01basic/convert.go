package _1basic

// 各种类型的转换

import (
	"encoding/binary"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// ArrayToString 数组转 string = explode
func ArrayToString(s []string) string {
	return strings.Join(s, "")
}

// IntToBin 转换成2进制的字符串
func IntToBin(n int) (result string) {
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return
}

// IntToString int转换字符串
func IntToString(i int) string {
	return strconv.Itoa(i)
}

// IntToString2 int转换字符串
func IntToString2(i int) string {
	return fmt.Sprintf("%d", i)
}

// Int64ToString int64 转换字符串
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func ByteToString(b []byte) string {
	return string(b)
}

// ByteToUInt64 go里大小端的问题： binary.BigEndian/ binary.LettleEndian
func ByteToUInt64(n []byte) uint64 {
	//buffer := bytes.NewBuffer(n)
	//var x int32
	//binary.Read(buffer, binary.BigEndian, &x)
	//return int(x)

	return binary.BigEndian.Uint64(n)
}

// StringToRune string 转 int32
func StringToRune(s string) (r []rune) {
	for _, i := range s {
		r = append(r, i)
	}
	return r
}

// StringSplit string 转 string array
func StringSplit(s string) []string {
	return strings.SplitAfter(s, "")
}

// StringToByte string转换byte
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

func SetUrl(base string, params map[string]string) (str string, err error) {
	myUrl, err := url.Parse(base)
	if err != nil {
		return
	}

	if len(params) > 0 {
		urlParams := url.Values{}
		for k, i := range params {
			urlParams.Set(k, i)
		}
		myUrl.RawQuery = urlParams.Encode()
	}

	str = myUrl.String()
	return
}

func ParseUrl(u string) (*url.URL, error) {
	parse, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Scheme: %s ", parse.Scheme)
	values, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Values: %+v \n", values)

	return parse, nil
}

func JsonEncodeAndDecode(s interface{}, save interface{}) error {
	// j 是encode后的[]byte
	j, err := json.Marshal(s)
	if err != nil {
		return err
	}
	fmt.Printf("%s \n %+v\n", j, s)

	// 将j decode 到 save
	json.Unmarshal(j, save)
	fmt.Printf("%+v\n", save)

	return nil
}

func XmlEncodeAndDecode(s interface{}, save interface{}) error {
	j, err := xml.Marshal(s)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", j)

	xml.Unmarshal(j, save)

	fmt.Printf("%+v\n", save)

	return nil
}

func UrlEncode(s string) string {
	return ""
}

func UrlDecode(s string) string {
	return ""
}
