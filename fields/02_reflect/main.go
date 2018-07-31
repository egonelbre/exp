package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/egonelbre/exp/fields/testdata"
)

func main() {
	//gistsnip:start:main
	var example struct {
		Alpha float64
		Gamma float64
		Beta  uint
	}
	err := Unmarshal(strings.NewReader(testdata.Basic), &example)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(example.Alpha + example.Gamma)
	//gistsnip:end:main
}
