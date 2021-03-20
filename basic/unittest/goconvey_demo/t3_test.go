package goconvey_demo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAdd(t *testing.T) {
	// 使用Convey申明一个测试，第二个参数需要传一个*testing.T
	Convey("将两树相加", t, func() {

		// 用So进行一次测试，参数分别为： 测试方法、断言、断言的值
		So(Add(1, 2), ShouldEqual, 3)
	})
}

func TestDivision(t *testing.T) {
	Convey("将两数相除", t, func() {

		// 子Convey第二个参数不需要再传*testing.T了
		Convey("normal ", func() {
			num, err := Division(10, 2)
			So(err, ShouldBeNil)
			So(num, ShouldEqual, 5)
		})

		Convey("exist zero ", func() {
			_, err := Division(10, 0)
			So(err, ShouldNotBeNil)
		})
	})
}

func TestMultiply(t *testing.T) {
	Convey("将两数相乘", t, func() {
		So(Multiply(3, 2), ShouldEqual, 6)
	})
}

func TestSubtract(t *testing.T) {
	Convey("相减", t, func() {
		So(Subtract(2, 1), ShouldEqual, 1)
	})
}
