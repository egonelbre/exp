package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	deferslice "github.com/egonelbre/exp/check-deferslice"
)

func main() {
	singlechecker.Main(deferslice.Analyzer)
}
