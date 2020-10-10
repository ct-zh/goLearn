package request

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

type Curl struct {
	Url     string
	Method  string
	Header  []string
	Cookie  string
	Timeout int
}

// 使用 [http/url 包] 拼凑url
func (c *Curl) SetUrl(base string, params map[string]string) (err error) {
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

	c.Url = myUrl.String()
	return
}

func (c *Curl) Exec() ([]byte, error) {
	switch c.Method {
	default:
		resp, err := http.Get(c.Url)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		all, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return all, nil
	}
}
