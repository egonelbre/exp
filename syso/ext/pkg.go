package ext

//go:generate cc -c -fPIC -o ext.syso ext.c

// extern int f(int);
import "C"

func F(v int) int {
	return int(C.f(C.int(v)))
}
