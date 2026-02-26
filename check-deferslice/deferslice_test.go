package deferslice_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	deferslice "github.com/egonelbre/exp/check-deferslice"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, deferslice.Analyzer, "example")
}
