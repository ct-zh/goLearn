package convery

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSetUrl(t *testing.T) {
	Convey("set url test", t, func() {
		params := make(map[string]string)
		params["get"] = "1"
		params["name"] = "aa"

		url, err := SetUrl("http://baidu.com", params)
		So(err, ShouldBeNil)
		So(url, ShouldEqual, "http://baidu.com?get=1&name=aa")
	})
}
