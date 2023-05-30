package _1basic

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// convey用法看这个文件就够了
func TestInt(t *testing.T) {
	Convey("Int test", t, func() { // 1. 用Convey申明一个测试开始
		Convey("test 2", t, func() { // 2. Convey里面可以包另外一个convey
			So(100, ShouldEqual, 100) // 3. So是断言，如果不符合，会输出错误信息
			// 有哪些断言？点一下ShouldEqual看看吧，或者这个 https://github.com/smartystreets/goconvey/wiki/Assertions

			So("aaa", shouldLetMeTest, "aaa") // 4. 也可以自定义断言
		})

		// 5. skip，可以不改变import内容，实际不执行。
		// 在测试写完后，某次比较小的变动，只想单独执行某些测试，但是又不想大块大块屏蔽代码时使用
		// 比如
		//Convey("not skip", t, func() {
		//	Println("aaaa")
		//})
		// 直接把方法修改为SkipConvery即可
		SkipConvey("skip", t, func() {
			Println("aaaa")
			SkipSo(1, ShouldBeLessThan, 10) // skipSo一样的用法
		})

		// 6.同理focus与skip正相反
		//FocusConvey("focus", t, func() {})

	})
}

func shouldLetMeTest(actual interface{}, expected ...interface{}) string {
	if actual.(string) == expected[0].(string) {
		return "" // 返回空字符串代表成功
	}
	return "error"
}
