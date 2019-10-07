// The shadow command runs the shadow analyzer.
package main

import (
	"github.com/yndai/tools/go/analysis/passes/shadow"
	"github.com/yndai/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(shadow.Analyzer) }
