// The lostcancel command applies the golang.org/x/tools/go/analysis/passes/lostcancel
// analysis to the specified packages of Go source code.
package main

import (
	"github.com/yndai/tools/go/analysis/passes/lostcancel"
	"github.com/yndai/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(lostcancel.Analyzer) }
