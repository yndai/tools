// The findcall command runs the findcall analyzer.
package main

import (
	"github.com/yndai/tools/go/analysis/passes/findcall"
	"github.com/yndai/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(findcall.Analyzer) }
