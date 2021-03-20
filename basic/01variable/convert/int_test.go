package convert

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInt(t *testing.T) {
	Convey("Int test", t, func() {
		// IntToBin
		So(IntToBin(4), ShouldEqual, "100")

		// 转字符串
		So(IntToString(5), ShouldEqual, "5")

		// 转字符串
		So(IntToString2(5), ShouldEqual, "5")

		var a int64
		a = 10
		So(Int64ToString(a), ShouldEqual, "10")

		So(ByteToString([]byte{'a', 'b', 'c'}), ShouldEqual, "abc")

	})
}
