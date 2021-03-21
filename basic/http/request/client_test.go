package request

import (
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"testing"
)

func TestClient(t *testing.T) {
	Convey("client 请求 test", t, func() {
		var result []byte
		result, err := client("http://www.imooc.com/")
		So(err, ShouldBeNil)
		So(result, ShouldNotBeNil)
		log.Println(result)
	})
}

// 先启动 ../server/server.go:simpleServe 服务
func TestSimpleGet(t *testing.T) {
	Convey("get demo", t, func() {
		data, err := simpleGet("http://127.0.0.1:18999")
		So(err, ShouldBeNil)
		So(data, ShouldResemble, []byte("hello world"))
	})
}
