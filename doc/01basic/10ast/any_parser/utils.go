package any_parser

import (
	"io/fs"
	"strings"
)

func isGoFile(info fs.FileInfo) bool {
	return !info.IsDir() && strings.HasSuffix(info.Name(), ".go")
}
