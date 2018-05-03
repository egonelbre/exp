package runtime2

import (
	"runtime"
	"testing"
)

func BenchmarkRuntimePID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = runtime_pid()
	}
}

func TestAsmPID(t *testing.T) {
	t.Skip("crash")
	for i := 0; i < 100; i++ {
		if int(asm_pid()) != runtime_pid() {
			t.Fatalf("%v %v", int(asm_pid()), runtime_pid())
		}
		runtime.Gosched()
	}
}

func BenchmarkAsmPID(b *testing.B) {
	b.Skip("crash")
	for i := 0; i < b.N; i++ {
		_ = asm_pid()
	}
}

func BenchmarkRDTSCPPID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = rdtscp_pid()
	}
}

func BenchmarkRDPID(b *testing.B) {
	b.Skip("unsupported")
	for i := 0; i < b.N; i++ {
		_ = rdpid_pid()
	}
}

func BenchmarkCPUIDPID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cpuid_pid()
	}
}
