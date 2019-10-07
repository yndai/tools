// The unmarshal command runs the unmarshal analyzer.
package main

import (
	"github.com/yndai/tools/go/analysis/passes/unmarshal"
	"github.com/yndai/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(unmarshal.Analyzer) }
