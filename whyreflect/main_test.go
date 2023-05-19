package main

import (
	"bytes"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	data, err := os.ReadFile("testdata/dump.txt")
	if err != nil {
		t.Fatal(err)
	}

	deps := Parse(bytes.NewReader(data))
	if len(deps) != 12034 {
		t.Fatal("expected 12034 entries")
	}
}
