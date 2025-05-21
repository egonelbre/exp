package testtag

import "fmt"

func PrintValue() {
	fmt.Println(IsTesting)
}

var IsTesting = false