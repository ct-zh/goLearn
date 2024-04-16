package mapstructure

import (
	"errors"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

var DataTypeErr = errors.New("invalid data type")

// Parser 使用mapstructure包实现单函数解自定义类型json
// data 实际需要解的结构体指针
// originData 原始json数据用map解开
func Parser(result interface{}, originData interface{}) error {
	if reflect.TypeOf(result).Kind() != reflect.Ptr {
		return DataTypeErr
	}

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		WeaklyTypedInput: true, // 不用强类型匹配
		Result:           result,
		TagName:          "json", // 不使用mapstructure专有标签，而是使用json标签
	})
	if err != nil {
		return err
	}

	if err := decoder.Decode(originData); err != nil {
		return err
	}

	return nil
}
