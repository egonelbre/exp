package queue

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"testing"
	"time"

	"github.com/egonelbre/async"
)

const TestProcs = 4

var BatchSizes = [...]int{1, 8, 32, 64}
var TestSizes = [...]int{1, 2, 3, 7, 8, 9, 15, 127, 128, 129}

func flush(q Queue) {
	if qf, ok := q.(Flusher); ok {
		qf.FlushSend()
	}
}

func mustSend(q NonblockingSPSC, v Value) bool {
	for {
		if q.TrySend(v) {
			return true
		}
		runtime.Gosched()
	}
}

func mustRecv(q NonblockingSPSC, v *Value) bool {
	for {
		if q.TryRecv(v) {
			return true
		}
		runtime.Gosched()
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
		for _, size := range TestSizes {
			q := ctor(size)
			count := Value(size/2) + 1

			result := async.All(func() error {
				for i := Value(0); i < count; i++ {
					if !q.Send(i) {
						return fmt.Errorf("%v: failed to send %v", size, i)
					}
				}
				flush(q)
				return nil
			}, func() error {
				for exp := Value(0); exp < Value(exp); exp++ {
					var got Value
					if !q.Recv(&got) || got != exp {
						return fmt.Errorf("%v: invalid value got %v, expected %v", size, got, exp)
					}
				}
				return nil
			})

			if errs := result.Wait(); errs != nil {
				t.Fatal(errs)
			}
		}
	})

	t.Run("BlockOnFull", func(t *testing.T) {
		for _, size := range TestSizes {
			q := ctor(size)
			capacity := q.Cap()
			for i := 0; i < capacity; i++ {
				if !q.Send(0) {
					t.Fatal("failed to send")
				}
			}
			flush(q)
			sent := uint32(0)
			go func() {
				if !q.Send(0) {
					t.Fatal("failed to send")
				}
				flush(q)
				atomic.StoreUint32(&sent, 1)
			}()
			runtime.Gosched()
			time.Sleep(time.Millisecond)
			if atomic.LoadUint32(&sent) != 0 {
				t.Fatalf("send to full chan")
			}

			var v Value
			if !q.Recv(&v) {
				t.Fatal("failed to recv")
			}
			runtime.Gosched()
			if atomic.LoadUint32(&sent) != 1 {
				runtime.Gosched()
				time.Sleep(time.Millisecond)
				if atomic.LoadUint32(&sent) != 1 {
					t.Fatalf("did not unblock")
				}
			}
		}
	})
}
func testBlockingMPSC(t *testing.T, ctor func(int) BlockingMPSC) {
	t.Run("Basic", func(t *testing.T) {
		for _, size := range TestSizes {
			q := ctor(size)
			count := Value(size/2) + 1

			result := async.All(func() error {
				return async.SpawnWithResult(TestProcs, func(id int) error {
					for i := Value(0); i < count; i++ {
						if !q.Send(Value(id)<<32 | i) {
							return fmt.Errorf("%v: failed to send %v", size, i)
						}
					}
					flush(q)
					return nil
				}).WaitError()
			}, func() error {
				var exps [TestProcs]Value
				for i := Value(0); i < count*TestProcs; i++ {
					var val Value
					if !q.Recv(&val) {
						return fmt.Errorf("%v: failed to get", size)
					}
					id, got := val>>32, val&0xFFFFFFFF
					exp := exps[id]
					exps[id]++
					if exp != got {
						return fmt.Errorf("%v: invalid order got %v, expected %v", size, got, exp)
					}
				}
				return nil
			})

			if errs := result.Wait(); errs != nil {
				t.Fatal(errs)
			}
		}
	})
}
func testBlockingSPMC(t *testing.T, ctor func(int) BlockingSPMC) {
	t.Run("Basic", func(t *testing.T) {
		for _, size := range TestSizes {
			q := ctor(size)
			count := Value(size/2) + 1

			result := async.All(func() error {
				for i := Value(0); i < count*TestProcs; i++ {
					if !q.Send(i) {
						return fmt.Errorf("%v: failed to send %v", size, i)
					}
				}
				flush(q)
				return nil
			}, func() error {
				return async.SpawnWithResult(TestProcs, func(int) error {
					var lastexp Value = -1
					for i := Value(0); i < count; i++ {
						var got Value
						if !q.Recv(&got) {
							return fmt.Errorf("%v: failed to get", size)
						}
						exp := lastexp
						lastexp = got
						if got <= exp {
							return fmt.Errorf("%v: invalid order got %v, expected %v", size, got, exp)
						}
					}
					return nil
				}).WaitError()
			})

			if errs := result.Wait(); errs != nil {
				t.Fatal(errs)
			}
		}
	})
}
func testBlockingMPMC(t *testing.T, ctor func(int) BlockingMPMC) {
	t.Run("Basic", func(t *testing.T) {
		for _, size := range TestSizes {
			q := ctor(size)
			count := Value(size/2) + 1

			result := async.All(func() error {
				return async.SpawnWithResult(TestProcs, func(id int) error {
					for i := Value(0); i < count; i++ {
						if !q.Send(Value(id)<<32 | i) {
							return fmt.Errorf("%v: failed to send %v", size, i)
						}
					}
					flush(q)
					return nil
				}).WaitError()
			}, func() error {
				return async.SpawnWithResult(TestProcs, func(int) error {
					var exps [TestProcs]Value
					for i := range exps {
						exps[i] = -1
					}
					for i := Value(0); i < count; i++ {
						var val Value
						if !q.Recv(&val) {
							return fmt.Errorf("%v: failed to get", size)
						}
						id, got := val>>32, val&0xFFFFFFFF
						exp := exps[id]
						exps[id] = got
						if got <= exp {
							return fmt.Errorf("%v: invalid order got %v, expected %v", size, got, exp)
						}
					}
					return nil
				}).WaitError()
			})

			if errs := result.Wait(); errs != nil {
				t.Fatal(errs)
			}
		}
	})
}

