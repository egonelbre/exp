package untangle_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/egonelbre/exp/untangle"
)

func TestSink(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, untangle.Analyzer, "example.test/...")
}
