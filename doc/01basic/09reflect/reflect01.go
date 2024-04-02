package reflect

import (
	"fmt"
	"reflect"

	jsoniter "github.com/json-iterator/go"
)

// JsonMapDecode 将 original 中的每一项解json 到target中
// target必须为指向map[string]interface{}的指针，其中key必须是string类型，value可以是任意类型
func JsonMapDecode(original map[string]string, target interface{}) error {
	destMapValue := reflect.ValueOf(target)

	// 检查 target 的类型是否为指针，且其指向的值是 map
	if destMapValue.Kind() != reflect.Ptr || destMapValue.Elem().Kind() != reflect.Map {
		return fmt.Errorf("target must be a pointer to a map")
	}

	// 创建一个新的 map 用于存储结果
	destMap := reflect.MakeMap(destMapValue.Elem().Type())

	for namespace, item := range original {
		// 创建一个新的结构体类型用于解码 JSON
		structType := reflect.New(destMapValue.Elem().Type().Elem()).Interface()

		if err := jsoniter.UnmarshalFromString(item, structType); err != nil {
			return err
		}
		// 将结果放入新 map 中
		destMap.SetMapIndex(reflect.ValueOf(namespace), reflect.ValueOf(structType).Elem())
	}
	// 将新 map 赋值给传入的 target
	destMapValue.Elem().Set(destMap)

	return nil
}
