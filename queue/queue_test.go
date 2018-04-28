package queue

import "testing"

const TestSize = 1024

func flush(q Queue) {
	if qf, ok := q.(Flusher); ok {
		qf.FlushSend()
	}
}

func test(t *testing.T, ctor Ctor) {
	t.Run("b", func(t *testing.T) {
		if ctor := ctor.BlockingSPSC(); ctor != nil {
			t.Run("SPSC", func(t *testing.T) { testBlockingSPSC(t, ctor) })
		}
		if ctor := ctor.BlockingSPMC(); ctor != nil {
			t.Run("SPMC", func(t *testing.T) { testBlockingSPMC(t, ctor) })
		}
		if ctor := ctor.BlockingMPSC(); ctor != nil {
			t.Run("MPSC", func(t *testing.T) { testBlockingMPSC(t, ctor) })
		}
		if ctor := ctor.BlockingMPMC(); ctor != nil {
			t.Run("MPMC", func(t *testing.T) { testBlockingMPMC(t, ctor) })
		}
	})

	t.Run("n", func(t *testing.T) {
		if ctor := ctor.NonblockingSPSC(); ctor != nil {
			t.Run("SPSC", func(t *testing.T) { testNonblockingSPSC(t, ctor) })
		}
		if ctor := ctor.NonblockingSPMC(); ctor != nil {
			t.Run("SPMC", func(t *testing.T) { testNonblockingSPMC(t, ctor) })
		}
		if ctor := ctor.NonblockingMPSC(); ctor != nil {
			t.Run("MPSC", func(t *testing.T) { testNonblockingMPSC(t, ctor) })
		}
		if ctor := ctor.NonblockingMPMC(); ctor != nil {
			t.Run("MPMC", func(t *testing.T) { testNonblockingMPMC(t, ctor) })
		}
	})
}

func testBlockingSPSC(t *testing.T, ctor func(int) BlockingSPSC) {
	t.Run("Basic", func(t *testing.T) {
		for _, size := range []int{1, 2, 3, 7, 8, 9, 15, 127, 128, 129} {
			q := ctor(size)
			for i := Value(0); i < Value(size); i++ {
				if !q.Send(i) {
					t.Fatalf("%v: failed to send %v", size, i)
				}
			}
			flush(q)
			for exp := Value(0); exp < Value(exp); exp++ {
				var got Value
				if !q.Recv(&got) || got != exp {
					t.Fatalf("%v: failed to get %v; got %v", size, exp, got)
				}
			}
		}
	})
}
func testBlockingMPSC(t *testing.T, ctor func(int) BlockingMPSC) {

}
func testBlockingSPMC(t *testing.T, ctor func(int) BlockingSPMC) {

}
func testBlockingMPMC(t *testing.T, ctor func(int) BlockingMPMC) {

}

func testNonblockingSPSC(t *testing.T, ctor func(int) NonblockingSPSC) {

}
func testNonblockingMPSC(t *testing.T, ctor func(int) NonblockingMPSC) {

}
func testNonblockingSPMC(t *testing.T, ctor func(int) NonblockingSPMC) {

}
func testNonblockingMPMC(t *testing.T, ctor func(int) NonblockingMPMC) {

}

func bench(b *testing.B, ctor Ctor) {

}
