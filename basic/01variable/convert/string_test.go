package convert

import (
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	Convey("string test ", t, func() {
		r1 := StringToRune("abc")
		for _, r := range r1 {
			So(reflect.TypeOf(r).Kind(), ShouldEqual, reflect.Int32)
		}

		s1 := StringSplit("abc")
		So(s1, ShouldResemble, []string{"a", "b", "c"})

	})
}
