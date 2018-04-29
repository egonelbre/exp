// +build !amd64

package runtime2

func GOID() int64 { return goidslow() }
