package webConvert

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

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
