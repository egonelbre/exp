package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/egonelbre/exp/fields/testdata"
)

func main() {
	//gistsnip:start:main
	fields, err := ParseFields(strings.NewReader(testdata.Basic))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#+v\n", Add(fields["Alpha"], fields["Beta"]))
	fmt.Printf("%#+v\n", Add(fields["Alpha"], fields["Gamma"]))
	//gistsnip:end:main
}
