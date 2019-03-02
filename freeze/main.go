package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

/*
#include <sys/types.h>
#include <sys/mman.h>
#include <unistd.h>
*/
import "C"

const PageSize = 4096

type Example struct {
	A, B, C int64
}

func (example *Example) Freeze() *Example {
	// note this must be later explicitly freed
	page := C.malloc(PageSize)

	result := (*Example)(unsafe.Pointer(page))
	*result = *example

	C.mprotect(page, PageSize, syscall.PROT_READ)

	return result
}

func main() {
	a := &Example{1, 2, 3}
	a = a.Freeze()

	a.B = 123
	fmt.Println(a)
}
