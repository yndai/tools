package links

import (
	"fmt" //@link(re`".*"`,"https://godoc.org/fmt")

	"github.com/yndai/tools/internal/lsp/foo" //@link(re`".*"`,"https://godoc.org/golang.org/x/tools/internal/lsp/foo")
)

var (
	_ fmt.Formatter
	_ foo.StructFoo
)
