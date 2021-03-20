package convert

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/url"
)

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
