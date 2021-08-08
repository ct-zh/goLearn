package chrome

import (
	"fmt"
	"net/url"
	"testing"
)

func TestNewChrome(t *testing.T) {
	c, err := NewChrome("/Applications/Google\\ Chrome.app/Contents/MacOS/Google\\ Chrome")
	if err != nil {
		panic(err)
	}
	c.Args = map[string]string{
		"--headless":    "",
		"--disable-gpu": "",
		"-print-to-pdf": "",
	}
	c.Target, err = url.Parse("http://www.baidu.com")

	output, err := c.HtmlToPdf()
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}
