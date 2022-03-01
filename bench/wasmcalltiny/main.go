package main

func main(){}

//export add
func Add(a, b float64) float64 {
	return a + b
}

//export sub
func Sub(a, b float64) float64 {
	return a - b
}