package any_parser

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_FindObject(t *testing.T) {
	var objectsName = map[string]struct{}{
		"AnyDir":    {},
		"AnyObject": {},
		"AnyFunc":   {},
		"AnyFile":   {},
	}

	Convey("find object function", t, func() {
		objects, err := FindObject("./")
		So(err, ShouldBeNil)

		for s, object := range objects {
			t.Logf("s: %s, object: %+v", s, object)
			_, ok := objectsName[s]
			if !ok {
				continue
			}

		}
	})
}
