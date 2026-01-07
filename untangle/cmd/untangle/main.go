package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/egonelbre/exp/untangle"
)

func main() {
	singlechecker.Main(untangle.Analyzer)
}
