// The nilness command applies the golang.org/x/tools/go/analysis/passes/nilness
// analysis to the specified packages of Go source code.
package main

import (
	"github.com/yndai/tools/go/analysis/passes/nilness"
	"github.com/yndai/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(nilness.Analyzer) }
