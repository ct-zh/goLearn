package webConvert

import (
	"fmt"
	"net/url"
)

func ParseUrl(u string) (interface{}, error) {
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
