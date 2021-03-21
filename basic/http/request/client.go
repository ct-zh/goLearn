package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

// 发起get请求的简单例子
func simpleGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// 创建request，对url发起一次请求
// 这个例子里header的UA部分写了手机标示，所以可能会发生redirect操作
func client(url string) (result []byte, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}

	// 在UA里写入手机的head，访问部分网站可能会跳到m站
	req.Header.Add("User-Agent", ": Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")

	cli := http.Client{
		// 检测是否发生重定向
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect: ", req)
			return nil
		},
	}

	resp, err := cli.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	result, err = httputil.DumpResponse(resp, true)
	return
}
