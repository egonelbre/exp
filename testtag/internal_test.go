package testtag_test

import (
	"testing"

	"github.com/egonelbre/exp/testtag"
)

func TestX(t *testing.T) {
	t.Log(testtag.Expose())
}
