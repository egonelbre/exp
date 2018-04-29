package runtime2

import (
	"bytes"
	"runtime"
	"strconv"
)

func GOID() int64
func g() uintptr

func goidslow() int64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseInt(string(b), 10, 64)
	return n
}
