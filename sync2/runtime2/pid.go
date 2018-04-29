package runtime2

import _ "unsafe"

func ProcessorHint() int { return int(asm_pid()) }

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

func asm_pid() uint32

func rdtscp_pid() int
func rdpid_pid() int
func cpuid_pid() int
