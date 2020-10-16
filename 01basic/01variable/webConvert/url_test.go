package webConvert

import (
	"fmt"
	"testing"
)

func TestParseUrl(t *testing.T) {
	tests := []struct {
		url string
	}{
		{url: "http://m.baidu.com"},
		{url: "aaabbbccc"},
		{url: "http://localhost:9200/product/all?id=1&name=aaa"},
	}

	for k, i := range tests {
		result, err := ParseUrl(i.url)
		if err != nil {
			t.Errorf("%d Error: %s", k, err.Error())
		} else {
			fmt.Printf("%d Url: %s Result: %+v \n", k, i.url, result)
		}
	}
}
