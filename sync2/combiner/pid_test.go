package combiner

import "testing"

func BenchmarkRuntimePID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = runtime_pid()
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
