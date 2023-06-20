package main

import (
	"github.com/jinzhu/copier"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCopier_struct(t *testing.T) {
	Convey("copy struct", t, func() {
		type A struct {
			A int
			B string
			C interface{}
			D []int
			E map[string]string
		}
		type B struct {
			A int
			E map[string]string
		}

		bb := B{}
		aa := A{
			A: 1,
			B: "2",
			C: 3,
			E: map[string]string{"1": "2"},
		}

		copier.Copy(&bb, &aa)
		So(bb.A, ShouldEqual, aa.A)
		So(bb.E, ShouldResemble, aa.E)
	})
}