func testNonblockingSPSC(t *testing.T, ctor func(int) NonblockingSPSC) {
	t.Run("Basic", func(t *testing.T) {
		for _, size := range TestSizes {
			q := ctor(size)
			count := Value(size/2) + 1

			result := async.All(func() error {
				for i := Value(0); i < count; i++ {
					if !mustSend(q, i) {
						return fmt.Errorf("%v: failed to send %v", size, i)
					}
				}
				flush(q)
				return nil
			}, func() error {
				for exp := Value(0); exp < Value(exp); exp++ {
					var got Value
					if !mustRecv(q, &got) || got != exp {
						return fmt.Errorf("%v: invalid value got %v, expected %v", size, got, exp)
					}
				}
				return nil
			})

			if errs := result.Wait(); errs != nil {
				t.Fatal(errs)
			}
		}
	})
	t.Run("NonblockOnFull", func(t *testing.T) {
		for _, size := range TestSizes {
			q := ctor(size)
			capacity := q.Cap()
			for i := 0; i < capacity; i++ {
				if !q.TrySend(0) {
					t.Fatal("failed to send")
				}
			}
			if q.TrySend(0) {
				t.Fatal("send succeeded")
			}
			flush(q)
		}
	})
}
func testNonblockingMPSC(t *testing.T, ctor func(int) NonblockingMPSC) {
	t.Run("Basic", func(t *testing.T) {
		for _, size := range TestSizes {
			q := ctor(size)
			count := Value(size/2) + 1

			result := async.All(func() error {
				return async.SpawnWithResult(TestProcs, func(id int) error {
					for i := Value(0); i < count; i++ {
						if !mustSend(q, Value(id)<<32|i) {
							return fmt.Errorf("%v: failed to send %v", size, i)
						}
					}
					flush(q)
					return nil
				}).WaitError()
			}, func() error {
				var exps [TestProcs]Value
				for i := Value(0); i < count*TestProcs; i++ {
					var val Value
					if !mustRecv(q, &val) {
						return fmt.Errorf("%v: failed to get", size)
					}
					id, got := val>>32, val&0xFFFFFFFF
					exp := exps[id]
					exps[id]++
					if exp != got {
						return fmt.Errorf("%v: invalid order got %v, expected %v", size, got, exp)
					}
				}
				return nil
			})

			if errs := result.Wait(); errs != nil {
				t.Fatal(errs)
			}
		}
	})
}
func testNonblockingSPMC(t *testing.T, ctor func(int) NonblockingSPMC) {
	t.Run("Basic", func(t *testing.T) {
		for _, size := range TestSizes {
			q := ctor(size)
			count := Value(size/2) + 1

			result := async.All(func() error {
				for i := Value(0); i < count*TestProcs; i++ {
					if !mustSend(q, i) {
						return fmt.Errorf("%v: failed to send %v", size, i)
					}
				}
				flush(q)
				return nil
			}, func() error {
				return async.SpawnWithResult(TestProcs, func(int) error {
					var lastexp Value = -1
					for i := Value(0); i < count; i++ {
						var got Value
						if !mustRecv(q, &got) {
							return fmt.Errorf("%v: failed to get", size)
						}
						exp := lastexp
						lastexp = got
						if got <= exp {
							return fmt.Errorf("%v: invalid order got %v, expected %v", size, got, exp)
						}
					}
					return nil
				}).WaitError()
			})

			if errs := result.Wait(); errs != nil {
				t.Fatal(errs)
			}
		}
	})
}
func testNonblockingMPMC(t *testing.T, ctor func(int) NonblockingMPMC) {
	t.Run("Basic", func(t *testing.T) {
		for _, size := range TestSizes {
			q := ctor(size)
			count := Value(size/2) + 1

			result := async.All(func() error {
				return async.SpawnWithResult(TestProcs, func(id int) error {
					for i := Value(0); i < count; i++ {
						if !mustSend(q, Value(id)<<32|i) {
							return fmt.Errorf("%v: failed to send %v", size, i)
						}
					}
					flush(q)
					return nil
				}).WaitError()
			}, func() error {
				return async.SpawnWithResult(TestProcs, func(int) error {
					var exps [TestProcs]Value
					for i := range exps {
						exps[i] = -1
					}
					for i := Value(0); i < count; i++ {
						var val Value
						if !mustRecv(q, &val) {
							return fmt.Errorf("%v: failed to get", size)
						}
						id, got := val>>32, val&0xFFFFFFFF
						exp := exps[id]
						exps[id] = got
						if got <= exp {
							return fmt.Errorf("%v: invalid order got %v, expected %v", size, got, exp)
						}
					}
					return nil
				}).WaitError()
			})

			if errs := result.Wait(); errs != nil {
				t.Fatal(errs)
			}
		}
	})
}
