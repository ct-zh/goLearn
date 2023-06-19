package any_parser

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_NewAnyDir(t *testing.T) {
	Convey("")
	anyDir, err := NewAnyDir("../")
	So(err, ShouldBeNil)

	t.Logf("anyDir: %+v", anyDir)
}
