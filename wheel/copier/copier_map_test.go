package main

import (
	"testing"

	"github.com/jinzhu/copier"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCopier_map(t *testing.T) {
	Convey("copy map", t, func() {
		Convey("copy map int", func() {
			var (
				mapIntA = map[int]int{1: 2, 3: 4}
				mapIntB = map[int]int{}
				err     error
			)
			err = copier.Copy(&mapIntB, &mapIntA)
			So(err, ShouldBeNil)
			So(mapIntB, ShouldResemble, mapIntA)
		})

		Convey("copy map string", func() {
			var (
				mapStringA = map[string]string{"1": "2", "3": "4"}
				mapStringB = map[string]string{}
				err        error
			)
			err = copier.Copy(&mapStringB, &mapStringA)
			So(err, ShouldBeNil)
			So(mapStringB, ShouldResemble, mapStringA)
		})

		Convey("copy map interface", func() {
			var (
				mapInterfaceA = map[interface{}]interface{}{1: "2", "3": 4}
				mapInterfaceB = map[interface{}]interface{}{}
				err           error
			)
			err = copier.Copy(&mapInterfaceB, &mapInterfaceA)
			So(err, ShouldBeNil)
			So(mapInterfaceB, ShouldResemble, mapInterfaceA)
		})
	})
}
