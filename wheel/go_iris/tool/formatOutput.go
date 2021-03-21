package tool

import (
	"encoding/json"
	"strconv"
)

type Output struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 输出的data 是一个字符串
func FormatOutput(code int, data interface{}, msg string) interface{} {
	j, err := json.Marshal(data)
	if err != nil {
		return map[string]string{
			"code": "-1",
			"msg":  "format error"}
	}

	return map[string]string{
		"code": strconv.Itoa(code),
		"msg":  msg,
		"data": string(j)}
}
