package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/egonelbre/exp/fields/testdata"
)

func main() {
	var err error

	//gistsnip:start:main
	config := &jsonConfig{}
	err = json.NewDecoder(strings.NewReader(testdata.Basic)).Decode(config)
	if err != nil {
		log.Fatal(err)
	}

	var (
		alpha float64
		gamma float64
		beta  uint
	)

	err = config.Scan(
		Float{"Alpha", &alpha},
		Float{"Gamma", &gamma},
		Uint{"Beta", &beta},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(alpha + gamma)
	//gistsnip:end:main
}
