package call

//void nop() { }
import "C"

//go:noinline
func Nop() {}

func CNop() { C.nop() }
