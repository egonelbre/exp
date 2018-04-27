package queue

import _ "unsafe"

// func pid() uint32 {
// 	return getg().m.p.ptr().id
// }

//go:linkname procPin runtime.procPin
//go:nosplit
func procPin() int

//go:linkname procUnpin runtime.procUnpin
//go:nosplit
func procUnpin()

func runtime_pid() int {
	pid := procPin()
	procUnpin()
	return pid
}

func rdtscp_pid() int
func rdpid_pid() int
func cpuid_pid() int
