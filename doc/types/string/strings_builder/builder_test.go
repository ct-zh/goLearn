package strings_builder

import (
	"strings"
	"testing"
)

func TestStringBuilder(t *testing.T) {
	b := strings.Builder{}
	b.WriteString("hello world")
	t.Logf("%s", b.String())
}
